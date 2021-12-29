package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/actions-go/toolkit/core"
)

func runMain() {
	secrets := os.Getenv("secrets")
	if secrets == "" {
		core.Error("secrets is not passed")
		return
	}
	branch := os.Getenv("branch")
	if branch == "" {
		core.Error("branch is not passed")
		return
	}
	var secretsMap map[string]string
	if err := json.Unmarshal([]byte(secrets), &secretsMap); err != nil {
		core.Error(fmt.Sprintf("error reading in secrets map %s", err.Error()))
		return
	}
	segmentIOKey := fmt.Sprintf("SEGMENT_IO_KEY_%s", strings.ToUpper(branch))
	core.SetOutput("SEGMENT_IO_KEY", segmentIOKey)
}

func main() {
	runMain()
}
