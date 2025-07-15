package monitoring

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/atomic"
)

var ready = atomic.NewBool(false)

func RunMonitoringServer(port int) error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	mux.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		if ready.Load() {
			w.WriteHeader(200)
			w.Write([]byte("ready"))
		} else {
			w.WriteHeader(503)
			w.Write([]byte("not ready"))
		}
	})
	mux.HandleFunc("/live", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("live"))
	})

	// pprof endpoints
	mux.HandleFunc("/debug/pprof/", http.DefaultServeMux.ServeHTTP)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}

func SetReady() {
	ready.Store(true)
}
