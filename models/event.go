package models

import (
	"cloud.google.com/go/datastore"
	"context"
	"github.com/gofrs/uuid"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"log"
	"time"
)

type Event struct {
	Uid      string    `json:"uid"`
	Title    string    `json:"title"`
	Describe string    `json:"describe"`
	Deadline time.Time `json:"deadline"`
	CrateAt time.Time `json:"crate_at"`
}

func New() Event {
	return Event{
		Uid:      uuid.Must(uuid.NewV4()).String(),
		Title:    "",
		Describe: "",
		Deadline: time.Time{},
		CrateAt:time.Now(),
	}
}

const (
	projectId    = "web-todo-list"
	namespace    = "Entity"
	kind         = "Entity"
	emulatorHost = "localhost:8081"
)

// 新增事件
func (e *Event)Creator() bool {
	o := []option.ClientOption{
		option.WithEndpoint(emulatorHost),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithInsecure()),
		option.WithGRPCConnectionPool(50),
	}

	ctx := context.Background()

	dsClient, err := datastore.NewClient(ctx, projectId, o...)
	if err != nil {
		log.Println(err)
		return false
	}

	key := datastore.IncompleteKey("Event", nil)

	_, err = dsClient.Put(ctx, key, e)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

// 取得所有事件
func (e *Event) All() []Event {
	o := []option.ClientOption{
		option.WithEndpoint(emulatorHost),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithInsecure()),
		option.WithGRPCConnectionPool(50),
	}

	ctx := context.Background()

	dsClient, err := datastore.NewClient(ctx, projectId, o...)
	if err != nil {
		// TODO
		return []Event{}
	}

	query := datastore.NewQuery("Event")
	it := dsClient.Run(ctx, query)

	var list []Event
	for {
		var e  Event
		_, err := it.Next(&e)
		if err == iterator.Done {
			break
		} else if err != nil {
			// TODO
		}

		list = append(list, e)
	}
	return list
}

// 刪除事件
func (e *Event) Delete()bool  {
	o := []option.ClientOption{
		option.WithEndpoint(emulatorHost),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithInsecure()),
		option.WithGRPCConnectionPool(50),
	}

	ctx := context.Background()

	dsClient, err := datastore.NewClient(ctx, projectId, o...)
	if err != nil {
		// TODO
		log.Fatal(err)
		return false
	}
	log.Println(&e.Uid)

	query := datastore.NewQuery("Event").Filter("Uid = ", e.Uid)
	it := dsClient.Run(ctx, query)
	for {
		var e  Event
		k, err := it.Next(&e)
		if err == iterator.Done {
			log.Println("aa")
			break
		} else if err != nil {
			// TODO
			log.Fatal(err)
			return false
		}

		log.Println(e)
		log.Println(k)
		err = dsClient.Delete(ctx, k)
		if err != nil {
			// TODO
			log.Fatal(err)
			return false
		}
	}

	return true
}
