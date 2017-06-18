package main

import (
	"github.com/lovecust/EcustLibraryScraper/status"
	"log"
)

func main() {
	status, err := LibraryStatus.GetStatus()
	if err != nil {
		log.Fatal("Failed to get library status.", err)
	}
	log.Println(status.ToJson())
}
