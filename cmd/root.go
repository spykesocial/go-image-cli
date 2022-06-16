/*
Copyright © 2022 Spyke Social Private Limited.
Author: yashdiniz

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spykesocial/go-image-cli/img"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-image-cli",
	Short: "Prints the dimensions of an image from a URL",
	Long: `go-image-cli uses the image package of go to download the image from a URL 
and print the dimensions of that image.
	example: go-image-cli -u URL	
`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		flagVal, err := cmd.Flags().GetString("url")
		if err != nil {
			log.Fatal("Need a URL to fetch image from")
		}

		config, err := img.GetImageConfig(flagVal)
		if err != nil {
			log.Fatal("Failed to download image from URL", err)
		}
		fmt.Println("Image width:", config.Width)
		fmt.Println("Image Height:", config.Height)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-image-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("url", "u", "", "The URL to get the image from")

}
