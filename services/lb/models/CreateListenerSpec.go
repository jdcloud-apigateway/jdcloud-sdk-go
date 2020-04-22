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


type CreateListenerSpec struct {

    /* Listener的名字,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符  */
    ListenerName string `json:"listenerName"`

    /* 监听协议, 取值为Tcp, Tls, Http, Https <br>【alb】支持Http, Https，Tcp和Tls <br>【nlb】支持Tcp  <br>【dnlb】支持Tcp  */
    Protocol string `json:"protocol"`

    /* 监听端口，取值范围为[1, 65535]  */
    Port int `json:"port"`

    /* 默认的后端服务Id  */
    BackendId string `json:"backendId"`

    /* Listener所属loadBalancer的Id  */
    LoadBalancerId string `json:"loadBalancerId"`

    /* 【alb Https和Http协议】转发规则组Id (Optional) */
    UrlMapId string `json:"urlMapId"`

    /* 默认后端服务的转发策略,取值为Forward或Redirect, 现只支持Forward, 默认为Forward (Optional) */
    Action string `json:"action"`

    /* 【alb Https和Tls协议】ssl server证书列表，现只支持一个证书 (Optional) */
    CertificateSpecs []CertificateSpec `json:"certificateSpecs"`

    /* 【alb、nlb】空闲连接超时时间, 范围为[1,86400]。 <br>（Tcp和Tls协议）默认为：1800s <br>（Http和Https协议）默认为：60s <br>【dnlb】不支持 (Optional) */
    ConnectionIdleTimeSeconds int `json:"connectionIdleTimeSeconds"`

    /* 描述,允许输入UTF-8编码下的全部字符，不超过256字符 (Optional) */
    Description string `json:"description"`
}