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

package models


type VideoPlayInfo struct {

    /* 生成播放信息的转码任务ID (Optional) */
    TaskId string `json:"taskId"`

    /* 清晰度规格标记。取值范围：
  SD - 标清
  HD - 高清
  FHD - 超清
  2K
  4K
 (Optional) */
    Definition string `json:"definition"`

    /* 媒体类型 (Optional) */
    MediaType int `json:"mediaType"`

    /* 播放信息状态，目前只有正常状态(normal) (Optional) */
    Status string `json:"status"`

    /* CDN地址，原始地址或者鉴权地址 (Optional) */
    Url string `json:"url"`

    /*  (Optional) */
    Size int64 `json:"size"`

    /* 视频时长 (Optional) */
    Duration int64 `json:"duration"`

    /* 码率 (Optional) */
    Bitrate int64 `json:"bitrate"`

    /* 编码格式 (Optional) */
    Codec string `json:"codec"`

    /* 封装格式 (Optional) */
    Format string `json:"format"`

    /* 视频宽度 (Optional) */
    Width int `json:"width"`

    /* 视频高度 (Optional) */
    Height int `json:"height"`

    /* 视频帧率 (Optional) */
    Fps string `json:"fps"`

    /*  (Optional) */
    CreateTime string `json:"createTime"`

    /*  (Optional) */
    UpdateTime string `json:"updateTime"`
}