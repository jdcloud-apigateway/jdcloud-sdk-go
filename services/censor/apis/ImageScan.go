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
    censor "github.com/jdcloud-api/jdcloud-sdk-go/services/censor/models"
)

type ImageScanRequest struct {

    core.JDCloudRequest

    /* 机审策略，eg: default (Optional) */
    BizType *string `json:"bizType"`

    /* 指定检测场景 (Optional) */
    Scenes []string `json:"scenes"`

    /* 检测任务列表，包含一个或多个元素。每个元素是个结构体，最多可添加10个元素，即最多对10段文本进行检测。每个元素的具体结构描述见ImageTask。 (Optional) */
    Tasks []censor.ImageTask `json:"tasks"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewImageScanRequest(
) *ImageScanRequest {

	return &ImageScanRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/image:scan",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param bizType: 机审策略，eg: default (Optional)
 * param scenes: 指定检测场景 (Optional)
 * param tasks: 检测任务列表，包含一个或多个元素。每个元素是个结构体，最多可添加10个元素，即最多对10段文本进行检测。每个元素的具体结构描述见ImageTask。 (Optional)
 */
func NewImageScanRequestWithAllParams(
    bizType *string,
    scenes []string,
    tasks []censor.ImageTask,
) *ImageScanRequest {

    return &ImageScanRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/image:scan",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        BizType: bizType,
        Scenes: scenes,
        Tasks: tasks,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewImageScanRequestWithoutParam() *ImageScanRequest {

    return &ImageScanRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/image:scan",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param bizType: 机审策略，eg: default(Optional) */
func (r *ImageScanRequest) SetBizType(bizType string) {
    r.BizType = &bizType
}

/* param scenes: 指定检测场景(Optional) */
func (r *ImageScanRequest) SetScenes(scenes []string) {
    r.Scenes = scenes
}

/* param tasks: 检测任务列表，包含一个或多个元素。每个元素是个结构体，最多可添加10个元素，即最多对10段文本进行检测。每个元素的具体结构描述见ImageTask。(Optional) */
func (r *ImageScanRequest) SetTasks(tasks []censor.ImageTask) {
    r.Tasks = tasks
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ImageScanRequest) GetRegionId() string {
    return ""
}

type ImageScanResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ImageScanResult `json:"result"`
}

type ImageScanResult struct {
    Data []censor.ImageResult `json:"data"`
}