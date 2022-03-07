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
    dms "github.com/jdcloud-api/jdcloud-sdk-go/services/dms/models"
)

type QueryPersonalSqlsRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* 数据源id (Optional) */
    DataSourceId *int `json:"dataSourceId"`

    /* 显示数据的页码，取值范围：[1,∞)。pageNumber为Null时，返回所有数据页码；超过总页数时，无数据。 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 每页显示的数据条数，用于查询列表的接口。 (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryPersonalSqlsRequest(
    regionId string,
) *QueryPersonalSqlsRequest {

	return &QueryPersonalSqlsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/personalSql:query",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param dataSourceId: 数据源id (Optional)
 * param pageNumber: 显示数据的页码，取值范围：[1,∞)。pageNumber为Null时，返回所有数据页码；超过总页数时，无数据。 (Optional)
 * param pageSize: 每页显示的数据条数，用于查询列表的接口。 (Optional)
 */
func NewQueryPersonalSqlsRequestWithAllParams(
    regionId string,
    dataSourceId *int,
    pageNumber *int,
    pageSize *int,
) *QueryPersonalSqlsRequest {

    return &QueryPersonalSqlsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/personalSql:query",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DataSourceId: dataSourceId,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryPersonalSqlsRequestWithoutParam() *QueryPersonalSqlsRequest {

    return &QueryPersonalSqlsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/personalSql:query",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *QueryPersonalSqlsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param dataSourceId: 数据源id(Optional) */
func (r *QueryPersonalSqlsRequest) SetDataSourceId(dataSourceId int) {
    r.DataSourceId = &dataSourceId
}

/* param pageNumber: 显示数据的页码，取值范围：[1,∞)。pageNumber为Null时，返回所有数据页码；超过总页数时，无数据。(Optional) */
func (r *QueryPersonalSqlsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 每页显示的数据条数，用于查询列表的接口。(Optional) */
func (r *QueryPersonalSqlsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryPersonalSqlsRequest) GetRegionId() string {
    return r.RegionId
}

type QueryPersonalSqlsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryPersonalSqlsResult `json:"result"`
}

type QueryPersonalSqlsResult struct {
    Count int `json:"count"`
    PersonalSqls []dms.PersonalSql `json:"personalSqls"`
}