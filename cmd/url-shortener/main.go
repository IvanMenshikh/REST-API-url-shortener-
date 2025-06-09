package main

import (
	"fmt"
	"url-shortener/internal/config"
)

func main() {
	// TODO: init config: cleanenv
	cfg := config.MustLoad()

	fmt.Printf("%+v\n", cfg)

	// TODO: init Logger: slog

	// TODO: init storage: sqlite

	// TODO: init router: chi, "chi render"

	// TODO: init server
}
