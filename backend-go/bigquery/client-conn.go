package bigquery

import (
	"cloud.google.com/go/bigquery"
	storage "cloud.google.com/go/storage"
	"context"
	"github.com/github-user-behavior-analysis/backend-go/logs"
	"google.golang.org/api/option"
)

type Client struct {
	*bigquery.Client
}

var ConnClient *Client

func SetUpClient(ctx context.Context, projectID string, jsonPath string)(*bigquery.Client, error)  {

	client, err := storage.NewClient(ctx, option.WithCredentialsFile(jsonPath))
	if err != nil {
		logs.PrintLogger().Error(err)
		return  nil ,err
	}

	return client, err
}

func (client *Client) ReadQuery(sqlQuery string) (*bigquery.RowIterator, error) {

	query := client.Query(sqlQuery)

	return query.Read(context.Background())
}

func (client *Client) SaveResult(iter *bigquery.RowIterator)  {

}