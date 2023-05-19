package cmd

import (
	"go-ethereum/internal/delivery/index"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	start  int64
	window int64
	end    int64
	sleep  int64
	// indexCmd represents the index command
	indexCmd = &cobra.Command{
		Use:   "index",
		Short: "Run blocks indexer",
		Run: func(cmd *cobra.Command, args []string) {
			indexer, err := index.NewEthIndexer()
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			indexer.Run(start, window, end, sleep)
		},
	}
)

func init() {
	rootCmd.AddCommand(indexCmd)
	indexCmd.Flags().Int64Var(&start, "start", 0, "Start is where the block starts")
	indexCmd.Flags().Int64Var(&window, "window", 5, "Window is the subset of block numbers ( (5 by default)")
	indexCmd.Flags().Int64Var(&end, "end", 0, "End is Where the block ends (optional, default value is obtained from the JSON-RPC API)")
	indexCmd.Flags().Int64Var(&sleep, "sleep", 3, "How many seconds you want service to sleep (3 by default)")
}
