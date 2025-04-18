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
    openjrtc "github.com/jdcloud-api/jdcloud-sdk-go/services/openjrtc/models"
)

type DescribeStreamInfosByUserIdRequest struct {

    core.JDCloudRequest

    /* 应用ID  */
    AppId string `json:"appId"`

    /* 业务接入方定义的且在JRTC系统内注册过的房间号  */
    UserRoomId string `json:"userRoomId"`

    /* 业务接入方用户体系定义的且在JRTC系统内注册过的userId  */
    UserId string `json:"userId"`
}

/*
 * param appId: 应用ID (Required)
 * param userRoomId: 业务接入方定义的且在JRTC系统内注册过的房间号 (Required)
 * param userId: 业务接入方用户体系定义的且在JRTC系统内注册过的userId (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeStreamInfosByUserIdRequest(
    appId string,
    userRoomId string,
    userId string,
) *DescribeStreamInfosByUserIdRequest {

	return &DescribeStreamInfosByUserIdRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/describeStreamInfosByUserId/{appId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        AppId: appId,
        UserRoomId: userRoomId,
        UserId: userId,
	}
}

/*
 * param appId: 应用ID (Required)
 * param userRoomId: 业务接入方定义的且在JRTC系统内注册过的房间号 (Required)
 * param userId: 业务接入方用户体系定义的且在JRTC系统内注册过的userId (Required)
 */
func NewDescribeStreamInfosByUserIdRequestWithAllParams(
    appId string,
    userRoomId string,
    userId string,
) *DescribeStreamInfosByUserIdRequest {

    return &DescribeStreamInfosByUserIdRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeStreamInfosByUserId/{appId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        AppId: appId,
        UserRoomId: userRoomId,
        UserId: userId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeStreamInfosByUserIdRequestWithoutParam() *DescribeStreamInfosByUserIdRequest {

    return &DescribeStreamInfosByUserIdRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/describeStreamInfosByUserId/{appId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param appId: 应用ID(Required) */
func (r *DescribeStreamInfosByUserIdRequest) SetAppId(appId string) {
    r.AppId = appId
}
/* param userRoomId: 业务接入方定义的且在JRTC系统内注册过的房间号(Required) */
func (r *DescribeStreamInfosByUserIdRequest) SetUserRoomId(userRoomId string) {
    r.UserRoomId = userRoomId
}
/* param userId: 业务接入方用户体系定义的且在JRTC系统内注册过的userId(Required) */
func (r *DescribeStreamInfosByUserIdRequest) SetUserId(userId string) {
    r.UserId = userId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeStreamInfosByUserIdRequest) GetRegionId() string {
    return ""
}

type DescribeStreamInfosByUserIdResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeStreamInfosByUserIdResult `json:"result"`
}

type DescribeStreamInfosByUserIdResult struct {
    Content []openjrtc.StreamInfo `json:"content"`
}