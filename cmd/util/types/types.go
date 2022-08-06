package types

// userInfo encapsulates user information
type userInfo struct {
	Name	    string `json:"name"`
	Location    string `json:"location"`
}

// User represents list of user information
type User []userInfo

// Follower encapsulates follower meta
type follower struct {
	Name    string `json:"login"`
	HTMLURL string `json:"html_url"`
}

// Followers represents list of followers
type Followers []follower
