package view

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Page - simple html page structure
type Page struct {
	Title string
	Body  []byte
}

const (
	// Htmlpath - the folder to keep dynamic html pages
	Htmlpath string = "/views/"
)

// Save - Save a wiki page on disk
func (p *Page) Save() error {
	_, err := os.Stat(Htmlpath)
	if err != nil {
		log.Printf("html folder doesn't exist - %s", err)
		log.Printf("creating html folder at - %s", Htmlpath)
		os.MkdirAll(Htmlpath, 0600)
	}
	filename := Htmlpath + p.Title + ".html"
	log.Println("Writing to file", filename)
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// Copy - copy a wiki page from source location to a destination location
func Copy(dstname, srcname string) (written int64, err error) {
	src, err := os.Open(Htmlpath + srcname)
	if err != nil {
		log.Printf("Not able to open source file, error - %s", err)
		return 0, err
	}
	defer func() {
		if src != nil {
			src.Close()
		}
	}()

	dst, err := os.OpenFile(Htmlpath+dstname, os.O_RDWR, 0644)
	if err != nil {
		log.Printf("Not able to open destination file, error - %s", err)
		return 0, err
	}

	defer func() {
		if src != nil {
			src.Close()
		}
	}()

	written, err = io.Copy(dst, src)
	return
}

// Load - to load a html page view
func Load(title string) (*Page, error) {
	filename := Htmlpath + title + ".html"
	p := Page{Title: title, Body: []byte("")}
	//defer errorhandler(&p)
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Cannot read file, error - %s", err)
		p.Title = "Error"
	} else {
		p.Body = body
	}
	return &p, err
}

func errorhandler(p *Page) {
	rmsg := recover()
	if rmsg != nil {
		p.Body = []byte(fmt.Sprintf("%s", rmsg))
		log.Printf("Recovered with error msg %s", rmsg)
	}
}
