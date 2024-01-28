package db

import (
	"context"
	"github.com/patrickmn/go-cache"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MongoConnection struct {
	DatabaseName string
	Uri          string
	AppName      string

	Client  *mongo.Client
	DB      *mongo.Database
	Session *cache.Cache
	//Timeout      int
	//MaxPool      int
}

func (c *MongoConnection) Connect() {
	to := 60 * time.Second
	clientOpts := options.ClientOptions{ConnectTimeout: &to}
	clientOpts.SetMaxPoolSize(uint64(100)).ApplyURI(c.Uri).SetAppName(c.AppName)

	log.Printf("trying to connect mongodb with uri: %s", c.Uri)

	// connect mongo db
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	client, err := mongo.Connect(ctx, &clientOpts)

	if err != nil {
		log.Fatalf("error while connecting to db: %v", err)
	}
	cancel()

	//client, err := mongo.NewClient(&clientOpts)
	//if err != nil {
	//	log.Fatal(err)
	//}

	// give 10s timeout to connect to database
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	//if err := client.Connect(ctx); err != nil {
	//	log.Fatalf("error while connecting to db: %v", err)
	//}
	//cancel()

	// ping db
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*3)
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	cancel()

	c.Client = client
	log.Println("Connected to mongoDB")

	// save to db
	c.DB = client.Database(c.DatabaseName)

	// set session
	c.Session = cache.New(3*time.Minute, 0)

	return
}

func (c *MongoConnection) Database() *mongo.Database {
	return c.DB
}

func (c *MongoConnection) StartSession(id string, opts ...*options.SessionOptions) error {
	// start session
	ss, err := c.Client.StartSession(opts...)

	if err != nil {
		return err
	}

	// add to in memory cache
	return c.Session.Add(id, ss, 3*time.Minute)
}

func (c *MongoConnection) GetSession(id any) mongo.Session {
	if id == nil {
		return nil
	}

	// get session
	ss, ok := c.Session.Get(id.(string))

	if !ok {
		return nil
	}

	return ss.(mongo.Session)
}

func (c *MongoConnection) EndSession(ctx context.Context, id string) {
	if ss := c.GetSession(id); ss != nil {
		ss.EndSession(ctx)
	}
	c.Session.Delete(id)
}

func SoftDelete() bson.M {
	return bson.M{
		"$set": bson.M{
			"deleted_at": time.Now(),
		},
	}
}
