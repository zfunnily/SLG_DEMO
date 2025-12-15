package internal

import timer "slg_sever/timer"

var t *timer.TickTimer

func GetTimer() *timer.TickTimer {
	return t
}

func init() {
	t = timer.NewTickTimer()
}
