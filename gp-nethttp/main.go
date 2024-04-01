package main

import (
	"log"
	"net/http"

	"github.com/zeindevs/nethttp/handler"
	"github.com/zeindevs/nethttp/middleware"
	"github.com/zeindevs/nethttp/monster"
)

func handleOther(w http.ResponseWriter, _ *http.Request) {
	log.Println("Received a non domain request")
	w.Write([]byte("Hello, stranger..."))
}

func handle(w http.ResponseWriter, _ *http.Request) {
	log.Println("Received a request at my domain")
	w.Write([]byte("Hello, Domain name!"))
}

func loadRoutes(router *http.ServeMux) {
	router.HandleFunc("/", handleOther)
	router.HandleFunc("zeindevs.foo/", handle)

	router.HandleFunc("POST /monster", monster.Create)
	router.HandleFunc("PUT /monster/{id}", monster.UpdateByID)
	router.HandleFunc("GET /monster/{id}", monster.FindByID)
	router.HandleFunc("DELETE /monster/{id}", monster.DeleteByID)

}

func main() {
	router := http.NewServeMux()
	// loadRoutes(router)

	adminRouter := http.NewServeMux()
	adminRouter.HandleFunc("POST /invoice", handler.Create)
	adminRouter.HandleFunc("PUT /invoice/{id}", handler.UpdateByID)
	adminRouter.HandleFunc("DELETE /invoice/{id}", handler.DeleteByID)

	router.Handle("/", middleware.EnsureAdmin(adminRouter))

	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.AllowCors,
		middleware.IsAuthed,
		middleware.CheckPermissions,
	)

	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", router))

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}

	log.Println("Starting server on port :8080")
	server.ListenAndServe()
}
