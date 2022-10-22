package main

import (
	"github.com/sirupsen/logrus"
	"github.com/weeabook/src/university-platform-backend/internal/app"
	"github.com/weeabook/src/university-platform-backend/internal/config"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	cfg := config.GetConfig()

	a, err := app.NewApp(cfg)
	if err != nil {
		logrus.Fatalf("Error with initializing app: %s", err)
	}

	a.Run()
}
