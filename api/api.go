package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func API(key string) {

	url := "https://letznav.webhook.office.com/webhookb2/11249524-44bb-4b99-a5b0-d49b04ad3a89@59e2b12b-071c-4970-9f68-fcdfe50a242a/IncomingWebhook/0747b0a0540344e9b52556942a381c04/f4e7e8b1-5640-4a9b-accc-70aa124eb0b8"
	method := "POST"

	lol := `{
		"@type": "MessageCard",
		"@context": "http://schema.org/extensions",
		"correlationId": "0d59eaca1c8c6a78bb87060d2e9486738b21f493",
		"themeColor": "#999",
		"title": "@apty/apty-session-tracker development Deployed",
		"summary": "@apty/apty-session-tracker published by saiumesh-apty",
		"sections": [
		  {
			"activityTitle": "@apty/apty-session-tracker ->> development",
			"activitySubtitle": "Published by saiumesh-apty",
			"activityImage": "https://avatars.githubusercontent.com/u/35692863?v=4",
			"facts": [
			  {
				"name": "Published By:",
				"value": "%key%"
			  },
			  {
				"name": "Branch",
				"value": "development-0d59eaca1c8c6a78bb87060d2e9486738b21f493"
			  }
			]
		  }
		],
		"potentialAction": [
		  {
			"@type": "OpenUri",
			"name": "Repository",
			"targets": [
			  {
				"os": "default",
				"uri": "https://github.com/aptyInc/apty-session-tracker"
			  }
			]
		  },
		  {
			"@type": "OpenUri",
			"name": "Compare",
			"targets": [
			  {
				"os": "default",
				"uri": "https://github.com/aptyInc/apty-session-tracker/compare/07807e5c3e84...0d59eaca1c8c"
			  }
			]
		  }
		],
		"text": "**Changelog:**\n\n+ renamed folder"
	  }`

	lol = strings.Replace(lol, "%key%", key, 1)

	payload := strings.NewReader(lol)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
