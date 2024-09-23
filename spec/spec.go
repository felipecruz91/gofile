package spec

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Gofile defines the Gofile YAML representation
type Gofile struct {
	GitRepo string `yaml:"repo"`
	GitRef  string `yaml:"ref"`
	Path    string `yaml:"path"`
	Scratch bool   `yaml:"scratch"`
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
	if gf.GitRepo == "" {
		return false, "repo field is required"
	}
	if gf.Path == "" {
		return false, "path field is required"
	}
	return true, ""
}
