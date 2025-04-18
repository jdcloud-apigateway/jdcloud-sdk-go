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


type TableBriefInfo struct {

    /* 表名  */
    TableName string `json:"tableName"`

    /* 表操作类型 CREATE:新建 ALTER:修改 GRANT_PRIV:赋权 (Optional) */
    OperationType string `json:"operationType"`

    /* 发布阶段 UNPUBLISH:未发布、PUBLISHING:待发布、PUBLISHED:已发布 (Optional) */
    PublishStage string `json:"publishStage"`

    /* 环境信息 dev:开发 prod:生产 (Optional) */
    Env string `json:"env"`

    /* 表负责人 (Optional) */
    Owner string `json:"owner"`

    /* 表协作人，多个以逗号分隔 (Optional) */
    Managers string `json:"managers"`
}
