package ui

import (
	"log"

	"sql-in-go/config"
	"sql-in-go/database"
	"sql-in-go/internal/models"
	"sql-in-go/internal/repository"
	"sql-in-go/internal/services"
)

type Store struct {
	contactService *services.ContactService
	contactRepo    *repository.ContactRepository
}

func (s *Store) Init() error {
	cfg := config.LoadConfig()
	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatal("Error conectando a la base de datos:", err)
	}

	s.contactRepo = repository.NewContactRepository(db)
	s.contactService = services.NewContactService(s.contactRepo)

	return nil
}

func (s *Store) GetContacts() ([]models.Contact, error) {
	contacts, err := s.contactRepo.GetAllContacts()
	if err != nil {
		return nil, err
	}

	return contacts, nil
}

func (s *Store) GetContactByID(id int) (*models.Contact, error) {
	contact, err := s.contactRepo.GetContactByID(id)
	if err != nil {
		return nil, err
	}

	return contact, nil
}

func (s *Store) CreateContact(contact models.Contact) error {
	return s.contactRepo.CreateContact(contact)
}

func (s *Store) UpdateContact(contact models.Contact) error {
	return s.contactRepo.UpdateContact(contact)
}

func (s *Store) DeleteContact(id int) error {
	return s.contactRepo.DeleteContact(id)
}
