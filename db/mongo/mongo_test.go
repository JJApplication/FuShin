/*
Create: 2022/8/17
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package mongo
package mongo

import (
	"context"
	"testing"
)

type TestModel struct {
	Name      string `json:"name" bson:"name"`
	MetaModel `bson:",inline"`
}

func (tt TestModel) CollectionName() string {
	return "TestColl"
}

func TestMongo(t *testing.T) {
	m := Mongo{
		ContextTimeout: 1,
		DBName:         "Test",
		URL:            "mongodb://127.0.0.1:27017",
	}
	err := m.Init()
	if err != nil {
		t.Error(err)
	}
	var data TestModel
	name := m.Coll(&TestModel{}).Name()
	t.Log(name)

	// find
	err = m.Coll(&TestModel{}).FindOne(context.Background(), M{"name": "test"}).Decode(&data)
	t.Log(data)

	// insert
	data.Name = "test normal insert"
	insres, err := m.Coll(&TestModel{}).InsertOne(context.Background(), data)
	t.Log(insres, err)

	// m insert
	data.Name = "test short insert"
	err = m.Insert(&TestModel{}, data)
	t.Log(err)

	var ud TestModel
	err = m.Coll(&TestModel{}).FindOne(context.Background(), M{"name": "test"}).Decode(&ud)
	t.Log(ud, err)

	// update
	ud.Name = "test normal update"
	err = m.Coll(&TestModel{}).Update(&ud)
	t.Log(err)
	// update short
	ud.Name = "test normal update"
	err = m.Update(&TestModel{}, &ud)
	t.Log(err)
}

func TestPersist(t *testing.T) {
	m := Mongo{
		ContextTimeout: 1,
		DBName:         "Test",
		URL:            "mongodb://127.0.0.1:27017",
	}
	err := m.Init()
	if err != nil {
		t.Error(err)
	}

	err = m.Persist(&TestModel{})
	t.Log(err)
}
