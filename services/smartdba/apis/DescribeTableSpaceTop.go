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
    smartdba "github.com/jdcloud-api/jdcloud-sdk-go/services/smartdba/models"
)

type DescribeTableSpaceTopRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* RDS 实例ID，唯一标识一个RDS实例。  */
    InstanceGid string `json:"instanceGid"`

    /* 排序字段：表空间(totalSize)、表数据空间(dataSize)、索引空间(idxSize)、碎片率(fragment)、行数(dataRows) (Optional) */
    SortField *string `json:"sortField"`

    /* 排序类型，desc 降序、asc 升序，默认降序 (Optional) */
    SortType *string `json:"sortType"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceGid: RDS 实例ID，唯一标识一个RDS实例。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeTableSpaceTopRequest(
    regionId string,
    instanceGid string,
) *DescribeTableSpaceTopRequest {

	return &DescribeTableSpaceTopRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instance/{instanceGid}/tableSpaceTop",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        InstanceGid: instanceGid,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceGid: RDS 实例ID，唯一标识一个RDS实例。 (Required)
 * param sortField: 排序字段：表空间(totalSize)、表数据空间(dataSize)、索引空间(idxSize)、碎片率(fragment)、行数(dataRows) (Optional)
 * param sortType: 排序类型，desc 降序、asc 升序，默认降序 (Optional)
 */
func NewDescribeTableSpaceTopRequestWithAllParams(
    regionId string,
    instanceGid string,
    sortField *string,
    sortType *string,
) *DescribeTableSpaceTopRequest {

    return &DescribeTableSpaceTopRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance/{instanceGid}/tableSpaceTop",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        InstanceGid: instanceGid,
        SortField: sortField,
        SortType: sortType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeTableSpaceTopRequestWithoutParam() *DescribeTableSpaceTopRequest {

    return &DescribeTableSpaceTopRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instance/{instanceGid}/tableSpaceTop",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *DescribeTableSpaceTopRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceGid: RDS 实例ID，唯一标识一个RDS实例。(Required) */
func (r *DescribeTableSpaceTopRequest) SetInstanceGid(instanceGid string) {
    r.InstanceGid = instanceGid
}

/* param sortField: 排序字段：表空间(totalSize)、表数据空间(dataSize)、索引空间(idxSize)、碎片率(fragment)、行数(dataRows)(Optional) */
func (r *DescribeTableSpaceTopRequest) SetSortField(sortField string) {
    r.SortField = &sortField
}

/* param sortType: 排序类型，desc 降序、asc 升序，默认降序(Optional) */
func (r *DescribeTableSpaceTopRequest) SetSortType(sortType string) {
    r.SortType = &sortType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeTableSpaceTopRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeTableSpaceTopResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeTableSpaceTopResult `json:"result"`
}

type DescribeTableSpaceTopResult struct {
    Data []smartdba.TableStorageInfo `json:"data"`
}