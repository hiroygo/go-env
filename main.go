package main

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v6"
)

type config struct {
	Home     string `env:"HOME"`
	Port     int    `env:"PORT" envDefault:"3000"`
	Password string `env:"PASSWORD,unset"`

	// While required demands the environment variable to be check,
	// it doesn't check its value. If you want to make sure the environment is set and not empty,
	// you need to use the notEmpty tag option instead
	IsProduction bool `env:"PRODUCTION,required"`

	Hosts      []string      `env:"HOSTS" envSeparator:":"`
	Duration   time.Duration `env:"DURATION"`
	TempFolder string        `env:"TEMP_FOLDER" envDefault:"${HOME}/tmp" envExpand:"true"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", cfg)
}
