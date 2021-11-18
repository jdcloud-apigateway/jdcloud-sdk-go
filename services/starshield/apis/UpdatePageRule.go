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
    starshield "github.com/jdcloud-api/jdcloud-sdk-go/services/starshield/models"
)

type UpdatePageRuleRequest struct {

    core.JDCloudRequest

    /*   */
    Zone_identifier string `json:"zone_identifier"`

    /*   */
    Identifier string `json:"identifier"`

    /* 根据请求评估的目标 (Optional) */
    Targets []starshield.Target `json:"targets"`

    /* 如果此规则的目标与请求匹配，则要执行的操作集。操作可以将url重定向到另一个url或覆盖设置（但不能同时覆盖两者） (Optional) */
    Actions []starshield.Action `json:"actions"`

    /* 一个数字，表示一个页面规则优先于另一个页面规则。
如果您可能有一个全面的页面规则（例如#1 “/images/”）
但是想要更具体的规则优先（例如#2 '/images/special/'），
您需要在后者（#2）上指定更高的优先级，以便它将覆盖第一个优先级。
 (Optional) */
    Priority *int `json:"priority"`

    /* 页面规则的状态 (Optional) */
    Status *string `json:"status"`
}

/*
 * param zone_identifier:  (Required)
 * param identifier:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdatePageRuleRequest(
    zone_identifier string,
    identifier string,
) *UpdatePageRuleRequest {

	return &UpdatePageRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/zones/{zone_identifier}/pagerules/{identifier}",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        Zone_identifier: zone_identifier,
        Identifier: identifier,
	}
}

/*
 * param zone_identifier:  (Required)
 * param identifier:  (Required)
 * param targets: 根据请求评估的目标 (Optional)
 * param actions: 如果此规则的目标与请求匹配，则要执行的操作集。操作可以将url重定向到另一个url或覆盖设置（但不能同时覆盖两者） (Optional)
 * param priority: 一个数字，表示一个页面规则优先于另一个页面规则。
如果您可能有一个全面的页面规则（例如#1 “/images/”）
但是想要更具体的规则优先（例如#2 '/images/special/'），
您需要在后者（#2）上指定更高的优先级，以便它将覆盖第一个优先级。
 (Optional)
 * param status: 页面规则的状态 (Optional)
 */
func NewUpdatePageRuleRequestWithAllParams(
    zone_identifier string,
    identifier string,
    targets []starshield.Target,
    actions []starshield.Action,
    priority *int,
    status *string,
) *UpdatePageRuleRequest {

    return &UpdatePageRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/pagerules/{identifier}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        Zone_identifier: zone_identifier,
        Identifier: identifier,
        Targets: targets,
        Actions: actions,
        Priority: priority,
        Status: status,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdatePageRuleRequestWithoutParam() *UpdatePageRuleRequest {

    return &UpdatePageRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/pagerules/{identifier}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param zone_identifier: (Required) */
func (r *UpdatePageRuleRequest) SetZone_identifier(zone_identifier string) {
    r.Zone_identifier = zone_identifier
}

/* param identifier: (Required) */
func (r *UpdatePageRuleRequest) SetIdentifier(identifier string) {
    r.Identifier = identifier
}

/* param targets: 根据请求评估的目标(Optional) */
func (r *UpdatePageRuleRequest) SetTargets(targets []starshield.Target) {
    r.Targets = targets
}

/* param actions: 如果此规则的目标与请求匹配，则要执行的操作集。操作可以将url重定向到另一个url或覆盖设置（但不能同时覆盖两者）(Optional) */
func (r *UpdatePageRuleRequest) SetActions(actions []starshield.Action) {
    r.Actions = actions
}

/* param priority: 一个数字，表示一个页面规则优先于另一个页面规则。
如果您可能有一个全面的页面规则（例如#1 “/images/”）
但是想要更具体的规则优先（例如#2 '/images/special/'），
您需要在后者（#2）上指定更高的优先级，以便它将覆盖第一个优先级。
(Optional) */
func (r *UpdatePageRuleRequest) SetPriority(priority int) {
    r.Priority = &priority
}

/* param status: 页面规则的状态(Optional) */
func (r *UpdatePageRuleRequest) SetStatus(status string) {
    r.Status = &status
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdatePageRuleRequest) GetRegionId() string {
    return ""
}

type UpdatePageRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdatePageRuleResult `json:"result"`
}

type UpdatePageRuleResult struct {
    Data starshield.PageRule `json:"data"`
}