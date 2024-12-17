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

type DeductOweFeeRequest struct {

    core.JDCloudRequest

    /* 地域编码，参考OpenAPI公共说明  */
    RegionId string `json:"regionId"`

    /* 抵扣订单号 (Optional) */
    OrderIds *string `json:"orderIds"`
}

/*
 * param regionId: 地域编码，参考OpenAPI公共说明 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeductOweFeeRequest(
    regionId string,
) *DeductOweFeeRequest {

	return &DeductOweFeeRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/deductOweFee",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域编码，参考OpenAPI公共说明 (Required)
 * param orderIds: 抵扣订单号 (Optional)
 */
func NewDeductOweFeeRequestWithAllParams(
    regionId string,
    orderIds *string,
) *DeductOweFeeRequest {

    return &DeductOweFeeRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/deductOweFee",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        OrderIds: orderIds,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeductOweFeeRequestWithoutParam() *DeductOweFeeRequest {

    return &DeductOweFeeRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/deductOweFee",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域编码，参考OpenAPI公共说明(Required) */
func (r *DeductOweFeeRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param orderIds: 抵扣订单号(Optional) */
func (r *DeductOweFeeRequest) SetOrderIds(orderIds string) {
    r.OrderIds = &orderIds
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeductOweFeeRequest) GetRegionId() string {
    return r.RegionId
}

type DeductOweFeeResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeductOweFeeResult `json:"result"`
}

type DeductOweFeeResult struct {
    Code int `json:"code"`
    Msg string `json:"msg"`
}