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


type Listener struct {

    /* Listener的Id (Optional) */
    ListenerId string `json:"listenerId"`

    /* Listener的名称 (Optional) */
    ListenerName string `json:"listenerName"`

    /* Listener状态, 取值为On或者为Off (Optional) */
    Status string `json:"status"`

    /* Listener所属loadBalancer的Id (Optional) */
    LoadBalancerId string `json:"loadBalancerId"`

    /* Listener所属负载均衡类型，取值为：alb、nlb、dnlb (Optional) */
    LoadBalancerType string `json:"loadBalancerType"`

    /* 监听协议, 取值为Tcp, Tls, Http, Https, Udp <br>【alb】支持Http, Https, Tcp, Tls和Udp <br>【nlb】支持Tcp, Udp  <br>【dnlb】支持Tcp, Udp (Optional) */
    Protocol string `json:"protocol"`

    /* 【alb使用https时支持】是否开启HSTS，True(开启)， False(关闭) (Optional) */
    HstsEnable bool `json:"hstsEnable"`

    /* 【alb使用https时支持】HSTS过期时间(秒)，取值范围为[1, 94608000(3年)] (Optional) */
    HstsMaxAge int `json:"hstsMaxAge"`

    /* 监听端口，取值范围为[1, 65535] (Optional) */
    Port int `json:"port"`

    /* 默认后端服务的转发策略,取值为Forward或Redirect, 现只支持Forward (Optional) */
    Action string `json:"action"`

    /* 默认的后端服务Id (Optional) */
    BackendId string `json:"backendId"`

    /* 【alb Https和Http协议】转发规则组Id (Optional) */
    UrlMapId string `json:"urlMapId"`

    /* 【alb、nlb】空闲连接超时时间, 范围为[1,86400]。 <br>（Tcp和Tls协议）默认为：1800s <br>（Udp协议）默认为：300s <br>（Http和Https协议）默认为：60s <br>【dnlb】不支持 (Optional) */
    ConnectionIdleTimeSeconds int `json:"connectionIdleTimeSeconds"`

    /* 【alb Https和Tls协议】Listener绑定的默认证书，最多支持两个，两个证书的加密算法需要不同 (Optional) */
    CertificateSpecs []CertificateSpec `json:"certificateSpecs"`

    /* 【仅ALB支持】限速配置 (Optional) */
    Limitation LimitationSpec `json:"limitation"`

    /* Listener的描述信息 (Optional) */
    Description string `json:"description"`

    /* Listener的创建时间 (Optional) */
    CreatedTime string `json:"createdTime"`

    /* 【alb Https和Tls协议】Listener绑定的扩展证书列表 (Optional) */
    ExtensionCertificateSpecs []ExtensionCertificateSpec `json:"extensionCertificateSpecs"`

    /* 绑定的安全策略id，仅支持应用负载均衡的HTTPS、TLS监听配置 (Optional) */
    SecurityPolicyId string `json:"securityPolicyId"`
}
