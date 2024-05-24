package main

import (
	"fmt"
	"log"

	"Gin-Api/config"
	"Gin-Api/internal/database"
	"Gin-Api/internal/database/postgres"
	"Gin-Api/internal/repo"
	docRepo "Gin-Api/internal/repo/document"
	userRepo "Gin-Api/internal/repo/user"
	docServ "Gin-Api/internal/service/document"
	userServ "Gin-Api/internal/service/user"

	docHand "Gin-Api/internal/handlers/document"
	userHand "Gin-Api/internal/handlers/user"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Println(err, nil)
	}
	db, err := postgres.OpenDB(*cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = database.CheckAndCreateTable(db)
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	err = repo.Init(db)
	if err != nil {
		log.Println(err)
		return
	}
	repo := userRepo.NewUserRepository(db)     // возможность создавать user в бд
	service := userServ.NewUserService(repo)   // логика
	handler := userHand.NewUserClient(service) // пользователь
	documentrepo := docRepo.NewDocumentRepository(db)
	documentservice := docServ.NewDocumentService(documentrepo)
	document := docHand.NewDocumentClient(documentservice)
	handler.Routes(router, cfg)
	document.Routes(router, cfg)
	err = router.Run(fmt.Sprintf(":%v", cfg.Port))
	if err != nil {
		log.Fatal(err)
	}
}
