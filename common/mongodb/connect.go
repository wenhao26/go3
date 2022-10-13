package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Mgo struct {
	DB *mongo.Database
}

func ConnectDB(uri, name string, timeout time.Duration, poolSize uint64) (*Mgo, error) {
	var mgo Mgo

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	opt := options.Client().ApplyURI(uri)
	opt.SetMaxPoolSize(poolSize)
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		return &mgo, err
	}
	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		return &mgo, err
	}

	mgo.DB = client.Database(name)
	return &mgo, nil
}

func (mgo *Mgo) SetCollection(name string) *mongo.Collection {
	return mgo.DB.Collection(name)
}
