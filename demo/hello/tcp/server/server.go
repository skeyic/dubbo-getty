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
	"flag"
)

import (
	gxsync "github.com/dubbogo/gost/sync"
)

import (
	"github.com/skeyic/dubbo-getty"
	"github.com/skeyic/dubbo-getty/demo/hello/tcp"
	"github.com/skeyic/dubbo-getty/demo/util"
)

var (
	taskPoolMode = flag.Bool("taskPool", false, "task pool mode")
	taskPoolSize = flag.Int("task_pool_size", 2000, "task poll size")
	pprofPort    = flag.Int("pprof_port", 65432, "pprof http port")
)

var taskPool gxsync.GenericTaskPool

func main() {
	flag.Parse()

	util.SetLimit()

	util.Profiling(*pprofPort)

	options := []getty.ServerOption{getty.WithLocalAddress(":8090")}

	if *taskPoolMode {
		taskPool = gxsync.NewTaskPoolSimple(*taskPoolSize)
		options = append(options, getty.WithServerTaskPool(taskPool))
	}

	server := getty.NewTCPServer(options...)

	go server.RunEventLoop(NewHelloServerSession)

	util.WaitCloseSignals(server)
}

func NewHelloServerSession(session getty.Session) (err error) {
	err = tcp.InitialSession(session)
	if err != nil {
		return
	}
	return
}
