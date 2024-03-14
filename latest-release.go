package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/go-github/v60/github"
)

func main() {

	var fileName string
	var numDays int

	flag.StringVar(&fileName, "file", "repositories.txt", "File containing repositories")
	flag.IntVar(&numDays, "days", 10, "Number of days to look back")
	flag.Parse()

	if numDays < 0 {
		fmt.Println("Days must be >= 0")
		os.Exit(1)
	}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
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
		getLatestTag(owner, repository, numDays)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func getLatestTag(owner, repo string, days int) {

	// Crea un client GitHub
	client := github.NewClient(nil)

	// Ottieni l'ultima release del repository
	release, _, err := client.Repositories.GetLatestRelease(context.Background(), owner, repo)
	if err != nil {
		fmt.Println(err)
		return
	}
	releaseDate := release.GetPublishedAt().Time

	if days > 0 {
		today := time.Now()
		diff := today.Sub(releaseDate).Hours() / 24

		if diff < float64(days) {
			fmt.Printf("REPO %s %s ==> Release in the last %d days: %s on %s", owner, repo, days, release.GetTagName(), releaseDate)
		} else {
			fmt.Println("REPO", owner, repo, release.GetTagName(), releaseDate.Format("2006-01-02"))
		}
	} else {
		fmt.Println("REPO", owner, repo, release.GetTagName(), releaseDate.Format("2006-01-02"))
	}

}
