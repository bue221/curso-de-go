package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	appNameStyle = lipgloss.NewStyle().Background(lipgloss.Color("99")).Padding(0, 1)

	faint = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Faint(true)

	listEnumeratorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("99")).MarginRight(1)

	// Detail view styles
	detailTitleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("99")).
		Bold(true).
		MarginBottom(1)

	detailFieldStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("255")).
		MarginBottom(1)

	detailValueStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("99")).
		MarginLeft(2)

	// Edit view styles
	editFieldStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("255")).
		MarginBottom(1)

	editLabelStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("99")).
		Bold(true)

	errorStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("196")).
		Bold(true)
)

func (m model) View() string {
	s := appNameStyle.Render("CONTACTS APP") + "\n\n"

	switch m.state {
	case titleView:
		s += m.renderTitleView()
	case bodyView:
		s += m.renderDetailView()
	case editView:
		s += m.renderEditView()
	case listView:
		s += m.renderListView()
	}

	return s
}

func (m model) renderTitleView() string {
	s := "New Contact\n\n"
	s += editLabelStyle.Render("Name:") + "\n"
	s += m.textinput.View() + "\n\n"
	s += faint.Render("enter - continue • esc - cancel")
	return s
}

func (m model) renderDetailView() string {
	if m.currentContact == nil {
		return "No contact selected"
	}

	s := detailTitleStyle.Render("Contact Details") + "\n\n"
	
	s += detailFieldStyle.Render("ID:") + "\n"
	s += detailValueStyle.Render(fmt.Sprintf("%d", m.currentContact.ID)) + "\n\n"
	
	s += detailFieldStyle.Render("Name:") + "\n"
	s += detailValueStyle.Render(m.currentContact.Name) + "\n\n"
	
	s += detailFieldStyle.Render("Email:") + "\n"
	s += detailValueStyle.Render(m.currentContact.Email) + "\n\n"
	
	s += detailFieldStyle.Render("Phone:") + "\n"
	s += detailValueStyle.Render(m.currentContact.Phone) + "\n\n"
	
	s += faint.Render("e - edit • d - delete • esc - back to list")
	return s
}

func (m model) renderEditView() string {
	if m.currentContact == nil {
		return "No contact selected"
	}

	s := detailTitleStyle.Render("Edit Contact") + "\n\n"
	
	// Name field
	s += editLabelStyle.Render("Name:") + "\n"
	if m.editField == "name" {
		s += m.textinput.View() + "\n"
	} else {
		s += editFieldStyle.Render(m.currentContact.Name) + "\n"
	}
	s += "\n"
	
	// Email field
	s += editLabelStyle.Render("Email:") + "\n"
	if m.editField == "email" {
		s += m.textinput.View() + "\n"
	} else {
		s += editFieldStyle.Render(m.currentContact.Email) + "\n"
	}
	s += "\n"
	
	// Phone field
	s += editLabelStyle.Render("Phone:") + "\n"
	if m.editField == "phone" {
		s += m.textinput.View() + "\n"
	} else {
		s += editFieldStyle.Render(m.currentContact.Phone) + "\n"
	}
	s += "\n"
	
	// Show validation errors if any
	if m.validationError != "" {
		s += errorStyle.Render("Error: " + m.validationError) + "\n\n"
	}
	
	s += faint.Render("↑/↓ - navigate fields • enter - edit field • ctrl+s - save • esc - cancel")
	return s
}

func (m model) renderListView() string {
	s := ""
	for i, contact := range m.contacts {
		prefix := " "
		if i == m.listIndex {
			prefix = ">"
		}
		s += listEnumeratorStyle.Render(prefix) + contact.Name + "\n"
	}
	s += "\n" + faint.Render("n - new contact • q - quit • ↑/↓ - navigate • enter - view details")
	return s
}
