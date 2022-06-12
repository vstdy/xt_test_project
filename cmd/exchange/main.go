package main

import (
	"os"

	"github.com/rs/zerolog/log"

	"github.com/vstdy/xt_test_project/cmd/exchange/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("Application crashed")
		os.Exit(1)
	}
}
