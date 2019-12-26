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

type DescribeQualityDetectionBindingRequest struct {

    core.JDCloudRequest

    /* 质量检测模板  */
    Template string `json:"template"`
}

/*
 * param template: 质量检测模板 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeQualityDetectionBindingRequest(
    template string,
) *DescribeQualityDetectionBindingRequest {

	return &DescribeQualityDetectionBindingRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/qualityDetectionTemplates/{template}:binding",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Template: template,
	}
}

/*
 * param template: 质量检测模板 (Required)
 */
func NewDescribeQualityDetectionBindingRequestWithAllParams(
    template string,
) *DescribeQualityDetectionBindingRequest {

    return &DescribeQualityDetectionBindingRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/qualityDetectionTemplates/{template}:binding",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Template: template,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeQualityDetectionBindingRequestWithoutParam() *DescribeQualityDetectionBindingRequest {

    return &DescribeQualityDetectionBindingRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/qualityDetectionTemplates/{template}:binding",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param template: 质量检测模板(Required) */
func (r *DescribeQualityDetectionBindingRequest) SetTemplate(template string) {
    r.Template = template
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeQualityDetectionBindingRequest) GetRegionId() string {
    return ""
}

type DescribeQualityDetectionBindingResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeQualityDetectionBindingResult `json:"result"`
}

type DescribeQualityDetectionBindingResult struct {
    BindingList []live.TemplateBinding `json:"bindingList"`
}