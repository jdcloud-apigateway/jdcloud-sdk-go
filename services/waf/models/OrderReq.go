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


type OrderReq struct {

    /* 地域信息, hb_bgp, hn, hd_bgp 企业版支持两个，旗舰版支持3个，多个以 , 分隔  */
    Region string `json:"region"`

    /* 购买类型, 1:创建 2:续费 3:升配  */
    BuyType int `json:"buyType"`

    /* 购买时长  */
    TimeSpan int `json:"timeSpan"`

    /* 时间单位，month, year  */
    TimeUnit string `json:"timeUnit"`

    /* 创建时间  */
    StartTime int64 `json:"startTime"`

    /* 实例id，除新建必传 (Optional) */
    WafInstanceId *string `json:"wafInstanceId"`

    /* 套餐类型 1:高级版, 2:企业版 3:旗舰版  */
    PackageType int `json:"packageType"`

    /* 额外的域名扩展包  */
    ExtraDomainsNum int `json:"extraDomainsNum"`

    /* 实例名，新建订单时必传，只能包含汉字、英文字母、下划线、破折号、数字，且长度不能超过16 (Optional) */
    NickName *string `json:"nickName"`

    /* 下单成功后返回的url, eg:http://abc.com  */
    ReturnURL string `json:"returnURL"`

    /* 额外的qps扩展包,单位为M 该值为50M的倍数  */
    ExtraBitsLimit int `json:"extraBitsLimit"`

    /* 云鼎的appCode (Optional) */
    AppCode *string `json:"appCode"`

    /* 云鼎的serviceCode (Optional) */
    ServiceCode *string `json:"serviceCode"`

    /* 购物车活动参数 (Optional) */
    BuyScenario *string `json:"buyScenario"`

    /* true表示支持autoPay (Optional) */
    NeedPay *bool `json:"needPay"`
}
