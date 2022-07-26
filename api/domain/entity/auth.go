package entity

type Auth struct {
	Key1     string `dynamo:"key1,hash"`
	Key2     string `dynamo:"key2,range"`
	Payload  string `dynamo:"payload"`
	Disabled bool   `dynamo:"disabled"`
	Ttl      int    `dynamo:"ttl"`
}
