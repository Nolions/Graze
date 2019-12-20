package models

import (
	"cloud.google.com/go/datastore"
	"github.com/gofrs/uuid"
	"google.golang.org/api/iterator"
	"graze/errors"
	"time"
)

type Incident struct {
	Uid      string    `json:"uid"`
	Title    string    `json:"title" validate:"required"`
	Describe string    `json:"describe" validate:"required"`
	Deadline string    `json:"deadline" validate:"datetime"`
	CrateAt  time.Time `json:"crate_at"`
}

func (i *Incident) TableName() string {
	return "Incident"
}

// 新增事件
func (d *Datastore) NewIncident(title, describe, deadline string) (bool, errors.Errors) {
	i := new(Incident)
	i.Uid = uuid.Must(uuid.NewV4()).String()
	i.CrateAt = time.Now()
	i.Title = title
	i.Describe = describe
	i.Deadline = deadline

	k := d.setDatastroeKey(i.Uid, new(Incident).TableName())
	_, err := d.Conn.Put(d.Ctx, k, i)
	if err != nil {
		e := errors.InsertErrors{}
		e.Error()
		return false, &e
	}

	return true, nil
}

// 取得所有事件
func (d *Datastore) AllIncident() []Incident {
	query := datastore.NewQuery(new(Incident).TableName()).Order("-CrateAt")
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
func (d *Datastore) DeleteIncident(uid string) (bool, errors.Errors) {
	k := d.setDatastroeKey(uid, new(Incident).TableName())
	err := d.Conn.Delete(d.Ctx, k)
	if err != nil {
		e := errors.InsertErrors{}
		e.Error()
		return false, &e
	}

	return true, nil
}

func (d *Datastore) MultiDeleteIncident(uids []string) bool {
	var keys []*datastore.Key
	for _, uid := range uids {
		k := d.setDatastroeKey(uid, new(Incident).TableName())
		keys = append(keys, k)
	}
	err := d.Conn.DeleteMulti(d.Ctx, keys)
	if err != nil {
		// TODO Handle error.

		return false
	}

	return true
}

// 編輯事件
func (d *Datastore) EditIncident(uid, title, describe, deadline string) (bool, errors.Errors) {
	k := d.setDatastroeKey(uid, new(Incident).TableName())

	i := new(Incident)
	err := d.Conn.Get(d.Ctx, k, i)
	if err != nil {
		e := errors.ModelNoFoundError{}
		e.Error()
		return false, &e
	}
	i.Title = title
	i.Describe = describe
	i.Deadline = deadline

	if _, err := d.Conn.Put(d.Ctx, k, i); err != nil {
		e := errors.UpdateErrors{}
		e.Error()
		return false, &e
	}

	return true, nil
}

func (i Incident) FieldTrans() errors.ModelFieldTran {
	return errors.ModelFieldTran{
		"Title":    "事件名稱",
		"Describe": "事件描述",
		"Deadline": "事件期限",
	}
}
