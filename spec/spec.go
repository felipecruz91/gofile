package spec

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Gofile defines the Gofile YAML representation
type Gofile struct {
	Repo    string `yaml:"repo"`
	Version string `yaml:"version"`
	Path    string `yaml:"path"`
}

func NewGofile(filename string) *Gofile {
	var goFile Gofile

	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	if err := yaml.Unmarshal(b, &goFile); err != nil {
		log.Fatal(err)
	}

	valid, msg := goFile.IsValid()
	if !valid {
		log.Fatalf("invalid Gofile: %s\n", msg)
	}

	return &goFile
}

func (gf *Gofile) IsValid() (bool, string) {
	if gf.Repo == "" {
		return false, "repo field is required"
	}
	if gf.Path == "" {
		return false, "path field is required"
	}
	return true, ""
}
