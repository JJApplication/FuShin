/*
Create: 2022/8/17
Project: FuShin
Github: https://github.com/landers1037
Copyright Renj
*/

// Package mongo
package mongo

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

// Model 自定义表名 增加结构体函数
// CollectionName()
type Model mgm.Model

// MetaModel 不使用bson: inline时 无法正确获取到id
// 查询返回的结果 不能使用Model decode
type MetaModel mgm.DefaultModel

// Mongo的类型封装

// M bson字典类型
type M bson.M

// A bson数组类型
type A bson.A

// D bson有序字典
type D bson.D

// E bson element
type E bson.E

// Raw bson raw data
type Raw bson.Raw
