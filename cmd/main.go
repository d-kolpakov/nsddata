package main

import (
	"github.com/d-kolpakov/nsddata"
	"net/http"
)

func main() {
	clientOptions := nsddata.Options{
		BaseURL: "",
		APIKey:  "DEMO",
		Client:  http.Client{},
		Debug:   true,
	}

	cl := nsddata.NewClient(clientOptions)
	filter := map[string]map[string]interface{}{
		"ca_type": map[string]interface{}{
			nsddata.FilterEQ: "PRIO",
		},
	}
	cl.GetNews(10, 0, filter)
}
