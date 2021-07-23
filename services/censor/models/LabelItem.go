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


type LabelItem struct {

    /* 分类信息，100：色情，200：广告，400：违禁，500：涉政，600：谩骂，700：灌水，900：其他 (Optional) */
    Label int `json:"label"`

    /* 分类级别，1：不确定，2：确定 (Optional) */
    Level int `json:"level"`

    /*  (Optional) */
    Details LabelItemDetail `json:"details"`

    /*  (Optional) */
    SubLabels []SubLabelItem `json:"subLabels"`
}