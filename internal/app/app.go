package app

import (
	"dictionary/internal/config"
	"dictionary/internal/logger"
)

var (
	cf  config.Config
	err error
)

func Run() {
	if cf, err = config.New(); err != nil {
		panic(err)
	}

	if err = logger.Run(cf.File); err != nil {
		panic(err)
	}

}
