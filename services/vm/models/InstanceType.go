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


type InstanceType struct {

    /* 实例规格族。 (Optional) */
    Family string `json:"family"`

    /* 实例规格。 (Optional) */
    InstanceType string `json:"instanceType"`

    /* cpu个数。 (Optional) */
    Cpu int `json:"cpu"`

    /* 镜像架构。取值范围：`x86_64、arm64`。 (Optional) */
    Architecture string `json:"architecture"`

    /* 内存大小。 (Optional) */
    MemoryMB int `json:"memoryMB"`

    /* 支持绑定的弹性网卡数量，包括主网卡。 (Optional) */
    NicLimit int `json:"nicLimit"`

    /* 支持挂载的云硬盘数量，包括云盘系统盘。 (Optional) */
    CloudDiskCountLimit int `json:"cloudDiskCountLimit"`

    /* 实例规格描述。 (Optional) */
    Desc string `json:"desc"`

    /* 实例规格售卖状态。已售罄的实例规格无法使用。 (Optional) */
    State []InstanceTypeState `json:"state"`

    /* GPU配置，针对GPU类型的实例规格有效。 (Optional) */
    Gpu Gpu `json:"gpu"`

    /* 本地数据盘配置（缓存盘），针对GPU类型、或本地存储型的实例规格有效。 (Optional) */
    LocalDisks []LocalDisk `json:"localDisks"`

    /* 实例规格代数。 (Optional) */
    Generation int `json:"generation"`

    /* 突发型规格信息 (Optional) */
    BurstInfo InstanceTypeBurstInfo `json:"burstInfo"`

    /* 支持的云盘类型 (Optional) */
    CloudDiskTypes []string `json:"cloudDiskTypes"`
}
