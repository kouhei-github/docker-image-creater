/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"nagamatsu-cobra/services"
	"nagamatsu-cobra/services/progressBar"
	"os"
	"sync"
	"time"
)

// catCmd represents the cat command
var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "A brief description of your command",
	Long: `
		A longer description that spans multiple lines and likely contains examples
		and usage of using your command. For example:

		Cobra is a CLI library for Go that empowers applications.
		This application is a tool to generate the needed files
		to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup

		pool := poolImplement("MyProgressBar")

		first := progressBar.NewManagerProgress(100, 100, 80)
		second := progressBar.NewManagerProgress(100, 100, 80)
		third := progressBar.NewManagerProgress(100, 100, 80)

		pool.AddPool(first.Bar, second.Bar, third.Bar)

		if err := pool.Start(); err != nil {
			fmt.Println(err.Error())
			return
		}

		var stepImp services.SteBarImp

		for _, progress := range []*progressBar.ManageProgress{first, second, third} {
			wg.Add(1)
			time.Sleep(time.Millisecond * 800)
			stepImp = progress
			go func() {
				stepImp.Progress()
				wg.Done()
			}()
		}
		wg.Wait()

		b, err := os.ReadFile("./ASCII_ART/cat.txt")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(string(b))
	},
}

func poolImplement(value string) services.PoolImp {
	switch value {
	case "MyProgressBar":
		return progressBar.NewManagePool()
	}
	return progressBar.NewManagePool()
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
