package main

import (
	"os"
	"log"
	"net/http"

	"github.com/Neil-uli/Restful-go/encryptServiceWithLogging/helpers"
	kitlog "github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := kitlog.NewLogfmtLogger(os.Stderr)
	var svc helpers.EncryptService
	svc = helpers.EncryptServiceInstance{}
	svc = helpers.LoggingMiddleware{Logger: logger, Next: svc}
	encryptHandler := httptransport.NewServer(helpers.MakeEncryptEndpoint(svc),
		helpers.DecodeEncryptRequest,
		helpers.EncodeResponse)

	decryptHandler := httptransport.NewServer(helpers.MakeDecryptEndpoint(svc),
		helpers.DecodeDecryptRequest,
		helpers.EncodeResponse)

	http.Handle("/encrypt", encryptHandler)
	http.Handle("/decrypt", decryptHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// curl -X POST -d '{"key":"111023043350789514532147", "text": "I am A Message"}' localhost:8080/encrypt
