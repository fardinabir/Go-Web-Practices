package finders

import (
	"context"
	"encoding/json"
	"find-a-repo/models"
	"fmt"
	googlesearch "github.com/rocketlaunchr/google-search"
	"io/ioutil"
	"net/http"
	"strings"
)

func FindRepos(keywords []string) {
	ctx := context.Background()
	name := strings.Join(keywords, " ")
	site := "github.com"

	result, _ := googlesearch.Search(ctx, name+" site:"+site)
	printRepos(&result, name)
}

func printRepos(result *[]googlesearch.Result, name string) {
	for i := 0; i < len(*result); i++ {
		url := (*result)[i].URL
		repoUserSuffix := url[19:len(url)]
		newApiUrl := "https://api.github.com/repos/" + repoUserSuffix
		fmt.Println(name)

		resp, err := http.Get(newApiUrl)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		//fmt.Printf("Body : %s\n", body)

		var respRepo models.RepoDetails
		json.Unmarshal(body, &respRepo)

		fmt.Println(" https://github.com/"+respRepo.FullName, "\t>>>>>>>>>>>>>>>\t", respRepo.StargazersCount)
	}
}
