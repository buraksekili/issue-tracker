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

func main() {
	flag.Parse()
	args := githubapi.Flags{Repo: *repo, Status: *status, Count: *count}

	result, err := githubapi.SearchIssues(&args)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues found in %s\n", result.TotalCount, *repo)
	if *disable {
		for _, item := range result.Items {
			fmt.Printf("#%-5d %.55s\n",
				item.Number, item.Title)
		}
	} else {
		for _, item := range result.Items {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}
