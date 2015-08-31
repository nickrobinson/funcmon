package funcmon

import (
	"time"
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
	return Config{
		Port: 8086,
	}
}

type Client struct {
	host string
	port int
	db string
	metricMap map[string]time.Time
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
func (c *Client) startMonitoring(key string) {
	c.metricMap[key] = time.Now()
}

// End the timer for the provided key and add the time metric to 
// the batch
func (c *Client) endMonitoring(key string) {

}

func (c *Client) flushMetrics() {

}
