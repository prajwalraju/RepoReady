package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Result struct {
	HTMLURL string `json:"html_url"`
}

type RepoInput struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Private     bool     `json:"private"`
	HasWiki     bool     `json:"has_wiki"`
	Topics      []string `json:"topics"`
	Owner       string   `json:"owner"`
}

var gitHubToken string = "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"

func CreateRemoteRepo(repoInput RepoInput) (Result, error) {

	var result Result

	url := "https://api.github.com/user/repos"

	payloadJSON, err := json.Marshal(repoInput)
	if err != nil {
		fmt.Println("Error serializing JSON:", err)
		return result, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payloadJSON))

	if err != nil {
		fmt.Println(err)
		return result, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+gitHubToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	defer res.Body.Close()

	fmt.Println("Github CreateRemoteRepo returned :", res.StatusCode)

	if res.StatusCode == 201 {

		err = json.NewDecoder(res.Body).Decode(&result)
		if err != nil {
			fmt.Println(err)
			return result, err
		}
		fmt.Println("Repository url :", result.HTMLURL)
		return result, nil
	}

	// Error API handling
	type ErrorResponse struct {
		ErrorsList []struct {
			Message string `json:"message"`
		} `json:"errors"`
	}

	var errorResponse ErrorResponse

	err = json.NewDecoder(res.Body).Decode(&errorResponse)
	if err != nil {
		fmt.Println(err)
		return result, err
	}

	fmt.Println("Error response :", errorResponse.ErrorsList[0].Message)

	return result, errors.New(errorResponse.ErrorsList[0].Message)
}

func UpdateTopics(repoInput RepoInput) error {

	url := "https://api.github.com/repos/prajwalraju/" + repoInput.Name + "/topics"

	payload := map[string]interface{}{
		"names": repoInput.Topics,
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error serializing JSON:", err)
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(payloadJSON))

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+gitHubToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	fmt.Println("Github UpdateTopics returned :", res.StatusCode)

	if res.StatusCode != 200 {
		return errors.New("Error in updating topics")
	}

	return nil
}
