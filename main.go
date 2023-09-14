package main

import (
	"errors"
	"fmt"
	"os"

	"example.com/fetch-web/cmd"
	customErrors "example.com/fetch-web/pkg/errors"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		switch {
		case errors.Is(err, customErrors.ErrInvalidURL):
			fmt.Println("Invalid URL:", err)
		case errors.Is(err, customErrors.ErrCreateFileError):
			fmt.Println("Error creating file:", err)
		case errors.Is(err, customErrors.ErrFetchError):
			fmt.Println("Error fetching web page:", err)
		default:
			fmt.Println("Unknown error:", err)
		}
		os.Exit(1)
	}
}
