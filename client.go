package secretdb

import (
	"fmt"
	"io"
	"net/http"

	"github.com/mikerybka/constants"
)

func NewClient() *Client {
	return &Client{}
}

type Client struct {
}

func (c *Client) Email() (string, error) {
	res, err := http.Get(fmt.Sprintf("http://%s:2222/secrets/email", constants.BackendIP))
	if err != nil {
		return "", err
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
