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
)

type DescribePlansSoldOutStatusRequest struct {

    core.JDCloudRequest

    /* regionId
  */
    RegionId string `json:"regionId"`

    /* 方案Ids, jsonArray的string 例如 `[\"plan-id1\", \"plan-id2\"]`
  */
    PlanIds string `json:"planIds"`
}

/*
 * param regionId: regionId
 (Required)
 * param planIds: 方案Ids, jsonArray的string 例如 `[\"plan-id1\", \"plan-id2\"]`
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribePlansSoldOutStatusRequest(
    regionId string,
    planIds string,
) *DescribePlansSoldOutStatusRequest {

	return &DescribePlansSoldOutStatusRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/describePlansSoldOutStatus",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        PlanIds: planIds,
	}
}

/*
 * param regionId: regionId
 (Required)
 * param planIds: 方案Ids, jsonArray的string 例如 `[\"plan-id1\", \"plan-id2\"]`
 (Required)
 */
func NewDescribePlansSoldOutStatusRequestWithAllParams(
    regionId string,
    planIds string,
) *DescribePlansSoldOutStatusRequest {

    return &DescribePlansSoldOutStatusRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/describePlansSoldOutStatus",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PlanIds: planIds,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribePlansSoldOutStatusRequestWithoutParam() *DescribePlansSoldOutStatusRequest {

    return &DescribePlansSoldOutStatusRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/describePlansSoldOutStatus",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: regionId
(Required) */
func (r *DescribePlansSoldOutStatusRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param planIds: 方案Ids, jsonArray的string 例如 `[\"plan-id1\", \"plan-id2\"]`
(Required) */
func (r *DescribePlansSoldOutStatusRequest) SetPlanIds(planIds string) {
    r.PlanIds = planIds
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribePlansSoldOutStatusRequest) GetRegionId() string {
    return r.RegionId
}

type DescribePlansSoldOutStatusResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribePlansSoldOutStatusResult `json:"result"`
}

type DescribePlansSoldOutStatusResult struct {
    StatusMap interface{} `json:"statusMap"`
}