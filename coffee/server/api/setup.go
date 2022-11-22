package api

import (
	"fmt"
	"net/http"

	"github.com/4nte/meetup/coffee/proto/protoconnect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func Setup() {
	server := &CoffeServer{
		sugar: 10,
		milk:  250,
	}
	mux := http.NewServeMux()
	path, handler := protoconnect.NewCoffeServiceHandler(server)
	mux.Handle(path, handler)

	fmt.Println("gRPC server is ready!")

	http.ListenAndServe(
		"localhost:9090",
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
