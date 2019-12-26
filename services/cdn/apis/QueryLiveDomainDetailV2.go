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
    cdn "github.com/jdcloud-api/jdcloud-sdk-go/services/cdn/models"
)

type QueryLiveDomainDetailV2Request struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`
}

/*
 * param domain: 用户域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryLiveDomainDetailV2Request(
    domain string,
) *QueryLiveDomainDetailV2Request {

	return &QueryLiveDomainDetailV2Request{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/liveDomain/{domain}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 用户域名 (Required)
 */
func NewQueryLiveDomainDetailV2RequestWithAllParams(
    domain string,
) *QueryLiveDomainDetailV2Request {

    return &QueryLiveDomainDetailV2Request{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveDomain/{domain}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryLiveDomainDetailV2RequestWithoutParam() *QueryLiveDomainDetailV2Request {

    return &QueryLiveDomainDetailV2Request{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveDomain/{domain}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *QueryLiveDomainDetailV2Request) SetDomain(domain string) {
    r.Domain = domain
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryLiveDomainDetailV2Request) GetRegionId() string {
    return ""
}

type QueryLiveDomainDetailV2Response struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryLiveDomainDetailV2Result `json:"result"`
}

type QueryLiveDomainDetailV2Result struct {
    Id int `json:"id"`
    Domain string `json:"domain"`
    DomainType string `json:"domainType"`
    PublishDomain string `json:"publishDomain"`
    CreateTime string `json:"createTime"`
    Cname string `json:"cname"`
    SiteType string `json:"siteType"`
    PlayProtocol []string `json:"playProtocol"`
    Status string `json:"status"`
    Source cdn.SnowLeopardBackSourceInfo `json:"source"`
    SourceType string `json:"sourceType"`
    BackSourceType string `json:"backSourceType"`
    VideoType string `json:"videoType"`
    AudioType string `json:"audioType"`
    Type string `json:"type"`
    DefaultSourceHost string `json:"defaultSourceHost"`
    ArchiveNo string `json:"archiveNo"`
    ForwardCustomVhost string `json:"forwardCustomVhost"`
    FlvUrls []string `json:"flvUrls"`
    HlsUrls []string `json:"hlsUrls"`
    RtmpUrls []string `json:"rtmpUrls"`
    ProtocolConverts []cdn.ProtocolConvert `json:"protocolConverts"`
    Certificate string `json:"certificate"`
    RsaKey string `json:"rsaKey"`
    AccesskeyType int `json:"accesskeyType"`
    AccesskeyKey string `json:"accesskeyKey"`
    PlayAuthLifeTime int `json:"playAuthLifeTime"`
    AuthLifeTime int `json:"authLifeTime"`
    ForwardAccessKeyType int `json:"forwardAccessKeyType"`
    ForwardPrivateKey string `json:"forwardPrivateKey"`
    OriginAccessKeyType int `json:"originAccessKeyType"`
    OriginPrivateKey string `json:"originPrivateKey"`
    AllowApps []string `json:"allowApps"`
    Ips []string `json:"ips"`
    BlackIpsEnable string `json:"blackIpsEnable"`
    IgnoreQueryString string `json:"ignoreQueryString"`
    ReferType string `json:"referType"`
    ReferList []string `json:"referList"`
    AllowNoReferHeader string `json:"allowNoReferHeader"`
    AllowNullReferHeader string `json:"allowNullReferHeader"`
    PublishNormalTimeout int `json:"publishNormalTimeout"`
    NotifyCustomUrl string `json:"notifyCustomUrl"`
    NotifyCustomAuthKey string `json:"notifyCustomAuthKey"`
    CertFrom string `json:"certFrom"`
    SslCertId string `json:"sslCertId"`
    CertName string `json:"certName"`
    CertType string `json:"certType"`
    SslCertStartTime string `json:"sslCertStartTime"`
    SslCertEndTime string `json:"sslCertEndTime"`
    Digest string `json:"digest"`
    AccelerateRegion string `json:"accelerateRegion"`
}