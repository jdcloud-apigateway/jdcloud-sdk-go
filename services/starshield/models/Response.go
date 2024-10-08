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


type Response struct {

    /* 仅当规则action为block时可用。
允许您定义由于速率限制而阻止请求时,返回的HTTP响应体。
最大大小为30 KB。
 (Optional) */
    Content *string `json:"content"`

    /* 仅当规则action为block时可用。
允许您定义阻止请求时,响应的内容类型。
有效值application/json, text/html, text/xml, text/plain
 (Optional) */
    Content_type *string `json:"content_type"`

    /* 仅当规则action为block时可用。
允许您定义阻止请求时,返回给访问者的HTTP状态代码。
您必须输入一个介于400和499之间的值。
 (Optional) */
    Status_code *int `json:"status_code"`
}
