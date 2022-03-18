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

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    clickhouse "github.com/jdcloud-api/jdcloud-sdk-go/services/clickhouse/models"
)

type DescribeNodeClassesRequest struct {

    core.JDCloudRequest

    /* 地域代码  */
    RegionId string `json:"regionId"`

    /* 节点类型 CH ZK Monitor (Optional) */
    NodeType *string `json:"nodeType"`

    /* 存储类型： EBS_SSD (Optional) */
    NodeStorageType *string `json:"nodeStorageType"`
}

/*
 * param regionId: 地域代码 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeNodeClassesRequest(
    regionId string,
) *DescribeNodeClassesRequest {

	return &DescribeNodeClassesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances:describeNodeClasses",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域代码 (Required)
 * param nodeType: 节点类型 CH ZK Monitor (Optional)
 * param nodeStorageType: 存储类型： EBS_SSD (Optional)
 */
func NewDescribeNodeClassesRequestWithAllParams(
    regionId string,
    nodeType *string,
    nodeStorageType *string,
) *DescribeNodeClassesRequest {

    return &DescribeNodeClassesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances:describeNodeClasses",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        NodeType: nodeType,
        NodeStorageType: nodeStorageType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeNodeClassesRequestWithoutParam() *DescribeNodeClassesRequest {

    return &DescribeNodeClassesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances:describeNodeClasses",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码(Required) */
func (r *DescribeNodeClassesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param nodeType: 节点类型 CH ZK Monitor(Optional) */
func (r *DescribeNodeClassesRequest) SetNodeType(nodeType string) {
    r.NodeType = &nodeType
}

/* param nodeStorageType: 存储类型： EBS_SSD(Optional) */
func (r *DescribeNodeClassesRequest) SetNodeStorageType(nodeStorageType string) {
    r.NodeStorageType = &nodeStorageType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeNodeClassesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeNodeClassesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeNodeClassesResult `json:"result"`
}

type DescribeNodeClassesResult struct {
    ChSpec clickhouse.ChSpec `json:"chSpec"`
    ZkSpec clickhouse.ZkSpec `json:"zkSpec"`
    MonitorSpec clickhouse.MonitorSpec `json:"monitorSpec"`
}