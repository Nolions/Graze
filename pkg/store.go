package pkg

import (
	"cloud.google.com/go/datastore"
	"context"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"graze/config"
	"log"
)

type Store interface {
	Client()
}

type Datastore struct {
	Conn *datastore.Client
	Ctx  context.Context
	err  error
}

func (d *Datastore) Client() {
	d.Ctx = context.Background()
	if config.Conf.Debug {
		o := []option.ClientOption{
			option.WithEndpoint("localhost:8081"),
			option.WithoutAuthentication(),
			option.WithGRPCDialOption(grpc.WithInsecure()),
			option.WithGRPCConnectionPool(50),
		}

		d.Conn, d.err = datastore.NewClient(d.Ctx, config.Conf.DatastoreProjectId, o...)
	} else {
		d.Conn, d.err = datastore.NewClient(d.Ctx, config.Conf.DatastoreProjectId)
	}

	if d.err != nil {
		log.Println(d.err)
	}
}
