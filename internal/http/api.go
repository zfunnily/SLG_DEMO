package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slg_sever/internal/battle"
	"slg_sever/internal/global"
	"slg_sever/internal/march"
	"slg_sever/internal/node"
)

// 定义请求和响应结构

// StartMarchRequest 开始行军请求
type StartMarchRequest struct {
	UnitID int64         `json:"unit_id"`
	Path   []node.NodeID `json:"path"`
}

// StartMarchResponse 开始行军响应
type StartMarchResponse struct {
	MarchID  int64  `json:"march_id"`
	ArriveAt int64  `json:"arrive_at"`
	Success  bool   `json:"success"`
	Message  string `json:"message"`
}

// SpeedUpMarchRequest 加速行军请求
type SpeedUpMarchRequest struct {
	MarchID      int64 `json:"march_id"`
	SpeedUpTicks int64 `json:"speed_up_ticks"`
}

// SpeedUpMarchResponse 加速行军响应
type SpeedUpMarchResponse struct {
	NewArriveAt int64  `json:"new_arrive_at"`
	Success     bool   `json:"success"`
	Message     string `json:"message"`
}

// StartHTTPServer 启动HTTP服务器
func StartHTTPServer(port int) {
	http.HandleFunc("/api/march/start", handleStartMarch)
	http.HandleFunc("/api/march/speedup", handleSpeedUpMarch)

	addr := fmt.Sprintf(":%d", port)
	log.Printf("HTTP server starting on %s", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("HTTP server failed to start: %v", err)
	}
}

// handleStartMarch 处理开始行军请求
func handleStartMarch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req StartMarchRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 获取当前世界和tick
	world := global.GetWorld()
	currentTick := world.GetNowTick() // 这里需要实现获取当前tick的方法

	// 调用MarchMgr创建行军
	marchID := world.MarchMgr.CreateMarch(
		battle.BattleUnitID(req.UnitID),
		req.Path,
		currentTick,
	)

	// 获取行军信息
	marchInfo, exists := world.MarchMgr.GetMarch(marchID)
	if !exists {
		http.Error(w, "March creation failed", http.StatusInternalServerError)
		return
	}

	// 构造响应
	resp := StartMarchResponse{
		MarchID:  int64(marchID),
		ArriveAt: marchInfo.ArriveAt,
		Success:  true,
		Message:  "March started successfully",
	}

	// 发送响应
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// handleSpeedUpMarch 处理加速行军请求
func handleSpeedUpMarch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SpeedUpMarchRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 获取当前世界和tick
	world := global.GetWorld()
	currentTick := world.GetNowTick()

	// 调用MarchMgr加速行军
	world.MarchMgr.SpeedUpMarch(
		march.MarchID(req.MarchID),
		req.SpeedUpTicks,
		currentTick,
	)

	// 获取更新后的行军信息
	marchInfo, exists := world.MarchMgr.GetMarch(march.MarchID(req.MarchID))
	if !exists {
		http.Error(w, "March not found", http.StatusNotFound)
		return
	}

	// 构造响应
	resp := SpeedUpMarchResponse{
		NewArriveAt: marchInfo.ArriveAt,
		Success:     true,
		Message:     "March speed up successfully",
	}

	// 发送响应
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
