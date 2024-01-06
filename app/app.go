package app

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func Start() {
	r := chi.NewRouter()

	r.Get("/", greeting)
	r.Get("/customers", getAllcustomers)

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", r)
}
