package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/net/context"
)

const remaingThreshold = 1

func downloadFromURL(url, dirName, fileName string) (err error) {
	filePath := fmt.Sprintf("%s/%s", dirName, fileName)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error while opening file", filePath, err)
		return
	}

	_, err = io.Copy(out, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	return
}

func main() {

	t := &github.UnauthenticatedRateLimitedTransport{
		ClientID:     "npateriya",
		ClientSecret: "ae35d520d69faf2b41d252c76313266fc72a3360",
	}
	ctx := context.Background()
	client := github.NewClient(t.Client())

	fmt.Println("Repos that contain cisco.")

	filename := "./ciscorepolist.csv"

	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		//this is a fatal error, quit
		return
	}
	defer f.Close()

	// slice the queries into batches to get around the API limit of 1000

	queries := []string{
		"\"2001-06-01 .. 2010-12-31\"",
		"\"2013-01-01 .. 2013-12-31\"",
		"\"2014-01-01 .. 2014-12-31\"",
		"\"2015-01-01 .. 2015-12-31\"",
		"\"2016-01-01 .. 2016-06-30\"",
		"\"2016-07-01 .. 2016-12-31\"",
		"\"2017-01-01 .. 2017-06-30\"",
		"\"2017-07-01 .. 2017-12-31\"",
	}
	header := fmt.Sprint("repoURL, repoName, createdAt, updatedAt, username, language," +
		"forksCount, subscribersCount, watchersCount, stargazersCount," +
		"orgName, orgCompany\n")
	_, err = io.WriteString(f, header)
	for _, q := range queries {

		query := fmt.Sprintf("cisco created:" + q)

		page := 1
		maxPage := math.MaxInt32

		opts := &github.SearchOptions{
			Sort:  "updated",
			Order: "desc",
			ListOptions: github.ListOptions{
				PerPage: 100,
			},
		}

		for ; page <= maxPage; page++ {
			opts.Page = page
			result, response, err := client.Search.Repositories(ctx, query, opts)
			if err != nil {
				log.Fatal("FindRepos:", err)
				break
			}

			maxPage = response.LastPage

			msg := fmt.Sprintf("page: %v/%v, size: %v, total: %v",
				page, maxPage, len(result.Repositories), *result.Total)
			log.Println(msg)

			for _, repo := range result.Repositories {

				repoName := *repo.FullName
				username := *repo.Owner.Login
				repoURL := *repo.HTMLURL
				readmeURL := fmt.Sprintf("%s/master/README.md", repoURL)
				readmeURL = strings.Replace(readmeURL, "github.com", "raw.githubusercontent.com", 1)
				repoNormName := strings.Replace(repoName, "/", "__", 1)
				language := ""
				if repo.Language != nil {
					language = *repo.Language
				}
				createdAt := repo.CreatedAt.String()
				updatedAt := repo.UpdatedAt.String()
				forksCount := 0
				if repo.ForksCount != nil {
					forksCount = *repo.ForksCount
				}
				subscribersCount := 0
				if repo.SubscribersCount != nil {
					subscribersCount = *repo.SubscribersCount
				}
				watchersCount := 0
				if repo.WatchersCount != nil {
					watchersCount = *repo.WatchersCount
				}
				stargazersCount := 0
				if repo.StargazersCount != nil {
					stargazersCount = *repo.StargazersCount
				}
				orgName := ""
				if repo.Organization != nil && repo.Organization.Name != nil {
					orgName = *repo.Organization.Name
				}
				orgCompany := ""
				if repo.Organization != nil && repo.Organization.Company != nil {
					orgCompany = *repo.Organization.Company
				}
				//fmt.Printf("repo: %s, owner:%s, created:%s", repoName, username, createdAt)

				row := fmt.Sprintf("%s,%s,%s,%s,%s,%s,%d,%d,%d,%d,%s,%s,%s,%s\n",
					repoURL, repoName, createdAt, updatedAt, username, language,
					forksCount, subscribersCount, watchersCount, stargazersCount,
					orgName, orgCompany, readmeURL, repoNormName)
				n, err := io.WriteString(f, row)
				if err != nil {
					fmt.Println(n, err)
				}
				err = downloadFromURL(readmeURL, "Readmes", repoNormName)

			}
			time.Sleep(time.Millisecond * 7000)

		}
	}

}
