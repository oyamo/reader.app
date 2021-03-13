package _const

import "os"

const (
	Port string = ":8080"
	Host string = "0.0.0.0"
)

const Address = Host + Port

// Global Variables
var (
	MongoUri = os.Getenv("mongo_uri")
)
