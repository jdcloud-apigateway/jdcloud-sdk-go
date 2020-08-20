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
    lb "github.com/jdcloud-api/jdcloud-sdk-go/services/lb/models"
)

type AddListenerCertificatesRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 监听器ID  */
    ListenerId string `json:"listenerId"`

    /* 【alb Https和Tls协议】ssl server证书列表  */
    Certificates []lb.ExtCertificateSpec `json:"certificates"`
}

/*
 * param regionId: Region ID (Required)
 * param listenerId: 监听器ID (Required)
 * param certificates: 【alb Https和Tls协议】ssl server证书列表 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddListenerCertificatesRequest(
    regionId string,
    listenerId string,
    certificates []lb.ExtCertificateSpec,
) *AddListenerCertificatesRequest {

	return &AddListenerCertificatesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/listeners/{listenerId}:addListenerCertificates",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ListenerId: listenerId,
        Certificates: certificates,
	}
}

/*
 * param regionId: Region ID (Required)
 * param listenerId: 监听器ID (Required)
 * param certificates: 【alb Https和Tls协议】ssl server证书列表 (Required)
 */
func NewAddListenerCertificatesRequestWithAllParams(
    regionId string,
    listenerId string,
    certificates []lb.ExtCertificateSpec,
) *AddListenerCertificatesRequest {

    return &AddListenerCertificatesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/listeners/{listenerId}:addListenerCertificates",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ListenerId: listenerId,
        Certificates: certificates,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddListenerCertificatesRequestWithoutParam() *AddListenerCertificatesRequest {

    return &AddListenerCertificatesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/listeners/{listenerId}:addListenerCertificates",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *AddListenerCertificatesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param listenerId: 监听器ID(Required) */
func (r *AddListenerCertificatesRequest) SetListenerId(listenerId string) {
    r.ListenerId = listenerId
}

/* param certificates: 【alb Https和Tls协议】ssl server证书列表(Required) */
func (r *AddListenerCertificatesRequest) SetCertificates(certificates []lb.ExtCertificateSpec) {
    r.Certificates = certificates
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddListenerCertificatesRequest) GetRegionId() string {
    return r.RegionId
}

type AddListenerCertificatesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddListenerCertificatesResult `json:"result"`
}

type AddListenerCertificatesResult struct {
}