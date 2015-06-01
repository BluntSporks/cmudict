package cmudict

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// LoadDict loads the CMU dictionary file and returns it as a map.
func LoadDict(file string) map[string]string {
	// Open file.
	hdl, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer hdl.Close()

	// Scan file line by line.
	dict := make(map[string]string)
	scanner := bufio.NewScanner(hdl)
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		// Skip comments
		if line[:3] == ";;;" {
			continue
		}
		flds := strings.Split(line, "  ")
		word := flds[0]
		pron := flds[1]
		dict[word] = pron
	}
	return dict
}
