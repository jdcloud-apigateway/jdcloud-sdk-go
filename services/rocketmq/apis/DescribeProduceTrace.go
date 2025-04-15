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
    rocketmq "github.com/jdcloud-api/jdcloud-sdk-go/services/rocketmq/models"
)

type DescribeProduceTraceRequest struct {

    core.JDCloudRequest

    /* 区域ID  */
    RegionId string `json:"regionId"`

    /* 实例id  */
    InstanceId string `json:"instanceId"`

    /* topic 名称  */
    Topic string `json:"topic"`

    /* message Id  */
    MessageId string `json:"messageId"`

    /* 取决于客户端使用，如果未指定traceTopic,则不填 (Optional) */
    TraceTopic *string `json:"traceTopic"`
}

/*
 * param regionId: 区域ID (Required)
 * param instanceId: 实例id (Required)
 * param topic: topic 名称 (Required)
 * param messageId: message Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeProduceTraceRequest(
    regionId string,
    instanceId string,
    topic string,
    messageId string,
) *DescribeProduceTraceRequest {

	return &DescribeProduceTraceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/topics/{topic}/produceTrace/{messageId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        Topic: topic,
        MessageId: messageId,
	}
}

/*
 * param regionId: 区域ID (Required)
 * param instanceId: 实例id (Required)
 * param topic: topic 名称 (Required)
 * param messageId: message Id (Required)
 * param traceTopic: 取决于客户端使用，如果未指定traceTopic,则不填 (Optional)
 */
func NewDescribeProduceTraceRequestWithAllParams(
    regionId string,
    instanceId string,
    topic string,
    messageId string,
    traceTopic *string,
) *DescribeProduceTraceRequest {

    return &DescribeProduceTraceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/topics/{topic}/produceTrace/{messageId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        Topic: topic,
        MessageId: messageId,
        TraceTopic: traceTopic,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeProduceTraceRequestWithoutParam() *DescribeProduceTraceRequest {

    return &DescribeProduceTraceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/topics/{topic}/produceTrace/{messageId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域ID(Required) */
func (r *DescribeProduceTraceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param instanceId: 实例id(Required) */
func (r *DescribeProduceTraceRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}
/* param topic: topic 名称(Required) */
func (r *DescribeProduceTraceRequest) SetTopic(topic string) {
    r.Topic = topic
}
/* param messageId: message Id(Required) */
func (r *DescribeProduceTraceRequest) SetMessageId(messageId string) {
    r.MessageId = messageId
}
/* param traceTopic: 取决于客户端使用，如果未指定traceTopic,则不填(Optional) */
func (r *DescribeProduceTraceRequest) SetTraceTopic(traceTopic string) {
    r.TraceTopic = &traceTopic
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeProduceTraceRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeProduceTraceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeProduceTraceResult `json:"result"`
}

type DescribeProduceTraceResult struct {
    ProduceInfo rocketmq.ProduceInfo `json:"produceInfo"`
}