package models

import (
	"encoding/json"
	"fmt"
	"github.com/fj-onathan/crupi/src/db"
	bolt "go.etcd.io/bbolt"
)

const crudBucket = "CRUDS"

type Crud struct {
	ID   int
	Name string
}

func init() {
	// Create Bucket if not exists.
	db.Local.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(crudBucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
}

func (c *Crud) AddCrud() {
	db.Local.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(crudBucket))
		id, _ := b.NextSequence()
		c.ID = int(id)

		// Marshal user data into bytes.
		buf, err := json.Marshal(c)
		if err != nil {
			return err
		}

		// Persist bytes to users bucket.
		return b.Put(Itob(c.ID), buf)
	})
}

func ListCruds() []string {
	var crudsList []string
	db.Local.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys.
		b := tx.Bucket([]byte(crudBucket))
		c := b.Cursor()

		var crud Crud
		for k, v := c.First(); k != nil; k, v = c.Next() {
			_ = json.Unmarshal(v, &crud)
			crudsList = append(crudsList, crud.Name)
		}
		return nil
	})
	return crudsList
}

func IfExistCrud(name string) bool {
	ls := ListCruds()
	for _, item := range ls {
		if item == name {
			return true
		}
	}
	return false
}
