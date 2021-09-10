package main

import (
	"fmt"
	"log"

	"github.com/buildkite/go-buildkite/v2/buildkite"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	apiToken = kingpin.Flag("token", "API token").Required().OverrideDefaultFromEnvar("BUILDKITE_API_TOKEN").String()
	branch   = kingpin.Flag("branch", "A branch name").Required().String()
	debug    = kingpin.Flag("debug", "Enable debugging").Bool()
)

func main() {
	kingpin.Parse()

	config, err := buildkite.NewTokenConfig(*apiToken, *debug)
	if err != nil {
		log.Fatalf("client config failed: %s", err)
	}
	client := buildkite.NewClient(config.Client())

	opt := buildkite.BuildsListOptions{Branch: *branch, State: []string{"passed"}}
	builds, _, err := client.Builds.List(&opt)
	if err != nil {
		log.Fatalf("fetch builds failed: %s", err)
	}

	for _, build := range builds {
		fmt.Printf("%v %v %v\n", build.CreatedAt.Format("2006-01-02 15:04:05"), build.FinishedAt.Sub(build.StartedAt.Time), *build.WebURL)
	}
}
