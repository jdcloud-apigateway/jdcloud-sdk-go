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

type CreateCCProtectRuleRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`

    /* null (Optional) */
    Uri *string `json:"uri"`

    /* null (Optional) */
    DetectPeriod *int `json:"detectPeriod"`

    /* null (Optional) */
    SingleIpLimit *int `json:"singleIpLimit"`

    /* null (Optional) */
    BlockType *int `json:"blockType"`

    /* null (Optional) */
    BlockTime *int `json:"blockTime"`
}

/*
 * param domain: 用户域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateCCProtectRuleRequest(
    domain string,
) *CreateCCProtectRuleRequest {

	return &CreateCCProtectRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/domain/{domain}/ccProtectRule",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 用户域名 (Required)
 * param uri: null (Optional)
 * param detectPeriod: null (Optional)
 * param singleIpLimit: null (Optional)
 * param blockType: null (Optional)
 * param blockTime: null (Optional)
 */
func NewCreateCCProtectRuleRequestWithAllParams(
    domain string,
    uri *string,
    detectPeriod *int,
    singleIpLimit *int,
    blockType *int,
    blockTime *int,
) *CreateCCProtectRuleRequest {

    return &CreateCCProtectRuleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/ccProtectRule",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
        Uri: uri,
        DetectPeriod: detectPeriod,
        SingleIpLimit: singleIpLimit,
        BlockType: blockType,
        BlockTime: blockTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateCCProtectRuleRequestWithoutParam() *CreateCCProtectRuleRequest {

    return &CreateCCProtectRuleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/domain/{domain}/ccProtectRule",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *CreateCCProtectRuleRequest) SetDomain(domain string) {
    r.Domain = domain
}

/* param uri: null(Optional) */
func (r *CreateCCProtectRuleRequest) SetUri(uri string) {
    r.Uri = &uri
}

/* param detectPeriod: null(Optional) */
func (r *CreateCCProtectRuleRequest) SetDetectPeriod(detectPeriod int) {
    r.DetectPeriod = &detectPeriod
}

/* param singleIpLimit: null(Optional) */
func (r *CreateCCProtectRuleRequest) SetSingleIpLimit(singleIpLimit int) {
    r.SingleIpLimit = &singleIpLimit
}

/* param blockType: null(Optional) */
func (r *CreateCCProtectRuleRequest) SetBlockType(blockType int) {
    r.BlockType = &blockType
}

/* param blockTime: null(Optional) */
func (r *CreateCCProtectRuleRequest) SetBlockTime(blockTime int) {
    r.BlockTime = &blockTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateCCProtectRuleRequest) GetRegionId() string {
    return ""
}

type CreateCCProtectRuleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateCCProtectRuleResult `json:"result"`
}

type CreateCCProtectRuleResult struct {
}