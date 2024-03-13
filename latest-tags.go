package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/v60/github"
)

func main() {
	file, err := os.Open("repositories.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		repo := strings.Split(scanner.Text(), "/")
		if len(repo) != 2 {
			fmt.Printf("Invalid repository format: %s\n", scanner.Text())
			continue
		}

		owner, repository := repo[0], repo[1]
		getLatestTag(owner, repository)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func getLatestTag(owner, repo string) {
	client := github.NewClient(nil)

	opt := &github.ListOptions{Page: 1, PerPage: 1}
	tags, _, err := client.Repositories.ListTags(_, owner, repo, opt)
	if err != nil {
		fmt.Printf("Error getting tags for repository %s/%s: %v\n", owner, repo, err)
		return
	}

	if len(tags) == 0 {
		fmt.Printf("No tags found for repository %s/%s\n", owner, repo)
	} else {
		fmt.Printf("Latest tag for repository %s/%s: %s, created at %s\n", owner, repo, *tags[0].Name, tags[0].Commit.Author.GetDate().Format("2006-01-02"))
	}
}
