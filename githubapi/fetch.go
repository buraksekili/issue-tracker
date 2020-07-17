package githubapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func SearchIssues(args *Flags) (*IssuesSearchResult, error) {

	if strings.Count(args.Repo, "/") != 1 {
		return nil, fmt.Errorf("fetch.go: wrong repo name as %s\nCheck --help for proper format of 'repo' name.\n", args.Repo)
	}

	status := "is:"
	if args.Status {
		status += "open"
	} else {
		status += "closed"
	}

	sorting := "sort:created-"
	// sorting := "sort:created&direction="

	if args.Asc {
		fmt.Println("asc")
		sorting += "asc"
		fmt.Println(sorting)
	} else {
		fmt.Println("desc")
		sorting += "desc"
		fmt.Println(sorting)
	}
	// + " page=3"
	query := url.QueryEscape(args.Repo + " " + status + " " + sorting + " json " + "decoder")
	resp, err := http.Get(ApiURL + "?q=" + query)
	fmt.Println("fetching: ", ApiURL+"?q="+query)
	if err != nil {
		fmt.Printf("fetch.go: fetching query\n%s\n", err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("fetch.go: resp status code\n%s\n", err)
		resp.Body.Close()
		return nil, err
	}
	var results IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &results, nil
}

// func displayData(issues *IssuesSearchResult, count int) (IssuesSearchResult, error) {
//
// }
