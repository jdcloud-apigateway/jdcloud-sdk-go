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
    vm "github.com/jdcloud-api/jdcloud-sdk-go/services/vm/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeInstanceTemplatesCustomdataRequest struct {

    core.JDCloudRequest

    /* 地域ID。  */
    RegionId string `json:"regionId"`

    /* <b>filters 中支持使用以下关键字进行过滤</b>
`instanceTemplateId`: 实例模板ID，精确匹配，最多支持10个
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域ID。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeInstanceTemplatesCustomdataRequest(
    regionId string,
) *DescribeInstanceTemplatesCustomdataRequest {

	return &DescribeInstanceTemplatesCustomdataRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instanceTemplatesCustomData",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID。 (Required)
 * param filters: <b>filters 中支持使用以下关键字进行过滤</b>
`instanceTemplateId`: 实例模板ID，精确匹配，最多支持10个
 (Optional)
 */
func NewDescribeInstanceTemplatesCustomdataRequestWithAllParams(
    regionId string,
    filters []common.Filter,
) *DescribeInstanceTemplatesCustomdataRequest {

    return &DescribeInstanceTemplatesCustomdataRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instanceTemplatesCustomData",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeInstanceTemplatesCustomdataRequestWithoutParam() *DescribeInstanceTemplatesCustomdataRequest {

    return &DescribeInstanceTemplatesCustomdataRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instanceTemplatesCustomData",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID。(Required) */
func (r *DescribeInstanceTemplatesCustomdataRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param filters: <b>filters 中支持使用以下关键字进行过滤</b>
`instanceTemplateId`: 实例模板ID，精确匹配，最多支持10个
(Optional) */
func (r *DescribeInstanceTemplatesCustomdataRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeInstanceTemplatesCustomdataRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeInstanceTemplatesCustomdataResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeInstanceTemplatesCustomdataResult `json:"result"`
}

type DescribeInstanceTemplatesCustomdataResult struct {
    InstanceTemplatesCustomData []vm.InstanceTemplateCustomData `json:"instanceTemplatesCustomData"`
    TotalCount int `json:"totalCount"`
}