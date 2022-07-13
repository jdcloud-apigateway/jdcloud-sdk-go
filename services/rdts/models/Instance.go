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


type Instance struct {

    /* 迁移任务ID (Optional) */
    InstanceId string `json:"instanceId"`

    /* 迁移任务名 (Optional) */
    MigrationName string `json:"migrationName"`

    /* 迁移状态（creating：创建中，validated：验证通过，invalid：验证失败，starting：启动中，waitMigrate：待迁移，migrating：迁移中，migrated：迁移成功，createFailed：创建失败，startFailed：启动失败，migrateFailed：迁移失败，error：异常错误，deleting：删除中，migrateTimeout：迁移超时) (Optional) */
    MigrationStatus string `json:"migrationStatus"`

    /* 迁移代理域名 (Optional) */
    ProxyDomain string `json:"proxyDomain"`

    /* 迁移代理连接密码 (Optional) */
    ProxyPassword string `json:"proxyPassword"`

    /* 当前迁移步骤 (Optional) */
    CurrentStep string `json:"currentStep"`

    /* 当前迁移步骤状态（init：初始化，running：运行中，success：成功，fail：失败） (Optional) */
    CurrentStepStatus string `json:"currentStepStatus"`

    /* 迁移失败原因 (Optional) */
    FailedReason string `json:"failedReason"`

    /* 源端redis实例配置信息 (Optional) */
    Source RedisConfig `json:"source"`

    /* 目的端redis实例配置信息 (Optional) */
    Target RedisConfig `json:"target"`

    /* 地域 (Optional) */
    Region string `json:"region"`

    /* 创建时间(ISO 8601标准的UTC时间，格式为：YYYY-MM-DDTHH:mm:ssZ) (Optional) */
    CreatedTime string `json:"createdTime"`
}