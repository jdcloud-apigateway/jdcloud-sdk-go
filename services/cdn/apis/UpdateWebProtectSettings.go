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

type UpdateWebProtectSettingsRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`

    /* 0：拦截模式 (阻断forbidden 493跳到自定义页面) ，1-检测模式(观察notice) (Optional) */
    WafMode *string `json:"wafMode"`

    /* 规则策略等级 0为宽松, 1为正常, 2为严格 (Optional) */
    WafLevel *int `json:"wafLevel"`

    /* 拦截模式跳转的自定义页面名称, 缺省或default返回默认页面 (Optional) */
    Redirection *string `json:"redirection"`
}

/*
 * param domain: 用户域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateWebProtectSettingsRequest(
    domain string,
) *UpdateWebProtectSettingsRequest {

	return &UpdateWebProtectSettingsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domain/{domain}/wafWebProtectSettings",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 用户域名 (Required)
 * param wafMode: 0：拦截模式 (阻断forbidden 493跳到自定义页面) ，1-检测模式(观察notice) (Optional)
 * param wafLevel: 规则策略等级 0为宽松, 1为正常, 2为严格 (Optional)
 * param redirection: 拦截模式跳转的自定义页面名称, 缺省或default返回默认页面 (Optional)
 */
func NewUpdateWebProtectSettingsRequestWithAllParams(
    domain string,
    wafMode *string,
    wafLevel *int,
    redirection *string,
) *UpdateWebProtectSettingsRequest {

    return &UpdateWebProtectSettingsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/wafWebProtectSettings",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
        WafMode: wafMode,
        WafLevel: wafLevel,
        Redirection: redirection,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateWebProtectSettingsRequestWithoutParam() *UpdateWebProtectSettingsRequest {

    return &UpdateWebProtectSettingsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/wafWebProtectSettings",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *UpdateWebProtectSettingsRequest) SetDomain(domain string) {
    r.Domain = domain
}
/* param wafMode: 0：拦截模式 (阻断forbidden 493跳到自定义页面) ，1-检测模式(观察notice)(Optional) */
func (r *UpdateWebProtectSettingsRequest) SetWafMode(wafMode string) {
    r.WafMode = &wafMode
}
/* param wafLevel: 规则策略等级 0为宽松, 1为正常, 2为严格(Optional) */
func (r *UpdateWebProtectSettingsRequest) SetWafLevel(wafLevel int) {
    r.WafLevel = &wafLevel
}
/* param redirection: 拦截模式跳转的自定义页面名称, 缺省或default返回默认页面(Optional) */
func (r *UpdateWebProtectSettingsRequest) SetRedirection(redirection string) {
    r.Redirection = &redirection
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateWebProtectSettingsRequest) GetRegionId() string {
    return ""
}

type UpdateWebProtectSettingsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateWebProtectSettingsResult `json:"result"`
}

type UpdateWebProtectSettingsResult struct {
}