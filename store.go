package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

var db *bolt.DB

func init() {
	dbCreate, err := bolt.Open("azbrowse.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	db = dbCreate

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("cache"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func put(key, value string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("cache"))
		err := b.Put([]byte(key), []byte(value))
		return err
	})
}

func get(key string) (string, error) {
	var s []byte
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("cache"))
		v := b.Get([]byte(key))
		s = v
		return nil
	})
	if err != nil || s == nil {
		return "", fmt.Errorf("Failed to find item: %v", err)
	}
	return string(s), nil
}