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

type PushRequest struct {

    core.JDCloudRequest

    /* 日志主题uid  */
    LogtopicUID string `json:"logtopicUID"`
}

/*
 * param logtopicUID: 日志主题uid (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewPushRequest(
    logtopicUID string,
) *PushRequest {

	return &PushRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/logtopics/{logtopicUID}:push",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        LogtopicUID: logtopicUID,
	}
}

/*
 * param logtopicUID: 日志主题uid (Required)
 */
func NewPushRequestWithAllParams(
    logtopicUID string,
) *PushRequest {

    return &PushRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/logtopics/{logtopicUID}:push",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        LogtopicUID: logtopicUID,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewPushRequestWithoutParam() *PushRequest {

    return &PushRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/logtopics/{logtopicUID}:push",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param logtopicUID: 日志主题uid(Required) */
func (r *PushRequest) SetLogtopicUID(logtopicUID string) {
    r.LogtopicUID = logtopicUID
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r PushRequest) GetRegionId() string {
    return ""
}

type PushResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result PushResult `json:"result"`
}

type PushResult struct {
}