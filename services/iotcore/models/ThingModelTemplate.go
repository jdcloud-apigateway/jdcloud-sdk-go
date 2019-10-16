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


type ThingModelTemplate struct {

    /* 物模型模板ID (Optional) */
    ThingModelTemplateId string `json:"thingModelTemplateId"`

    /* 模型模板名称 (Optional) */
    ThingModelTemplateName string `json:"thingModelTemplateName"`

    /* 产品分类名称数组，索引0为一级产品分类名称 (Optional) */
    ProductCategoryNames []string `json:"productCategoryNames"`

    /* 产品分类ID数组，索引0为一级产品分类ID (Optional) */
    ProductCategoryIds []string `json:"productCategoryIds"`

    /* 更新时间，时间为东八区（UTC/GMT+08:00） (Optional) */
    UpdatedTime int64 `json:"updatedTime"`

    /* 操作人 (Optional) */
    UserPin string `json:"userPin"`

    /* 物模型文件在oss上的存储路径 (Optional) */
    OssPath string `json:"ossPath"`
}