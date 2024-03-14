# Latest Release

A simple CLI program for checking the latest release of a GitHub repository.

## Usage

- add the name of the repositories you want to check in a text file, one per line
- run the program with the following command:

```bash
go run latest-release.go --file <filename> --days <numOfDays>
```

If the file flag is not provided, the program will look for a file called `repositories.txt` in the current directory.

If the days flag is not provided, the program will not check if the latest release of the repository is newer then numOfDays.