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
    cloudsign "github.com/jdcloud-api/jdcloud-sdk-go/services/cloudsign/models"
)

type GetSaveReportRequest struct {

    core.JDCloudRequest

    /* 业务流水号  */
    BusinessId string `json:"businessId"`

    /* 业务代码 (Optional) */
    ChainCode *string `json:"chainCode"`
}

/*
 * param businessId: 业务流水号 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetSaveReportRequest(
    businessId string,
) *GetSaveReportRequest {

	return &GetSaveReportRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/evidence:evidenceGetSaveReport",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        BusinessId: businessId,
	}
}

/*
 * param businessId: 业务流水号 (Required)
 * param chainCode: 业务代码 (Optional)
 */
func NewGetSaveReportRequestWithAllParams(
    businessId string,
    chainCode *string,
) *GetSaveReportRequest {

    return &GetSaveReportRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/evidence:evidenceGetSaveReport",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        BusinessId: businessId,
        ChainCode: chainCode,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetSaveReportRequestWithoutParam() *GetSaveReportRequest {

    return &GetSaveReportRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/evidence:evidenceGetSaveReport",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param businessId: 业务流水号(Required) */
func (r *GetSaveReportRequest) SetBusinessId(businessId string) {
    r.BusinessId = businessId
}
/* param chainCode: 业务代码(Optional) */
func (r *GetSaveReportRequest) SetChainCode(chainCode string) {
    r.ChainCode = &chainCode
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetSaveReportRequest) GetRegionId() string {
    return ""
}

type GetSaveReportResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetSaveReportResult `json:"result"`
}

type GetSaveReportResult struct {
    Code string `json:"code"`
    Message string `json:"message"`
    Success bool `json:"success"`
    Data cloudsign.EvidenceFile `json:"data"`
}