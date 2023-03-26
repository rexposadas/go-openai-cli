package require

import (
	"log"
	"os"
)

func File(name string) string {
	if len(name) == 0 {
		log.Fatalln("a file is required to run this command, use the --file flag.")
	}

	if _, err := os.Stat(name); err != nil {
		log.Fatalf("file %q does not exist", name)
	}

	return name
}
