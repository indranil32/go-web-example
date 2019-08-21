package view

import (
	"log"
	"os"
	"testing"
)

func Test_save(t *testing.T) {
	var tests = []Page{
		Page{"Test Title", []byte("This is a test title!")},
		Page{"", []byte("")},
		Page{"Test Tile with empty body", nil},
	}

	for i, p := range tests {
		log.Printf("Testing Save for - #%d <%s> -> (%s)",i, p.Title, string(p.Body))
		err := p.Save()
		if err != nil {
			t.Errorf("Save failed %s", err)
		}
		// cleanup
		defer os.Remove(Htmlpath + p.Title + ".html")
	}
}

func Test_Copy(t *testing.T) {
	var src, dest string = "src.html", "dest.html"
	defer os.Remove(Htmlpath + src)
	defer os.Remove(Htmlpath + dest)

	p1 := Page{"src", []byte("This is test for copy")}
	p2 := Page{"dest", []byte("")}
	err1 := p1.Save()
	if err1 != nil {
		t.Errorf("Save failed %s", err1)
	}
	err2 := p2.Save()
	if err2 != nil {
		t.Errorf("Save failed %s", err2)
	}

	count, err := Copy(dest, src)
	if err != nil {
		t.Errorf("Copy failed %s", err)
	}
	log.Printf("Byte count %d", count)
}

func Test_load(t *testing.T) {
	p := Page{"test3", []byte("This is test for load")}
	p.Save()
	err := p.Save()
	if err != nil {
		t.Errorf("Save failed %s", err)
	}
	// cleanup
	defer os.Remove(Htmlpath + p.Title + ".html")

	pl, err := Load(p.Title)
	if err != nil || string(pl.Body) != string(p.Body) {
		t.Errorf("Page Load failed %s", err)
	}
}

func Test_load_negative(t *testing.T) {
	p, err := Load("NA")
	if err == nil && p.Title != "Error" {
		t.Errorf("Negative Load failed %s", err)
	}
}
