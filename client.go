package secretdb

import (
	"github.com/mikerybka/util"
)

func NewClient(host, user, pass string) *Client {
	return &Client{
		Host:     host,
		Username: user,
		Password: pass,
	}
}

type Client struct {
	Host     string
	Username string
	Password string
}

func (c *Client) Email() (string, error) {
	path := "data/secrets/email"
	return util.ReadRemoteFile(c.Username, c.Password, c.Host, 22, path)
}
