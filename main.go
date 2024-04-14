package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
	"github.com/robfig/cron/v3"
	"github.com/walnuts1018/2024-golang-linebot/common"
	"github.com/walnuts1018/2024-golang-linebot/common/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to load config: %v", err))
		os.Exit(1)
	}

	logger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{
		TimeFormat: time.RFC3339,
		Level:      cfg.LogLevel,
	}))
	slog.SetDefault(logger)

	router, err := common.NewRouter(cfg)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to create router: %v", err))
		os.Exit(1)
	}

	// サーバー起動
	go func() {
		if err := router.Run(fmt.Sprintf(":%s", cfg.ServerPort)); err != nil {
			slog.Error(fmt.Sprintf("Failed to run router: %v", err))
			os.Exit(1)
		}
	}()

	c := cron.New()
	// 毎朝7時に実行
	c.AddFunc("0 7 * * *", func() {
		slog.Info("7時になりました")
	})

	c.Run()
}
