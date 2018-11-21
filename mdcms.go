package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
        "log"
	"github.com/russross/blackfriday"
	"strings"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type site struct {
	Title	string
}

func (s site) Config() string {
	configfile , err := ioutil.ReadFile("config.txt")
	handle_err(err)
	return string(configfile)
}
func (s site) SiteName() string {
	config := s.Config()
	sitename := strings.Split((strings.Split(config, "sitename=")[1]), "\n")[0]
	return sitename
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Error handler
func handle_err(err error) {
        if err != nil {
                panic(err.Error())
        }
}

///Read all files from the relative path "markdown"
func mdparse() ([]string) {

  files, err := ioutil.ReadDir("markdown")
  handle_err(err)
  var mdpaths []string
  for _, file := range files {
    mdfile := file.Name()
    mdpath := "markdown/"+ mdfile
    mdpaths = append(mdpaths, mdpath)
  }
  return mdpaths
}


// Parsing markdown
func mdtohtml(markdownfile string) string {
	file, err := ioutil.ReadFile(markdownfile)
	handle_err(err)
        output := string(blackfriday.MarkdownCommon(file))
	md := "<div class=\"markdown\">" + output + "</div>"
        return  md
}
//General URL handler function
func handler(w http.ResponseWriter, r *http.Request, html_input string, data site) {
	content := template.New("content")
	content.Parse(html_input)
	content.ParseFiles("templates/base.gohtml")
	content.ExecuteTemplate(w, "base", data)
}

//Index Handler. Parse the markdown files read y mdparse() with the base layout template(templates/base.gohtml) and display at /
func index(w http.ResponseWriter, r *http.Request) {
	mysite := site{"Home"}
	var mdhtml  string
        for _, mdpath := range mdparse() {
	  contenttitle := strings.Split((strings.Split(mdpath, "/")[1]), ".md")[0]
	  contenttitleuri := "/content/" + contenttitle
	  mdhtml = mdhtml + "<div class=\"small\">" + mdtohtml(mdpath) + "</div> <p> <a class=\"home\" href=\"" + contenttitleuri +"\">Read More</a> </p>"
        }
	handler(w, r, mdhtml, mysite)
}

//About handler. Parse about.html(template/about.gohtml) with base.gohtml and display at /about.
func about(w http.ResponseWriter, r *http.Request) {
	mysite := site{"About"}
        abouthtml := mdtohtml("templates/about.md")
	handler(w, r, abouthtml, mysite)
}

//Content Handler
func contenthandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/content/"):]
	mysite := site{title}
	contenthtml := mdtohtml("markdown/"+ title + ".md")
	handler(w, r, contenthtml, mysite)
}

//Downloads page handler.
func downloads(w http.ResponseWriter, r *http.Request) {
	mysite := site{"Downloads"}
	downloadhtml := mdtohtml("templates/downloads.md")
	handler(w, r, downloadhtml, mysite)
}


///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func main() {

	http.HandleFunc("/", index)
        http.HandleFunc("/about", about)
	http.HandleFunc("/content/", contenthandler)
	http.HandleFunc("/downloads/", downloads)
	//static serving
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("files"))))
	log.Fatal(http.ListenAndServe(":80", nil))

}
