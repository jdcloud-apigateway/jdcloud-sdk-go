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

type QueryCompanyTransferRequest struct {

    core.JDCloudRequest

    /* 订单号  */
    OrderNumber string `json:"orderNumber"`
}

/*
 * param orderNumber: 订单号 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryCompanyTransferRequest(
    orderNumber string,
) *QueryCompanyTransferRequest {

	return &QueryCompanyTransferRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/company:transferStatus",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        OrderNumber: orderNumber,
	}
}

/*
 * param orderNumber: 订单号 (Required)
 */
func NewQueryCompanyTransferRequestWithAllParams(
    orderNumber string,
) *QueryCompanyTransferRequest {

    return &QueryCompanyTransferRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/company:transferStatus",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        OrderNumber: orderNumber,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryCompanyTransferRequestWithoutParam() *QueryCompanyTransferRequest {

    return &QueryCompanyTransferRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/company:transferStatus",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param orderNumber: 订单号(Required) */
func (r *QueryCompanyTransferRequest) SetOrderNumber(orderNumber string) {
    r.OrderNumber = orderNumber
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryCompanyTransferRequest) GetRegionId() string {
    return ""
}

type QueryCompanyTransferResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryCompanyTransferResult `json:"result"`
}

type QueryCompanyTransferResult struct {
    Success bool `json:"success"`
    HasException bool `json:"hasException"`
    Code string `json:"code"`
    Message string `json:"message"`
    Detail string `json:"detail"`
}