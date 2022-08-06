package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/lucasafonsokremer/github-go-cli/cmd/utils/types"
)

// githubHost declare the base url
const githubHost = "https://api.github.com"

// Client is interface of service
type Client interface {
	GetUser() (types.UserInfo, error)
	GetFollowers() (types.Followers, error)
}

// config encapsulates url information
type config struct {
	BaseURL string
	URL	string
}

// CreateClient creates the client for make the request
func CreateClient(path string) Client {
	return config{
		URL: githubHost + path,
	}
}

// GetUser fetches user information from github.com
func (config config) GetUser() (types.UserInfo, error) {
	resp, err := makeRequest(http.MethodGet, config.URL, nil)
	if err != nil {
		return types.UserInfo{}, err
	}

	var userInfo types.UserInfo

	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return types.UserInfo{}, fmt.Errorf("error decoding response %v", err)
	}

	return userInfo, nil
}

// GetFollowers fetches followers list of user
func (config config) GetFollowers() (types.Followers, error) {
	resp, err := makeRequest(http.MethodGet, config.URL, nil)

	if err != nil {
		return types.Followers{}, err
	}

	var followers types.Followers

	if err := json.NewDecoder(resp.Body).Decode(&followers); err != nil {
		return types.Followers{}, err
	}

	return followers, nil
}

func makeRequest(method string, URL string, body io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(method, URL, body)

	if err != nil {
		return nil, fmt.Errorf("error creating new HTTP request %v", err)
	}

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, fmt.Errorf("error getting response from service %v", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s", resp.Status)
	}

	return resp, nil
}
