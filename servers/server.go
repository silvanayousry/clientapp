package servers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/clientapp/configs"
	"github.com/clientapp/routes"
)

// StartServer starts the http server
func StartServer() {
	e := routes.Router()

	appconfigs := configs.App()

	portStr := fmt.Sprintf(":%d", appconfigs.Port)

	server := &http.Server{
		ReadTimeout: appconfigs.ReadTimeout,
		Addr:        portStr,
		Handler:     e,
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGKILL, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	log.Printf("Listening on port::%d...\n", appconfigs.Port)
	<-sig

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_ = server.Shutdown(ctx)

	log.Println("Server shutdown ")

}
