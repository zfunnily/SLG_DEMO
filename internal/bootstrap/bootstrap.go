package bootstrap

import (
	"slg_sever/internal/global"
	"slg_sever/internal/http"
	"slg_sever/internal/test"
	"slg_sever/pkg/timer"
	"slg_sever/pkg/uuid"
)

func Run() {

	println("bootstrap Run...")

	global.InitWorld()
	data := global.InitMarchUnit()
	uuid.Init(0, 0)

	t := timer.NewTickTimer()
	t.Start(func(tick int64) {
		global.GetWorld().Tick(tick)
	})

	// 启动行军单位
	go test.StartMarch(data)

	// 启动HTTP服务器
	http.StartHTTPServer(8080)
}
