package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Page struct {
	Title string
	Body  string
}

var disableCache = true

// Métodos del struct Page
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile("./data/"+filename, []byte(p.Body), 0600)
}

// Función para convertir markdown básico a HTML
func markdownToHTML(markdown string) template.HTML {
	// Procesar por líneas para manejar headers y listas correctamente
	lines := strings.Split(markdown, "\n")
	var result []string
	inList := false

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Headers
		if strings.HasPrefix(trimmed, "### ") {
			if inList {
				result = append(result, "</ul>")
				inList = false
			}
			title := strings.TrimPrefix(trimmed, "### ")
			result = append(result, "<h3>"+title+"</h3>")
		} else if strings.HasPrefix(trimmed, "## ") {
			if inList {
				result = append(result, "</ul>")
				inList = false
			}
			title := strings.TrimPrefix(trimmed, "## ")
			result = append(result, "<h2>"+title+"</h2>")
		} else if strings.HasPrefix(trimmed, "# ") {
			if inList {
				result = append(result, "</ul>")
				inList = false
			}
			title := strings.TrimPrefix(trimmed, "# ")
			result = append(result, "<h1>"+title+"</h1>")
		} else if strings.HasPrefix(trimmed, "- ") || strings.HasPrefix(trimmed, "* ") {
			// Lista
			if !inList {
				result = append(result, "<ul>")
				inList = true
			}
			item := strings.TrimPrefix(trimmed, "- ")
			item = strings.TrimPrefix(item, "* ")
			result = append(result, "<li>"+item+"</li>")
		} else {
			// Línea normal
			if inList {
				result = append(result, "</ul>")
				inList = false
			}
			if trimmed != "" {
				result = append(result, "<p>"+trimmed+"</p>")
			} else {
				result = append(result, "")
			}
		}
	}

	if inList {
		result = append(result, "</ul>")
	}

	// Unir todas las líneas
	html := strings.Join(result, "\n")

	// Procesar elementos inline
	// Bold **text**
	html = regexp.MustCompile(`\*\*(.*?)\*\*`).ReplaceAllString(html, `<strong>$1</strong>`)

	// Italic *text*
	html = regexp.MustCompile(`\*(.*?)\*`).ReplaceAllString(html, `<em>$1</em>`)

	// Code `code`
	html = regexp.MustCompile("`([^`]+)`").ReplaceAllString(html, `<code>$1</code>`)

	// Links [text](url)
	html = regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`).ReplaceAllString(html, `<a href="$2">$1</a>`)

	return template.HTML(html)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile("./data/" + filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: string(body)}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	var t *template.Template
	var err error

	// Crear template con funciones personalizadas
	funcMap := template.FuncMap{
		"markdown": markdownToHTML,
	}

	if disableCache {
		// Parsear el template cada vez para evitar cache
		t, err = template.New(tmpl + ".html").Funcs(funcMap).ParseFiles("html/" + tmpl + ".html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		// Usar cache (modo producción)
		t = template.Must(template.New(tmpl + ".html").Funcs(funcMap).ParseFiles("html/" + tmpl + ".html"))
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Handlers

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: body}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	files, _ := os.ReadDir("./data")
	var pages []Page

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".txt") {
			title := strings.TrimSuffix(file.Name(), ".txt")
			page, err := loadPage(title)
			if err == nil {
				pages = append(pages, *page)
			}
		}
	}

	renderTemplate(w, "index", pages)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	files, _ := os.ReadDir("./data")
	var pages []Page

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".txt") {
			title := strings.TrimSuffix(file.Name(), ".txt")
			page, err := loadPage(title)
			if err == nil {
				pages = append(pages, *page)
			}
		}
	}

	renderTemplate(w, "list", map[string]interface{}{
		"Pages":      pages,
		"TotalPages": len(pages),
	})
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "create", nil)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/delete/"):]
	if title == "" {
		http.Error(w, "Título vacío", http.StatusBadRequest)
		return
	}

	filename := "./data/" + title + ".txt"
	err := os.Remove(filename)
	if err != nil {
		http.Error(w, "Error al eliminar archivo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/list", http.StatusFound)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "404", nil)
}

func main() {
	os.MkdirAll("./data", 0755)

	// Servir archivos estáticos (CSS, JS, etc.)
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("./styles/"))))

	// Rutas de la aplicación
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/delete/", deleteHandler)
	http.HandleFunc("/404", notFoundHandler)

	log.Println("🚀 Servidor iniciado en http://localhost:8081")
	log.Println("📚 Wiki disponible en:")
	log.Println("   - Página principal: http://localhost:8081/")
	log.Println("   - Lista de páginas: http://localhost:8081/list")
	log.Println("   - Crear página: http://localhost:8081/create")
	log.Println("   - Ver página: http://localhost:8081/view/[titulo]")
	log.Println("   - Editar página: http://localhost:8081/edit/[titulo]")
	log.Println("⚠️  Cache de templates DESHABILITADO - Los cambios en HTML se cargarán automáticamente")

	log.Fatal(http.ListenAndServe(":8081", nil))
}
