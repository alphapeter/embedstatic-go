package main

import(
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	embeddedHandler := CreateHandler(Data, "index.html")
	router.Get("/*}", embeddedHandler.ServeHTTP)

	http.ListenAndServe(":8000", router)
}
