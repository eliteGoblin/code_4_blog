package main

import (
	"context"
	"log_demo/log"
)

func main() {
	lg := log.NewFluentd(false)
	lg.Info("hello log")

	// add fields to logger
	flog := lg.WithField("request_id", "fake-request-id")
	// log carried by ctx
	ctx := context.Background()
	logCtx := log.NewContext(ctx, flog)
	logFromCtx, _ := log.FromContext(logCtx)
	logFromCtx.Info("hello")
}
