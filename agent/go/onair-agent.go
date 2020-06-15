package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/mitchellh/go-ps"
	"gopkg.in/yaml.v2"
)

const configFile = "config.yaml"

func main() {
	config := loadConfigFromYaml(configFile)

	for {
		running := false

		ps, _ := ps.Processes()

		for i := range ps {
			name := ps[i].Executable()

			for j := range config.Processes {
				if name == config.Processes[j] {
					running = true
				}
			}
		}

		if running {
			fmt.Println("At least one monitored process is running.")
			invokeWebhook(config.Webhook, "on")
		} else {
			fmt.Println("No monitored processes are running.")
			invokeWebhook(config.Webhook, "off")
		}

		time.Sleep(time.Millisecond * time.Duration(config.Interval))
	}
}

func invokeWebhook(webhook string, state string) {
	client := http.Client{}

	url := fmt.Sprintf("%s/%s", webhook, state)
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response)
	response.Body.Close()
}

func loadConfigFromYaml(file string) (config Config) {
	file = filepath.FromSlash(file)

	log.Printf("Loading config from %s...", file)

	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed to read file %s: %s", file, err)
	}

	config = Config{}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal yaml contents: %s", err)
	}

	return config
}

type Config struct {
	Processes []string
	Interval  int
	Webhook   string
}
