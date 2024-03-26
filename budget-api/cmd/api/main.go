package main

import (
	"go.uber.org/zap"
	"jemyishai/budget-app/cmd/api/server"
	"runtime"
)

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewDevelopment()))
	zap.L().Info("logger is initialized")
}

func main() {
	numCPU := runtime.NumCPU()
	zap.L().Info("Number of CPU", zap.Int("numCPU", numCPU))
	if runtime.NumCPU() > 2 {
		runtime.GOMAXPROCS(numCPU / 2)
	}
	app, err := server.NewApp()
	if err != nil {
		zap.L().Fatal("failed to initialize app", zap.Error(err))
	}
	if err := app.Run(); err != nil {
		zap.L().Fatal("failed to run app", zap.Error(err))
	}
}
