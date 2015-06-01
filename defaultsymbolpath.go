package cmudict

import (
	"log"
	"os"
	"path"
)

// DefaultSymbolPath gets the CMU symbols file location from the CMUDICT_DATA environment variable.
func DefaultSymbolPath() string {
	dir := os.Getenv("CMUDICT_DATA")
	if dir == "" {
		log.Fatal("Set CMUDICT_DATA variable to directory of symbols file")
	}
	return path.Join(dir, "cmudict.0.7a.symbols")
}
