package models

import (
	"github.com/google/go-github/github"
	"time"
)

type UserInfo struct {
    GithubUser github.User `json:"github_user"`
	StatusCode *int `json:"status_code"`
	FetchTime *time.Time `json:"fetch_time"`
}

type RepositoryInfo struct {
	GithubRepo github.Repository `json:"github_repo"`
	StatusCode *int `json:"status_code"`
	FetchTime *time.Time `json:"fetch_time"`
}
