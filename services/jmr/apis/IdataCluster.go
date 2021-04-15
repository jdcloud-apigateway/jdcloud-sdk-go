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

type IdataClusterRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 地域信息 (Optional) */
    DataCenter *string `json:"dataCenter"`
}

/*
 * param regionId: 地域ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewIdataClusterRequest(
    regionId string,
) *IdataClusterRequest {

	return &IdataClusterRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/idata",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param dataCenter: 地域信息 (Optional)
 */
func NewIdataClusterRequestWithAllParams(
    regionId string,
    dataCenter *string,
) *IdataClusterRequest {

    return &IdataClusterRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/idata",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DataCenter: dataCenter,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewIdataClusterRequestWithoutParam() *IdataClusterRequest {

    return &IdataClusterRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/idata",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *IdataClusterRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param dataCenter: 地域信息(Optional) */
func (r *IdataClusterRequest) SetDataCenter(dataCenter string) {
    r.DataCenter = &dataCenter
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r IdataClusterRequest) GetRegionId() string {
    return r.RegionId
}

type IdataClusterResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result IdataClusterResult `json:"result"`
}

type IdataClusterResult struct {
    Data interface{} `json:"data"`
    Status bool `json:"status"`
}