package main

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type event struct {
	Source      string `env:"EVENT_SOURCE"`
	Detail      string `env:"EVENT_DETAIL"`
	DetailType  string `env:"EVENT_DETAIL_TYPE"`
	BucketName  string `env:"BUCKET_NAME"`
	S3ObjectKey string `env:"S3_OBJECT_KEY"`
}

func main() {
	// parse env vars
	ev := &event{}
	if err := env.Parse(ev); err != nil {
		panic(err)
	}

	log.Printf("env vars: %+v", ev)
}
