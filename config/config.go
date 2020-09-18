// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import (
	//"time"
	)

type Script struct{
	Command       string  `script:"command"`
	Period        string  `script:"period"`
	Args          string  `script:"args"`
	ScheduleType  string  `script:"scheduletype"`
}

type Config struct {
	Schedule  []Script  `config:"schedule"`
}

var DefaultConfig = Config{
	Schedule: []Script{
	{ Command: ".", Period: ".", Args: ".", ScheduleType: "."},
	},
}
