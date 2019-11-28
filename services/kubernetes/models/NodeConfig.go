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


type NodeConfig struct {

    /* 实例类型 (Optional) */
    InstanceType string `json:"instanceType"`

    /* 镜像信息 (Optional) */
    ImageId string `json:"imageId"`

    /* 云主机SSH密钥对名称 (Optional) */
    KeyNames []string `json:"keyNames"`

    /* 云主机磁盘类型 (Optional) */
    SystemDiskCategory string `json:"systemDiskCategory"`

    /* 云主机云盘系统盘大小  单位(GB) (Optional) */
    SystemDiskSize int `json:"systemDiskSize"`

    /* 云主机云盘系统盘类型 (Optional) */
    SystemDiskType string `json:"systemDiskType"`

    /* 云主机云盘 iops，仅限 ssd 类型云盘有效 (Optional) */
    SystemDiskIops int `json:"systemDiskIops"`

    /* 工作节点组标签 (Optional) */
    Labels []LabelSpec `json:"labels"`
}
