package repository

import (
	"database/sql"

	"sql-in-go/models"
)

type ContactRepository struct {
	db *sql.DB
}

func NewContactRepository(db *sql.DB) *ContactRepository {
	return &ContactRepository{db: db}
}

func (r *ContactRepository) GetAllContacts() ([]models.Contact, error) {
	query := "SELECT id, name, email, phone FROM contacts"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []models.Contact
	for rows.Next() {
		var contact models.Contact
		err := rows.Scan(&contact.ID, &contact.Name, &contact.Email, &contact.Phone)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func (r *ContactRepository) GetContactByID(id int) (*models.Contact, error) {
	query := "SELECT id, name, email, phone FROM contacts WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var contact models.Contact
	err := row.Scan(&contact.ID, &contact.Name, &contact.Email, &contact.Phone)
	if err != nil {
		return nil, err
	}

	return &contact, nil
}

func (r *ContactRepository) CreateContact(contact models.Contact) error {
	query := "INSERT INTO contacts (name, email, phone) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, contact.Name, contact.Email, contact.Phone)
	return err
}

func (r *ContactRepository) UpdateContact(contact models.Contact) error {
	query := "UPDATE contacts SET name = ?, email = ?, phone = ? WHERE id = ?"
	_, err := r.db.Exec(query, contact.Name, contact.Email, contact.Phone, contact.ID)
	return err
}

func (r *ContactRepository) DeleteContact(id int) error {
	query := "DELETE FROM contacts WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}
