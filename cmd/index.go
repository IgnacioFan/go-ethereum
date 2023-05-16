package cmd

import (
	"fmt"
	"go-ethereum/internal/delivery/index"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// indexCmd represents the index command
var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "Run blocks indexer",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Blocks index called")
		indexer, err := index.NewEthIndexer()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		indexer.Run(1, 5)
	},
}

func init() {
	rootCmd.AddCommand(indexCmd)
}
