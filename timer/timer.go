package timer

import "time"

type TickTimer struct {
	ticker *time.Ticker
	quit   chan struct{}
	tick   int64
}

func (t *TickTimer) Start(onTick func(tick int64)) {
	t.ticker = time.NewTicker(time.Second)

	go func() {
		for {
			select {
			case <-t.ticker.C:
				t.tick++
				onTick(t.tick)

			case <-t.quit:
				t.ticker.Stop()
				return
			}
		}
	}()
}

func (t *TickTimer) Stop() {
	close(t.quit)
}

func NewTickTimer() *TickTimer {
	return &TickTimer{
		quit: make(chan struct{}),
	}
}
