package render

import (
	"bytes"
	"github.com/justinas/nosurf"

	"github.com/aville22/greeneats/pkg/config"
	"github.com/aville22/greeneats/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}
var app *config.AppConfig

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}
func NewTemplate(a *config.AppConfig) {
	app = a
}
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCash {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCash()
	}
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("error")
	}
	buf := new(bytes.Buffer)
	td = AddDefaultData(td, r)
	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
	/*parsedTemplate, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.html")
	if err != nil {
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		fmt.Println("Ошибка при парсинге шаблона:", err)
		return
	}

	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		fmt.Println("Ошибка при выполнении шаблона:", err)
		return
	}*/
}
func CreateTemplateCash() (map[string]*template.Template, error) {
	myCashe := make(map[string]*template.Template)
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCashe, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCashe, err
		}
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCashe, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCashe, err
			}
		}
		myCashe[name] = ts
	}
	return myCashe, nil
}

/*var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error
	_, inMap := tc[t]
	if !inMap {
		fmt.Println("creating template")
		err = createTemplateCash(t)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("using cashed template")
	}
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}
func createTemplateCash(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.html",
	}
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	tc[t] = tmpl
	return nil
}
*/
