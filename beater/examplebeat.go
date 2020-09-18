package beater

import (
	"fmt"
	"os/exec"
	"time"
	"bytes"
  "strconv"
	
	"github.com/gorhill/cronexpr"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/PavelStupnitsky/examplebeat/config"
)

type Schedule cronexpr.Expression

// Examplebeat configuration.
type Examplebeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// New creates an instance of examplebeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Examplebeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

func CheckSchedule(c config.Script) *time.Ticker{
	var ticker *time.Ticker
	if c.ScheduleType == "cron"{
		Schedule, errPars := cronexpr.Parse(c.Period)
		if errPars != nil{
			panic(errPars)
		}
		temp := Schedule.Next(time.Now())
		ticker = time.NewTicker(temp.Sub(time.Now()))
	}else{
		temp, errPars := time.ParseDuration(c.Period)
		if errPars != nil{
			panic(errPars)
		}
		ticker = time.NewTicker(temp)
	}
	return ticker
}

func CreateGorutinScript(c config.Script, chanalOne chan string, chanalTwo chan string, i string){
  for{
    ticker := CheckSchedule(c)
    output := OutputScript(c)
    <-ticker.C
    chanalOne <- output
    chanalTwo <- i
  }
}

func PublishScript(bt *Examplebeat, chanalOne chan string, chanalTwo chan string){
  for{
    output := <-chanalOne
    index, err := strconv.Atoi(<-chanalTwo)
    if err != nil{
      panic(err)
    }
    event := beat.Event{ 
      Timestamp: time.Now(),
      Fields: common.MapStr{
        "output": output,
        "input_script": bt.config.Schedule[index],
      },
    }
    fmt.Println(output)
    bt.client.Publish(event)
    logp.Info("Event sent")
  }
}


func OutputScript(c config.Script) string{
		cmd := exec.Command(c.Command, c.Args)
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		errr := cmd.Run()
		if errr != nil{
			fmt.Println(fmt.Sprint(errr) + ": " + stderr.String())
		}
   
    return out.String()
}



// Run starts examplebeat.
func (bt *Examplebeat) Run(b *beat.Beat) error {
	logp.Info("examplebeat is running! Hit CTRL-C to stop it.")
	fmt.Println(bt.config.Schedule)

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}
	
  outputScr := make(chan string, len(bt.config.Schedule))
  indexScr := make(chan string, len(bt.config.Schedule))
   
  for i:=0; i<len(bt.config.Schedule); i++{
    go CreateGorutinScript(bt.config.Schedule[i], outputScr, indexScr, strconv.Itoa(i))
  }
  
  go PublishScript(bt, outputScr, indexScr)
  
  select{
  case <-bt.done:
    return nil
  }
}



// Stop stops examplebeat.
func (bt *Examplebeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
