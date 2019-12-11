package models

import (
	"cloud.google.com/go/datastore"
	"github.com/gofrs/uuid"
	"google.golang.org/api/iterator"
	"graze/pkg"
	"time"
)

const EntityIncident = "Incident"

type Incident struct {
	Uid      string    `json:"uid"`
	Title    string    `json:"title"`
	Describe string    `json:"describe"`
	Deadline time.Time `json:"deadline"`
	CrateAt  time.Time `json:"crate_at"`
}

func New() Incident {
	return Incident{
		Uid:      uuid.Must(uuid.NewV4()).String(),
		Title:    "",
		Describe: "",
		Deadline: time.Time{},
		CrateAt:  time.Now(),
	}
}

// 新增事件
func (e *Incident) Creator() bool {
	d := new(pkg.Datastore)
	d.Client()

	k := datastore.NameKey(EntityIncident, e.Uid, nil)

	_, err := d.Conn.Put(d.Ctx, k, e)
	if err != nil {
		// TODO Handle error.
		return false
	}

	return true
}

// 取得所有事件
func (e *Incident) All() []Incident {
	d := new(pkg.Datastore)
	d.Client()

	query := datastore.NewQuery(EntityIncident)
	it := d.Conn.Run(d.Ctx, query)

	var list []Incident
	for {
		var e Incident
		_, err := it.Next(&e)
		if err == iterator.Done {
			break
		} else if err != nil {
			// TODO Handle error.
		}

		list = append(list, e)
	}
	return list
}

// 刪除事件
func (e *Incident) Delete() bool {
	d := new(pkg.Datastore)
	d.Client()

	k := datastore.NameKey(EntityIncident, e.Uid, nil)
	err := d.Conn.Delete(d.Ctx, k)
	if err != nil {
		// TODO Handle error.
		return false
	}

	return true
}

func (e *Incident) Edit() bool {
	d := new(pkg.Datastore)
	d.Client()

	k := datastore.NameKey(EntityIncident, e.Uid, nil)
	event := new(Incident)
	d.Conn.Get(d.Ctx, k, event)

	if err := d.Conn.Get(d.Ctx, k, e); err != nil {
		// TODO Handle error.
		return false
	}

	event.Title = e.Title
	event.Describe = e.Describe
	event.Deadline = e.Deadline

	if _, err := d.Conn.Put(d.Ctx, k, event); err != nil {
		// TODO Handle error.
		return false
	}

	return true
}
