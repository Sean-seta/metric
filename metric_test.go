package main

import (
	"github.com/go-resty/resty/v2"
	"testing"
	"time"
)

func TestCallAPI(t *testing.T) {
	client := resty.New()
	client.SetBaseURL("http://localhost:8080")

	go func() {
		for {
			_, err := client.R().
				Get("")
			if err != nil {
				break
			}
		}
		time.Sleep(5 * time.Second)
	}()
	go func() {
		for {
			_, err := client.R().
				Get("ping")
			if err != nil {
				break
			}
		}
		time.Sleep(1 * time.Second)
	}()
	go func() {
		for {
			_, err := client.R().
				Get("health")
			if err != nil {
				break
			}
		}
	}()
	for {
		_, err := client.R().
			Get("health")
		if err != nil {
			break
		}
	}

}
