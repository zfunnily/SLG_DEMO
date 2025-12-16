package bootstrap

import "slg_sever/pkg/timer"

func Run() {

	println("bootstrap Run...")

	InitWorld()

	t := timer.NewTickTimer()
	t.Start(func(tick int64) {
		GetSLGMap().Tick(tick)
	})

	err := make(chan error)
	<-err
}
