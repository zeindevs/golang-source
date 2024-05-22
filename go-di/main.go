package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"go.uber.org/fx"
)

// -10
// 0

func main() {
	// this is way to simple! We need to make it more complx. YEAH YIKES.
	router := http.NewServeMux()

	// this is comming from configuration / flag / environment / yadayada
	dbn := "mongo"
	db := getDBByName(dbn)

	echoHandler := NewEchoHandler(db)
	router.Handle("GET /echo", echoHandler)
	http.ListenAndServe(":3000", router)

	// fx.New(
	// 	fx.Provide(
	// 		NewHTTPServer,
	// 		NewEchoHandler,
	// 		NewServeMux,
	// 		NewMongoDB,
	// 	),
	// 	fx.Invoke(func(*http.Server) {}),
	// ).Run()
}

func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux) *http.Server {
	srv := &http.Server{
		Addr:    ":3348",
		Handler: mux,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			fmt.Println("Starting HTTP server at", srv.Addr)
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

// EchoHandler is an http.Handler that copies its request body
// back to the response.
type EchoHandler struct {
	db DBER
}

// NewEchoHandler builds a new EchoHandler.
func NewEchoHandler(db DBER) *EchoHandler {
	return &EchoHandler{
		db: db,
	}
}

// ServeHTTP handles an HTTP request to the /echo endpoint.
func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// I want to get access to my MongoDB??
	w.Write([]byte(h.db.GetUsernameByID(2)))
}

func NewServeMux(echo *EchoHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/echo", echo)
	return mux
}

type DBER interface {
	GetUsernameByID(int) string
}

type MongoDB struct{}

func NewMongoDB() *MongoDB {
	return &MongoDB{}
}

func (db *MongoDB) GetUsernameByID(id int) string {
	return "Anthony from MongoDB"
}

type PostgresDB struct{}

func NewPostgresDB() *PostgresDB {
	return &PostgresDB{}
}

func (db *PostgresDB) GetUsernameByID(id int) string {
	return "Anthony from PostgresDB"
}

func getDBByName(name string) DBER {
	if name == "mongo" {
		return NewMongoDB()
	} else {
		return NewPostgresDB()
	}
}
