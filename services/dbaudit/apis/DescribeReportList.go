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
    dbaudit "github.com/jdcloud-api/jdcloud-sdk-go/services/dbaudit/models"
)

type DescribeReportListRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 任务ID  */
    TaskId string `json:"taskId"`

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为10；取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 实例ID  */
    InstanceId string `json:"instanceId"`
}

/*
 * param regionId: 地域 Id (Required)
 * param taskId: 任务ID (Required)
 * param instanceId: 实例ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeReportListRequest(
    regionId string,
    taskId string,
    instanceId string,
) *DescribeReportListRequest {

	return &DescribeReportListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/tasks/{taskId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        TaskId: taskId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param taskId: 任务ID (Required)
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为10；取值范围[10, 100] (Optional)
 * param instanceId: 实例ID (Required)
 */
func NewDescribeReportListRequestWithAllParams(
    regionId string,
    taskId string,
    pageNumber *int,
    pageSize *int,
    instanceId string,
) *DescribeReportListRequest {

    return &DescribeReportListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/tasks/{taskId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        TaskId: taskId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        InstanceId: instanceId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeReportListRequestWithoutParam() *DescribeReportListRequest {

    return &DescribeReportListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/tasks/{taskId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *DescribeReportListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param taskId: 任务ID(Required) */
func (r *DescribeReportListRequest) SetTaskId(taskId string) {
    r.TaskId = taskId
}

/* param pageNumber: 页码；默认为1(Optional) */
func (r *DescribeReportListRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认为10；取值范围[10, 100](Optional) */
func (r *DescribeReportListRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param instanceId: 实例ID(Required) */
func (r *DescribeReportListRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeReportListRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeReportListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeReportListResult `json:"result"`
}

type DescribeReportListResult struct {
    TotalCount int `json:"totalCount"`
    List []dbaudit.ReportInfo `json:"list"`
}