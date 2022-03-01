package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/nextzeus/kratos-shop/app/cart/service/internal/conf"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewMongo, NewCartRepo)

// Data .
type Data struct {
	db  *mongo.Database
	log *log.Helper
}

// NewData .
func NewData(database *mongo.Database, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "cart-service/data"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	d := &Data{
		db:  database,
		log: log,
	}
	cleanup := func() {
		if err := d.db.Client().Disconnect(ctx); err != nil {
			log.Errorf(`closing the data resources %v`, err)
		}
	}
	return d, cleanup, nil
}

// NewMongo .
func NewMongo(conf *conf.Data) *mongo.Database {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.Mongodb.Uri))
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	return client.Database(conf.Mongodb.Database)
}
