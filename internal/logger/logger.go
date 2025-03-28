package logger

import (
	"log"
	"os"
)

var Info = New("/Users/serhiichornenkyi/www/open/textwire/lsp/log.txt")

func New(filename string) *log.Logger {
	const fileMode = 0666

	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, fileMode)
	if err != nil {
		log.Panicf("The filepath %s is missing a file", filename)
	}

	return log.New(logfile, "[textwire lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
