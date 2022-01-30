package logger

import (
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestRunLog(t *testing.T) {
	os.Chdir("../../configs")
	testConfigFile := "config.local.yml"
	var cfgFile *string = &testConfigFile

	err := Run(cfgFile)
	if err != nil {
		panic(err)
	}

	log.Warn("Test log message")

	actual := err
	var expected error

	assert.Equal(t, expected, actual)
}
