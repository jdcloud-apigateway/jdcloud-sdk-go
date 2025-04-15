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

type SendMessageRequest struct {

    core.JDCloudRequest

    /* 区域ID  */
    RegionId string `json:"regionId"`

    /* 实例id  */
    InstanceId string `json:"instanceId"`

    /* topic 名称  */
    Topic string `json:"topic"`

    /* 消息内容  */
    MessageBody string `json:"messageBody"`

    /* 消息tag (Optional) */
    Tag *string `json:"tag"`

    /* key id (Optional) */
    Key *string `json:"key"`

    /* 是否延时  */
    EnableDelay bool `json:"enableDelay"`

    /* 延时秒数 (Optional) */
    DelaySecond *int `json:"delaySecond"`
}

/*
 * param regionId: 区域ID (Required)
 * param instanceId: 实例id (Required)
 * param topic: topic 名称 (Required)
 * param messageBody: 消息内容 (Required)
 * param enableDelay: 是否延时 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSendMessageRequest(
    regionId string,
    instanceId string,
    topic string,
    messageBody string,
    enableDelay bool,
) *SendMessageRequest {

	return &SendMessageRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/topics/{topic}/messages",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        Topic: topic,
        MessageBody: messageBody,
        EnableDelay: enableDelay,
	}
}

/*
 * param regionId: 区域ID (Required)
 * param instanceId: 实例id (Required)
 * param topic: topic 名称 (Required)
 * param messageBody: 消息内容 (Required)
 * param tag: 消息tag (Optional)
 * param key: key id (Optional)
 * param enableDelay: 是否延时 (Required)
 * param delaySecond: 延时秒数 (Optional)
 */
func NewSendMessageRequestWithAllParams(
    regionId string,
    instanceId string,
    topic string,
    messageBody string,
    tag *string,
    key *string,
    enableDelay bool,
    delaySecond *int,
) *SendMessageRequest {

    return &SendMessageRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/topics/{topic}/messages",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        Topic: topic,
        MessageBody: messageBody,
        Tag: tag,
        Key: key,
        EnableDelay: enableDelay,
        DelaySecond: delaySecond,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSendMessageRequestWithoutParam() *SendMessageRequest {

    return &SendMessageRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/topics/{topic}/messages",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域ID(Required) */
func (r *SendMessageRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param instanceId: 实例id(Required) */
func (r *SendMessageRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}
/* param topic: topic 名称(Required) */
func (r *SendMessageRequest) SetTopic(topic string) {
    r.Topic = topic
}
/* param messageBody: 消息内容(Required) */
func (r *SendMessageRequest) SetMessageBody(messageBody string) {
    r.MessageBody = messageBody
}
/* param tag: 消息tag(Optional) */
func (r *SendMessageRequest) SetTag(tag string) {
    r.Tag = &tag
}
/* param key: key id(Optional) */
func (r *SendMessageRequest) SetKey(key string) {
    r.Key = &key
}
/* param enableDelay: 是否延时(Required) */
func (r *SendMessageRequest) SetEnableDelay(enableDelay bool) {
    r.EnableDelay = enableDelay
}
/* param delaySecond: 延时秒数(Optional) */
func (r *SendMessageRequest) SetDelaySecond(delaySecond int) {
    r.DelaySecond = &delaySecond
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SendMessageRequest) GetRegionId() string {
    return r.RegionId
}

type SendMessageResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SendMessageResult `json:"result"`
}

type SendMessageResult struct {
    MsgId string `json:"msgId"`
    SendStatus string `json:"sendStatus"`
}