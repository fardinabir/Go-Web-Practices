package main

//package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rocketlaunchr/google-search"
	"io/ioutil"
	"net/http"
)

type Result struct {

	// Rank is the order number of the search result.
	Rank int `json:"rank"`

	// URL of result.
	URL string `json:"url"`

	// Title of result.
	Title string `json:"title"`

	// Description of the result.
	Description string `json:"description"`
}

func main() {
	ctx := context.Background()
	name := "code builder"
	site := "github.com"

	result, _ := googlesearch.Search(ctx, name+" site:"+site)
	url := result[0].URL
	repoUserSuffix := url[19:len(url)]
	newApiUrl := "https://api.github.com/repos/" + repoUserSuffix
	fmt.Println(result, newApiUrl)

	resp, err := http.Get(newApiUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	//fmt.Printf("Body : %s\n", body)

	var respRepo RepoDetails
	json.Unmarshal(body, &respRepo)

	fmt.Println(respRepo.FullName, "\t>>>>>>>>>>>>>>>\t", respRepo.StargazersCount)
}
