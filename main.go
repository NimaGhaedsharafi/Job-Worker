package main

import (
	"coroner/config"
	"coroner/database"
	"coroner/worker"
)

func main() {
	config.Init()

	dispatcher := worker.NewDispatcher(config.Cfg.Dispatcher.MaxWorker)
	dispatcher.Run()
	// TODO: add close functionality to dispatcher

	// fetch drivers
	database.Init()
	defer database.Close()

	// loop through all drivers and dispatch a job

	// once everything is done, create a report in csv

	// send an email
}
