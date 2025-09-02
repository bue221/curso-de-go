package book

import "fmt"

type Printable interface {
	PrintBook()
}

func Print(p Printable) {
	p.PrintBook()
}

type Book struct {
	ID     int
	Title  string
	Author string
	Pages  int
}

func (b *Book) PrintBook() {
	fmt.Printf("ID: %d, Title: %s, Author: %s, Pages: %d\n", b.ID, b.Title, b.Author, b.Pages)
}

// Metodos setter y getter en la POO

func (b *Book) GetBook() Book {
	return *b
}

func (b *Book) SetBook(id int, title string, author string, pages int) {
	b.ID = id
	b.Title = title
	b.Author = author
	b.Pages = pages
}

func (b *Book) GetTitle() string {
	return b.Title
}

func (b *Book) SetTitle(title string) {
	b.Title = title
}

// Herencia de structs
type TextBook struct {
	Book
	editorial string
	level     string
}

// Composición de structs
func NewTextBook(title, author, subject, editorial, level string, id, pages int) *TextBook {
	return &TextBook{
		Book: Book{
			ID:     id,
			Title:  title,
			Author: author,
			Pages:  pages,
		},
		editorial: editorial,
		level:     level,
	}
}

func (t *TextBook) PrintTextBook() {
	fmt.Printf("ID: %d, Title: %s, Author: %s, Pages: %d, Editorial: %s, Level: %s\n", t.ID, t.Title, t.Author, t.Pages, t.editorial, t.level)
}
