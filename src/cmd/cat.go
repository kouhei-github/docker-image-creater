/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/spf13/cobra"
	"os"
	"sync"
	"time"
)

// catCmd represents the cat command
var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		for i := 0; i < 20; i++ {
			wg.Add(1)
			myProgress := MyProgressManager{
				count: 100,
				bar:   pb.Simple.Start(100),
			}
			go func(myProgress MyProgressManager) {
				myProgress.progress()
				wg.Done()
			}(myProgress)
		}
		wg.Wait()

		b, err := os.ReadFile("./ASCII_ART/cat.txt")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(string(b))
	},
}

type MyProgressManager struct {
	count int
	bar   *pb.ProgressBar
}

func (progress *MyProgressManager) progress() {
	progress.bar.SetMaxWidth(80)
	for i := 0; i < progress.count; i++ {
		progress.bar.Increment()
		time.Sleep(time.Millisecond * 30)
	}
	progress.bar.Finish()
}

func init() {
	rootCmd.AddCommand(catCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// catCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// catCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
