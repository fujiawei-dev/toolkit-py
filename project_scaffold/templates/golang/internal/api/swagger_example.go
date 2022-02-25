/*
 * @Date: 2022.02.25 10:47
 * @Description: Omit
 * @LastEditors: Rustle Karl
 * @LastEditTime: 2022.02.25 10:47
 */

package api

import (
	"{{GOLANG_MODULE}}/internal/entity"
	"{{GOLANG_MODULE}}/internal/form"
)

// @Summary 创建示例
// @Description 创建示例
// @Tags 示例
// @Accept json
// @Security ApiKeyAuth
// @Param object body form.ExampleCreate true "参数"
// @Produce json
// @Success 200 {object} httpResponseBody "操作成功"
// @Router /example [post]
func postExample() {}

// @Summary 修改示例
// @Description 修改示例
// @Tags 示例
// @Accept json
// @Security ApiKeyAuth
// @Param id path int true "Example ID"
// @Param object body form.ExampleUpdate true "参数"
// @Produce json
// @Success 200 {object} httpResponseBody "操作成功"
// @Router /example/{id} [put]
func putExample() {}

// @Summary 删除示例
// @Description 删除示例
// @Tags 示例
// @Accept json
// @Security ApiKeyAuth
// @Param id path int true "Example ID"
// @Produce json
// @Success 200 {object} httpResponseBody "操作成功"
// @Router /example/{id} [delete]
func deleteExample() {}

type examplesResponseBody struct {
	httpResponseBody
	Result struct {
		form.Pager `json:"pager"`
		List       entity.Examples `json:"list"`
	} `json:"result"`
}

// @Summary 获取示例
// @Description 获取示例
// @Tags 示例
// @Accept json
// @Security ApiKeyAuth
// @Param id path int true "Example ID"
// @Produce json
// @Success 200 {object} entity.Example "操作成功"
// @Router /example/{id} [get]
func getExample() {}

// @Summary 获取示例列表
// @Description 获取示例列表
// @Tags 示例
// @Accept json
// @Security ApiKeyAuth
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" Enums(10, 20) default(10)
// @Produce json
// @Success 200 {object} examplesResponseBody "操作成功"
// @Router /examples [get]
func getExamples() {}