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
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo 全局唯一的单例客户端
// 只能初始化一次使用
type Mongo struct {
	ContextTimeout int    // 超时时间 单位:s
	DBName         string // 连接的库名称
	URL            string // 数据库连接URI 例如mongodb://localhost:27017
}

const (
	DefaultTimeout = 5
)

// Init 初始化连接
// 调用后客户端连接即建立可以进行操作
// 对于无法连接时 不会显式的返回错误 只有在第一次执行时报错
func (m *Mongo) Init() error {
	if m.check() != nil {
		return m.check()
	}

	return mgm.SetDefaultConfig(&mgm.Config{
		CtxTimeout: time.Duration(m.ContextTimeout) * time.Second,
	}, m.DBName, options.Client().ApplyURI(m.URL))
}

func (m *Mongo) check() error {
	if m.DBName == "" || m.URL == "" {
		return errors.New(ErrCheckError)
	}
	if m.ContextTimeout <= 0 {
		m.ContextTimeout = DefaultTimeout
	}
	return nil
}

// 根据模型存取

// Coll 取得当前集合
func (m *Mongo) Coll(mod Model) *Collection {
	return &Collection{mgm.Coll(mod)}
}

// Get 语法糖
// Coll(model).FindOne(context.Background(), bson.M{filter: key})
func (m *Mongo) Get(mod Model, filter interface{}) *SingleResult {
	res := (mgm.Coll(mod)).FindOne(context.Background(), filter)
	return (*SingleResult)(res)
}

// GetFilter 通过filter获取
// filter 可以为M D E A Raw
func (m *Mongo) GetFilter(mod Model, filter interface{}) *SingleResult {
	res := (mgm.Coll(mod)).FindOne(context.Background(), filter)
	return (*SingleResult)(res)
}

// GetAll r应该为指针
func (m *Mongo) GetAll(mod Model, r interface{}, filter interface{}) error {
	return mgm.Coll(mod).SimpleFind(r, filter)
}

// Update 更新
// data 需要传递指针
func (m *Mongo) Update(mod Model, data Model) error {
	return mgm.Coll(mod).Update(data)
}

// Insert 插入一条数据 返回错误
func (m *Mongo) Insert(mod Model, data interface{}) error {
	_, err := mgm.Coll(mod).InsertOne(context.Background(), data)
	return err
}

func (m *Mongo) Delete(mod Model, data Model) error {
	return mgm.Coll(mod).Delete(data)
}

// Persist 持久化
// 保存文件为bson数据 默认保存在当前目录下的tableName.bson
func (m *Mongo) Persist(mod Model) error {
	var res []bson.M
	err := mgm.Coll(mod).SimpleFind(&res, M{})
	if err != nil {
		return err
	}

	_, b, err := bson.MarshalValue(res)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return ioutil.WriteFile(genName(mgm.Coll(mod).Name()), b, 0644)
}

func genName(s string) string {
	return fmt.Sprintf("%s.bson", s)
}
