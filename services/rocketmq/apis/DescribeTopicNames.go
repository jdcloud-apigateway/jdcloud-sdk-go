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

type DescribeTopicNamesRequest struct {

    core.JDCloudRequest

    /* 区域id  */
    RegionId string `json:"regionId"`

    /* 实例id  */
    InstanceId string `json:"instanceId"`

    /* topic名称的过滤条件 (Optional) */
    TopicFilter *string `json:"topicFilter"`
}

/*
 * param regionId: 区域id (Required)
 * param instanceId: 实例id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeTopicNamesRequest(
    regionId string,
    instanceId string,
) *DescribeTopicNamesRequest {

	return &DescribeTopicNamesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/topicNames",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 区域id (Required)
 * param instanceId: 实例id (Required)
 * param topicFilter: topic名称的过滤条件 (Optional)
 */
func NewDescribeTopicNamesRequestWithAllParams(
    regionId string,
    instanceId string,
    topicFilter *string,
) *DescribeTopicNamesRequest {

    return &DescribeTopicNamesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/topicNames",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        TopicFilter: topicFilter,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeTopicNamesRequestWithoutParam() *DescribeTopicNamesRequest {

    return &DescribeTopicNamesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/topicNames",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域id(Required) */
func (r *DescribeTopicNamesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param instanceId: 实例id(Required) */
func (r *DescribeTopicNamesRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}
/* param topicFilter: topic名称的过滤条件(Optional) */
func (r *DescribeTopicNamesRequest) SetTopicFilter(topicFilter string) {
    r.TopicFilter = &topicFilter
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeTopicNamesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeTopicNamesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeTopicNamesResult `json:"result"`
}

type DescribeTopicNamesResult struct {
    TotalCount int `json:"totalCount"`
    TopicNameList []string `json:"topicNameList"`
}