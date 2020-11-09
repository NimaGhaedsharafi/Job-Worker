package main

import (
	"coroner/config"
	"coroner/worker"
)

func main() {
	config.Init()

	dispatcher := worker.NewDispatcher(config.Cfg.Dispatcher.MaxWorker)
	dispatcher.Run()

	// TODO: add close functionality to dispatcher

	// fetch drivers

	// loop through all drivers and dispatch a job

	// once everything is done, create a report in csv

	// send an email
}
