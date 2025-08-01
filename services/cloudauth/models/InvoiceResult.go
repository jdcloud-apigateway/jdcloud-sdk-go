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


type InvoiceResult struct {

    /* 返回多页pdf或者tiff的情况下所在页码，若是单页返回1 (Optional) */
    Page_no int64 `json:"page_no"`

    /* 发票置信度 (Optional) */
    Score float64 `json:"score"`

    /* 增值税普 (Optional) */
    Page_name string `json:"page_name"`

    /* 发票用途码对应的中文名 (Optional) */
    Expense_type_msg string `json:"expense_type_msg"`

    /* 1收费，0不收费 (Optional) */
    Error_msg string `json:"error_msg"`

    /* 采集页面URL (Optional) */
    Qrcode_warning int64 `json:"qrcode_warning"`

    /* 增值税发票号码不一致告警 (Optional) */
    Invoice_no_warning int64 `json:"invoice_no_warning"`

    /* 发票的位置,依次为左上，右上，右下，左下顺时针四个点的x,y值 (Optional) */
    Quad []float64 `json:"quad"`

    /* 发票类型 (Optional) */
    Type string `json:"type"`

    /* 每个子图的识别结果 (Optional) */
    Recognize_result string `json:"recognize_result"`

    /* 发票用途 (Optional) */
    Expense_type string `json:"expense_type"`

    /* 如果设置do_query参数为1，将返回发票验证详细信息 (Optional) */
    Query_result string `json:"query_result"`

    /* 发票类型码对应的中文名 (Optional) */
    Type_msg string `json:"type_msg"`

    /* 错误码。 (Optional) */
    Error_code int64 `json:"error_code"`

    /* 根据post方式传入的参数决定是否返回子图，格式是base64。 (Optional) */
    Image string `json:"image"`

    /* 发票特殊标记 (Optional) */
    Pass_fee int64 `json:"pass_fee"`

    /* 合并标记 (Optional) */
    Merged_id float64 `json:"merged_id"`
}
