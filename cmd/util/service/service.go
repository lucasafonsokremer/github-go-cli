package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/lucasafonsokremer/github-go-cli/cmd/util/types"
)

// githubHost declare the base url
const githubHost = "https://api.github.com"

// Client is interface of service
type Client interface {
	GetUser() (types.User, error)
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
func (config config) GetUser() (types.User, error) {
	resp, err := makeRequest(http.MethodGet, config.URL, nil)

	if err != nil {
		return types.User{}, err
	}

	var user types.User

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		fmt.Println("aqui")
		return types.User{}, err
	}

	fmt.Println(user)
	return user, nil
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

// makeRequest create the request and deal with some possible errors
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
