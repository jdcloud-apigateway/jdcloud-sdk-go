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

type DescribeDeploymentRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 分组ID  */
    ApiGroupId string `json:"apiGroupId"`

    /* 部署ID  */
    DeploymentId string `json:"deploymentId"`
}

/*
 * param regionId: 地域ID (Required)
 * param apiGroupId: 分组ID (Required)
 * param deploymentId: 部署ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeDeploymentRequest(
    regionId string,
    apiGroupId string,
    deploymentId string,
) *DescribeDeploymentRequest {

	return &DescribeDeploymentRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apiGroups/{apiGroupId}/deployments/{deploymentId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ApiGroupId: apiGroupId,
        DeploymentId: deploymentId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param apiGroupId: 分组ID (Required)
 * param deploymentId: 部署ID (Required)
 */
func NewDescribeDeploymentRequestWithAllParams(
    regionId string,
    apiGroupId string,
    deploymentId string,
) *DescribeDeploymentRequest {

    return &DescribeDeploymentRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apiGroups/{apiGroupId}/deployments/{deploymentId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ApiGroupId: apiGroupId,
        DeploymentId: deploymentId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeDeploymentRequestWithoutParam() *DescribeDeploymentRequest {

    return &DescribeDeploymentRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apiGroups/{apiGroupId}/deployments/{deploymentId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *DescribeDeploymentRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param apiGroupId: 分组ID(Required) */
func (r *DescribeDeploymentRequest) SetApiGroupId(apiGroupId string) {
    r.ApiGroupId = apiGroupId
}

/* param deploymentId: 部署ID(Required) */
func (r *DescribeDeploymentRequest) SetDeploymentId(deploymentId string) {
    r.DeploymentId = deploymentId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeDeploymentRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeDeploymentResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeDeploymentResult `json:"result"`
}

type DescribeDeploymentResult struct {
    DeploymentId string `json:"deploymentId"`
    Revision string `json:"revision"`
    Path string `json:"path"`
    Environment string `json:"environment"`
    BackendServiceType string `json:"backendServiceType"`
    BackendUrl string `json:"backendUrl"`
    Description string `json:"description"`
    CreateTime int64 `json:"createTime"`
    JdsfName string `json:"jdsfName"`
    JdsfRegistryName string `json:"jdsfRegistryName"`
    JdsfId string `json:"jdsfId"`
}