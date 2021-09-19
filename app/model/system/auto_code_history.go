package system

import (
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/flipped-aurora/gf-vue-admin/library/types"
)

type AutoCodeHistory struct {
	global.Model
	Flag          int    `json:"flag" gorm:"column:flag;comment:表示对应状态 0 代表创建, 1 代表回滚 ..."`
	ApiIDs        string `json:"apiIDs" gorm:"column:api_ids;comment:api表注册内容"`
	TableName     string `json:"table_name" gorm:"column:table_name;comment:表名"`
	StructName    string `json:"structName" gorm:"column:struct_name;comment:结构体名"`
	StructCNName  string `json:"structCNName" gorm:"column:struct_cn_name;comment:结构体中文名"`
	RequestMeta   string `json:"requestMeta" gorm:"column:request_meta;comment:前端传入的结构化信息"`
	AutoCodePath  string `json:"autoCodePath" gorm:"column:auto_code_path;comment:其他meta信息 path;path"`
	InjectionMeta string `json:"injectionMeta" gorm:"column:injection_meta;comment:注入的内容 RouterPath@functionName@RouterString;"`

	Apis types.UintSlice `json:"apis" gorm:"column:apis;comment:api表注册内容"`
}

type Injection struct {
	Path         string `json:"path" gorm:"column:path;comment:注入文件路径"`
	FunctionName string `json:"function_name" gorm:"column:function_name;comment:函数名"`
	InsertString string `json:"insert_string" gorm:"column:insert_string;comment:插入字符串"`
}

func (a *AutoCodeHistory) ToCommonGetByID() *common.GetByIDs {
	if a == nil {
		return nil
	}
	return &common.GetByIDs{Ids: a.Apis.Uints}
}
