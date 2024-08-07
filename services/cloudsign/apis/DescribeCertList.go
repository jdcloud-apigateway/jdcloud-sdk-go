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

type DescribeCertListRequest struct {

    core.JDCloudRequest

    /* 页码, 默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小, 默认为10, 取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 证书渠道 (Optional) */
    CaType *string `json:"caType"`

    /* 个人用户姓名或企业名 (Optional) */
    Name *string `json:"name"`

    /* 证书序列号 (Optional) */
    SerialNo *string `json:"serialNo"`

    /* 证书算法 (Optional) */
    KeyAlg *string `json:"keyAlg"`

    /* 证书状态 (Optional) */
    CertStatus *int `json:"certStatus"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeCertListRequest(
) *DescribeCertListRequest {

	return &DescribeCertListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/smqCert:list",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param pageNumber: 页码, 默认为1 (Optional)
 * param pageSize: 分页大小, 默认为10, 取值范围[10, 100] (Optional)
 * param caType: 证书渠道 (Optional)
 * param name: 个人用户姓名或企业名 (Optional)
 * param serialNo: 证书序列号 (Optional)
 * param keyAlg: 证书算法 (Optional)
 * param certStatus: 证书状态 (Optional)
 */
func NewDescribeCertListRequestWithAllParams(
    pageNumber *int,
    pageSize *int,
    caType *string,
    name *string,
    serialNo *string,
    keyAlg *string,
    certStatus *int,
) *DescribeCertListRequest {

    return &DescribeCertListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/smqCert:list",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageNumber: pageNumber,
        PageSize: pageSize,
        CaType: caType,
        Name: name,
        SerialNo: serialNo,
        KeyAlg: keyAlg,
        CertStatus: certStatus,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeCertListRequestWithoutParam() *DescribeCertListRequest {

    return &DescribeCertListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/smqCert:list",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageNumber: 页码, 默认为1(Optional) */
func (r *DescribeCertListRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}
/* param pageSize: 分页大小, 默认为10, 取值范围[10, 100](Optional) */
func (r *DescribeCertListRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}
/* param caType: 证书渠道(Optional) */
func (r *DescribeCertListRequest) SetCaType(caType string) {
    r.CaType = &caType
}
/* param name: 个人用户姓名或企业名(Optional) */
func (r *DescribeCertListRequest) SetName(name string) {
    r.Name = &name
}
/* param serialNo: 证书序列号(Optional) */
func (r *DescribeCertListRequest) SetSerialNo(serialNo string) {
    r.SerialNo = &serialNo
}
/* param keyAlg: 证书算法(Optional) */
func (r *DescribeCertListRequest) SetKeyAlg(keyAlg string) {
    r.KeyAlg = &keyAlg
}
/* param certStatus: 证书状态(Optional) */
func (r *DescribeCertListRequest) SetCertStatus(certStatus int) {
    r.CertStatus = &certStatus
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeCertListRequest) GetRegionId() string {
    return ""
}

type DescribeCertListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeCertListResult `json:"result"`
}

type DescribeCertListResult struct {
    CertList []cloudsign.CertInfo `json:"certList"`
    TotalCount int `json:"totalCount"`
}