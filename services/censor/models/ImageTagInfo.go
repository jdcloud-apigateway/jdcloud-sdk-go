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


type ImageTagInfo struct {

    /* hinInfo中可返回的图片中包含的可识别内容 (Optional) */
    TagName string `json:"tagName"`

    /* tagName对应的分组名称，用于对tageName的解释 (Optional) */
    TagGroup string `json:"tagGroup"`
}