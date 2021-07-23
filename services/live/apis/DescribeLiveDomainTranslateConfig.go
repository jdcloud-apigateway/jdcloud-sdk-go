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

type DescribeLiveDomainTranslateConfigRequest struct {

    core.JDCloudRequest

    /* 模板配置查询过滤条件:
  - name:   publishDomain 必填(推流域名)
  - value:  参数
  - name:   level 非必填(Level)
  - value:  参数，取值：domain,app,stream
 (Optional) */
    Filters []live.Filter `json:"filters"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeLiveDomainTranslateConfigRequest(
) *DescribeLiveDomainTranslateConfigRequest {

	return &DescribeLiveDomainTranslateConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/translateTemplates:domain",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param filters: 模板配置查询过滤条件:
  - name:   publishDomain 必填(推流域名)
  - value:  参数
  - name:   level 非必填(Level)
  - value:  参数，取值：domain,app,stream
 (Optional)
 */
func NewDescribeLiveDomainTranslateConfigRequestWithAllParams(
    filters []live.Filter,
) *DescribeLiveDomainTranslateConfigRequest {

    return &DescribeLiveDomainTranslateConfigRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/translateTemplates:domain",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeLiveDomainTranslateConfigRequestWithoutParam() *DescribeLiveDomainTranslateConfigRequest {

    return &DescribeLiveDomainTranslateConfigRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/translateTemplates:domain",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param filters: 模板配置查询过滤条件:
  - name:   publishDomain 必填(推流域名)
  - value:  参数
  - name:   level 非必填(Level)
  - value:  参数，取值：domain,app,stream
(Optional) */
func (r *DescribeLiveDomainTranslateConfigRequest) SetFilters(filters []live.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeLiveDomainTranslateConfigRequest) GetRegionId() string {
    return ""
}

type DescribeLiveDomainTranslateConfigResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeLiveDomainTranslateConfigResult `json:"result"`
}

type DescribeLiveDomainTranslateConfigResult struct {
    TemplateList []live.DomainTranslateTemplateConfig `json:"templateList"`
}