package system

import (
	"github.com/flipped-aurora/gf-vue-admin/app/model/system/request"
	"github.com/flipped-aurora/gf-vue-admin/app/service/system"
	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"github.com/flipped-aurora/gf-vue-admin/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Dictionary = new(dictionary)

type dictionary struct{}

// Create
// @Tags SystemDictionary
// @Summary 创建系统字典
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DictionaryCreate true "请求参数"
// @Success 200 {object} response.Response{} "创建成功!"
// @Router /sysDictionary/createSysDictionary [post]
func (a *dictionary) Create(r *ghttp.Request) *response.Response {
	var info request.DictionaryCreate
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	if err := system.Dictionary.Create(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorCreated}
	}
	return &response.Response{MessageCode: response.SuccessCreated}
}

// First
// @Tags SystemDictionary
// @Summary 用id查询系统字典
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.DictionaryFirst true "请求参数"
// @Success 200 {object} response.Response{} "获取数据成功!"
// @Router /sysDictionary/findSysDictionary [get]
func (a *dictionary) First(r *ghttp.Request) *response.Response {
	var info request.DictionaryFirst
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	}
	data, err := system.Dictionary.First(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorFirst}
	}
	return &response.Response{Data: g.Map{"resysDictionary": data}, MessageCode: response.SuccessFirst}
}

// Update
// @Tags SystemDictionary
// @Summary 更新系统字典
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DictionaryUpdate true "请求参数"
// @Success 200 {object} response.Response{data=system.Dictionary} "更新成功!"
// @Router /sysDictionary/updateSysDictionary [put]
func (a *dictionary) Update(r *ghttp.Request) *response.Response {
	var info request.DictionaryUpdate
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	if err := system.Dictionary.Update(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorUpdated}
	}
	return &response.Response{MessageCode: response.SuccessUpdated}
}

// Delete
// @Tags SystemDictionary
// @Summary 删除系统字典
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body common.GetByID true "请求参数"
// @Success 200 {object} response.Response{} "删除成功!"
// @Router /sysDictionary/deleteSysDictionary [delete]
func (a *dictionary) Delete(r *ghttp.Request) *response.Response {
	var info common.GetByID
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	if err := system.Dictionary.Delete(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorDeleted}
	}
	return &response.Response{MessageCode: response.SuccessDeleted}
}

// GetList
// @Tags SystemDictionary
// @Summary 分页获取系统字典列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.DictionarySearch true "请求参数"
// @Success 200 {object} response.Response{data=[]system.Dictionary} "获取列表数据成功!"
// @Router /sysDictionary/getSysDictionaryList [get]
func (a *dictionary) GetList(r *ghttp.Request) *response.Response {
	var info request.DictionarySearch
	if err := r.Parse(&info); err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	list, total, err := system.Dictionary.GetList(&info)
	if err != nil {
		return &response.Response{Error: err, MessageCode: response.ErrorGetList}
	}
	return &response.Response{Data: common.NewPageResult(list, total, info.PageInfo), MessageCode: response.SuccessGetList}
}
