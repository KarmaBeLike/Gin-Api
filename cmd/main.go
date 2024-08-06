package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"Gin-Api/config"
	"Gin-Api/internal/database/postgres"
	userHand "Gin-Api/internal/handlers/user"
	"Gin-Api/internal/repo"
	docRepo "Gin-Api/internal/repo/document"
	userRepo "Gin-Api/internal/repo/user"
	docServ "Gin-Api/internal/service/document"
	userServ "Gin-Api/internal/service/user"

	docHand "Gin-Api/internal/handlers/document"

	"Gin-Api/internal/database/redis"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.OpenDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	redisClient, err := redis.NewRedisClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	err = repo.Init(db)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := userRepo.NewUserRepository(db)              // возможность создавать user в бд
	service := userServ.NewUserService(userRepo)            // логика
	handler := userHand.NewUserClient(service, redisClient) // пользователь

	documentRepo := docRepo.NewDocumentRepository(db)
	documentService := docServ.NewDocumentService(documentRepo)
	documentHandler := docHand.NewDocumentClient(documentService)

	handler.Routes(router, cfg)
	documentHandler.Routes(router, cfg)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Channel to signal the server has gracefully shut down
	done := make(chan bool, 1)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router.Handler(),
	}
	go func() {
		<-stop
		log.Println("Shutting down the server...")

		// Context with timeout to ensure server shuts down within 5 seconds
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("Server shutdown failed:%+v", err)
		}
		close(done)
	}()
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
	// err = router.Run(fmt.Sprintf(":%v", cfg.Port))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	<-done
	log.Println("Server stopped")
}

// type closure struct {
// 	toClose []func
// }

// func (c *closure) add(f func) {

// }

// func (c *closure) CloseAll() {}
