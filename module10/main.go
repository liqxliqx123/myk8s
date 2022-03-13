package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"k8s-homework/module10/metrics"
)

func delay(w http.ResponseWriter, r *http.Request) {
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	randInt := rand.Intn(2000)
	time.Sleep(time.Millisecond * time.Duration(randInt))
	w.Write([]byte(fmt.Sprintf("<h1>%d<h1>", randInt)))
}

func main() {

	metrics.Register()
	mux := http.NewServeMux()

	mux.HandleFunc("/delay", delay)
	mux.Handle("/metrics", promhttp.Handler())

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
