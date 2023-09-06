package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Result struct {
	HtmlUrl string `json:"html_url"`
	GitUrl  string `json:"git_url"`
	SshUrl  string `json:"ssh_url"`
	Owner   struct {
		Login string `json:"login"`
	} `json:"owner"`
}

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
	req.Header.Add("Authorization", "Bearer "+GetEnvVar("GITHUB_TOKEN"))

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
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+GetEnvVar("GITHUB_TOKEN"))

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

func UserUnauth() {
	fmt.Println("You are not authenticated. \n" +
		"Please visit https://github.com/settings/tokens?type=beta to generate a new token.\n" +
		"Then run the command 'export GITHUB_TOKEN=<GithubToken>' to add the token.")
}
