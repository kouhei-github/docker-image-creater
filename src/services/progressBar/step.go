package progressBar

import (
	"github.com/cheggaaa/pb/v3"
	"math/rand"
	"time"
)

type ManageProgress struct {
	Count int
	Bar   *pb.ProgressBar
}

func NewManagerProgress(count int, total int, maxWidth int) *ManageProgress {
	return &ManageProgress{
		Count: count,
		Bar:   pb.New(total).SetMaxWidth(maxWidth),
	}
}

func (progress *ManageProgress) Progress() {
	progress.Bar.Start()
	for i := 0; i < progress.Count; i++ {
		progress.Bar.Increment()
		waitTime := rand.Intn(501)
		time.Sleep(time.Millisecond * time.Duration(waitTime))
	}
	progress.Bar.Finish()
}
