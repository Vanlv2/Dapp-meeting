package main

import (
	"log"

	"Dapp-meeting/bandwidthSaving/config"
	"Dapp-meeting/bandwidthSaving/handlers"
	"Dapp-meeting/bandwidthSaving/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	cfg := config.LoadConfig()

	// Tạo instance CloudflareService với account ID và API token từ cấu hình.
	cloudflareService := services.NewCloudflareService(cfg.CloudflareAppID, cfg.CloudflareToken)

	// Tạo instance của MeetingHandler với CloudflareService được inject.
	meetingHandler := handlers.NewMeetingHandler(cloudflareService)

	e.POST("/optimize", meetingHandler.OptimizeMeetingHandler)

	e.Logger.Info("Server đang chạy trên cổng 8081...")
	log.Fatal(e.Start(":8081"))
}
