package main

import(
	"log"
	"GO_LANG/handlers/monitoring"
)


func main(){

	metrics, err := monitoring.GetSysMetrics()
	if err != nil {
		log.Fatal(err)
	}

	monitoring.PrintSystemMetrics(metrics)
}

