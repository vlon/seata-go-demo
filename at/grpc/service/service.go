/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/seata/seata-go-samples/at/grpc/unary"
	"github.com/seata/seata-go-samples/util"
	"sync"
	"time"
)

type GreetServer struct {
	lock     sync.Mutex
	alive    bool
	downTime time.Time
	db       *sql.DB
}

func NewGreetServer() *GreetServer {
	return &GreetServer{
		alive: true,
		db:    util.GetAtMySqlDb(),
	}
}

type OrderTbl struct {
	ID            int    `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	UserID        string `gorm:"column:user_id;NOT NULL"`
	CommodityCode string `gorm:"column:commodity_code"`
	Count         int    `gorm:"column:count"`
	Money         int64  `gorm:"column:money"`
	Descs         string `gorm:"column:descs"`
}

func (m *OrderTbl) TableName() string {
	return "order_tbl"
}

func (s *GreetServer) Greet(ctx context.Context, req *unary.Request) (*unary.Response, error) {
	fmt.Println("=>", req)

	//gormDB, err := gorm.Open(mysql.New(mysql.Config{
	//	Conn: util.GetAtMySqlDb(),
	//}), &gorm.Config{})
	//
	//data := &OrderTbl{
	//	ID:            26,
	//	UserID:        "NO-1000068",
	//	CommodityCode: "C100001",
	//	Count:         101,
	//	Money:         11,
	//	Descs:         "insert desc",
	//}
	//
	//err = gormDB.WithContext(ctx).Model(&OrderTbl{}).Create(data).Error
	//if err != nil {
	//	return nil, err
	//}
	////fmt.Println(row)
	//
	//data = &OrderTbl{
	//	ID:            1024,
	//	UserID:        "NO-1000067",
	//	CommodityCode: "C100001",
	//	Count:         101,
	//	Money:         11,
	//	Descs:         "insert desc",
	//}
	//
	//err = gormDB.WithContext(ctx).Model(&OrderTbl{}).Create(data).Error
	//if err != nil {
	//	return nil, err
	//}

	sql := "insert into order_tbl (id, user_id, commodity_code, count, money, descs) values (?, ?, ?, ?, ?, ?) "
	ret, err := s.db.ExecContext(ctx, sql, 36, "NO-100002", "C100000", 100, nil, "init desc")
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return nil, err
	}

	sql2 := "insert into order_tbl (id, user_id, commodity_code, count, money, descs) values (?, ?, ?, ?, ?, ?) "
	ret, err = s.db.ExecContext(ctx, sql2, 1024, "NO-100002", "C100000", 100, nil, "init desc")
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return nil, err
	}

	rows, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return nil, err
	}
	fmt.Printf("update success： %d.\n", rows)

	return &unary.Response{
		Greet: "hello from " + req.GetName(),
	}, nil
}
