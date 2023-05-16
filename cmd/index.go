package cmd

import (
	"fmt"
	"go-ethereum/internal/delivery/index"
	"go-ethereum/pkg/postgres"

	"github.com/spf13/cobra"
)

// indexCmd represents the index command
var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "Run blocks indexer",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Blocks index called")
		pg, err := postgres.NewPostgres()
		if err != nil {
			fmt.Println("postgres:", err)
		}
		fmt.Println(pg)
		indexer := index.NewEthIndexer()
		indexer.Run()
	},
}

func init() {
	rootCmd.AddCommand(indexCmd)
}
