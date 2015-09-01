package funcmon

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/nickrobinson/funcmon"
)

func TestStub(t *testing.T) {
    assert.True(t, true, "This is good. Canary test passing")
}

func TestNewConfig(t *testing.T) {
	config := funcmon.NewConfig()
	assert.Equal(t, config.Port, 8086)
	assert.NotNil(t, config.Host, "Host field not present")
	assert.NotNil(t, config.DB, "DB field not present")
}

func TestNewClient(t *testing.T) {
	config := funcmon.Config{
		Host: "127.0.0.1",
		Port: 8086,
		DB: "funcmon",
	}
	client,err := funcmon.NewClient(config)
	assert.Nil(t, err, "Error is not nil")
	assert.NotNil(t, client, "Client is not initialized")
	
}

func TestStartMonitoring(t *testing.T) {
	config := funcmon.Config{
		Host: "127.0.0.1",
		Port: 8086,
		DB: "funcmon",
	}
	client,err := funcmon.NewClient(config)
	client.StartMonitoring("test")
	assert.Nil(t, err, "Error is not nil")
}
