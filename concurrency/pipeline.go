package main

import (
	"log"
	"time"
)

func main() {
	testComplete := make(chan string, 2)
	deployComplete := make(chan string, 2)
	emailComplete := make(chan string, 2)

	apps := []string{
		"application-service",
		"accounts-service",
		"persons-service",
		"transactions-service",
		"deposits-service",
		"accounts-ledger-service",
	}

	go test(apps, testComplete)
	go deploy(testComplete, deployComplete)
	go sendCompleteEmail(deployComplete, emailComplete)
	terminatePipeline(emailComplete)
}

func test(apps []string, out chan<- string) {
	for _, app := range apps {
		time.Sleep(1000 * time.Millisecond)
		log.Printf("Running tests for %s\n", app)
		out <- app
	}
	close(out)
}

func deploy(in <-chan string, out chan<- string) {
	for app := range in {
		time.Sleep(1000 * time.Millisecond)
		log.Printf("Deploying %s\n", app)
		out <- app
	}
	close(out)
}

func sendCompleteEmail(in <-chan string, out chan<- string) {
	for app := range in {
		time.Sleep(1000 * time.Millisecond)
		log.Printf("Sending complete email notification for %s\n", app)
		out <- app
	}
	close(out)
}

func terminatePipeline(in <-chan string) {
	for app := range in {
		log.Printf("Pipeline complete for %s\n", app)
	}
}
