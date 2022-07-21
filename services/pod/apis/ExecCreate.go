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

type ExecCreateRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Pod ID  */
    PodId string `json:"podId"`

    /* container name  */
    ContainerName string `json:"containerName"`

    /* 执行的命令  */
    Command []string `json:"command"`

    /* 执行命令是否分配tty。默认不分配 (Optional) */
    Tty *bool `json:"tty"`

}

/*
 * param regionId: Region ID (Required)
 * param podId: Pod ID (Required)
 * param containerName: container name (Required)
 * param command: 执行的命令 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewExecCreateRequest(
    regionId string,
    podId string,
    containerName string,
    command []string,
) *ExecCreateRequest {

	return &ExecCreateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/pods/{podId}/containers/{containerName}:execCreate",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        PodId: podId,
        ContainerName: containerName,
        Command: command,
	}
}

/*
 * param regionId: Region ID (Required)
 * param podId: Pod ID (Required)
 * param containerName: container name (Required)
 * param command: 执行的命令 (Required)
 * param tty: 执行命令是否分配tty。默认不分配 (Optional)
 */
func NewExecCreateRequestWithAllParams(
    regionId string,
    podId string,
    containerName string,
    command []string,
    tty *bool,
) *ExecCreateRequest {

    return &ExecCreateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/pods/{podId}/containers/{containerName}:execCreate",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PodId: podId,
        ContainerName: containerName,
        Command: command,
        Tty: tty,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewExecCreateRequestWithoutParam() *ExecCreateRequest {

    return &ExecCreateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/pods/{podId}/containers/{containerName}:execCreate",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *ExecCreateRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param podId: Pod ID(Required) */
func (r *ExecCreateRequest) SetPodId(podId string) {
    r.PodId = podId
}

/* param containerName: container name(Required) */
func (r *ExecCreateRequest) SetContainerName(containerName string) {
    r.ContainerName = containerName
}

/* param command: 执行的命令(Required) */
func (r *ExecCreateRequest) SetCommand(command []string) {
    r.Command = command
}

/* param tty: 执行命令是否分配tty。默认不分配(Optional) */
func (r *ExecCreateRequest) SetTty(tty bool) {
    r.Tty = &tty
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ExecCreateRequest) GetRegionId() string {
    return r.RegionId
}

type ExecCreateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ExecCreateResult `json:"result"`
}

type ExecCreateResult struct {
    ExecId string `json:"execId"`
}