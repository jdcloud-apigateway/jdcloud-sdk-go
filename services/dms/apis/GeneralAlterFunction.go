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
    dms "github.com/jdcloud-api/jdcloud-sdk-go/services/dms/models"
)

type GeneralAlterFunctionRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* 数据源id。 (Optional) */
    DataSourceId *int `json:"dataSourceId"`

    /* 数据库名称。 (Optional) */
    DbName *string `json:"dbName"`

    /* 函数名称。 (Optional) */
    FunctionName *string `json:"functionName"`

    /* 函数名称。 (Optional) */
    OriginFunctionName *string `json:"originFunctionName"`

    /* 安全性，DEFAULT("DEFAULT", 1),DEFINER("DEFINER", 2), INVOKER("INVOKER", 3); (Optional) */
    FunctionSecurity *string `json:"functionSecurity"`

    /* 数据访问，DEFAULT("DEFAULT", 1),NO_SQL("NO_SQL", 2), CONTAINS_SQL("CONTAINS_SQL", 3), READS_SQL_DATA("READS_SQL_DATA", 4), MODIFIES_SQL_DATA("MODIFIES_SQL_DATA", 5); (Optional) */
    DataAccess *string `json:"dataAccess"`

    /* 确定性。 (Optional) */
    Deterministic *bool `json:"deterministic"`

    /* 函数定义SQL。 (Optional) */
    DefinitionSql *string `json:"definitionSql"`

    /* 函数定义SQL。 (Optional) */
    CompleteSql *string `json:"completeSql"`

    /* 参数列表。 (Optional) */
    Parameters []dms.Parameter `json:"parameters"`

    /* 定义者。 (Optional) */
    Definer *string `json:"definer"`

    /* 注释。 (Optional) */
    Comment *string `json:"comment"`

    /* 返回值类型，TINYINT("TINYINT", 0), SMALLINT("SMALLINT", 1), MEDIUMINT("MEDIUMINT", 2), INT("INT", 3), BIGINT("BIGINT", 4), INTEGER("INTEGER", 5), FLOAT("FLOAT", 6), DOUBLE("DOUBLE", 7), REAL("REAL", 8), DECIMAL("DECIMAL", 9), CHAR("CHAR", 10), VARCHAR("VARCHAR", 11), TINYTEXT("TINYTEXT", 12), TEXT("TEXT", 13), MEDIUMTEXT("MEDIUMTEXT", 14), LONGTEXT("LONGTEXT", 15), DATE("DATE", 16), DATETIME("DATETIME", 17), TIMESTAMP("TIMESTAMP", 18), TIME("TIME", 19), YEAR("YEAR", 19), BINARY("BINARY", 20), VARBINARY("VARBINARY", 21), TINYBLOB("TINYBLOB", 22), BLOB("BLOB", 23), MEDIUMBLOB("MEDIUMBLOB", 24), LONGBLOB("LONGBLOB", 25); (Optional) */
    ReturnType *string `json:"returnType"`

    /* 返回值长度。 (Optional) */
    ReturnLength *int `json:"returnLength"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGeneralAlterFunctionRequest(
    regionId string,
) *GeneralAlterFunctionRequest {

	return &GeneralAlterFunctionRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/function:generalAlter",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param dataSourceId: 数据源id。 (Optional)
 * param dbName: 数据库名称。 (Optional)
 * param functionName: 函数名称。 (Optional)
 * param originFunctionName: 函数名称。 (Optional)
 * param functionSecurity: 安全性，DEFAULT("DEFAULT", 1),DEFINER("DEFINER", 2), INVOKER("INVOKER", 3); (Optional)
 * param dataAccess: 数据访问，DEFAULT("DEFAULT", 1),NO_SQL("NO_SQL", 2), CONTAINS_SQL("CONTAINS_SQL", 3), READS_SQL_DATA("READS_SQL_DATA", 4), MODIFIES_SQL_DATA("MODIFIES_SQL_DATA", 5); (Optional)
 * param deterministic: 确定性。 (Optional)
 * param definitionSql: 函数定义SQL。 (Optional)
 * param completeSql: 函数定义SQL。 (Optional)
 * param parameters: 参数列表。 (Optional)
 * param definer: 定义者。 (Optional)
 * param comment: 注释。 (Optional)
 * param returnType: 返回值类型，TINYINT("TINYINT", 0), SMALLINT("SMALLINT", 1), MEDIUMINT("MEDIUMINT", 2), INT("INT", 3), BIGINT("BIGINT", 4), INTEGER("INTEGER", 5), FLOAT("FLOAT", 6), DOUBLE("DOUBLE", 7), REAL("REAL", 8), DECIMAL("DECIMAL", 9), CHAR("CHAR", 10), VARCHAR("VARCHAR", 11), TINYTEXT("TINYTEXT", 12), TEXT("TEXT", 13), MEDIUMTEXT("MEDIUMTEXT", 14), LONGTEXT("LONGTEXT", 15), DATE("DATE", 16), DATETIME("DATETIME", 17), TIMESTAMP("TIMESTAMP", 18), TIME("TIME", 19), YEAR("YEAR", 19), BINARY("BINARY", 20), VARBINARY("VARBINARY", 21), TINYBLOB("TINYBLOB", 22), BLOB("BLOB", 23), MEDIUMBLOB("MEDIUMBLOB", 24), LONGBLOB("LONGBLOB", 25); (Optional)
 * param returnLength: 返回值长度。 (Optional)
 */
func NewGeneralAlterFunctionRequestWithAllParams(
    regionId string,
    dataSourceId *int,
    dbName *string,
    functionName *string,
    originFunctionName *string,
    functionSecurity *string,
    dataAccess *string,
    deterministic *bool,
    definitionSql *string,
    completeSql *string,
    parameters []dms.Parameter,
    definer *string,
    comment *string,
    returnType *string,
    returnLength *int,
) *GeneralAlterFunctionRequest {

    return &GeneralAlterFunctionRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/function:generalAlter",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DataSourceId: dataSourceId,
        DbName: dbName,
        FunctionName: functionName,
        OriginFunctionName: originFunctionName,
        FunctionSecurity: functionSecurity,
        DataAccess: dataAccess,
        Deterministic: deterministic,
        DefinitionSql: definitionSql,
        CompleteSql: completeSql,
        Parameters: parameters,
        Definer: definer,
        Comment: comment,
        ReturnType: returnType,
        ReturnLength: returnLength,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGeneralAlterFunctionRequestWithoutParam() *GeneralAlterFunctionRequest {

    return &GeneralAlterFunctionRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/function:generalAlter",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *GeneralAlterFunctionRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param dataSourceId: 数据源id。(Optional) */
func (r *GeneralAlterFunctionRequest) SetDataSourceId(dataSourceId int) {
    r.DataSourceId = &dataSourceId
}
/* param dbName: 数据库名称。(Optional) */
func (r *GeneralAlterFunctionRequest) SetDbName(dbName string) {
    r.DbName = &dbName
}
/* param functionName: 函数名称。(Optional) */
func (r *GeneralAlterFunctionRequest) SetFunctionName(functionName string) {
    r.FunctionName = &functionName
}
/* param originFunctionName: 函数名称。(Optional) */
func (r *GeneralAlterFunctionRequest) SetOriginFunctionName(originFunctionName string) {
    r.OriginFunctionName = &originFunctionName
}
/* param functionSecurity: 安全性，DEFAULT("DEFAULT", 1),DEFINER("DEFINER", 2), INVOKER("INVOKER", 3);(Optional) */
func (r *GeneralAlterFunctionRequest) SetFunctionSecurity(functionSecurity string) {
    r.FunctionSecurity = &functionSecurity
}
/* param dataAccess: 数据访问，DEFAULT("DEFAULT", 1),NO_SQL("NO_SQL", 2), CONTAINS_SQL("CONTAINS_SQL", 3), READS_SQL_DATA("READS_SQL_DATA", 4), MODIFIES_SQL_DATA("MODIFIES_SQL_DATA", 5);(Optional) */
func (r *GeneralAlterFunctionRequest) SetDataAccess(dataAccess string) {
    r.DataAccess = &dataAccess
}
/* param deterministic: 确定性。(Optional) */
func (r *GeneralAlterFunctionRequest) SetDeterministic(deterministic bool) {
    r.Deterministic = &deterministic
}
/* param definitionSql: 函数定义SQL。(Optional) */
func (r *GeneralAlterFunctionRequest) SetDefinitionSql(definitionSql string) {
    r.DefinitionSql = &definitionSql
}
/* param completeSql: 函数定义SQL。(Optional) */
func (r *GeneralAlterFunctionRequest) SetCompleteSql(completeSql string) {
    r.CompleteSql = &completeSql
}
/* param parameters: 参数列表。(Optional) */
func (r *GeneralAlterFunctionRequest) SetParameters(parameters []dms.Parameter) {
    r.Parameters = parameters
}
/* param definer: 定义者。(Optional) */
func (r *GeneralAlterFunctionRequest) SetDefiner(definer string) {
    r.Definer = &definer
}
/* param comment: 注释。(Optional) */
func (r *GeneralAlterFunctionRequest) SetComment(comment string) {
    r.Comment = &comment
}
/* param returnType: 返回值类型，TINYINT("TINYINT", 0), SMALLINT("SMALLINT", 1), MEDIUMINT("MEDIUMINT", 2), INT("INT", 3), BIGINT("BIGINT", 4), INTEGER("INTEGER", 5), FLOAT("FLOAT", 6), DOUBLE("DOUBLE", 7), REAL("REAL", 8), DECIMAL("DECIMAL", 9), CHAR("CHAR", 10), VARCHAR("VARCHAR", 11), TINYTEXT("TINYTEXT", 12), TEXT("TEXT", 13), MEDIUMTEXT("MEDIUMTEXT", 14), LONGTEXT("LONGTEXT", 15), DATE("DATE", 16), DATETIME("DATETIME", 17), TIMESTAMP("TIMESTAMP", 18), TIME("TIME", 19), YEAR("YEAR", 19), BINARY("BINARY", 20), VARBINARY("VARBINARY", 21), TINYBLOB("TINYBLOB", 22), BLOB("BLOB", 23), MEDIUMBLOB("MEDIUMBLOB", 24), LONGBLOB("LONGBLOB", 25);(Optional) */
func (r *GeneralAlterFunctionRequest) SetReturnType(returnType string) {
    r.ReturnType = &returnType
}
/* param returnLength: 返回值长度。(Optional) */
func (r *GeneralAlterFunctionRequest) SetReturnLength(returnLength int) {
    r.ReturnLength = &returnLength
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GeneralAlterFunctionRequest) GetRegionId() string {
    return r.RegionId
}

type GeneralAlterFunctionResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GeneralAlterFunctionResult `json:"result"`
}

type GeneralAlterFunctionResult struct {
    DmsSqls []dms.DmsSql `json:"dmsSqls"`
}