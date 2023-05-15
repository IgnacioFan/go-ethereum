package cmd

import (
	"fmt"
	"go-ethereum/internal/delivery/http"

	"github.com/spf13/cobra"
)

var port int

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Run http server",
	Run: func(cmd *cobra.Command, args []string) {
		server := http.NewHttpServer()
		if err := server.Start(port); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
	httpCmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "expose port number")
}
