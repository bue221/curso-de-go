package main

import (
	"log"

	"sql-in-go/config"
	"sql-in-go/database"
	"sql-in-go/repository"
	"sql-in-go/services"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg := config.LoadConfig()
	
	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatal("Error conectando a la base de datos:", err)
	}
	defer db.Close()

	contactRepo := repository.NewContactRepository(db)
	contactService := services.NewContactService(contactRepo)

	if err := contactService.DisplayAllContacts(); err != nil {
		log.Fatal("Error mostrando contactos:", err)
	}
}