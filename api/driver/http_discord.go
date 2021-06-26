package driver

import (
	"net/http"
	"time"
)

type HttpDriver struct {
	client *http.Client
}

func NewHttpDriver() *HttpDriver {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 20,
		},
		Timeout: 60 * time.Second,
	}

	return &HttpDriver{client: client}
}

func (d *HttpDriver) GetUserInfo(token string) {

}
