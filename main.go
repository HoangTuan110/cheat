package main

import (
	"fmt"
	"os"
	"time"
	"log"

	"github.com/gocolly/colly/v2"
	"github.com/spf13/cobra"
)

func fetchData() {
	// First, check if the CLI argument has more than one element, if not, we will told the user to provide the topic
	if len(os.Args) < 2 {
		log.Fatalf("Please provide a topic")
	}
	// Take the name as the first argument, then combine it with the base URL ("https://cht.sh/") to create the URL
	name := os.Args[1]
	url := "https://cht.sh/" + name

	// Start scraping
	c := colly.NewCollector()
	c.Clone().SetRequestTimeout(64 * time.Second) // Set timeout

	// Find the pre (where all the contents are met) and extract the contents
	c.OnHTML("pre", func(e *colly.HTMLElement) {
		content := e.Text
		content = content[:len(content)-3] // Remove the last character ("$" and some "\n")

		fmt.Println(content)
	})

	c.OnError(func(r *colly.Response, e error) {
		log.Fatalf("Request URL: %s failed with response: %s\nStatus code: %d", r.Request.URL, e, r.StatusCode)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Printf("Visiting %s", r.URL.String())
	})

	c.Visit(url)
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "cheat <topic>",
		Short: "A CLI tool that fetches data from cht.sh",
		Long:  `'cheat' is a program that fetches data from cht.sh based on the input`,
		Run: func(cmd *cobra.Command, args []string) {
			fetchData()
		},
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
