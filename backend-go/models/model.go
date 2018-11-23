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

type Ranking struct {
	RepoNum int64 `json:"repo_num"`
	TimeStamp string`json:"timestamp"`
	N1lang string	`json:"n1lang"`
	N1num int64		`json:"n1num"`
	N2lang string	`json:"n2lang"`
	N2num int64		`json:"n2num"`
	N3lang string	`json:"n3lang"`
	N3num int64		`json:"n3num"`
	N4lang string	`json:"n4lang"`
	N4num int64		`json:"n4num"`
	N5lang string	`json:"n5lang"`
	N5num int64		`json:"n5num"`
	N6lang string	`json:"n6lang"`
	N6num int64		`json:"n6num"`
	N7lang string	`json:"n7lang"`
	N7num int64		`json:"n7num"`
	N8lang string	`json:"n8lang"`
	N8num int64		`json:"n8num"`
	N9lang string	`json:"n9lang"`
	N9num int64		`json:"n9num"`
	N10lang string	`json:"n10lang"`
	N10num int64	`json:"n10num"`
}

type RankingsJSON struct {
	Rankings []Ranking `json:"rankings"`
}

type LangaugeRank struct {
	Amount int64 `json:"amount"`
	TimeStamp string `json:"time_stamp"`
	Rank int64 `json:"rank"`
}

type UserFollower struct {
	User string `json:"user"`
	Follower int64 `json:"follower"`
	Rank int64 `json:"rank"`
} 