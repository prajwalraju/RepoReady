package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/prajwalraju/RepoReady/dto"
)

// The above type represents a result object containing URLs and owner information.
// @property {string} HtmlUrl - The `HtmlUrl` property is a string that represents the URL of the
// repository's HTML page on GitHub. This URL can be used to view the repository in a web browser.
// @property {string} GitUrl - The `GitUrl` property is a string that represents the clone URL of a Git
// repository. This URL can be used to clone the repository using Git commands.
// @property {string} SshUrl - The `SshUrl` property is a string that represents the SSH URL for
// cloning the repository.
// @property Owner - The `Owner` property is a nested struct that contains information about the owner
// of a repository. It has a single property called `Login`, which represents the username of the
// owner.
type Result struct {
	HtmlUrl string `json:"html_url"`
	GitUrl  string `json:"clone_url"`
	SshUrl  string `json:"ssh_url"`
	Owner   struct {
		Login string `json:"login"`
	} `json:"owner"`
}

// The function `CreateRemoteRepo` creates a remote repository on GitHub using the provided input and
// returns the result or an error.
func CreateRemoteRepo(repoInput dto.RepoInput) (Result, error) {

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

	if res.StatusCode == 201 {

		err = json.NewDecoder(res.Body).Decode(&result)
		if err != nil {
			fmt.Println(err)
			return result, err
		}
		fmt.Println("Remote GitHub repo is created successfully, you can vist by clicking ", result.HtmlUrl)
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

	fmt.Println("Github CreateRemoteRepo returned :", res.StatusCode)

	return result, errors.New(errorResponse.ErrorsList[0].Message)
}

// The function `UpdateTopics` updates the topics of a GitHub repository using the GitHub API.
func UpdateTopics(repoInput dto.RepoInput) error {

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

	if res.StatusCode != 200 {
		if res.StatusCode == 401 {
			UserUnauth()
			return errors.New("user not authenticated")
		}
		fmt.Println("Github UpdateTopics returned :", res.StatusCode)
		return errors.New("error in updating topics")
	}

	return nil
}

// The `License` type represents a software license with properties such as key, name, SPDX ID, URL,
// and node ID.
// @property {string} Key - The `Key` property represents the unique identifier for the license.
// @property {string} Name - The "Name" property represents the name of the license.
// @property {string} SpdxId - SpdxId is a unique identifier for the license, which follows the SPDX
// (Software Package Data Exchange) specification. It is used to uniquely identify and reference
// licenses in a standardized way.
// @property {string} Url - The `Url` property in the `License` struct represents the URL or web
// address associated with the license. It typically points to a webpage or document that provides more
// information about the license terms and conditions.
// @property {string} NodeId - The `NodeId` property is a unique identifier for the license. It is used
// to identify the license within a system or database.
type License struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	SpdxId string `json:"spdx_id"`
	Url    string `json:"url"`
	NodeId string `json:"node_id"`
}

// The function `GetLicenses` makes a GET request to the GitHub API to retrieve a list of licenses and
// returns the result as a slice of `License` structs or an error.
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

	fmt.Println("Github GetLicenses returned :", res.StatusCode)
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

	fmt.Println("Github GetLicenses returned :", res.StatusCode)
	return result, errors.New("error occured in getting licenses")
}

func UserUnauth() {
	fmt.Println("You are not authenticated.\n Please run the command 'repoready config' to config Tokens.")
}

func setHeaders(req *http.Request) (*http.Request, error) {

	config, err := GetEnvVar()
	if err != nil {
		fmt.Println("Error in reading config file:", err)
		return req, err
	}

	if config.Github.Token == "" {
		UserUnauth()
		return req, errors.New("user not authenticated")
	}

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + config.Github.Token,
	}

	for headerName, headerValue := range headers {
		req.Header.Add(headerName, headerValue)
	}

	return req, nil
}
