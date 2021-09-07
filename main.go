package main

import (
	"net/http"

	"github.com/cmllmd/bookreview/routes"
)

func main() {

	routes.CarregaRotas()
	http.ListenAndServe(":8080", nil)

}
