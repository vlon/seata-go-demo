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
	"flag"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/seata/seata-go/pkg/constant"
	"github.com/seata/seata-go/pkg/tm"
	"github.com/seata/seata-go/pkg/util/log"
	"net/http"
	"time"

	"github.com/seata/seata-go/pkg/client"
)

var serverIpPort = "http://127.0.0.1:8080"

func main() {
	flag.Parse()
	client.InitPath("./conf/seatago.yml")

	bgCtx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()

	// sample update
	//sampleUpdate(bgCtx)
	//
	// sample insert on update
	//sampleInsertOnUpdate(bgCtx)
	insertData(bgCtx)

	//// sample select for update
	//sampleSelectForUpdate(bgCtx)
}

func insert(ctx context.Context) (re error) {
	request := gorequest.New()

	log.Infof("branch transaction begin")
	request.Post(serverIpPort+"/insertOnUpdateDataSuccess").
		Set(constant.XidKey, tm.GetXID(ctx)).
		End(func(response gorequest.Response, body string, errs []error) {
			if response.StatusCode != http.StatusOK {
				re = fmt.Errorf("insert on update data fail")
			}
		})

	request.Post(serverIpPort+"/insertData").
		Set(constant.XidKey, tm.GetXID(ctx)).
		End(func(response gorequest.Response, body string, errs []error) {
			if response.StatusCode != http.StatusOK {
				re = fmt.Errorf("insert on update data fail")
			}
		})

	log.Infof("branch transaction begin")

	return
}

func insertData(ctx context.Context) {
	if err := tm.WithGlobalTx(ctx, &tm.GtxConfig{
		Name:    "ATSampleLocalGlobalTx_Insert",
		Timeout: time.Second * 30,
	}, insert); err != nil {
		panic(fmt.Sprintf("tm insert on update data err, %v", err))
	}
}
