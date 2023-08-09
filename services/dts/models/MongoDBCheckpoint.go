// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package models


type MongoDBCheckpoint struct {

    /* 源端的名字 (Optional) */
    Replset *string `json:"replset"`

    /* 已经拉取的checkpoint位点（不一定写入） (Optional) */
    Lsn *MongoLSN `json:"lsn"`

    /* 已经成功写入目的端的checkpoint位点（已经成功写入，但是不一定持久化） (Optional) */
    LsnAck *MongoLSN `json:"lsnAck"`

    /* 已经成功持久化的checkpoint位点 (Optional) */
    LsnCkpt *MongoLSN `json:"lsnCkpt"`

    /* 当前的时间 (Optional) */
    Now *MongoLSN `json:"now"`
}