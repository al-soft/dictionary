package logger

import (
	"fmt"
	"os"

	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	gas "github.com/firstrow/goautosocket"
	log "github.com/sirupsen/logrus"
	yaml "github.com/tsuru/config"
)

// Log данные о лог файле
type Log struct {
	File     string
	Logstash string
	System   string
	Module   string
	Elastic  string
}

// Run создание и открытие файла логирования
func Run(configFile *string) error {
	yaml.ReadConfigFile(*configFile)

	l := Log{}
	l.Logstash, _ = yaml.GetString("logstash:host")
	l.System, _ = yaml.GetString("logstash:system")
	l.Module, _ = yaml.GetString("logstash:module")
	l.Elastic, _ = yaml.GetString("logstash:elastic_index")
	l.File, _ = yaml.GetString("log:file")

	p2, _ := os.Getwd()
	fmt.Println(p2)

	file, err := os.OpenFile(l.File, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	log.SetOutput(file)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetLevel(log.WarnLevel)

	if err := yaml.ReadConfigFile(*configFile); err != nil {
		return err
	}

	logstashconn, err := gas.Dial("tcp", l.Logstash)
	if err != nil {
		return err
	}

	hook := logrustash.New(logstashconn, logrustash.DefaultFormatter(log.Fields{
		"module":     l.Module,
		"system":     l.System,
		"index_name": l.Elastic}))
	log.AddHook(hook)

	log.Warn("...START LOGGING...")
	return nil
}
