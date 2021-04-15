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


type BindWebsiteInfo struct {

    /* 协议，http或https (Optional) */
    Scheme string `json:"scheme"`

    /* 域名，不含协议 (Optional) */
    Host string `json:"host"`

    /* 全站检测频率，1d-1天，7d-7天 (Optional) */
    HostInterval string `json:"hostInterval"`

    /* 完整的首页地址 (Optional) */
    HomePage string `json:"homePage"`

    /* 首页检测间隔，eg:1h表示一小时，必须为整小时 (Optional) */
    HomePageInterval string `json:"homePageInterval"`
}