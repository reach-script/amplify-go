package main

import (
	"backend/cmd/api"

	_ "github.com/lib/pq"
)

func main() {
	api.Execute()
}
