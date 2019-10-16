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

type CreateWatermarkRequest struct {

    core.JDCloudRequest

    /* 水印名称。只支持中英文、数字。长度不超过128个字符。UTF-8编码。
  */
    Name string `json:"name"`

    /* 图片地址  */
    ImgUrl string `json:"imgUrl"`

    /* 水印宽度。
当 sizeUnit = pixel 时，取值范围为 [8, 4096] 整数
当 sizeUnit = percent 时，取值范围为 [0, 100] 小数
  */
    Width string `json:"width"`

    /* 水印高度。
当 sizeUnit = pixel 时，取值范围为 [8, 4096] 整数
当 sizeUnit = percent 时，取值范围为 [0, 100] 小数
  */
    Height string `json:"height"`

    /* 尺寸单位。取值范围：
  pixel - 像素
  percent - 百分比
默认值为 pixel
 (Optional) */
    SizeUnit *string `json:"sizeUnit"`

    /* 水印位置。取值范围：
  LT - 左上
  RT - 右上
  LB - 左下
  RB - 右下
  */
    Position string `json:"position"`

    /* 水平偏移。
当 offsetUnit = pixel 时，取值范围为 [8, 4088] 整数
当 offsetUnit = percent 时，取值范围为 [0, 100] 小数
  */
    OffsetX string `json:"offsetX"`

    /* 竖直偏移。
当 offsetUnit = pixel 时，取值范围为 [8, 4088] 整数
当 offsetUnit = percent 时，取值范围为 [0, 100] 小数
  */
    OffsetY string `json:"offsetY"`

    /* 偏移单位。取值范围：
  pixel - 像素
  percent - 百分比
默认值为 pixel
 (Optional) */
    OffsetUnit *string `json:"offsetUnit"`
}

/*
 * param name: 水印名称。只支持中英文、数字。长度不超过128个字符。UTF-8编码。
 (Required)
 * param imgUrl: 图片地址 (Required)
 * param width: 水印宽度。
当 sizeUnit = pixel 时，取值范围为 [8, 4096] 整数
当 sizeUnit = percent 时，取值范围为 [0, 100] 小数
 (Required)
 * param height: 水印高度。
当 sizeUnit = pixel 时，取值范围为 [8, 4096] 整数
当 sizeUnit = percent 时，取值范围为 [0, 100] 小数
 (Required)
 * param position: 水印位置。取值范围：
  LT - 左上
  RT - 右上
  LB - 左下
  RB - 右下
 (Required)
 * param offsetX: 水平偏移。
当 offsetUnit = pixel 时，取值范围为 [8, 4088] 整数
当 offsetUnit = percent 时，取值范围为 [0, 100] 小数
 (Required)
 * param offsetY: 竖直偏移。
当 offsetUnit = pixel 时，取值范围为 [8, 4088] 整数
当 offsetUnit = percent 时，取值范围为 [0, 100] 小数
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateWatermarkRequest(
    name string,
    imgUrl string,
    width string,
    height string,
    position string,
    offsetX string,
    offsetY string,
) *CreateWatermarkRequest {

	return &CreateWatermarkRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/watermarks",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Name: name,
        ImgUrl: imgUrl,
        Width: width,
        Height: height,
        Position: position,
        OffsetX: offsetX,
        OffsetY: offsetY,
	}
}

/*
 * param name: 水印名称。只支持中英文、数字。长度不超过128个字符。UTF-8编码。
 (Required)
 * param imgUrl: 图片地址 (Required)
 * param width: 水印宽度。
当 sizeUnit = pixel 时，取值范围为 [8, 4096] 整数
当 sizeUnit = percent 时，取值范围为 [0, 100] 小数
 (Required)
 * param height: 水印高度。
当 sizeUnit = pixel 时，取值范围为 [8, 4096] 整数
当 sizeUnit = percent 时，取值范围为 [0, 100] 小数
 (Required)
 * param sizeUnit: 尺寸单位。取值范围：
  pixel - 像素
  percent - 百分比
默认值为 pixel
 (Optional)
 * param position: 水印位置。取值范围：
  LT - 左上
  RT - 右上
  LB - 左下
  RB - 右下
 (Required)
 * param offsetX: 水平偏移。
当 offsetUnit = pixel 时，取值范围为 [8, 4088] 整数
当 offsetUnit = percent 时，取值范围为 [0, 100] 小数
 (Required)
 * param offsetY: 竖直偏移。
当 offsetUnit = pixel 时，取值范围为 [8, 4088] 整数
当 offsetUnit = percent 时，取值范围为 [0, 100] 小数
 (Required)
 * param offsetUnit: 偏移单位。取值范围：
  pixel - 像素
  percent - 百分比
默认值为 pixel
 (Optional)
 */
func NewCreateWatermarkRequestWithAllParams(
    name string,
    imgUrl string,
    width string,
    height string,
    sizeUnit *string,
    position string,
    offsetX string,
    offsetY string,
    offsetUnit *string,
) *CreateWatermarkRequest {

    return &CreateWatermarkRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/watermarks",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Name: name,
        ImgUrl: imgUrl,
        Width: width,
        Height: height,
        SizeUnit: sizeUnit,
        Position: position,
        OffsetX: offsetX,
        OffsetY: offsetY,
        OffsetUnit: offsetUnit,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateWatermarkRequestWithoutParam() *CreateWatermarkRequest {

    return &CreateWatermarkRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/watermarks",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param name: 水印名称。只支持中英文、数字。长度不超过128个字符。UTF-8编码。
(Required) */
func (r *CreateWatermarkRequest) SetName(name string) {
    r.Name = name
}

/* param imgUrl: 图片地址(Required) */
func (r *CreateWatermarkRequest) SetImgUrl(imgUrl string) {
    r.ImgUrl = imgUrl
}

/* param width: 水印宽度。
当 sizeUnit = pixel 时，取值范围为 [8, 4096] 整数
当 sizeUnit = percent 时，取值范围为 [0, 100] 小数
(Required) */
func (r *CreateWatermarkRequest) SetWidth(width string) {
    r.Width = width
}

/* param height: 水印高度。
当 sizeUnit = pixel 时，取值范围为 [8, 4096] 整数
当 sizeUnit = percent 时，取值范围为 [0, 100] 小数
(Required) */
func (r *CreateWatermarkRequest) SetHeight(height string) {
    r.Height = height
}

/* param sizeUnit: 尺寸单位。取值范围：
  pixel - 像素
  percent - 百分比
默认值为 pixel
(Optional) */
func (r *CreateWatermarkRequest) SetSizeUnit(sizeUnit string) {
    r.SizeUnit = &sizeUnit
}

/* param position: 水印位置。取值范围：
  LT - 左上
  RT - 右上
  LB - 左下
  RB - 右下
(Required) */
func (r *CreateWatermarkRequest) SetPosition(position string) {
    r.Position = position
}

/* param offsetX: 水平偏移。
当 offsetUnit = pixel 时，取值范围为 [8, 4088] 整数
当 offsetUnit = percent 时，取值范围为 [0, 100] 小数
(Required) */
func (r *CreateWatermarkRequest) SetOffsetX(offsetX string) {
    r.OffsetX = offsetX
}

/* param offsetY: 竖直偏移。
当 offsetUnit = pixel 时，取值范围为 [8, 4088] 整数
当 offsetUnit = percent 时，取值范围为 [0, 100] 小数
(Required) */
func (r *CreateWatermarkRequest) SetOffsetY(offsetY string) {
    r.OffsetY = offsetY
}

/* param offsetUnit: 偏移单位。取值范围：
  pixel - 像素
  percent - 百分比
默认值为 pixel
(Optional) */
func (r *CreateWatermarkRequest) SetOffsetUnit(offsetUnit string) {
    r.OffsetUnit = &offsetUnit
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateWatermarkRequest) GetRegionId() string {
    return ""
}

type CreateWatermarkResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateWatermarkResult `json:"result"`
}

type CreateWatermarkResult struct {
    Id int64 `json:"id"`
    Name string `json:"name"`
    ImgUrl string `json:"imgUrl"`
    Width string `json:"width"`
    Height string `json:"height"`
    SizeUnit string `json:"sizeUnit"`
    Position string `json:"position"`
    OffsetX string `json:"offsetX"`
    OffsetY string `json:"offsetY"`
    OffsetUnit string `json:"offsetUnit"`
    CreateTime string `json:"createTime"`
    UpdateTime string `json:"updateTime"`
}