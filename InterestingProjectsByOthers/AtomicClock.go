package InterestingProjectsByOthers

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func AtomicClock() {
	var doc *goquery.Document
	var err error
	if doc, err = goquery.NewDocument("http://time.is/just"); err != nil {
		panic(err)
	}
	fmt.Printf("%s", doc.Find("#twd").Text())
}
