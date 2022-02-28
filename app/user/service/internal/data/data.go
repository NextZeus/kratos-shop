package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/nextzeus/kratos-shop/app/user/service/internal/conf"
	"github.com/nextzeus/kratos-shop/app/user/service/internal/data/ent"
	"github.com/nextzeus/kratos-shop/app/user/service/internal/data/ent/migrate"
	"time"
	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEntClient, NewRedisCmd, NewUserRepo, NewCardRepo, NewAddressRepo)

// Data .
type Data struct {
	db       *ent.Client
	redisCli redis.Cmdable
}

func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
	log := log.NewHelper(log.With(logger, "module", "server-service/data/ent"))
	client, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		log.Fatalf("failed opening connection to database: %v", err)
	}
	// Run the auto migration tool
	if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

func NewRedisCmd(conf *conf.Data, logger log.Logger) redis.Cmdable {
	log := log.NewHelper(log.With(logger, "module", "server-service/data/ent"))
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  time.Second * 2,
		PoolSize:     10,
	})
	timeout, cancelFuc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFuc()
	err := client.Ping(timeout).Err()
	if err != nil {
		log.Fatalf("redis connect error: %v", err)
	}
	return client
}

// NewData .
func NewData(entClient *ent.Client, redisCmd redis.Cmdable, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "server-service/data"))
	d := &Data{
		db:       entClient,
		redisCli: redisCmd,
	}
	cleanup := func() {
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}
	return d, cleanup, nil
}
