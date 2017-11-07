package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
	                  love by spf13 and friends in Go.
					                  Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var param string

func init() {
	RootCmd.AddCommand(startCmd)
	f := startCmd.Flags()
	f.StringVarP(&param, "param", "p", "", "Source directory to read from")
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "run Hugo",
	Long:  `run`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("param is ", param)
	},
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
