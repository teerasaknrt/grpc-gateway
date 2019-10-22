package main

import (
	"log"
	"net/http"

	"context"

	"github.com/golang/glog"
	runtime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	gw "github.com/teerasaknrt/grpc-gateway/api/v1"
	"google.golang.org/grpc"
)

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware", r.URL)
		h.ServeHTTP(w, r)
	})
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterProfileServiceHandlerFromEndpoint(ctx, mux, "localhost:9090", opts)
	if err != nil {
		return err
	}
	r := http.NewServeMux()
	r.Handle("/", mux)
	log.Println("listening to port *:8080")

	return http.ListenAndServe(":8080", r)
}

func main() {
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
