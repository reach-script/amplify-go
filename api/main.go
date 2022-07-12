package main

import (
	"backend/cmd/api"

	_ "github.com/lib/pq"
)

func main1() {
	api.Execute()
}
