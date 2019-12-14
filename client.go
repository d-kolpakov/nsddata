package nsddata

import (
	"fmt"
	"net/http"
)

type NSDDataAPIClient struct {
	O Options
}

// https://nsddata.ru/api/get/news?limit=10&apikey=DEMO
type Options struct {
	BaseURL string
	APIKey  string
	Client  http.Client
	Debug   bool
}

const BaseURL = "https://nsddata.ru/api/"

func NewClient(options Options) *NSDDataAPIClient {
	if options.BaseURL == "" {
		options.BaseURL = BaseURL
	}
	return &NSDDataAPIClient{
		O: options,
	}
}

func (n *NSDDataAPIClient) debug(s string) {
	if n.O.Debug {
		fmt.Println(s)
	}
}
