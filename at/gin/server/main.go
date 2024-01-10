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

package main

import (
	"database/sql"
	sql2 "github.com/seata/seata-go/pkg/datasource/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seata/seata-go-samples/util"
	"github.com/seata/seata-go/pkg/client"
	ginmiddleware "github.com/seata/seata-go/pkg/integration/gin"
	"github.com/seata/seata-go/pkg/util/log"
)

var db *sql.DB

func main() {
	client.InitPath("./conf/seatago.yml")
	db = util.GetAtMySqlDb()

	r := gin.Default()
	initDB()
	// NOTE: when use gin，must set ContextWithFallback true when gin version >= 1.8.1
	// r.ContextWithFallback = true

	r.Use(ginmiddleware.TransactionMiddleware())

	r.POST("/updateDataSuccess", updateDataSuccessHandler)
	r.POST("/selectForUpdateSuccess", selectForUpdateSuccHandler)

	r.POST("/insertData", insertData)

	r.POST("/insertOnUpdateDataSuccess", func(c *gin.Context) {
		log.Infof("get tm insertOnUpdateData")
		if err := insertOnUpdateDataSuccess(c); err != nil {
			c.JSON(http.StatusBadRequest, "insertOnUpdateData failure")
			return
		}
		c.JSON(http.StatusOK, "insertOnUpdateData ok")
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("start tcc server fatal: %v", err)
	}
}

func updateDataSuccessHandler(c *gin.Context) {
	log.Infof("get tm updateData")
	if err := updateDataSuccess(c); err != nil {
		c.JSON(http.StatusBadRequest, "updateData failure")
		return
	}
	c.JSON(http.StatusOK, "updateData ok")
}

func selectForUpdateSuccHandler(c *gin.Context) {
	log.Infof("execute select for update")
	if err := selectForUpdateSucc(c); err != nil {
		c.JSON(http.StatusBadRequest, "select for update failed")
		return
	}
	c.JSON(http.StatusOK, "select for update success")
}

func insertData(c *gin.Context) {
	data := &OrderTblModel{
		//Id:            1234,
		UserId:        "NO-100003",
		CommodityCode: "C100001",
		Count:         101,
		Money:         11,
		Descs:         "insert desc445566",
	}

	//orders := make([]*OrderTblModel, 0)
	//err := gormDB.Model(&OrderTblModel{}).Find(&orders).Error
	//if err != nil {
	//	panic(err)
	//}

	err := gormDB.WithContext(c.Request.Context()).Model(&OrderTblModel{}).Create(data).Error
	if err != nil {
		panic(err)
	}
}

var gormDB *gorm.DB

func initDB() {
	sqlDB, err := sql.Open(sql2.SeataATMySQLDriver, "root:root@tcp(127.0.0.1:33061)/seata_tbl?multiStatements=true&interpolateParams=true")
	if err != nil {
		panic("init service error")
	}

	gormDB, err = gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

type OrderTblModel struct {
	Id            int64  `gorm:"column:id" json:"id"`
	UserId        string `gorm:"column:user_id" json:"user_id"`
	CommodityCode string `gorm:"commodity_code" json:"commodity_code"`
	Count         int64  `gorm:"count" json:"count"`
	Money         int64  `gorm:"money" json:"money"`
	Descs         string `gorm:"descs" json:"descs"`
}

func (o *OrderTblModel) TableName() string {
	return "order_tbl"
}
