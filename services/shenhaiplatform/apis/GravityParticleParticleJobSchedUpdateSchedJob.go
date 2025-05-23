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
)

type GravityParticleParticleJobSchedUpdateSchedJobRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 作业名称 (Optional) */
    JobName *string `json:"jobName"`

    /* 客户作业名称 (Optional) */
    CstJobName *string `json:"cstJobName"`

    /* 租户code (Optional) */
    CompanyCode *string `json:"companyCode"`

    /* 工作空间code (Optional) */
    WorkspaceCode *string `json:"workspaceCode"`

    /* 计算资源组code (Optional) */
    ResourceCode *string `json:"resourceCode"`

    /* 作业描述 (Optional) */
    JobDesc *string `json:"jobDesc"`

    /* 最后一次运行状态，Pending、Ready、Running、Done、Failed、Clean (Optional) */
    LastStatus *string `json:"lastStatus"`

    /* 最后一次执行日期 (Optional) */
    LastTxDate *string `json:"lastTxDate"`

    /* 最后一次运行开始时间 (Optional) */
    LastStartTime *string `json:"lastStartTime"`

    /* 最后一次运行结束时间 (Optional) */
    LastEndTime *string `json:"lastEndTime"`

    /* 最后一次运行服务器 (Optional) */
    LastServer *string `json:"lastServer"`

    /* Session ID (Optional) */
    LastSessionid *int `json:"lastSessionid"`

    /* 最后返回编码 (Optional) */
    LastReturnCode *int `json:"lastReturnCode"`

    /* 当前作业状态信息 (Optional) */
    CurrentStatusMsg *string `json:"currentStatusMsg"`

    /* 触发类型:dependency 依赖、time 时间、file 文件、manual 手工、once 一次性 (Optional) */
    TriggerType *string `json:"triggerType"`

    /* 运行周期 ,F 分钟、H 小时、D 天、W 周、M 月、O 一次性运行、N 无周期 (Optional) */
    Cycle *string `json:"cycle"`

    /* 周期具体日期 (Optional) */
    Sequence *string `json:"sequence"`

    /* T+N，偏移量 (Optional) */
    TxDateOffset *int `json:"txDateOffset"`

    /* 作业优先级，数字越小优先级越高 (Optional) */
    Priority *int `json:"priority"`

    /* 作业超时时间，单位分钟 (Optional) */
    Timeout *int `json:"timeout"`

    /* 窗口期开始时间 (Optional) */
    WindowStartTime *string `json:"windowStartTime"`

    /* 窗口期结束时间 (Optional) */
    WindowEndTime *string `json:"windowEndTime"`

    /* 失败后次日是否自动运行，1是、0否 (Optional) */
    MorrowAutoExec *string `json:"morrowAutoExec"`

    /* 抽空之后的处理 0 无操作 1 作业失败 2发出警告 (Optional) */
    DataZeroKillEnable *int `json:"dataZeroKillEnable"`

    /* 重试次数 (Optional) */
    RetryCount *int `json:"retryCount"`

    /* 间隔/秒 (Optional) */
    RetryInterval *int `json:"retryInterval"`

    /* 当期已经跑成功过，是否可以再跑，1启用、0关闭 (Optional) */
    FlagAgain *string `json:"flagAgain"`

    /* 是否可以跨周期跑，1启用、0关闭 (Optional) */
    FlagAcross *string `json:"flagAcross"`

    /* 是否可以自身并行，1启用、0关闭 (Optional) */
    FlagParallel *string `json:"flagParallel"`

    /* 是否级联触发，父任务重跑后是否被强制触发重跑，1是、0否 (Optional) */
    FlagCascadedTrigger *string `json:"flagCascadedTrigger"`

    /* 需要的运行环境 (Optional) */
    RequiredRunEnv *string `json:"requiredRunEnv"`

    /* 运行脚本 (Optional) */
    RunScript *string `json:"runScript"`

    /* 是否启用，0未上线、1已上线、2已下线 (Optional) */
    Enable *string `json:"enable"`

    /* 负责人，不超过10个 (Optional) */
    Manager *string `json:"manager"`

    /* 共享人，不超过10个 (Optional) */
    ShareUser *string `json:"shareUser"`

    /* 失效时间 (Optional) */
    ExpireTime *string `json:"expireTime"`

    /* zip命令行 (Optional) */
    Commands *string `json:"commands"`

    /* 创建人 (Optional) */
    CreateUser *string `json:"createUser"`

    /* 创建时间 (Optional) */
    CreateTime *string `json:"createTime"`

    /* 更新时间 (Optional) */
    UpdateTime *string `json:"updateTime"`

    /* 实时结点名称 (Optional) */
    NodeName *string `json:"nodeName"`

    /* 实时主题 (Optional) */
    Topic *string `json:"topic"`

    /* 作业执行方式  0离线 1双写 2实时 (Optional) */
    JobRunWay *string `json:"jobRunWay"`

    /* 所属系统 (Optional) */
    BelongSys *string `json:"belongSys"`

    /* 主从同步延迟处理方式。0:警告无处理，1:作业延迟启动 (Optional) */
    MsDelayDealWay *string `json:"msDelayDealWay"`

    /* 周期开始时间(适用小时分钟) (Optional) */
    SequenceStartTime *string `json:"sequenceStartTime"`

    /* 周期结束时间(适用小时分钟) (Optional) */
    SequenceEndTime *string `json:"sequenceEndTime"`

    /* 周期间隔(适用小时分钟，当周期为小时，含义为间隔小时数，当周期为分钟，含义为间隔分钟数) (Optional) */
    SequenceInterval *int `json:"sequenceInterval"`

    /* 作业运行参数 (Optional) */
    RunParams *string `json:"runParams"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGravityParticleParticleJobSchedUpdateSchedJobRequest(
    regionId string,
    appName string,
) *GravityParticleParticleJobSchedUpdateSchedJobRequest {

	return &GravityParticleParticleJobSchedUpdateSchedJobRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/gravityParticleParticleJobSchedUpdateSchedJob",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        AppName: appName,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param jobName: 作业名称 (Optional)
 * param cstJobName: 客户作业名称 (Optional)
 * param companyCode: 租户code (Optional)
 * param workspaceCode: 工作空间code (Optional)
 * param resourceCode: 计算资源组code (Optional)
 * param jobDesc: 作业描述 (Optional)
 * param lastStatus: 最后一次运行状态，Pending、Ready、Running、Done、Failed、Clean (Optional)
 * param lastTxDate: 最后一次执行日期 (Optional)
 * param lastStartTime: 最后一次运行开始时间 (Optional)
 * param lastEndTime: 最后一次运行结束时间 (Optional)
 * param lastServer: 最后一次运行服务器 (Optional)
 * param lastSessionid: Session ID (Optional)
 * param lastReturnCode: 最后返回编码 (Optional)
 * param currentStatusMsg: 当前作业状态信息 (Optional)
 * param triggerType: 触发类型:dependency 依赖、time 时间、file 文件、manual 手工、once 一次性 (Optional)
 * param cycle: 运行周期 ,F 分钟、H 小时、D 天、W 周、M 月、O 一次性运行、N 无周期 (Optional)
 * param sequence: 周期具体日期 (Optional)
 * param txDateOffset: T+N，偏移量 (Optional)
 * param priority: 作业优先级，数字越小优先级越高 (Optional)
 * param timeout: 作业超时时间，单位分钟 (Optional)
 * param windowStartTime: 窗口期开始时间 (Optional)
 * param windowEndTime: 窗口期结束时间 (Optional)
 * param morrowAutoExec: 失败后次日是否自动运行，1是、0否 (Optional)
 * param dataZeroKillEnable: 抽空之后的处理 0 无操作 1 作业失败 2发出警告 (Optional)
 * param retryCount: 重试次数 (Optional)
 * param retryInterval: 间隔/秒 (Optional)
 * param flagAgain: 当期已经跑成功过，是否可以再跑，1启用、0关闭 (Optional)
 * param flagAcross: 是否可以跨周期跑，1启用、0关闭 (Optional)
 * param flagParallel: 是否可以自身并行，1启用、0关闭 (Optional)
 * param flagCascadedTrigger: 是否级联触发，父任务重跑后是否被强制触发重跑，1是、0否 (Optional)
 * param requiredRunEnv: 需要的运行环境 (Optional)
 * param runScript: 运行脚本 (Optional)
 * param enable: 是否启用，0未上线、1已上线、2已下线 (Optional)
 * param manager: 负责人，不超过10个 (Optional)
 * param shareUser: 共享人，不超过10个 (Optional)
 * param expireTime: 失效时间 (Optional)
 * param commands: zip命令行 (Optional)
 * param createUser: 创建人 (Optional)
 * param createTime: 创建时间 (Optional)
 * param updateTime: 更新时间 (Optional)
 * param nodeName: 实时结点名称 (Optional)
 * param topic: 实时主题 (Optional)
 * param jobRunWay: 作业执行方式  0离线 1双写 2实时 (Optional)
 * param belongSys: 所属系统 (Optional)
 * param msDelayDealWay: 主从同步延迟处理方式。0:警告无处理，1:作业延迟启动 (Optional)
 * param sequenceStartTime: 周期开始时间(适用小时分钟) (Optional)
 * param sequenceEndTime: 周期结束时间(适用小时分钟) (Optional)
 * param sequenceInterval: 周期间隔(适用小时分钟，当周期为小时，含义为间隔小时数，当周期为分钟，含义为间隔分钟数) (Optional)
 * param runParams: 作业运行参数 (Optional)
 */
func NewGravityParticleParticleJobSchedUpdateSchedJobRequestWithAllParams(
    regionId string,
    appName string,
    jobName *string,
    cstJobName *string,
    companyCode *string,
    workspaceCode *string,
    resourceCode *string,
    jobDesc *string,
    lastStatus *string,
    lastTxDate *string,
    lastStartTime *string,
    lastEndTime *string,
    lastServer *string,
    lastSessionid *int,
    lastReturnCode *int,
    currentStatusMsg *string,
    triggerType *string,
    cycle *string,
    sequence *string,
    txDateOffset *int,
    priority *int,
    timeout *int,
    windowStartTime *string,
    windowEndTime *string,
    morrowAutoExec *string,
    dataZeroKillEnable *int,
    retryCount *int,
    retryInterval *int,
    flagAgain *string,
    flagAcross *string,
    flagParallel *string,
    flagCascadedTrigger *string,
    requiredRunEnv *string,
    runScript *string,
    enable *string,
    manager *string,
    shareUser *string,
    expireTime *string,
    commands *string,
    createUser *string,
    createTime *string,
    updateTime *string,
    nodeName *string,
    topic *string,
    jobRunWay *string,
    belongSys *string,
    msDelayDealWay *string,
    sequenceStartTime *string,
    sequenceEndTime *string,
    sequenceInterval *int,
    runParams *string,
) *GravityParticleParticleJobSchedUpdateSchedJobRequest {

    return &GravityParticleParticleJobSchedUpdateSchedJobRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/gravityParticleParticleJobSchedUpdateSchedJob",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        JobName: jobName,
        CstJobName: cstJobName,
        CompanyCode: companyCode,
        WorkspaceCode: workspaceCode,
        ResourceCode: resourceCode,
        JobDesc: jobDesc,
        LastStatus: lastStatus,
        LastTxDate: lastTxDate,
        LastStartTime: lastStartTime,
        LastEndTime: lastEndTime,
        LastServer: lastServer,
        LastSessionid: lastSessionid,
        LastReturnCode: lastReturnCode,
        CurrentStatusMsg: currentStatusMsg,
        TriggerType: triggerType,
        Cycle: cycle,
        Sequence: sequence,
        TxDateOffset: txDateOffset,
        Priority: priority,
        Timeout: timeout,
        WindowStartTime: windowStartTime,
        WindowEndTime: windowEndTime,
        MorrowAutoExec: morrowAutoExec,
        DataZeroKillEnable: dataZeroKillEnable,
        RetryCount: retryCount,
        RetryInterval: retryInterval,
        FlagAgain: flagAgain,
        FlagAcross: flagAcross,
        FlagParallel: flagParallel,
        FlagCascadedTrigger: flagCascadedTrigger,
        RequiredRunEnv: requiredRunEnv,
        RunScript: runScript,
        Enable: enable,
        Manager: manager,
        ShareUser: shareUser,
        ExpireTime: expireTime,
        Commands: commands,
        CreateUser: createUser,
        CreateTime: createTime,
        UpdateTime: updateTime,
        NodeName: nodeName,
        Topic: topic,
        JobRunWay: jobRunWay,
        BelongSys: belongSys,
        MsDelayDealWay: msDelayDealWay,
        SequenceStartTime: sequenceStartTime,
        SequenceEndTime: sequenceEndTime,
        SequenceInterval: sequenceInterval,
        RunParams: runParams,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGravityParticleParticleJobSchedUpdateSchedJobRequestWithoutParam() *GravityParticleParticleJobSchedUpdateSchedJobRequest {

    return &GravityParticleParticleJobSchedUpdateSchedJobRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/gravityParticleParticleJobSchedUpdateSchedJob",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param jobName: 作业名称(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetJobName(jobName string) {
    r.JobName = &jobName
}
/* param cstJobName: 客户作业名称(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetCstJobName(cstJobName string) {
    r.CstJobName = &cstJobName
}
/* param companyCode: 租户code(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetCompanyCode(companyCode string) {
    r.CompanyCode = &companyCode
}
/* param workspaceCode: 工作空间code(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetWorkspaceCode(workspaceCode string) {
    r.WorkspaceCode = &workspaceCode
}
/* param resourceCode: 计算资源组code(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetResourceCode(resourceCode string) {
    r.ResourceCode = &resourceCode
}
/* param jobDesc: 作业描述(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetJobDesc(jobDesc string) {
    r.JobDesc = &jobDesc
}
/* param lastStatus: 最后一次运行状态，Pending、Ready、Running、Done、Failed、Clean(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetLastStatus(lastStatus string) {
    r.LastStatus = &lastStatus
}
/* param lastTxDate: 最后一次执行日期(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetLastTxDate(lastTxDate string) {
    r.LastTxDate = &lastTxDate
}
/* param lastStartTime: 最后一次运行开始时间(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetLastStartTime(lastStartTime string) {
    r.LastStartTime = &lastStartTime
}
/* param lastEndTime: 最后一次运行结束时间(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetLastEndTime(lastEndTime string) {
    r.LastEndTime = &lastEndTime
}
/* param lastServer: 最后一次运行服务器(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetLastServer(lastServer string) {
    r.LastServer = &lastServer
}
/* param lastSessionid: Session ID(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetLastSessionid(lastSessionid int) {
    r.LastSessionid = &lastSessionid
}
/* param lastReturnCode: 最后返回编码(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetLastReturnCode(lastReturnCode int) {
    r.LastReturnCode = &lastReturnCode
}
/* param currentStatusMsg: 当前作业状态信息(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetCurrentStatusMsg(currentStatusMsg string) {
    r.CurrentStatusMsg = &currentStatusMsg
}
/* param triggerType: 触发类型:dependency 依赖、time 时间、file 文件、manual 手工、once 一次性(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetTriggerType(triggerType string) {
    r.TriggerType = &triggerType
}
/* param cycle: 运行周期 ,F 分钟、H 小时、D 天、W 周、M 月、O 一次性运行、N 无周期(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetCycle(cycle string) {
    r.Cycle = &cycle
}
/* param sequence: 周期具体日期(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetSequence(sequence string) {
    r.Sequence = &sequence
}
/* param txDateOffset: T+N，偏移量(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetTxDateOffset(txDateOffset int) {
    r.TxDateOffset = &txDateOffset
}
/* param priority: 作业优先级，数字越小优先级越高(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetPriority(priority int) {
    r.Priority = &priority
}
/* param timeout: 作业超时时间，单位分钟(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetTimeout(timeout int) {
    r.Timeout = &timeout
}
/* param windowStartTime: 窗口期开始时间(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetWindowStartTime(windowStartTime string) {
    r.WindowStartTime = &windowStartTime
}
/* param windowEndTime: 窗口期结束时间(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetWindowEndTime(windowEndTime string) {
    r.WindowEndTime = &windowEndTime
}
/* param morrowAutoExec: 失败后次日是否自动运行，1是、0否(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetMorrowAutoExec(morrowAutoExec string) {
    r.MorrowAutoExec = &morrowAutoExec
}
/* param dataZeroKillEnable: 抽空之后的处理 0 无操作 1 作业失败 2发出警告(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetDataZeroKillEnable(dataZeroKillEnable int) {
    r.DataZeroKillEnable = &dataZeroKillEnable
}
/* param retryCount: 重试次数(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetRetryCount(retryCount int) {
    r.RetryCount = &retryCount
}
/* param retryInterval: 间隔/秒(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetRetryInterval(retryInterval int) {
    r.RetryInterval = &retryInterval
}
/* param flagAgain: 当期已经跑成功过，是否可以再跑，1启用、0关闭(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetFlagAgain(flagAgain string) {
    r.FlagAgain = &flagAgain
}
/* param flagAcross: 是否可以跨周期跑，1启用、0关闭(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetFlagAcross(flagAcross string) {
    r.FlagAcross = &flagAcross
}
/* param flagParallel: 是否可以自身并行，1启用、0关闭(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetFlagParallel(flagParallel string) {
    r.FlagParallel = &flagParallel
}
/* param flagCascadedTrigger: 是否级联触发，父任务重跑后是否被强制触发重跑，1是、0否(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetFlagCascadedTrigger(flagCascadedTrigger string) {
    r.FlagCascadedTrigger = &flagCascadedTrigger
}
/* param requiredRunEnv: 需要的运行环境(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetRequiredRunEnv(requiredRunEnv string) {
    r.RequiredRunEnv = &requiredRunEnv
}
/* param runScript: 运行脚本(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetRunScript(runScript string) {
    r.RunScript = &runScript
}
/* param enable: 是否启用，0未上线、1已上线、2已下线(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetEnable(enable string) {
    r.Enable = &enable
}
/* param manager: 负责人，不超过10个(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetManager(manager string) {
    r.Manager = &manager
}
/* param shareUser: 共享人，不超过10个(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetShareUser(shareUser string) {
    r.ShareUser = &shareUser
}
/* param expireTime: 失效时间(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetExpireTime(expireTime string) {
    r.ExpireTime = &expireTime
}
/* param commands: zip命令行(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetCommands(commands string) {
    r.Commands = &commands
}
/* param createUser: 创建人(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetCreateUser(createUser string) {
    r.CreateUser = &createUser
}
/* param createTime: 创建时间(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetCreateTime(createTime string) {
    r.CreateTime = &createTime
}
/* param updateTime: 更新时间(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetUpdateTime(updateTime string) {
    r.UpdateTime = &updateTime
}
/* param nodeName: 实时结点名称(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetNodeName(nodeName string) {
    r.NodeName = &nodeName
}
/* param topic: 实时主题(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetTopic(topic string) {
    r.Topic = &topic
}
/* param jobRunWay: 作业执行方式  0离线 1双写 2实时(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetJobRunWay(jobRunWay string) {
    r.JobRunWay = &jobRunWay
}
/* param belongSys: 所属系统(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetBelongSys(belongSys string) {
    r.BelongSys = &belongSys
}
/* param msDelayDealWay: 主从同步延迟处理方式。0:警告无处理，1:作业延迟启动(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetMsDelayDealWay(msDelayDealWay string) {
    r.MsDelayDealWay = &msDelayDealWay
}
/* param sequenceStartTime: 周期开始时间(适用小时分钟)(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetSequenceStartTime(sequenceStartTime string) {
    r.SequenceStartTime = &sequenceStartTime
}
/* param sequenceEndTime: 周期结束时间(适用小时分钟)(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetSequenceEndTime(sequenceEndTime string) {
    r.SequenceEndTime = &sequenceEndTime
}
/* param sequenceInterval: 周期间隔(适用小时分钟，当周期为小时，含义为间隔小时数，当周期为分钟，含义为间隔分钟数)(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetSequenceInterval(sequenceInterval int) {
    r.SequenceInterval = &sequenceInterval
}
/* param runParams: 作业运行参数(Optional) */
func (r *GravityParticleParticleJobSchedUpdateSchedJobRequest) SetRunParams(runParams string) {
    r.RunParams = &runParams
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GravityParticleParticleJobSchedUpdateSchedJobRequest) GetRegionId() string {
    return r.RegionId
}

type GravityParticleParticleJobSchedUpdateSchedJobResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GravityParticleParticleJobSchedUpdateSchedJobResult `json:"result"`
}

type GravityParticleParticleJobSchedUpdateSchedJobResult struct {
    Success int `json:"success"`
    Result string `json:"result"`
    Code string `json:"code"`
    Msg string `json:"msg"`
    _REQ_SEQ_NO_ string `json:"_REQ_SEQ_NO_"`
}