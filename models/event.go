package models

import (
	"cloud.google.com/go/datastore"
	"context"
	"github.com/gofrs/uuid"
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

func (e *Event)Store() bool {
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
