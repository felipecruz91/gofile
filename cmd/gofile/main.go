package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/felipecruz91/gofile"
	"github.com/felipecruz91/gofile/spec"
	"github.com/moby/buildkit/client/llb"
	"github.com/moby/buildkit/frontend/gateway/grpcclient"
	"github.com/moby/buildkit/util/appcontext"
)

var (
	filename string
	graph    bool
)

func main() {
	flag.StringVar(&filename, "f", "Gofile.yaml", "the file to read from")
	flag.BoolVar(&graph, "graph", false, "output a graph and exit")
	flag.Parse()

	if graph {
		goFile := spec.NewGofile(filename)

		state, _, err := gofile.Gofile2LLB(goFile)
		if err != nil {
			log.Fatal(err)
		}
		dt, err := state.Marshal(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		llb.WriteTo(dt, os.Stdout)
		os.Exit(0)
	}

	if err := grpcclient.RunFromEnvironment(appcontext.Context(), gofile.Build); err != nil {
		log.Fatal(err)
	}
}
