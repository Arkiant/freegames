package freegames

import (
	"embed"
	"io/fs"
)

//go:embed .env
var environmentFile embed.FS

func ReadEnvironmentFile() (fs.File, error) {
	return environmentFile.Open(".env")
}
