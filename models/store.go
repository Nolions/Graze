package models

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

func (*Datastore) NewClient() *Datastore {
	d := new(Datastore)
	d.Client()
	return d
}

// 建立Google Datastroe 連線
func (d *Datastore) Client() {
	d.Ctx = context.Background()
	if config.APIConf.Debug {
		o := []option.ClientOption{
			option.WithEndpoint("localhost:8081"),
			option.WithoutAuthentication(),
			option.WithGRPCDialOption(grpc.WithInsecure()),
			option.WithGRPCConnectionPool(50),
		}

		d.Conn, d.err = datastore.NewClient(d.Ctx, config.APIConf.DatastoreProjectId, o...)
	} else {
		d.Conn, d.err = datastore.NewClient(d.Ctx, config.APIConf.DatastoreProjectId)
	}

	if d.err != nil {
		log.Println(d.err)
	}
}

func (d *Datastore) setDatastroeKey(v, entityName string) *datastore.Key {
	return datastore.NameKey(entityName, v, nil)
}
