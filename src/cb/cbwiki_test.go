package cb

import (
	"testing"
	"time"
)

func Test_get(t *testing.T) {
	defer Destroy()
	InitCB()
	GetWikiPage("page1")
}

func Test_upsert(t *testing.T) {
	defer Destroy()
	InitCB()
	testdoc := Wikidoc{
		"page1",
		"This is a test page",
		time.Now().Format(time.RFC850),
	}
	UpsertWikiPage("page1", testdoc)
}
