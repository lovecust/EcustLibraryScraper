# Ecust Library Scraper

<!-- > 8:57 PM, Sunday, June 18, 2017. -->

A scraper for Ecust Library in *Golang*.

## Installation

```bash
go get github.com/lovecust/EcustLibraryScraper
```

## Examples


```golang
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
```
