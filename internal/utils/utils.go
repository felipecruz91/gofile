package utils

import (
	"path"
	"path/filepath"
	"strings"
)

func ConstructPath(p string) string {
	var upwardPath string

	cleanPath := path.Clean(p)
	if cleanPath == "." {
		upwardPath = "./"
	} else {
		// Add "../" for each level of subdirectory
		levels := strings.Split(cleanPath, string(filepath.Separator))
		upwardPath = strings.Repeat("../", len(levels))
	}

	return upwardPath + filepath.Join("bin", "server")
}
