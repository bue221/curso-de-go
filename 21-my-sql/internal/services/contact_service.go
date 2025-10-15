package services

import (
	"fmt"

	"sql-in-go/internal/models"
	"sql-in-go/internal/repository"
)

type ContactService struct {
	repo *repository.ContactRepository
}

func NewContactService(repo *repository.ContactRepository) *ContactService {
	return &ContactService{repo: repo}
}

func (s *ContactService) DisplayAllContacts() error {
	contacts, err := s.repo.GetAllContacts()
	if err != nil {
		return err
	}

	fmt.Println("=== Contactos encontrados ===")
	if len(contacts) == 0 {
		fmt.Println("No se encontraron contactos.")
		return nil
	}

	for _, contact := range contacts {
		s.displayContact(contact)
	}

	return nil
}

func (s *ContactService) DisplayContactByID(id int) error {
	contact, err := s.repo.GetContactByID(id)
	if err != nil {
		return err
	}

	s.displayContact(*contact)

	return nil
}

func (s *ContactService) displayContact(contact models.Contact) {
	fmt.Printf("ID: %d | Nombre: %s | Email: %s | Teléfono: %s\n",
		contact.ID, contact.Name, contact.Email, contact.Phone)
}
