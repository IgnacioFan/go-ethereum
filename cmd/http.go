package cmd

import (
	"go-ethereum/internal/delivery/http"
	"go-ethereum/pkg/postgres"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var port int

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Run http server",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		db, err := postgres.NewPostgres()
		if err = db.NewMirgate(); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		server, err := http.NewHttpServer()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		if err = server.Start(port); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
	httpCmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "expose port number")
}
