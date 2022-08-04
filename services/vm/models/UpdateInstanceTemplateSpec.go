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


type UpdateInstanceTemplateSpec struct {

    /* 实例规格，可查询 [DescribeInstanceTypes](https://docs.jdcloud.com/virtual-machines/api/describeinstancetypes) 接口获得指定地域或可用区的规格信息。 (Optional) */
    InstanceType *string `json:"instanceType"`

    /* 镜像ID，可查询 [DescribeImages](https://docs.jdcloud.com/virtual-machines/api/describeimages) 接口获得指定地域的镜像信息。 (Optional) */
    ImageId *string `json:"imageId"`

    /* 实例密码。可用于SSH登录和VNC登录。长度为8\~30个字符，必须同时包含大、小写英文字母、数字和特殊符号中的三类字符。特殊符号包括：\(\)\`~!@#$%^&\*\_-+=\|{}\[ ]:";'<>,.?/，更多密码输入要求请参见 [公共参数规范](https://docs.jdcloud.com/virtual-machines/api/general_parameters)。
如指定密钥且 `passwordAuth` 设置为 `true`，则密码不会生成注入，否则即使不指定密码系统也将默认自动生成随机密码，并以短信和邮件通知。
 (Optional) */
    Password *string `json:"password"`

    /* 密钥对名称。仅Linux系统下该参数生效，当前仅支持输入单个密钥。如指定了该参数则覆盖原有参数。 (Optional) */
    KeyNames []string `json:"keyNames"`

    /* 用户自定义元数据。以key-value键值对形式指定，可在实例系统内通过元数据服务查询获取。最多支持20对键值对，且key不超过256字符，value不超过16KB，不区分大小写。
注意：key以连字符(-)结尾，表示要删除该key。
 (Optional) */
    Metadata []Metadata `json:"metadata"`

    /* 自定义脚本。目前仅支持启动脚本，即 `launch-script`，须 `base64` 编码且编码前数据长度不能超过16KB。
**linux系统**：支持 `bash` 和 `python`，编码前须分别以 `#!/bin/bash` 和 `#!/usr/bin/env python` 作为内容首行。
**Windows系统**：支持 `bat` 和 `powershell`，编码前须分别以 `<cmd></cmd>和<powershell></powershell>` 作为内容首、尾行。
 (Optional) */
    Userdata []Userdata `json:"userdata"`

    /* 主网卡主IP关联的弹性公网IP配置。如指定了该参数则覆盖原有参数。 (Optional) */
    ElasticIp *InstanceTemplateElasticIpSpec `json:"elasticIp"`

    /* 主网卡配置。如指定了该参数则覆盖原有参数。 (Optional) */
    PrimaryNetworkInterface *InstanceTemplateNetworkInterfaceAttachmentSpec `json:"primaryNetworkInterface"`

    /* 系统盘配置。如指定了该参数则覆盖原有参数。 (Optional) */
    SystemDisk *InstanceTemplateDiskAttachmentSpec `json:"systemDisk"`

    /* 数据盘配置。单实例最多可挂载云硬盘（系统盘+数据盘）的数量受实例规格的限制。如指定了该参数则覆盖原有参数。 (Optional) */
    DataDisks []InstanceTemplateDiskAttachmentSpec `json:"dataDisks"`

    /* 停机不计费模式。该参数仅对按配置计费且系统盘为云硬盘的实例生效，并且不是专有宿主机中的实例。配置停机不计费且停机后，实例部分将停止计费，且释放实例自身包含的资源（CPU/内存/GPU/本地数据盘）。
可选值：
`keepCharging`（默认值）：停机后保持计费，不释放资源。
`stopCharging`：停机后停止计费，释放实例资源。
 (Optional) */
    ChargeOnStopped *string `json:"chargeOnStopped"`

    /* 自动任务策略ID。 (Optional) */
    AutoImagePolicyId *string `json:"autoImagePolicyId"`

    /* 是否允许SSH密码登录。
`yes`：允许SSH密码登录。
`no`：禁止SSH密码登录。
仅在指定密钥时此参数有效，指定此参数后密码即使输入也将被忽略，同时会在系统内禁用SSH密码登录。
 (Optional) */
    PassWordAuth *string `json:"passWordAuth"`

    /* 是否使用镜像中的登录凭证，不再指定密码或密钥。
`yes`：使用镜像登录凭证。
`no`（默认值）：不使用镜像登录凭证。
仅使用私有或共享镜像时此参数有效。若指定`imageInherit=yes`则指定的密码或密钥将无效。
 (Optional) */
    ImageInherit *string `json:"imageInherit"`

    /* 传 `true` 则会清空实例模板配置的密码。 (Optional) */
    NoPassword *bool `json:"noPassword"`

    /* 传 `true` 则会清空实例模板配置的公网IP。 (Optional) */
    NoElasticIp *bool `json:"noElasticIp"`

    /* 突发型实例参数配置。传入 `null` 表示忽略，否则以新传入的为准。如指定了该参数则覆盖原有参数。 (Optional) */
    BurstSpec *InstanceTemplateBurstSpec `json:"burstSpec"`
}
