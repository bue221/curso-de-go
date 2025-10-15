package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"

	"sql-in-go/internal/ui"

	_ "github.com/go-sql-driver/mysql"
)

// func main() {
// 	cfg := config.LoadConfig()

// 	db, err := database.ConnectDB(cfg)
// 	if err != nil {
// 		log.Fatal("Error conectando a la base de datos:", err)
// 	}
// 	defer db.Close()

// 	contactRepo := repository.NewContactRepository(db)
// 	contactService := services.NewContactService(contactRepo)

// 	p := tea.NewProgram(ui.NewModel())
// 	if _, err := p.Run(); err != nil {
// 		fmt.Printf("Alas, there's been an error: %v", err)
// 		os.Exit(1)
// 	}

// 	if _, err := p.Run(); err != nil {
// 		fmt.Printf("Alas, there's been an error: %v", err)
// 		os.Exit(1)
// 	}

// 	// Mostrar todos los contactos
// 	if err := contactService.DisplayAllContacts(); err != nil {
// 		log.Fatal("Error mostrando contactos:", err)
// 	}

// 	// Mostrar un contacto por ID
// 	if err := contactService.DisplayContactByID(1); err != nil {
// 		log.Fatal("Error mostrando contacto:", err)
// 	}
// }

func main() {

	store := new(ui.Store)
	if err := store.Init(); err != nil {
		log.Fatalf("unable to init store: %v", err)
	}

	m := ui.NewModel(store)

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		log.Fatalf("unable to run tui: %v", err)
	}
}
