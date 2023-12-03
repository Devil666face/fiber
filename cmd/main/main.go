package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/Devil666face/fiber/internal/web"
	// _ "github.com/joho/godotenv/autoload"
)

func main() {
	wa := web.New()
	if err := wa.Listen(); err != nil {
		slog.Error(fmt.Sprintf("Start programm: %s", err))
		os.Exit(1)
	}
}
