package progressBar

import (
	"github.com/cheggaaa/pb/v3"
	"time"
)

type ManageProgress struct {
	Count int
	Bar   *pb.ProgressBar
}

func NewManagerProgress(count int, tootal int, maxWidth int) *ManageProgress {
	return &ManageProgress{
		Count: count,
		Bar:   pb.New(tootal).SetMaxWidth(maxWidth),
	}
}

func (progress *ManageProgress) Progress() {
	progress.Bar.Start()
	for i := 0; i < progress.Count; i++ {
		progress.Bar.Increment()
		time.Sleep(time.Millisecond * 30)
	}
	progress.Bar.Finish()
}
