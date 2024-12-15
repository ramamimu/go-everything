package main

import (
	"log"

	elastic "github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"
)

func main() {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		log.Fatalf("Failed to create Elasticsearch client: %v", err)
	}

	logger := logrus.New()
	hook, err := elogrus.NewElasticHook(client, "localhost", logrus.DebugLevel, "go-logs")
	if err != nil {
		log.Fatalf("Failed to create Elasticsearch hook: %v", err)
	}

	logger.Hooks.Add(hook)

	logger.WithFields(logrus.Fields{
		"app":     "example-app",
		"version": "1.0",
	}).Info("Application started")

	logger.Error("This is an error log")
}
