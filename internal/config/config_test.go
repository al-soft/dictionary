package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigFile(t *testing.T) {
	os.Chdir("../../configs")

	testCF, err := New()
	if err != nil {
		fmt.Println(err)
	}

	actual := *testCF.File
	expected := "config.local.yml"
	assert.Equal(t, expected, actual)

	assert.Equal(t, nil, err)
}
