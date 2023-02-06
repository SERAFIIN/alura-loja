package main

import (
	"net/http"

	"github.com/alura-loja/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
