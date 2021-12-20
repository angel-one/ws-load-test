package processor

import (
	"context"
	"errors"
	"github.com/angel-one/go-utils/log"
	"io/ioutil"
	"net/http"
)

// ProcessMoxyResponse is used to process moxy response
func ProcessMoxyResponse(ctx context.Context, response *http.Response) (string, error) {
	if response.Body == nil {
		return "", errors.New("no response exists")
	}
	body, err := ioutil.ReadAll(response.Body)
	defer func() {
		err = response.Body.Close()
		if err != nil {
			log.Error(ctx).Stack().Err(err).Msg("error closing moxy response body")
		}
	}()
	if err != nil {
		return "", err
	}
	return string(body), err
}
