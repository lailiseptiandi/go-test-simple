package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lailiseptiandi/go-test-simple/internal/routers"
	"github.com/lailiseptiandi/go-test-simple/pkg/config"
	"github.com/lailiseptiandi/go-test-simple/pkg/database"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	dbs, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	defer database.DisconnectDB(dbs)

	// set routes
	r := gin.Default()
	routers.InitRoutes(dbs, r)

	serverAddr := fmt.Sprintf("%s:%d", cfg.HOST, cfg.PORT)
	r.Run(serverAddr)

	// srv := &http.Server{
	// 	Addr:    serverAddr,
	// 	Handler: r,
	// }

	// go func() {
	// 	err := srv.ListenAndServe()
	// 	if err != nil && err != http.ErrServerClosed {
	// 		log.Fatalf("server failed to run :%v", err)
	// 	}
	// }()

	// log.Println("Server is running", "address", serverAddr)

	// // handle shutdown
	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// <-quit
	// log.Println("Received shutdown signal, initiating server shutdown...")

	// // Graceful shutdown
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// if err := srv.Shutdown(ctx); err != nil {
	// 	log.Fatalf("Server shut down failed", err)
	// }

	// log.Println("Server successfully shut down")
}
