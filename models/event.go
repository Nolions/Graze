package models

import (
	"cloud.google.com/go/datastore"
	"github.com/gofrs/uuid"
	"google.golang.org/api/iterator"
	"graze/pkg"
	"log"
	"time"
)

type Event struct {
	Uid      string    `json:"uid"`
	Title    string    `json:"title"`
	Describe string    `json:"describe"`
	Deadline time.Time `json:"deadline"`
	CrateAt  time.Time `json:"crate_at"`
}

func New() Event {
	return Event{
		Uid:      uuid.Must(uuid.NewV4()).String(),
		Title:    "",
		Describe: "",
		Deadline: time.Time{},
		CrateAt:  time.Now(),
	}
}

// 新增事件
func (e *Event) Creator() bool {
	d := new(pkg.Datastore)
	d.Client()

	key := datastore.IncompleteKey("Event", nil)

	_, err := d.Conn.Put(d.Ctx, key, e)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

// 取得所有事件
func (e *Event) All() []Event {
	d := new(pkg.Datastore)
	d.Client()

	query := datastore.NewQuery("Event")
	it := d.Conn.Run(d.Ctx, query)

	var list []Event
	for {
		var e Event
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
func (e *Event) Delete() bool {
	d := new(pkg.Datastore)
	d.Client()

	query := datastore.NewQuery("Event").Filter("Uid = ", e.Uid)
	it := d.Conn.Run(d.Ctx, query)
	for {
		var e Event
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
		err = d.Conn.Delete(d.Ctx, k)
		if err != nil {
			// TODO
			log.Fatal(err)
			return false
		}
	}

	return true
}
