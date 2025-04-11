package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

var (
	msgCh     = make(chan Payload, 1)
	brokerURL = os.Getenv("BROKER_URL")  //"tcp://127.0.0.1:1883"
	mqttUser  = os.Getenv("BROKER_USER") //"demo"
	mqttPass  = os.Getenv("BROKER_PASS") //"demo"
)

type Payload struct {
	CPU    []float64 `json:"cpu"`
	Memory float64   `json:"memory"`
}

func getResourceUsage() ([]float64, float64, error) {
	cpuPercent, err := cpu.Percent(time.Second, true)
	if err != nil {
		return []float64{}, 0.0, err
	}
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return []float64{}, 0.0, err
	}
	return cpuPercent, memInfo.UsedPercent, nil
}

func runResourceMonitor() {
	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-ticker.C:
			c, m, err := getResourceUsage()
			if err != nil {
				log.Fatal(err)
			}
			msgCh <- Payload{
				CPU:    c,
				Memory: m,
			}
		}
	}
}

func runMQTT() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerURL)
	opts.SetClientID("go-mqtt-resource-monitor")
	opts.SetUsername(mqttUser)
	opts.SetPassword(mqttPass)
	opts.SetConnectRetry(true)
	opts.SetDefaultPublishHandler(messageHandler)

	opts.OnConnectionLost = func(cl mqtt.Client, err error) {
		log.Println("[MQTT] connection lost")
	}
	opts.OnConnect = func(mqtt.Client) {
		log.Println("[MQTT] connection established")
	}
	opts.OnReconnecting = func(mqtt.Client, *mqtt.ClientOptions) {
		log.Println("[MQTT] attempting to reconnect")
	}

	c := mqtt.NewClient(opts)
	if t := c.Connect(); t.Wait() && t.Error() != nil {
		log.Fatal(t.Error())
	}

	for {
		select {
		case msg := <-msgCh:
			payload, _ := json.MarshalIndent(msg, "", "")
			if t := c.Publish("monitor", 1, false, string(payload)); t.Wait() && t.Error() != nil {
				log.Printf("[MQTT] ERR: %s", t.Error())
				return
			}
			log.Printf("[MQTT] Send message: %s, to topic: monitor", string(payload))
		}
	}
}

func messageHandler(c mqtt.Client, msg mqtt.Message) {
	log.Printf("[MQTT] Received message: %s from topic: %s", msg.Payload(), msg.Topic())
}

func main() {
	go runResourceMonitor()
	go runMQTT()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Start(os.Getenv("APP_PORT"))
}
