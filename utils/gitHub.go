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
}

func CreateRemoteRepo(repoInput RepoInput) (Result, error) {

	var result Result

	url := "https://api.github.com/user/repos"
	method := "POST"

	payloadJSON, err := json.Marshal(repoInput)
	if err != nil {
		fmt.Println("Error serializing JSON:", err)
		return result, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payloadJSON))

	if err != nil {
		fmt.Println(err)
		return result, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer <GITHUB_TOKEN>")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	defer res.Body.Close()

	fmt.Println("Github returned :", res.StatusCode)

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
