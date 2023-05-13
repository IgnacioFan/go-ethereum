/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"go-ethereum/internal/http"

	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Run http server",
	Run: func(cmd *cobra.Command, args []string) {
		server := http.NewHttpServer()
		if err := server.Start(8081); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
