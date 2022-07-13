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


type SnapshotJobSummary struct {

    /* job ID (Optional) */
    JobId string `json:"jobId"`

    /* 视频ID (Optional) */
    VideoId string `json:"videoId"`

    /* 任务状态。
- submitted
- processing
- succeeded
- failed
 (Optional) */
    Status string `json:"status"`

    /* 任务列表 (Optional) */
    Tasks []TaskSummary `json:"tasks"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 修改时间 (Optional) */
    UpdateTime string `json:"updateTime"`
}