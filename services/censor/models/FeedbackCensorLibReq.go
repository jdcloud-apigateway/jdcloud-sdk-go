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


type FeedbackCensorLibReq struct {

    /* task任务信息，最多可批量100个  */
    TaskInfos []FeedbackTaskInfo `json:"taskInfos"`

    /* 加黑时场景。文本/语音支持 antispam-反垃圾，视频/图片支持 porn-涉黄，terrorism-涉政暴恐。加白时为空 (Optional) */
    Scenes []string `json:"scenes"`

    /* 结果，pass-正常(加白)，blocl-违规(加黑)  */
    Suggestion string `json:"suggestion"`

    /* 文件类型，text-文本，image-图片，audio-音频，video-视频  */
    ResourceType string `json:"resourceType"`

    /* 检测类型，api/oss/website  */
    CensorType string `json:"censorType"`
}
