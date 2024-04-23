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
    lb "github.com/jdcloud-api/jdcloud-sdk-go/services/lb/models"
)

type CreateCustomizedConfigurationRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 个性化配置名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符  */
    CustomizedConfigurationName string `json:"customizedConfigurationName"`

    /* 描述,允许输入UTF-8编码下的全部字符，不超过256字符 (Optional) */
    Description *string `json:"description"`

    /* 个性化配置内容 (Optional) */
    CustomizedConfigurationContent *lb.CustomizedConfigurationContentSpec `json:"customizedConfigurationContent"`
}

/*
 * param regionId: Region ID (Required)
 * param customizedConfigurationName: 个性化配置名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateCustomizedConfigurationRequest(
    regionId string,
    customizedConfigurationName string,
) *CreateCustomizedConfigurationRequest {

	return &CreateCustomizedConfigurationRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/customizedConfigurations/",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        CustomizedConfigurationName: customizedConfigurationName,
	}
}

/*
 * param regionId: Region ID (Required)
 * param customizedConfigurationName: 个性化配置名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符 (Required)
 * param description: 描述,允许输入UTF-8编码下的全部字符，不超过256字符 (Optional)
 * param customizedConfigurationContent: 个性化配置内容 (Optional)
 */
func NewCreateCustomizedConfigurationRequestWithAllParams(
    regionId string,
    customizedConfigurationName string,
    description *string,
    customizedConfigurationContent *lb.CustomizedConfigurationContentSpec,
) *CreateCustomizedConfigurationRequest {

    return &CreateCustomizedConfigurationRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/customizedConfigurations/",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        CustomizedConfigurationName: customizedConfigurationName,
        Description: description,
        CustomizedConfigurationContent: customizedConfigurationContent,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateCustomizedConfigurationRequestWithoutParam() *CreateCustomizedConfigurationRequest {

    return &CreateCustomizedConfigurationRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/customizedConfigurations/",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateCustomizedConfigurationRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param customizedConfigurationName: 个性化配置名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符(Required) */
func (r *CreateCustomizedConfigurationRequest) SetCustomizedConfigurationName(customizedConfigurationName string) {
    r.CustomizedConfigurationName = customizedConfigurationName
}
/* param description: 描述,允许输入UTF-8编码下的全部字符，不超过256字符(Optional) */
func (r *CreateCustomizedConfigurationRequest) SetDescription(description string) {
    r.Description = &description
}
/* param customizedConfigurationContent: 个性化配置内容(Optional) */
func (r *CreateCustomizedConfigurationRequest) SetCustomizedConfigurationContent(customizedConfigurationContent *lb.CustomizedConfigurationContentSpec) {
    r.CustomizedConfigurationContent = customizedConfigurationContent
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateCustomizedConfigurationRequest) GetRegionId() string {
    return r.RegionId
}

type CreateCustomizedConfigurationResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateCustomizedConfigurationResult `json:"result"`
}

type CreateCustomizedConfigurationResult struct {
    CustomizedConfigurationId string `json:"customizedConfigurationId"`
}