package timer

import "time"

type TickTimer struct {
	ticker   *time.Ticker
	quit     chan struct{}
	tick     int64
	lastTime int64
}

func (t *TickTimer) Start(onTick func(tick int64)) {
	t.ticker = time.NewTicker(time.Second)

	go func() {
		for {
			select {
			case <-t.ticker.C:
				now := time.Now().Unix()
				if t.lastTime == 0 {
					t.lastTime = now
				}

				delta := now - t.lastTime
				t.lastTime = now

				t.tick += delta
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
