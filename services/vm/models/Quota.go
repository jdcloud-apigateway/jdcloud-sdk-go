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


type Quota struct {

    /* 资源类型。支持范围：
`instance`：云主机。
`instance_cpu`：云主机的CPU。
`instance_memory`：云主机的内存。
`instance_local_disk`：云主机的本地盘。
`keypair`：密钥。
`image`：镜像。
`instanceTemplate`：实例模板。
`imageShare`：共享镜像。
 (Optional) */
    ResourceType string `json:"resourceType"`

    /* 配额上限。 (Optional) */
    Limit int `json:"limit"`

    /* 已用配额。 (Optional) */
    Used int `json:"used"`
}
