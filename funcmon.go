package funcmon

import (
	"fmt"
	"log"
	"time"
	"github.com/influxdb/influxdb/client"
)

//Config is used to specify what server to connect to.
//Host: The Hostname of the server to connect to
//Port: The port of the influx db to connect to
//DB: The name of the database to store the metrics in
type Config struct {
	Host string
	Port int
	DB string
}

func NewConfig() Config {
	return Config
}

type Client struct {
	host string
	port int
	db string
}

func NewClient(c Config) (*Client, error) {
	client := Client{
		host: c.Host,
		port: c.Port,
		db: c.DB,
	}

	return &client,nil
}

// Start a timer for the function name provided.
func startMonitoring(key string) {

}

// End the timer for the provided key and add the time metric to 
// the batch
func endMonitoring(key string) {

}

func flushMetrics() {

}
