package controllers

import (
	"IrisYouQiKangApi/models"
	"IrisYouQiKangApi/system"
	"github.com/kataras/iris"
)

/**
* @api {get} /plans/parent 获取所有的诊断方案
* @apiName 获取所有的诊断方案
* @apiGroup Plans
* @apiVersion 1.0.0
* @apiDescription 获取所有的诊断方案
* @apiSampleRequest /plans/parent
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func GetAllParentPlans(ctx iris.Context) {
	cp := system.Tools.ParseInt(ctx.FormValue("cp"), 1)
	mp := system.Tools.ParseInt(ctx.FormValue("mp"), 20)
	kw := ctx.FormValue("kw")
	plans := models.GetAllParentPlans(kw, cp, mp)

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(apiResource(true, plans, "操作成功"))
}

/**
* @api {get} /plans 获取所有的诊断方案
* @apiName 获取所有的诊断方案
* @apiGroup Plans
* @apiVersion 1.0.0
* @apiDescription 获取所有的诊断方案
* @apiSampleRequest /plans
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func GetAllPlans(ctx iris.Context) {
	cp := system.Tools.ParseInt(ctx.FormValue("cp"), 1)
	mp := system.Tools.ParseInt(ctx.FormValue("mp"), 20)
	kw := ctx.FormValue("kw")

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(models.GetAllPlans(kw, cp, mp))
}