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
    aisearch "github.com/jdcloud-api/jdcloud-sdk-go/services/aisearch/models"
)

type ExternalWebSearchRequest struct {

    core.JDCloudRequest

    /* api key编号  */
    ApiKey string `json:"apiKey"`

    /* 请求id，api key下唯一  */
    RequestId string `json:"requestId"`

    /* 查询内容  */
    Query string `json:"query"`

    /* page 默认1 (Optional) */
    Page *int `json:"page"`

    /* pageSize 默认10 (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param apiKey: api key编号 (Required)
 * param requestId: 请求id，api key下唯一 (Required)
 * param query: 查询内容 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewExternalWebSearchRequest(
    apiKey string,
    requestId string,
    query string,
) *ExternalWebSearchRequest {

	return &ExternalWebSearchRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/external:webSearch",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        ApiKey: apiKey,
        RequestId: requestId,
        Query: query,
	}
}

/*
 * param apiKey: api key编号 (Required)
 * param requestId: 请求id，api key下唯一 (Required)
 * param query: 查询内容 (Required)
 * param page: page 默认1 (Optional)
 * param pageSize: pageSize 默认10 (Optional)
 */
func NewExternalWebSearchRequestWithAllParams(
    apiKey string,
    requestId string,
    query string,
    page *int,
    pageSize *int,
) *ExternalWebSearchRequest {

    return &ExternalWebSearchRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/external:webSearch",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        ApiKey: apiKey,
        RequestId: requestId,
        Query: query,
        Page: page,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewExternalWebSearchRequestWithoutParam() *ExternalWebSearchRequest {

    return &ExternalWebSearchRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/external:webSearch",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param apiKey: api key编号(Required) */
func (r *ExternalWebSearchRequest) SetApiKey(apiKey string) {
    r.ApiKey = apiKey
}
/* param requestId: 请求id，api key下唯一(Required) */
func (r *ExternalWebSearchRequest) SetRequestId(requestId string) {
    r.RequestId = requestId
}
/* param query: 查询内容(Required) */
func (r *ExternalWebSearchRequest) SetQuery(query string) {
    r.Query = query
}
/* param page: page 默认1(Optional) */
func (r *ExternalWebSearchRequest) SetPage(page int) {
    r.Page = &page
}
/* param pageSize: pageSize 默认10(Optional) */
func (r *ExternalWebSearchRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ExternalWebSearchRequest) GetRegionId() string {
    return ""
}

type ExternalWebSearchResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ExternalWebSearchResult `json:"result"`
}

type ExternalWebSearchResult struct {
    RequestId string `json:"requestId"`
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data aisearch.CloudWebSearchResponseVo `json:"data"`
}