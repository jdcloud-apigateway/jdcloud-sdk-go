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


type CreditTaskMarketingDetail struct {

    /* 注册手机号，国内手机：11位手机号;海外手机：以+号开头，4位国家代码+5-11位手机号扩展位；手机注册，必填 (Optional) */
    Phone string `json:"phone"`

    /* 参与活动IP,用户领取奖励时的真实外网 IP（非服务端），IPV4 或 IPV6 (Optional) */
    Ip string `json:"ip"`

    /* 参与活动时间戳，参与活动时间戳，UNIX时间戳 (Optional) */
    Time int64 `json:"time"`

    /* 用户参加活动渠道或终端，1：PC端web浏览器注册 PC-Brower；2：PC客户端注册 PC-Client；3：移动设备各种APP注册 Mobile-APP；4 ：移动设备浏览器登录，移动端M页注册 Mobile-Brower；5：移动设备微信客户端中购物入口的注册操作 Mobile-WX；6： 移动设备QQ客户端中购物入口的注册操作 Mobile-QQ；7： 移动设备微信客户端中微信小程序注册操作- Mobile-WX；0：其他 (Optional) */
    Channel int `json:"channel"`

    /* 小写referUrl (Optional) */
    ReferUrlLower string `json:"referUrlLower"`

    /* 登录方式，1：手动帐号密码输入；2：动态短信密码登录；3：二维码扫描登录；0：其他 (Optional) */
    LoginType int `json:"loginType"`

    /* 登录时间，用户最近登录完成时间，UNIX时间戳 (Optional) */
    LastLoginTime int64 `json:"lastLoginTime"`

    /* 注册时间，UNIX时间戳 (Optional) */
    RegTime int `json:"regTime"`

    /* 注册来源的外网 IP，IPV4 或 IPV6 (Optional) */
    RegIp string `json:"regIp"`

    /* 注册类型，1：手机注册；2：邮箱注册；3：pin注册；0：其他。 (Optional) */
    RegType int `json:"regType"`

    /* 注册帐号名，用户注册使用名称 (Optional) */
    RegName string `json:"regName"`

    /* 注册渠道或注册终端，1： PC端web浏览器注册 PC-Brower；2：PC客户端注册 PC-Client；3：移动设备各种APP注册 Mobile-APP；4 ：移动设备浏览器登录，移动端M页注册 Mobile-Brower；5：移动设备微信客户端中购物入口的注册操作 Mobile-WX；6： 移动设备QQ客户端中购物入口的注册操作 Mobile-QQ；7： 移动设备微信客户端中微信小程序注册操作- Mobile-WX；0：其他 (Optional) */
    RegChannel int `json:"regChannel"`

    /* 参加活动设备uid，UID是使用iOS系统非隐私参数，用一套统一规则生成的用于标识苹果手机的ID (Optional) */
    Uid string `json:"uid"`

    /* 参与活动登录的设备号，设备指纹编码 (Optional) */
    Eid string `json:"eid"`

    /* MAC地址，MAC 地址或设备唯一标识。 (Optional) */
    MacAddress string `json:"macAddress"`

    /* 手机制造商 ID，手机制造商 ID，如果手机注册，请带上此信息。 (Optional) */
    VendorId string `json:"vendorId"`

    /* 手机设备号，Android：imei，IOS：idfa (Optional) */
    Imei string `json:"imei"`

    /* 手机设备号，Android：imei，IOS：idfa (Optional) */
    Idfa string `json:"idfa"`

    /* App 客户端版本，如果使用App操作，请带上此信息 (Optional) */
    AppVersion string `json:"appVersion"`

    /* 业务 ID， 网站或应用在多个业务中使用此服务，通过此 ID 区分统计数据 (Optional) */
    BusinessId int `json:"businessId"`

    /* 对于新人的类型做分类，用于新人权益互斥的场景做领券限制，不同业务场景的解释会有变化。例如：1001：新人188大礼包；1002：全链路新人礼包；1003：市场部新人；1004： 极速版拉新；1005：事业部拉新 (Optional) */
    NewPersonType int `json:"newPersonType"`

    /* 优惠券ID (Optional) */
    BatchId string `json:"batchId"`

    /* 活动key (Optional) */
    ActivityKeyRaw string `json:"activityKeyRaw"`

    /* 业务来源，基础账号侧配置的业务来源，用来识别和区分独立业务，枚举可根据客户具体需求调整。 (Optional) */
    Source int `json:"source"`

    /* cookie 的Hash值，用户 HTTP 请求中的 cookie 进行2次 hash 的值，只要保证相同 cookie 的 hash 值一致即可。 (Optional) */
    CookieHash string `json:"cookieHash"`

    /* 用户领取奖品邮寄地址 (Optional) */
    Address string `json:"address"`

    /* 用户 HTTP 请求的 userAgent (Optional) */
    UserAgent string `json:"userAgent"`

    /* 用户 HTTP 请求中的 x_forward_for。 (Optional) */
    XForwardedFor string `json:"xForwardedFor"`

    /* 用户操作过程中鼠标单击次数。 (Optional) */
    MouseClickCount int `json:"mouseClickCount"`

    /* 用户操作过程中键盘单击次数。 (Optional) */
    KeyboardClickCount int `json:"keyboardClickCount"`

    /* 登录耗时，从出登录页到登录提交之间的时间差（出登录视图埋点，提交时计算时间差），如果为免密码登录方式，可以在换取认证token时生成时间戳，验证token时计算时间差，单位ms (Optional) */
    LoginSpend int `json:"loginSpend"`

    /* 最后登录设备号（eid），设备指纹编码 (Optional) */
    LastLoginEid string `json:"lastLoginEid"`

    /* 登录成功后跳转页面。 (Optional) */
    JumpUrl string `json:"jumpUrl"`

    /* 注册占用时长，从用户进入注册页到点击注册提交之间的时间差，单位ms (Optional) */
    ElapsedTime string `json:"elapsedTime"`

    /* 注册结果，成功 or 失败；如直接做拦截校验可不填，1：注册成功；2：注册失败。 (Optional) */
    RegResult int `json:"regResult"`

    /* 用户注册邮箱 (Optional) */
    RegEmail string `json:"regEmail"`

    /* 单个红包允许领取的用户数量（分享红包） (Optional) */
    Share int `json:"share"`

    /* 单日内，单个账号每日领取奖励的最大次数。 (Optional) */
    DayTimes int `json:"dayTimes"`

    /* 整个活动周期内，单个账号能领取奖励的最大次数 (Optional) */
    TotalTimes int `json:"totalTimes"`

    /* 维度。浮点数，范围为-90 - 90 (Optional) */
    Atitude int `json:"atitude"`

    /* 经度。浮点数，范围为-180 - 180 (Optional) */
    Longitude int `json:"longitude"`
}