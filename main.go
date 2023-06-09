package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/honeycombio/honeycomb-opentelemetry-go"
	"github.com/honeycombio/otel-config-go/otelconfig"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if rand.Intn(100) >= 50 {
		w.WriteHeader(500)
	} else {
		fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	}

}

func main() {
	// enable multi-span attributes
	bsp := honeycomb.NewBaggageSpanProcessor()

	// use honeycomb distro to setup OpenTelemetry SDK
	otelShutdown, err := otelconfig.ConfigureOpenTelemetry(
		otelconfig.WithSpanProcessor(bsp),
	)
	if err != nil {
		log.Fatalf("error setting up OTel SDK - %e", err)
	}
	defer otelShutdown()

	createHandler()

	log.Fatal(http.ListenAndServe(determinePort(), nil))
}

func createHandler() {
	handler := http.HandlerFunc(handler)
	wrappedHandler := otelhttp.NewHandler(handler, "greet")
	http.Handle("/", wrappedHandler)
}

func determinePort() string {
	port := os.Getenv("TIPOLA_PORT")
	if port == "" {
		port = ":8999"
	}
	return port
}
