package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Response struct {
	Candidates []struct {
		Output string `json:"output"`
	} `json:"candidates"`
}

func GenerateText(prompt string, params map[string]string) (string, error) {

	if params["model"] == "" {
		params["model"] = "text-bison-001"
	}
	err := LoadEnvFromFile(".env")
	if err != nil {
		return "", err
	}
	apiKey := os.Getenv("PALM_API_KEY")
	endpoint := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta2/models/%s:generateText?key=%s", params["model"], apiKey)
	fmt.Printf("%s\n", endpoint)

	payload := `{
        "model": "` + params["model"] + `",
        "prompt": {"text": "` + prompt + `"},
    }`
	/*
	       "temperature": 1.0,
	       "candidate_count": 2,
	       "max_output_tokens": 200,
	       "top_p": null,
	       "top_k": null,
	       "safety_setting": null,
	       "stop_sequence": null,
	       "client": null
	   }`
	*/

	jsonPayload := []byte(payload)
	fmt.Printf("%s\n", jsonPayload)

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.StatusCode)
	var result Response
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}
	fmt.Println(result)
	return result.Candidates[0].Output, nil

}
