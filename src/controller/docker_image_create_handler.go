package controller

import (
	"fmt"
	"github.com/spf13/cobra"
	"math/rand"
	"nagamatsu-cobra/services"
	"nagamatsu-cobra/services/progressBar"
	"os"
	"sync"
	"time"
)

func DockerImageCreateHandler(cmd *cobra.Command, arg []string) {
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
		waitTime := rand.Intn(1001) + 100
		time.Sleep(time.Millisecond * time.Duration(waitTime))
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
}

func poolImplement(value string) services.PoolImp {
	switch value {
	case "MyProgressBar":
		return progressBar.NewManagePool()
	}
	return progressBar.NewManagePool()
}

// runPool 使い方
//
// -------------------
//
// bar := progressBar.NewManagePool()
//
//	if err := runPool(bar); err != nil {
//	   fmt.Println(err.Error())
//	   return
//	}
//
// -------------------
func runPool(imp services.PoolImp) error {
	first := progressBar.NewManagerProgress(100, 100, 80)
	second := progressBar.NewManagerProgress(100, 100, 80)
	third := progressBar.NewManagerProgress(100, 100, 80)

	imp.AddPool(first.Bar, second.Bar, third.Bar)

	return imp.Start()
}
