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
    live "github.com/jdcloud-api/jdcloud-sdk-go/services/live/models"
)

type DescribeLiveSnapshotsRequest struct {

    core.JDCloudRequest

    /* 当面页数
  */
    PageNum int `json:"pageNum"`

    /* 每页记录数
  */
    PageSize int `json:"pageSize"`

    /* 图片下载地址有效期，单位：秒，默认：3600
 (Optional) */
    AuthExpire *int `json:"authExpire"`

    /* 推流域名
  */
    PublishDomain string `json:"publishDomain"`

    /* 推流AppName
  */
    AppName string `json:"appName"`

    /* 流名称
  */
    StreamName string `json:"streamName"`

    /* 查询起始时间，UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'
  */
    StartTime string `json:"startTime"`

    /* 查询截至时间，UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'
  */
    EndTime string `json:"endTime"`
}

/*
 * param pageNum: 当面页数
 (Required)
 * param pageSize: 每页记录数
 (Required)
 * param publishDomain: 推流域名
 (Required)
 * param appName: 推流AppName
 (Required)
 * param streamName: 流名称
 (Required)
 * param startTime: 查询起始时间，UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'
 (Required)
 * param endTime: 查询截至时间，UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeLiveSnapshotsRequest(
    pageNum int,
    pageSize int,
    publishDomain string,
    appName string,
    streamName string,
    startTime string,
    endTime string,
) *DescribeLiveSnapshotsRequest {

	return &DescribeLiveSnapshotsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/snapshots",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        PageNum: pageNum,
        PageSize: pageSize,
        PublishDomain: publishDomain,
        AppName: appName,
        StreamName: streamName,
        StartTime: startTime,
        EndTime: endTime,
	}
}

/*
 * param pageNum: 当面页数
 (Required)
 * param pageSize: 每页记录数
 (Required)
 * param authExpire: 图片下载地址有效期，单位：秒，默认：3600
 (Optional)
 * param publishDomain: 推流域名
 (Required)
 * param appName: 推流AppName
 (Required)
 * param streamName: 流名称
 (Required)
 * param startTime: 查询起始时间，UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'
 (Required)
 * param endTime: 查询截至时间，UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'
 (Required)
 */
func NewDescribeLiveSnapshotsRequestWithAllParams(
    pageNum int,
    pageSize int,
    authExpire *int,
    publishDomain string,
    appName string,
    streamName string,
    startTime string,
    endTime string,
) *DescribeLiveSnapshotsRequest {

    return &DescribeLiveSnapshotsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/snapshots",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageNum: pageNum,
        PageSize: pageSize,
        AuthExpire: authExpire,
        PublishDomain: publishDomain,
        AppName: appName,
        StreamName: streamName,
        StartTime: startTime,
        EndTime: endTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeLiveSnapshotsRequestWithoutParam() *DescribeLiveSnapshotsRequest {

    return &DescribeLiveSnapshotsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/snapshots",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageNum: 当面页数
(Required) */
func (r *DescribeLiveSnapshotsRequest) SetPageNum(pageNum int) {
    r.PageNum = pageNum
}

/* param pageSize: 每页记录数
(Required) */
func (r *DescribeLiveSnapshotsRequest) SetPageSize(pageSize int) {
    r.PageSize = pageSize
}

/* param authExpire: 图片下载地址有效期，单位：秒，默认：3600
(Optional) */
func (r *DescribeLiveSnapshotsRequest) SetAuthExpire(authExpire int) {
    r.AuthExpire = &authExpire
}

/* param publishDomain: 推流域名
(Required) */
func (r *DescribeLiveSnapshotsRequest) SetPublishDomain(publishDomain string) {
    r.PublishDomain = publishDomain
}

/* param appName: 推流AppName
(Required) */
func (r *DescribeLiveSnapshotsRequest) SetAppName(appName string) {
    r.AppName = appName
}

/* param streamName: 流名称
(Required) */
func (r *DescribeLiveSnapshotsRequest) SetStreamName(streamName string) {
    r.StreamName = streamName
}

/* param startTime: 查询起始时间，UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'
(Required) */
func (r *DescribeLiveSnapshotsRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: 查询截至时间，UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'
(Required) */
func (r *DescribeLiveSnapshotsRequest) SetEndTime(endTime string) {
    r.EndTime = endTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeLiveSnapshotsRequest) GetRegionId() string {
    return ""
}

type DescribeLiveSnapshotsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeLiveSnapshotsResult `json:"result"`
}

type DescribeLiveSnapshotsResult struct {
    PageNumber int `json:"pageNumber"`
    PageSize int `json:"pageSize"`
    TotalCount int `json:"totalCount"`
    DomainDetails []live.Snapshot `json:"domainDetails"`
}