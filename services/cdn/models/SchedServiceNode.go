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


type SchedServiceNode struct {

    /*  (Optional) */
    Id int64 `json:"id"`

    /* 0=scdn node (Optional) */
    BusinessType int64 `json:"businessType"`

    /* 0:边缘节点，1:二级源节点 (Optional) */
    CdnNodeLevel int64 `json:"cdnNodeLevel"`

    /* 节点类型：1:大文件 2:小文件 3:直播 4:二级源大文件 5:二级源小文件 6:直播中转 (Optional) */
    NodeType int64 `json:"nodeType"`

    /* 节点状态 0：正常 1：异常 2：下线 (Optional) */
    CdnNodeStatus int64 `json:"cdnNodeStatus"`

    /* cdn节点Code (Optional) */
    CdnNodeName string `json:"cdnNodeName"`

    /* 数据中心code (Optional) */
    DatacenterCode string `json:"datacenterCode"`

    /* 机房code (Optional) */
    IdcCode string `json:"idcCode"`

    /* 机房名称 (Optional) */
    IdcName string `json:"idcName"`

    /* 所属运营商 (Optional) */
    Isp string `json:"isp"`

    /* 所属运营商名称 (Optional) */
    IspName string `json:"ispName"`

    /* 节点名称 (Optional) */
    Name string `json:"name"`

    /* 所属省份 (Optional) */
    Prov string `json:"prov"`

    /* vip (Optional) */
    Vips []VipItem `json:"vips"`
}