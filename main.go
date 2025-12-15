package main

import "slg_sever/internal"

func main() {
	internal.GetTimer().Start(func(tick int64) {
		internal.GetSLGMap().Tick(tick)
	})
}
