package main

import (
	"github.com/go-resty/resty/v2"
	"testing"
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
	}()
	go func() {
		for {
			_, err := client.R().
				Get("ping")
			if err != nil {
				break
			}
		}
	}()
	//go func() {
	//	for {
	//		_, err := client.R().
	//			Get("health")
	//		if err != nil {
	//			break
	//		}
	//	}
	//}()
	for {
		_, err := client.R().
			Get("health")
		if err != nil {
			break
		}
	}

}
