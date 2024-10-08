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

type InstanceTrafficDateHistogram4PaRequest struct {

    core.JDCloudRequest

    /* 实例标识  */
    InstanceId string `json:"instanceId"`

    /* 开始时间  */
    Since string `json:"since"`

    /* 结束时间  */
    Until string `json:"until"`
}

/*
 * param instanceId: 实例标识 (Required)
 * param since: 开始时间 (Required)
 * param until: 结束时间 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewInstanceTrafficDateHistogram4PaRequest(
    instanceId string,
    since string,
    until string,
) *InstanceTrafficDateHistogram4PaRequest {

	return &InstanceTrafficDateHistogram4PaRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/instances/{instanceId}/paTrafficDateHistogram",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        InstanceId: instanceId,
        Since: since,
        Until: until,
	}
}

/*
 * param instanceId: 实例标识 (Required)
 * param since: 开始时间 (Required)
 * param until: 结束时间 (Required)
 */
func NewInstanceTrafficDateHistogram4PaRequestWithAllParams(
    instanceId string,
    since string,
    until string,
) *InstanceTrafficDateHistogram4PaRequest {

    return &InstanceTrafficDateHistogram4PaRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/instances/{instanceId}/paTrafficDateHistogram",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        InstanceId: instanceId,
        Since: since,
        Until: until,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewInstanceTrafficDateHistogram4PaRequestWithoutParam() *InstanceTrafficDateHistogram4PaRequest {

    return &InstanceTrafficDateHistogram4PaRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/instances/{instanceId}/paTrafficDateHistogram",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param instanceId: 实例标识(Required) */
func (r *InstanceTrafficDateHistogram4PaRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}
/* param since: 开始时间(Required) */
func (r *InstanceTrafficDateHistogram4PaRequest) SetSince(since string) {
    r.Since = since
}
/* param until: 结束时间(Required) */
func (r *InstanceTrafficDateHistogram4PaRequest) SetUntil(until string) {
    r.Until = until
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r InstanceTrafficDateHistogram4PaRequest) GetRegionId() string {
    return ""
}

type InstanceTrafficDateHistogram4PaResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result InstanceTrafficDateHistogram4PaResult `json:"result"`
}

type InstanceTrafficDateHistogram4PaResult struct {
    Sum float64 `json:"sum"`
    Max float64 `json:"max"`
    MaxTimestamp int `json:"maxTimestamp"`
    DataSeries []float64 `json:"dataSeries"`
    TimeSeries []int `json:"timeSeries"`
}