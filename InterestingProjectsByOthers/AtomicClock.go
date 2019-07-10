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
	fmt.Println()

	var doc2 *goquery.Document
	var err2 error
	if doc2, err2 = goquery.NewDocument("https://scarborough.co.uk"); err2 != nil {
		panic(err2)
	}
	fmt.Println(doc2.Find(""))
}
