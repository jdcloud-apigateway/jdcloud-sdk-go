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


type SnapshotTemplateSampleConfigInfo struct {

    /* 截图开始时间，单位毫秒，缺省值为 0 (Optional) */
    StartTime int64 `json:"startTime"`

    /* 截图间隔时间，单位毫秒，缺省值为 0 (Optional) */
    IntervalTime int64 `json:"intervalTime"`

    /* 最大截图张数，取值必须大于 0，缺省值为 10
若 intervalTime 为 0，则从 startTime 开始，平均截取 number 张图片；
若 intervalTime 不为 0, 则从 startTime 开始，依次间隔截图，最多截取 number 张图片；
 (Optional) */
    Number int `json:"number"`

    /* 截图存储格式。取值范围：png、jpg (Optional) */
    Format string `json:"format"`

    /* 截图帧类型。取值范围：any、key (Optional) */
    FrameType string `json:"frameType"`

    /* 截图宽度，单位 px，缺省值为 0 (Optional) */
    Width int `json:"width"`

    /* 截图高度，单位 px，缺省值为 0
若 width、height 都不为 0，则截图分辨率为 width * height
若 width、height 某一项为 0，则该项随另一项等比缩放
若 width、height 两项都为 0, 则截图分辨率随源视频
 (Optional) */
    Height int `json:"height"`

    /* 填充方式。取值范围：
- stretch 伸缩
- black 留黑
- white 留白
- gauss 高斯模糊
 (Optional) */
    FillType string `json:"fillType"`
}
