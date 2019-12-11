package models

import (
	"cloud.google.com/go/datastore"
	"github.com/gofrs/uuid"
	"google.golang.org/api/iterator"
	"time"
)

type Incident struct {
	Uid      string    `json:"uid"`
	Title    string    `json:"title"`
	Describe string    `json:"describe"`
	Deadline time.Time `json:"deadline"`
	CrateAt  time.Time `json:"crate_at"`
}

func (i *Incident) TableName() string {
	return "Incident"
}

func (i *Incident) New() {
	i.Uid = uuid.Must(uuid.NewV4()).String()
	i.CrateAt = time.Now()
}

// 新增事件
func (d *Datastore) NewIncident(i *Incident) bool {
	k := d.setDatastroeKey(i.Uid, new(Incident).TableName())

	_, err := d.Conn.Put(d.Ctx, k, i)
	if err != nil {
		// TODO Handle error.
		return false
	}

	return true
}

// 取得所有事件
func (d *Datastore) AllIncident() []Incident {
	query := datastore.NewQuery(new(Incident).TableName())
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
func (d *Datastore) DeleteIncident(uid string) bool {
	k := d.setDatastroeKey(uid, new(Incident).TableName())
	err := d.Conn.Delete(d.Ctx, k)
	if err != nil {
		// TODO Handle error.
		return false
	}

	return true
}

// 編輯事件
func (d *Datastore) EditIncident(uid, title, describe string, deadline time.Time) bool {
	k := d.setDatastroeKey(uid, new(Incident).TableName())

	i := new(Incident)
	err := d.Conn.Get(d.Ctx, k, i)
	if err != nil {
		// TODO Handle error.
		return false
	}
	i.Title = title
	i.Describe = describe
	i.Deadline = deadline

	if _, err := d.Conn.Put(d.Ctx, k, i); err != nil {
		// TODO Handle error.
		return false
	}

	return true
}
