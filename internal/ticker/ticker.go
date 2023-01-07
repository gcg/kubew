package ticker

import (
	"time"

	"github.com/gcg/kubew/internal/screen"
)

const refreshInt = time.Second * 1

func RefreshScreen() {
	tick := time.NewTicker(refreshInt)
	for {
		select {
		case <-tick.C:
			screen.App.Draw()
		}
	}
}
