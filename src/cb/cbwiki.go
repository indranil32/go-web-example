package cb

import (
	"log"

	"gopkg.in/couchbase/gocb.v1"
)

// Wikidoc - the sample wiki page
type Wikidoc struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	Timestamp string `json:"timestamp"`
}

var globalbucket *gocb.Bucket
var globalcluster *gocb.Cluster

// InitCB -  Couchbase connection initialization
func InitCB() {
	log.Println("Couchbase connection initialization")
	cluster, err := gocb.Connect("couchbase://localhost")
	if err != nil {
		log.Fatalf("Failed to init CB %s", err)
	}
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "Administrator",
		Password: "Administrator",
	})
	bucket, err := cluster.OpenBucket("beer-sample", "") // password
	if err != nil {
		log.Fatalf("Failed to init CB %s", err)
	}
	log.Printf("Bucket opened %s", bucket.Name())
	globalcluster = cluster
	globalbucket = bucket
}

// Destroy - close all resources
func Destroy() {
	globalbucket.Close()
	globalcluster.Close()
	log.Println("Cluster and bucket closed!!")
}

func UpsertWikiPage(title string, doc Wikidoc) {
	_, err := globalbucket.Upsert(title, doc, 0)
	if err != nil {
		log.Printf("Error upserting %s", err)
	}
	//log.Printf("Cas %s", cas)
}

func GetWikiPage(title string) Wikidoc {
	var doc Wikidoc
	_, err := globalbucket.Get(title, &doc)
	if err != nil {
		log.Printf("CB Fetch error %s", err)
	}
	return doc
}
