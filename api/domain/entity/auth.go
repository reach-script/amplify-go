package entity

type Auth struct {
	Payload  string `dynamo:"payload,hash"`
	Disabled bool   `dynamo:"disabled"`
	Ttl      int    `dynamo:"ttl"`
}
