package main

import (
	"message-persist/internal/app"
)

func main() {
	application := app.NewApp()

	application.Run()
}
