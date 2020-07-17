package main

import (
	"./githubapi"
	"flag"
	"fmt"
	"log"
)

var repo = flag.String("r", "", "Github repo name (as 'golang/go')")
var status = flag.Bool("s", false, "Issue status\nTo see open issues, type -s; otherwise, do not type -s")
var count = flag.Int("n", 10, "Number of issues to fetch")
var disable = flag.Bool("d", false, "Disable to display username.")
var sorting = flag.Bool("a", true, "sort by ascending author date.")

func main() {

	flag.Parse()
	fmt.Println(*sorting)
	args := githubapi.Flags{Repo: *repo, Status: *status, Count: *count, Asc: *sorting}

	result, err := githubapi.SearchIssues(&args)
	if err != nil {
		log.Fatal(err)
	}

	stat := "closed"
	if *status {
		stat = "open"
	}

	fmt.Printf("%d %s issues found in %s\n", result.TotalCount, stat, *repo)
	if *disable {
		for _, item := range result.Items {
			if counter := 0; counter < *count {
				fmt.Printf("#%-5d %.55s\n",
					item.Number, item.Title)
				counter++
			}
		}
	} else {
		counter := 1
		for _, item := range result.Items {
			if counter > *count {
				break
			}
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
			counter++
		}
	}

	// tmp := result.Items[:args.Count]
}
