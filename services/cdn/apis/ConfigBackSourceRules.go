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
    cdn "github.com/jdcloud-api/jdcloud-sdk-go/services/cdn/models"
)

type ConfigBackSourceRulesRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`

    /* on/off，若为off则数组需为空，若为on则数组不可为空 (Optional) */
    Status *string `json:"status"`

    /*  (Optional) */
    Rules []cdn.BackSourceRule `json:"rules"`
}

/*
 * param domain: 用户域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewConfigBackSourceRulesRequest(
    domain string,
) *ConfigBackSourceRulesRequest {

	return &ConfigBackSourceRulesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domain/{domain}/configBackSourceRules",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 用户域名 (Required)
 * param status: on/off，若为off则数组需为空，若为on则数组不可为空 (Optional)
 * param rules:  (Optional)
 */
func NewConfigBackSourceRulesRequestWithAllParams(
    domain string,
    status *string,
    rules []cdn.BackSourceRule,
) *ConfigBackSourceRulesRequest {

    return &ConfigBackSourceRulesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/configBackSourceRules",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
        Status: status,
        Rules: rules,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewConfigBackSourceRulesRequestWithoutParam() *ConfigBackSourceRulesRequest {

    return &ConfigBackSourceRulesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/configBackSourceRules",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *ConfigBackSourceRulesRequest) SetDomain(domain string) {
    r.Domain = domain
}
/* param status: on/off，若为off则数组需为空，若为on则数组不可为空(Optional) */
func (r *ConfigBackSourceRulesRequest) SetStatus(status string) {
    r.Status = &status
}
/* param rules: (Optional) */
func (r *ConfigBackSourceRulesRequest) SetRules(rules []cdn.BackSourceRule) {
    r.Rules = rules
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ConfigBackSourceRulesRequest) GetRegionId() string {
    return ""
}

type ConfigBackSourceRulesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ConfigBackSourceRulesResult `json:"result"`
}

type ConfigBackSourceRulesResult struct {
}