package entity

type Auth struct {
	Sub      string `dynamo:"key1,hash"`
	Payload  string `dynamo:"key2,range"`
	Disabled bool   `dynamo:"disabled"`
	Ttl      int    `dynamo:"ttl"`
}
