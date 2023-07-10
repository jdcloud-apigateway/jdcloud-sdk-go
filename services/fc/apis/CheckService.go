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

type CheckServiceRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Service Name  */
    ServiceName string `json:"serviceName"`
}

/*
 * param regionId: Region ID (Required)
 * param serviceName: Service Name (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCheckServiceRequest(
    regionId string,
    serviceName string,
) *CheckServiceRequest {

	return &CheckServiceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/services/{serviceName}:check",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ServiceName: serviceName,
	}
}

/*
 * param regionId: Region ID (Required)
 * param serviceName: Service Name (Required)
 */
func NewCheckServiceRequestWithAllParams(
    regionId string,
    serviceName string,
) *CheckServiceRequest {

    return &CheckServiceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/services/{serviceName}:check",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ServiceName: serviceName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCheckServiceRequestWithoutParam() *CheckServiceRequest {

    return &CheckServiceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/services/{serviceName}:check",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CheckServiceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param serviceName: Service Name(Required) */
func (r *CheckServiceRequest) SetServiceName(serviceName string) {
    r.ServiceName = serviceName
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CheckServiceRequest) GetRegionId() string {
    return r.RegionId
}

type CheckServiceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CheckServiceResult `json:"result"`
}

type CheckServiceResult struct {
    Code int `json:"code"`
    Reason string `json:"reason"`
}