package cmd

import (
	"fmt"
	"log"

	fetch "example.com/fetch-web/internal/fetch"
	metadata "example.com/fetch-web/internal/metadata"
	"example.com/fetch-web/internal/regex"
	customErrors "example.com/fetch-web/pkg/errors"
	"github.com/spf13/cobra"
)

var printMetadata bool

func init() {
	RootCmd.AddCommand(fetchCmd)
	RootCmd.PersistentFlags().BoolVarP(&printMetadata, "metadata", "m", false, "print metadata")
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Serves Post List Synchronizer",
	RunE:  fetchUrls,
}

func fetchUrls(cmd *cobra.Command, args []string) error {
	for _, arg := range args {
		if urlIsValid := regex.UrlIsValid(arg); !urlIsValid {
			log.Printf("Error validating args")
			return fmt.Errorf("url is not valid. make sure url is proper pattern http[s]://xyz.com: %w", customErrors.ErrCreateFileError)
		}

		content, err := fetch.FetchWebPage(arg)
		if err != nil {
			return err
		}

		err = fetch.SaveFile(content, arg)
		if err != nil {
			return err
		}

		if printMetadata {
			fmt.Printf("Metadata for url: %s\n", arg)
			md := metadata.ParseMetadata(content)
			stringMd := metadata.FormatMetadata(md)
			fmt.Println(stringMd)
		}
	}

	return nil
}
