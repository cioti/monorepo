package approot

import (
	"sync"
	"time"

	"github.com/cioti/monorepo/pkg/logging"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

var (
	infraOnce sync.Once
	infra     InfraFactory
)

type InfraFactory interface {
	Logger() logging.Logger
	Mongo() *mongo.Client
}

type infraFactory struct {
	logger      logging.Logger
	loggerOnce  sync.Once
	mongoClient *mongo.Client
	mongoOnce   sync.Once
}

func (f *infraFactory) Logger() logging.Logger {
	f.loggerOnce.Do(func() {
		f.logger = logging.NewLogger("logger", logging.Debug)
	})

	return f.logger
}

func (f *infraFactory) Mongo() *mongo.Client {
	f.mongoOnce.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		opts := options.Client().ApplyURI("mongodb://root:123456@localhost:27017")
		client, err := mongo.Connect(ctx, opts)
		if err != nil {
			panic(err)
		}
		f.mongoClient = client
	})

	return f.mongoClient
}

func Infra() InfraFactory {
	infraOnce.Do(func() {
		infra = &infraFactory{}
	})

	return infra
}
