// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import (
	"fmt"
	"io/ioutil"
	"time"
)

type Config struct {
	Period time.Duration `config:"period"`
	Path string `config:"path"`
}

var DefaultConfig = Config{
	Period: 10 * time.Second,
	Path: ".",
}

type Examplebeat struct {
	done chan struct{}
	config config.Config
	client publisher.Client
	lastIndexTime time.Time
}

func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}
	ex := &Examplebeat{
		done: make(chan struct{})
		config: config
	}
	return ex, nil
}

func (bt *Lsbeat) listDir(dirFile string, beatname string) {
	files, _ := ioutil.ReadDir(dirFile)
	for _, f := range files {
		t := f.ModTime()
		path := filepath.Join(dirFile, f.Name())
		if t.After(bt.lastIndex) {
			event := common.MapStr{
				"@timestamp": common.Time(time.Now()),
				"type": beatname,
				"modtime": common.Time(t)
				"filename": f.Name(),
				"path": path,
				"directory": f.IsDir(),
				"filesize": f.Size(),
			}
			bt.client.PublishEvent(event)
		}
		if f.IsDir() {
			bt.listDir(path, beatname)
		}
	}
}

func (bt *Examplebeat) Run(b *beat.Beat) error {
	logp.Info("Examplebeat is running! Hit CTRL-C to stop it.")
	bt.client = b.Publisher.Connect()
	ticker := time.NewTicker(bt.config.Period)
	for {
		now := time.Now()
		bt.listDir(bt.config.Path, b.Name)
		bt.lastIndexTime = now
		logp.Info("Event sent")
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}
	}
}

func (bt *Examplebeat) Stop() {
	bt.client.Close()
	close(bt.done)
}