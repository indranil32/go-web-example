package main

import (
	"fmt"
	"go-web-example/cb"
	"go-web-example/view"
	"log"
	"net/http"
	"time"
)

func viewhandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := view.Load(title)
	if err != nil {
		fmt.Fprintln(w, "<h1>Not Found</h1><div>404</div>")
	} else {
		fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	}
}

func edithandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := view.Load(title)
	if err != nil {
		fmt.Fprintln(w, "<h1>Not Found</h1><div>404</div>")
	} else {
		fmt.Fprintf(w, "<h1>Editing %s</h1>"+
			"<form action=\"/save/%s\" method=\"POST\">"+
			"<textarea name=\"body\">%s</textarea><br>"+
			"<input type=\"submit\" value=\"Save\">"+
			"</form>",
			p.Title, p.Title, p.Body)
	}
}

func savehandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &view.Page{Title: title, Body: []byte(body)}
	p.Save()
	doc := cb.Wikidoc{Title: title, Body: body, Timestamp: time.Now().Format(time.RFC850)}
	cb.UpsertWikiPage(title, doc)
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	cb.InitCB()
	http.HandleFunc("/view/", viewhandler)
	http.HandleFunc("/edit/", edithandler)
	http.HandleFunc("/save/", savehandler)
	log.Fatal(http.ListenAndServe(":9080", nil))
}
