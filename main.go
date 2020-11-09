package main

import (
	"coroner/config"
	"coroner/database"
	"coroner/worker"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	config.Init()

	dispatcher := worker.NewDispatcher(config.Cfg.Dispatcher.MaxWorker)
	dispatcher.Run()
	// TODO: add close functionality to dispatcher

	// fetch drivers
	database.Init()
	defer database.Close()
	// set the correct date
	day, month, year := time.Now().AddDate(0, 0, -1).Date()
	query := fmt.Sprintf("select id, id_number from drivers where updated_at > \"%d-%d-%d\" and id_number != \"\"", day, month, year)
	rows, err := database.Db.Query(query)
	if err != nil {
		logrus.Error("failed to fetch drivers")
		return
	}
	defer rows.Close()

	// loop through all drivers and dispatch a job
	for rows.Next() {
		var id, ssn string
		if err := rows.Scan(&id, &ssn); err != nil {
			logrus.Error("Failed to scan a row")
		}

		work := worker.Job{SSN: ssn}
		worker.JobQueue <- work
	}
	// once everything is done, create a report in csv

	// send an email
}
