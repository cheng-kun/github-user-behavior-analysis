package bigquery

import (
	"context"
	"fmt"
	"testing"
)

func TestSetUpClient(t *testing.T) {
	client, e := SetUpClient(context.Background(), "analytical-camp-223000")
	fmt.Println(client)
	fmt.Println(e)

	q := `SELECT language, COUNT(*) amount
FROM [ghtorrent-bq:ght.projects]
WHERE LANGUAGE is not null
AND (created_at >='2009-02-01 00:00:00 UTC' and created_at <'2009-03-01 00:00:00 UTC')
GROUP BY 1
ORDER BY 2 DESC
LIMIT 100`

	iterator, i := client.Query(q).Read(context.Background())
	fmt.Println(iterator)

	fmt.Println(i)


}