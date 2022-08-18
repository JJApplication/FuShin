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
	"go.mongodb.org/mongo-driver/mongo"
)

// 查询结果集合

type Collection struct {
	*mgm.Collection
}

// SingleResult mongo 单条查询的别名
type SingleResult mongo.SingleResult
