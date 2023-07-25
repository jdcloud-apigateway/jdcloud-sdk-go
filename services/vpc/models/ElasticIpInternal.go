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

import charge "github.com/jdcloud-api/jdcloud-sdk-go/services/charge/models"

type ElasticIpInternal struct {

    /* 弹性公网IP的Id (Optional) */
    ElasticIpId string `json:"elasticIpId"`

    /* 弹性公网IP的地址 (Optional) */
    ElasticIpAddress string `json:"elasticIpAddress"`

    /* 弹性公网IP的限速（单位：Mbps) (Optional) */
    BandwidthMbps int `json:"bandwidthMbps"`

    /* 弹性公网IP的线路，标准公网IP的线路、取值为bgp或no_bgp；边缘公网IP的线路、可通过describeEdgeIpProviders接口获取 (Optional) */
    Provider string `json:"provider"`

    /* 私有IP的IPV4地址 (Optional) */
    PrivateIpAddress string `json:"privateIpAddress"`

    /* 配置弹性网卡Id (Optional) */
    NetworkInterfaceId string `json:"networkInterfaceId"`

    /* 实例Id (Optional) */
    InstanceId string `json:"instanceId"`

    /* 实例类型,取值为：compute、lb、container、pod、natgw (Optional) */
    InstanceType string `json:"instanceType"`

    /* 计费配置 (Optional) */
    Charge charge.Charge `json:"charge"`

    /* 弹性公网IP的创建时间 (Optional) */
    CreatedTime string `json:"createdTime"`

    /* 弹性公网IP的可用区属性，如果为空，表示全可用区 (Optional) */
    Az string `json:"az"`

    /* Tag信息 (Optional) */
    Tags []Tag `json:"tags"`

    /* 弹性公网IP的IP类型，取值：standard(标准弹性IP)、edge(边缘弹性IP) (Optional) */
    IpType string `json:"ipType"`

    /* 加入的共享带宽包ID，如果没有加入共享带宽包该值为空 (Optional) */
    BandwidthPackageId string `json:"bandwidthPackageId"`

    /* IP是否被绑定，取值：ASSOCIATED（被绑定）、NOT_ASSOCIATED（未被绑定） (Optional) */
    Status string `json:"status"`

    /* IP是否变更，取值：UP（正常使用）、DOWN（停止使用）、PROCESSING（配置变更中） (Optional) */
    State string `json:"state"`

    /* 资源所属资源组ID (Optional) */
    ResourceGroupId string `json:"resourceGroupId"`
}
