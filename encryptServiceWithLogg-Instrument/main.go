package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Neil-uli/Restful-go/encryptServiceWithLogg-Instrument/helpers"
	kitlog "github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	logger := kitlog.NewLogfmtLogger(os.Stderr)
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "encryption",
		Subsystem: "my_service",
		Name: "request_count",
		Help: "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "encryption",
		Subsystem: "my_service",
		Name: "request_latency_microseconds",
		Help: "Total duration of requests in microseconds.",
	}, fieldKeys)
	var svc helpers.EncryptService
	svc = helpers.EncryptServiceInstance{}
	svc = helpers.LoggingMiddleware{Logger: logger, Next: svc}
	svc = helpers.InstrumentingMiddleware{RequestCount: requestCount, RequestLatency: requestLatency, Next: svc}

	encryptHandler := httptransport.NewServer(helpers.MakeEncryptEndpoint(svc),
		helpers.DecodeEncryptRequest,
		helpers.EncodeResponse)

	decryptHandler := httptransport.NewServer(helpers.MakeDecryptEndpoint(svc),
		helpers.DecodeDecryptRequest,
		helpers.EncodeResponse)

	http.Handle("/encrypt", encryptHandler)
	http.Handle("/decrypt", decryptHandler)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// curl -X POST -d '{"key":"111023043350789514532147", "text": "I am A Message"}' localhost:8080/encrypt
// http://localhost:8080/metrics