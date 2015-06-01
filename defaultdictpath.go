package cmudict

import (
	"log"
	"os"
	"path"
)

// DefaultDictPath gets the CMU dictionary file location from the CMUDICT_DATA environment variable.
func DefaultDictPath() string {
	dir := os.Getenv("CMUDICT_DATA")
	if dir == "" {
		log.Fatal("Set CMUDICT_DATA variable to directory of dictionary file")
	}
	return path.Join(dir, "cmudict.0.7a")
}
