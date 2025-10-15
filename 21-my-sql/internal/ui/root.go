package ui

import (
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"sql-in-go/internal/models"
)

const (
	listView uint = iota
	titleView
	bodyView
	editView
)

type model struct {
	state           uint
	store           *Store
	contacts        []models.Contact
	currentContact  *models.Contact
	listIndex       int
	textarea        textarea.Model
	textinput       textinput.Model
	editField       string
	editFieldIndex  int
	validationError string
}

func NewModel(store *Store) model {
	contacts, err := store.GetContacts()
	if err != nil {
		log.Fatal("Error getting contacts:", err)
	}
	return model{
		store:     store,
		state:     listView,
		contacts:  contacts,
		textarea:  textarea.New(),
		textinput: textinput.New(),
	}
}

// MAIN METHODS
func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)

	m.textinput, cmd = m.textinput.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()

		if key == "ctrl+c" {
			return m, tea.Quit
		}

		switch m.state {
		case listView:
			switch key {
			case "q":
				return m, tea.Quit
			case "n":
				m.textinput.SetValue("")
				m.textinput.Focus()
				m.currentContact = &models.Contact{}
				m.state = titleView
			case "up", "k":
				if m.listIndex > 0 {
					m.listIndex--
				}
			case "down", "j":
				if m.listIndex < len(m.contacts)-1 {
					m.listIndex++
				}
			case "enter":
				m.currentContact = &m.contacts[m.listIndex]
				m.state = bodyView
			}

		case titleView:
			switch key {
			case "enter":
				name := m.textinput.Value()
				if name != "" {
					m.currentContact.Name = name
					m.currentContact.Email = ""
					m.currentContact.Phone = ""
					m.state = editView
					m.editField = "email"
					m.editFieldIndex = 0
					m.textinput.SetValue("")
					m.textinput.Focus()
				}
			case "esc":
				m.state = listView
			}

		case bodyView:
			switch key {
			case "e":
				m.state = editView
				m.editField = "name"
				m.editFieldIndex = 0
				m.textinput.SetValue(m.currentContact.Name)
				m.textinput.Focus()
			case "d":
				// Delete contact
				if m.currentContact.ID != 0 {
					err := m.store.DeleteContact(m.currentContact.ID)
					if err != nil {
						log.Printf("Error deleting contact: %v", err)
					} else {
						// Refresh contacts list
						contacts, err := m.store.GetContacts()
						if err != nil {
							log.Printf("Error refreshing contacts: %v", err)
						} else {
							m.contacts = contacts
							if m.listIndex >= len(m.contacts) {
								m.listIndex = len(m.contacts) - 1
							}
						}
					}
				}
				m.state = listView
			case "esc":
				m.state = listView
			}

		case editView:
			switch key {
			case "up", "k":
				if m.editFieldIndex > 0 {
					m.editFieldIndex--
					m.updateEditField()
				}
			case "down", "j":
				if m.editFieldIndex < 2 {
					m.editFieldIndex++
					m.updateEditField()
				}
			case "enter":
				m.updateCurrentField()
			case "ctrl+s":
				if m.validateContact() {
					if m.currentContact.ID == 0 {
						// Create new contact
						err := m.store.CreateContact(*m.currentContact)
						if err != nil {
							log.Printf("Error creating contact: %v", err)
						} else {
							// Refresh contacts list
							contacts, err := m.store.GetContacts()
							if err != nil {
								log.Printf("Error refreshing contacts: %v", err)
							} else {
								m.contacts = contacts
							}
						}
					} else {
						// Update existing contact
						err := m.store.UpdateContact(*m.currentContact)
						if err != nil {
							log.Printf("Error updating contact: %v", err)
						}
					}
					m.state = listView
				}
			case "esc":
				m.state = listView
			}
		}
	}

	return m, tea.Batch(cmds...)
}

func (m *model) updateEditField() {
	switch m.editFieldIndex {
	case 0:
		m.editField = "name"
		m.textinput.SetValue(m.currentContact.Name)
	case 1:
		m.editField = "email"
		m.textinput.SetValue(m.currentContact.Email)
	case 2:
		m.editField = "phone"
		m.textinput.SetValue(m.currentContact.Phone)
	}
	m.textinput.Focus()
}

func (m *model) updateCurrentField() {
	value := m.textinput.Value()
	switch m.editField {
	case "name":
		m.currentContact.Name = value
	case "email":
		m.currentContact.Email = value
	case "phone":
		m.currentContact.Phone = value
	}
}

func (m *model) validateContact() bool {
	m.validationError = ""
	
	if strings.TrimSpace(m.currentContact.Name) == "" {
		m.validationError = "Name is required"
		return false
	}
	
	if strings.TrimSpace(m.currentContact.Email) == "" {
		m.validationError = "Email is required"
		return false
	}
	
	// Basic email validation
	if !strings.Contains(m.currentContact.Email, "@") {
		m.validationError = "Email must contain @"
		return false
	}
	
	return true
}
