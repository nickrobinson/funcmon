package funcmon

import (
	"net/url"
	"fmt"
	"log"
	"os"
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
	return Config{
		Port: 8086,
	}
}

type MonitorInfo struct {
	host string
	port int
	db string
	metricMap map[string]time.Time
	influxClient *client.Client
}

func NewClient(c Config) (*MonitorInfo, error) {
	monClient := MonitorInfo{
		host: c.Host,
		port: c.Port,
		db: c.DB,
		metricMap: make(map[string]time.Time),
	}

	u, err := url.Parse(fmt.Sprintf("http://%s:%d", monClient.host, monClient.port))
	if err != nil {
		log.Fatal(err)
	}

	conf := client.Config{
		URL:      *u,
		Username: os.Getenv("INFLUX_USER"),
		Password: os.Getenv("INFLUX_PWD"),
	}

	con, err := client.NewClient(conf)
	if err != nil {
		log.Fatal(err)
	}

	monClient.influxClient = con

	return &monClient,nil
}

// Start a timer for the function name provided.
func (c *MonitorInfo) StartMonitoring(key string) {
	c.metricMap[key] = time.Now()
}

// End the timer for the provided key and add the time metric to 
// the batch
func (c *MonitorInfo) StopMonitoring(key string) {
	var pts = make([]client.Point, 1)

	pts[0] = client.Point {
		Measurement: key,
		Tags: map[string]string{
			"type": "function",
		},
		Time: time.Now(),
		Fields: map[string]interface{}{
			"value": time.Since(c.metricMap[key]),
		},
	}

	bps := client.BatchPoints{
		Points: pts,
		Database: c.db,
		RetentionPolicy: "default",
	}

	_, err := c.influxClient.Write(bps)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The call took %v to run.\n", time.Now().Sub(c.metricMap[key]))
}

func (c *MonitorInfo) flushMetrics() {

}
