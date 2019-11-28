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

type SetForwardAuthenticationRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`

    /* 鉴权类型，0表示无鉴权，1表示参数鉴权，2表示路径鉴权 (Optional) */
    AccesskeyType *int `json:"accesskeyType"`

    /* 密码，长度为8到32 (Optional) */
    AccesskeyKey *string `json:"accesskeyKey"`
}

/*
 * param domain: 用户域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSetForwardAuthenticationRequest(
    domain string,
) *SetForwardAuthenticationRequest {

	return &SetForwardAuthenticationRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/liveDomain/{domain}/forwardAuthentication",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 用户域名 (Required)
 * param accesskeyType: 鉴权类型，0表示无鉴权，1表示参数鉴权，2表示路径鉴权 (Optional)
 * param accesskeyKey: 密码，长度为8到32 (Optional)
 */
func NewSetForwardAuthenticationRequestWithAllParams(
    domain string,
    accesskeyType *int,
    accesskeyKey *string,
) *SetForwardAuthenticationRequest {

    return &SetForwardAuthenticationRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveDomain/{domain}/forwardAuthentication",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
        AccesskeyType: accesskeyType,
        AccesskeyKey: accesskeyKey,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSetForwardAuthenticationRequestWithoutParam() *SetForwardAuthenticationRequest {

    return &SetForwardAuthenticationRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveDomain/{domain}/forwardAuthentication",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *SetForwardAuthenticationRequest) SetDomain(domain string) {
    r.Domain = domain
}

/* param accesskeyType: 鉴权类型，0表示无鉴权，1表示参数鉴权，2表示路径鉴权(Optional) */
func (r *SetForwardAuthenticationRequest) SetAccesskeyType(accesskeyType int) {
    r.AccesskeyType = &accesskeyType
}

/* param accesskeyKey: 密码，长度为8到32(Optional) */
func (r *SetForwardAuthenticationRequest) SetAccesskeyKey(accesskeyKey string) {
    r.AccesskeyKey = &accesskeyKey
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SetForwardAuthenticationRequest) GetRegionId() string {
    return ""
}

type SetForwardAuthenticationResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SetForwardAuthenticationResult `json:"result"`
}

type SetForwardAuthenticationResult struct {
    IgnoreQueryString string `json:"ignoreQueryString"`
    PushIpWhiteList string `json:"pushIpWhiteList"`
    PublishNormalTimeout int `json:"publishNormalTimeout"`
    NotifyCustomUrl string `json:"notifyCustomUrl"`
    NotifyCustomAuthKey string `json:"notifyCustomAuthKey"`
    ForwardAccessKeyType int `json:"forwardAccessKeyType"`
    ForwardPrivateKey string `json:"forwardPrivateKey"`
    OriginAccessKeyType int `json:"originAccessKeyType"`
    OriginPrivateKey string `json:"originPrivateKey"`
}