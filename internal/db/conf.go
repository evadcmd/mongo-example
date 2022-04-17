package db

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
)

const NameDB = "test"
const NameCollectionCamera = "camera"

const (
	dbURI = "mongodb://localhost:27017"

	username = "root"
	password = "root"

	connectionTimeout = 30 * time.Second
	maxPoolSize       = 20
	minPoolSize       = 5

	replicaSet = "testreplica"
)

var credential = options.Credential{
	Username: username,
	Password: password,
}

var dbOpt = options.Client().
	SetAuth(credential).
	SetConnectTimeout(connectionTimeout).
	SetMaxPoolSize(maxPoolSize).
	SetMinPoolSize(minPoolSize).
	// SetReplicaSet
	ApplyURI(dbURI)
