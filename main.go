package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/actions-go/toolkit/core"
)

func runMain() {
	secrets := os.Getenv("SECRETS")
	if secrets == "" {
		core.Error("secrets is not passed")
		return
	}
	branch := os.Getenv("BRANCH_NAME")
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
	segmentIOValue := secretsMap[segmentIOKey]
	if segmentIOValue == "" {
		segmentIOKey_DEV := fmt.Sprintf("SEGMENT_IO_KEY_DEVELOPMENT")
		segmentIOValue = secretsMap[segmentIOKey_DEV] 
	}
	splitIOKey := fmt.Sprintf("SPLIT_IO_JS_%s", strings.ToUpper(branch))
	splitIOValue := secretsMap[splitIOKey]
	if segmentIOValue == "" {
		splitIOKey_DEV := fmt.Sprintf("SPLIT_IO_JS_DEVELOPMENT")
		splitIOValue = secretsMap[splitIOKey_DEV] 
	}
	core.SetOutput("FEATURE_FLAG_API_KEY", splitIOValue)
	core.SetOutput("SEGMENT_IO_KEY", segmentIOValue)
}

func main() {
	runMain()
}
