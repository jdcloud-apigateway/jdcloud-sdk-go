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


type ResourceCreateReq struct {

    /* 资源名称（支持中文、字母、数字、下划线，不超过50个字符） (Optional) */
    ResourceName string `json:"resourceName"`

    /* 原始资源名称（上传的原始文件在本地客户端的文件名称） (Optional) */
    OriginalName string `json:"originalName"`

    /* 父资源code（目录的根目录的父资源code为ROOT） (Optional) */
    ParentCode string `json:"parentCode"`

    /* 资源类型（DIRECTORY：目录；JAR：java的jar文件；FILE：其他普通文件；ARCHIVE：其他压缩文件/归档文件） (Optional) */
    ResourceType string `json:"resourceType"`

    /* 关联引擎（默认为JCW） (Optional) */
    RelativeEngine string `json:"relativeEngine"`

    /* 文件上传方式（默认为本地上传） (Optional) */
    UploadMode string `json:"uploadMode"`

    /* 环境信息（prod：生产环境；dev：开发环境），简单模式默认为prod (Optional) */
    Env string `json:"env"`

    /* 负责人 (Optional) */
    Managers []string `json:"managers"`
}
