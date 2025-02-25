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
    pod "github.com/jdcloud-api/jdcloud-sdk-go/services/pod/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribePodTemplatesRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 页码；默认为1。取值范围：[1,∞) (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为20；取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* podTemplateIds - podTemplate ID，精确匹配，支持多个
name - pod模板名称，模糊匹配，支持单个
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribePodTemplatesRequest(
    regionId string,
) *DescribePodTemplatesRequest {

	return &DescribePodTemplatesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/podTemplates",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param pageNumber: 页码；默认为1。取值范围：[1,∞) (Optional)
 * param pageSize: 分页大小；默认为20；取值范围[10, 100] (Optional)
 * param filters: podTemplateIds - podTemplate ID，精确匹配，支持多个
name - pod模板名称，模糊匹配，支持单个
 (Optional)
 */
func NewDescribePodTemplatesRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    filters []common.Filter,
) *DescribePodTemplatesRequest {

    return &DescribePodTemplatesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/podTemplates",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribePodTemplatesRequestWithoutParam() *DescribePodTemplatesRequest {

    return &DescribePodTemplatesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/podTemplates",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribePodTemplatesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param pageNumber: 页码；默认为1。取值范围：[1,∞)(Optional) */
func (r *DescribePodTemplatesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}
/* param pageSize: 分页大小；默认为20；取值范围[10, 100](Optional) */
func (r *DescribePodTemplatesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}
/* param filters: podTemplateIds - podTemplate ID，精确匹配，支持多个
name - pod模板名称，模糊匹配，支持单个
(Optional) */
func (r *DescribePodTemplatesRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribePodTemplatesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribePodTemplatesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribePodTemplatesResult `json:"result"`
}

type DescribePodTemplatesResult struct {
    PodTemplates []pod.PodTemplate `json:"podTemplates"`
    TotalCount int `json:"totalCount"`
}