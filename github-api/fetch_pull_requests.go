package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type IssueRequest struct {
	IncompleteResults bool `json:"incomplete_results"`
	TotalCount        int  `json:"total_count"`
}

func main() {
	token := os.Getenv("GITHUB_TOKEN")

	org := flag.String("organization", "default", "github organization")
	author := flag.String("author", "default", "author")
	start_date := flag.String("start-date", "2019-01-01", "start-date")
	end_date := flag.String("end-date", "2019-12-31", "end-date")
	flag.Parse()

	query := fmt.Sprintf("org:%s+author:%s+merged:%s..%s", *org, *author, *start_date, *end_date)
	end_point := "https://api.github.com/search/issues"
	url := end_point + "?q=" + query + "&access_token=" + token

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)

	var issue_request IssueRequest
	if err := json.Unmarshal(byteArray, &issue_request); err != nil {
		log.Fatal(err)
	}
	fmt.Println(issue_request.TotalCount)

}
