package main

import (
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
)

func wget(url string) {
	println("wget:", url)
	// fetch website content
	get, err := http.Get(url)
	if err != nil {
		println("error:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			println("error:", err)
		}
	}(get.Body)
	// print website content
	if _, err = io.Copy(io.Writer(os.Stdout), get.Body); err != nil {
		println("error:", err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "cli is a simple command line tool",
}

func init() {
	wgetCmd := &cobra.Command{
		Use:   "wget",
		Short: "wget website content",
		Long:  "wget website content by url",
		Run: func(cmd *cobra.Command, args []string) {
			url, err := cmd.Flags().GetString("url")
			if err != nil {
				println("error:", err)
				return
			}
			wget(url)
		},
	}
	wgetCmd.Flags().StringP("url", "u", "", "website url")
	if err := wgetCmd.MarkFlagRequired("url"); err != nil {
		panic(err)
	}
	rootCmd.AddCommand(wgetCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		return
	}
}
