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
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/seata/seata-go/pkg/tm"
	"github.com/seata/seata-go/pkg/util/log"
)

var gormDB *gorm.DB

func init() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:root@tcp(127.0.0.1:33061)/seata_tbl?charset=utf8mb4&parseTime=True&loc=Local",
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	gormDB = db
}

type RMService struct{}

func (b *RMService) Prepare(ctx context.Context, params interface{}) (bool, error) {
	log.Infof("TRMService Prepare, start........ params: %v", params)
	//data := &OrderTblModel{
	//	//Id:            1024,
	//	UserId:        "NO-1000034",
	//	CommodityCode: "C1000012",
	//	Count:         1017,
	//	Money:         119,
	//	Descs:         "insert descggggg",
	//}

	//orders := make([]*OrderTblModel, 0)
	//err := gormDB.Model(&OrderTblModel{}).Find(&orders).Error
	//if err != nil {
	//	panic(err)
	//}

	err := gormDB.WithContext(ctx).Model(&OrderTblModel{}).Create(params).Error
	if err != nil {
		return false, err
	}
	log.Infof("TRMService Prepare finish")
	return true, nil
}

func (b *RMService) Commit(ctx context.Context, businessActionContext *tm.BusinessActionContext) (bool, error) {
	log.Infof("RMService Commit start....., param %v", businessActionContext)

	//err := gormDB.WithContext().Commit().Error
	//if err != nil {
	//	panic(err)
	//}
	log.Infof("RMService Commit finish")
	return true, nil
}

func (b *RMService) Rollback(ctx context.Context, businessActionContext *tm.BusinessActionContext) (bool, error) {
	log.Infof("RMService Rollback, param %v", businessActionContext)
	return true, nil
}

func (b *RMService) GetActionName() string {
	return "ginTccRMService"
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

type RMService2 struct{}

func (b *RMService2) Prepare(ctx context.Context, params interface{}) (bool, error) {
	log.Infof("TRMService Prepare, start........ params: %v", params)
	//data := &OrderTblModel{
	//	//Id:            1024,
	//	UserId:        "NO-1000034",
	//	CommodityCode: "C1000012",
	//	Count:         1017,
	//	Money:         119,
	//	Descs:         "insert descggggg",
	//}

	//orders := make([]*OrderTblModel, 0)
	//err := gormDB.Model(&OrderTblModel{}).Find(&orders).Error
	//if err != nil {
	//	panic(err)
	//}

	err := gormDB.WithContext(ctx).Model(&OrderTblModel{}).Create(params).Error
	if err != nil {
		return false, err
	}
	log.Infof("TRMService Prepare finish")
	return true, nil
}

func (b *RMService2) Commit(ctx context.Context, businessActionContext *tm.BusinessActionContext) (bool, error) {
	log.Infof("RMService Commit start....., param %v", businessActionContext)

	//err := gormDB.WithContext().Commit().Error
	//if err != nil {
	//	panic(err)
	//}
	log.Infof("RMService Commit finish")
	return true, nil
}

func (b *RMService2) Rollback(ctx context.Context, businessActionContext *tm.BusinessActionContext) (bool, error) {
	log.Infof("RMService Rollback, param %v", businessActionContext)
	return true, nil
}

func (b *RMService2) GetActionName() string {
	return "ginTccRMService2"
}
