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
    refund "github.com/jdcloud-api/jdcloud-sdk-go/services/refund/models"
)

type DescirbeRefundStatusRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /*   */
    Pin string `json:"pin"`

    /*  (Optional) */
    RefundId *string `json:"refundId"`

    /*  (Optional) */
    ResourceId *string `json:"resourceId"`

    /* 第几页，默认值为0 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 每页条数，默认为20 (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: 地域ID (Required)
 * param pin:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescirbeRefundStatusRequest(
    regionId string,
    pin string,
) *DescirbeRefundStatusRequest {

	return &DescirbeRefundStatusRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/refund:status",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Pin: pin,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param pin:  (Required)
 * param refundId:  (Optional)
 * param resourceId:  (Optional)
 * param pageNumber: 第几页，默认值为0 (Optional)
 * param pageSize: 每页条数，默认为20 (Optional)
 */
func NewDescirbeRefundStatusRequestWithAllParams(
    regionId string,
    pin string,
    refundId *string,
    resourceId *string,
    pageNumber *int,
    pageSize *int,
) *DescirbeRefundStatusRequest {

    return &DescirbeRefundStatusRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/refund:status",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Pin: pin,
        RefundId: refundId,
        ResourceId: resourceId,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescirbeRefundStatusRequestWithoutParam() *DescirbeRefundStatusRequest {

    return &DescirbeRefundStatusRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/refund:status",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DescirbeRefundStatusRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pin: (Required) */
func (r *DescirbeRefundStatusRequest) SetPin(pin string) {
    r.Pin = pin
}

/* param refundId: (Optional) */
func (r *DescirbeRefundStatusRequest) SetRefundId(refundId string) {
    r.RefundId = &refundId
}

/* param resourceId: (Optional) */
func (r *DescirbeRefundStatusRequest) SetResourceId(resourceId string) {
    r.ResourceId = &resourceId
}

/* param pageNumber: 第几页，默认值为0(Optional) */
func (r *DescirbeRefundStatusRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 每页条数，默认为20(Optional) */
func (r *DescirbeRefundStatusRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescirbeRefundStatusRequest) GetRegionId() string {
    return r.RegionId
}

type DescirbeRefundStatusResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescirbeRefundStatusResult `json:"result"`
}

type DescirbeRefundStatusResult struct {
    PageInfos refund.PageInfos `json:"pageInfos"`
    Data []refund.RefundResult `json:"data"`
}