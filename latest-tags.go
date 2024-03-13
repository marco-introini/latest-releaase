package main

import (
	"bufio"
	"context"
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

	// Crea un client GitHub
	client := github.NewClient(nil)

	// Ottieni l'ultima release del repository
	release, _, err := client.Repositories.GetLatestRelease(context.Background(), owner, repo)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("REPO ", owner, repo, release.GetTagName(), release.GetPublishedAt().Format("2006-01-02"))

}
