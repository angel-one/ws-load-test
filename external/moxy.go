package external

import (
	"context"
	"github.com/angel-one/go-example-project/external/processor"
	"github.com/angel-one/go-example-project/utils/httpclient"
	"time"
)

// GetMoxy is used to get the response from the moxy service
func GetMoxy(ctx context.Context, url string) (string, error) {
	response, err := httpclient.GETWithTimeout(url, nil, 5*time.Second)
	if err != nil {
		return "", err
	}
	return processor.ProcessMoxyResponse(ctx, response)
}
