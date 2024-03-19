package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	utils "github.com/dwskme/dwSearchMe/utils"
)

func main() {
	var dumpPath, query string
	flag.StringVar(&dumpPath, "p", "temp/enwiki-latest-abstract1.xml.gz", "wiki abstract dump path")
	flag.Parse()

	log.Println("Running Full Text Search")

	start := time.Now()
	docs, err := utils.LoadDocument(dumpPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

	idx := make(utils.Index)
	idx.Add(docs)
	log.Printf("Indexed %d documents", len(docs))

	for {
		fmt.Print("\nEnter search query (type 'exit' to quit): ")
		fmt.Scanln(&query)
		if query == "exit" {
			break
		}

		start = time.Now()
		matchedIDs := idx.Search(query)
		log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))

		for _, id := range matchedIDs {
			doc := docs[id]
			log.Printf("%d\t%s\n", id, doc.Text)
		}
	}
}

