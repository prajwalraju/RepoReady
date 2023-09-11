package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// RepoInput is the Object that holds all the info required to create a remote repo
type RepoInput struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Private     bool     `json:"private"`
	Topics      []string `json:"topics"`
	Owner       string
	HtmlUrl     string
	GitUrl      string
	SshUrl      string
}

// Result is the Object that holds the response from the Github API to create a remote repo
type Result struct {
	HtmlUrl string `json:"html_url"`
	GitUrl  string `json:"clone_url"`
	SshUrl  string `json:"ssh_url"`
	Owner   struct {
		Login string `json:"login"`
	} `json:"owner"`
}

// CreateRemoteRepo creates a remote repo on Github
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

	setHeaders(req)

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
		fmt.Println("Visit the remote repo by clicking this ", result.HtmlUrl)
		return result, nil
	}

	// Error API handling

	if res.StatusCode == 401 {
		UserUnauth()
		return result, errors.New("user not authenticated")
	}

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

// UpdateTopics updates the topics of a remote repo on Github
func UpdateTopics(repoInput RepoInput) error {

	url := "https://api.github.com/repos/" + repoInput.Owner + "/" + repoInput.Name + "/topics"

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
	setHeaders(req)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	fmt.Println("Github UpdateTopics returned :", res.StatusCode)

	if res.StatusCode != 200 {
		if res.StatusCode == 401 {
			UserUnauth()
			return errors.New("user not authenticated")
		}
		return errors.New("error in updating topics")
	}

	return nil
}

type License struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	SpdxId string `json:"spdx_id"`
	Url    string `json:"url"`
	NodeId string `json:"node_id"`
}

// GetLicenses gets the list of licenses from Github
func GetLicenses() ([]License, error) {
	var result []License

	url := "https://api.github.com/licenses"

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		fmt.Println(err)
		return result, err
	}

	setHeaders(req)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	defer res.Body.Close()

	fmt.Println("Github GetLicenses returned :", res.StatusCode)

	if res.StatusCode == 200 {

		err = json.NewDecoder(res.Body).Decode(&result)
		if err != nil {
			fmt.Println(err)
			return result, err
		}
		return result, nil
	}

	// Error API handling

	if res.StatusCode == 401 {
		UserUnauth()
		return result, errors.New("user not authenticated")
	}

	return result, errors.New("error occured in getting licenses")
}

type LicenseBody struct {
	Body string `json:"body"`
}

// GetLicenses gets the list of licenses from Github
func GetLicenseContent(url string) (LicenseBody, error) {
	var result LicenseBody

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		fmt.Println(err)
		return result, err
	}

	setHeaders(req)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	defer res.Body.Close()

	fmt.Println("Github GetLicenses returned :", res.StatusCode)

	if res.StatusCode == 200 {

		err = json.NewDecoder(res.Body).Decode(&result)
		if err != nil {
			fmt.Println(err)
			return result, err
		}
		return result, nil
	}

	// Error API handling

	if res.StatusCode == 401 {
		UserUnauth()
		return result, errors.New("user not authenticated")
	}

	return result, errors.New("error occured in getting licenses")
}

func UserUnauth() {
	fmt.Println("You are not authenticated. \n" +
		"Please visit https://github.com/settings/tokens?type=beta to generate a new token.\n" +
		"Then run the command 'export GITHUB_TOKEN=<GithubToken>' to add the token.")
}

func setHeaders(req *http.Request) *http.Request {
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + GetEnvVar("GITHUB_TOKEN"),
	}

	for headerName, headerValue := range headers {
		req.Header.Add(headerName, headerValue)
	}

	return req
}
