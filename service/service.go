package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/sample-microservice/spmonitorlibrary"
)

var numDDL *monitoring.CounterMetric

func main() {
	http.HandleFunc("/", getRoot)

	var handler *monitoring.MetricHandler = monitoring.GetMetricHandler()
	http.HandleFunc("/metrics", handler.Metrics)
	numDDL = monitoring.NewCounterMetric(monitoring.MetricDescriptor{
		Name: "numDDL",
		Help: "Number of issued digital driving licenses",
	})

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET / -> Incrementing counter")
	numDDL.Increment()
	io.WriteString(w, "Simple GO Microservice!\n")
}
