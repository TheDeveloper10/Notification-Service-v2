package util

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var Logger = zerolog.New(os.Stdout).With().Int64("time", time.Now().Unix()).Logger()
