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
    jdccs "github.com/jdcloud-api/jdcloud-sdk-go/services/jdccs/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeCabinetsRequest struct {

    core.JDCloudRequest

    /* IDC机房ID  */
    Idc string `json:"idc"`

    /* 页码, 默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认为20 (Optional) */
    PageSize *int `json:"pageSize"`

    /* roomNo - 房间号，精确匹配，支持多个
cabinetId - 机柜ID，精确匹配，支持多个
 (Optional) */
    Filters []common.Filter `json:"filters"`

    /* 机柜开通状态 disabled:未开通 enabling:开通中 enabled:已开通 disabling:关电中 (Optional) */
    CabinetOpenStatus *string `json:"cabinetOpenStatus"`
}

/*
 * param idc: IDC机房ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeCabinetsRequest(
    idc string,
) *DescribeCabinetsRequest {

	return &DescribeCabinetsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/idcs/{idc}/cabinets",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Idc: idc,
	}
}

/*
 * param idc: IDC机房ID (Required)
 * param pageNumber: 页码, 默认为1 (Optional)
 * param pageSize: 分页大小，默认为20 (Optional)
 * param filters: roomNo - 房间号，精确匹配，支持多个
cabinetId - 机柜ID，精确匹配，支持多个
 (Optional)
 * param cabinetOpenStatus: 机柜开通状态 disabled:未开通 enabling:开通中 enabled:已开通 disabling:关电中 (Optional)
 */
func NewDescribeCabinetsRequestWithAllParams(
    idc string,
    pageNumber *int,
    pageSize *int,
    filters []common.Filter,
    cabinetOpenStatus *string,
) *DescribeCabinetsRequest {

    return &DescribeCabinetsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/idcs/{idc}/cabinets",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Idc: idc,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Filters: filters,
        CabinetOpenStatus: cabinetOpenStatus,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeCabinetsRequestWithoutParam() *DescribeCabinetsRequest {

    return &DescribeCabinetsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/idcs/{idc}/cabinets",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param idc: IDC机房ID(Required) */
func (r *DescribeCabinetsRequest) SetIdc(idc string) {
    r.Idc = idc
}

/* param pageNumber: 页码, 默认为1(Optional) */
func (r *DescribeCabinetsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小，默认为20(Optional) */
func (r *DescribeCabinetsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param filters: roomNo - 房间号，精确匹配，支持多个
cabinetId - 机柜ID，精确匹配，支持多个
(Optional) */
func (r *DescribeCabinetsRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

/* param cabinetOpenStatus: 机柜开通状态 disabled:未开通 enabling:开通中 enabled:已开通 disabling:关电中(Optional) */
func (r *DescribeCabinetsRequest) SetCabinetOpenStatus(cabinetOpenStatus string) {
    r.CabinetOpenStatus = &cabinetOpenStatus
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeCabinetsRequest) GetRegionId() string {
    return ""
}

type DescribeCabinetsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeCabinetsResult `json:"result"`
}

type DescribeCabinetsResult struct {
    Cabinets []jdccs.DescribeCabinet `json:"cabinets"`
    PageNumber int `json:"pageNumber"`
    PageSize int `json:"pageSize"`
    TotalCount int `json:"totalCount"`
}