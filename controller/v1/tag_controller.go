package v1

import (
	"gin-rest-api/controller"
	"gin-rest-api/service"
	"gin-rest-api/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get multiple article tags
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {object} controller.Response
// @Failure 500 {object} controller.Response
// @Router /api/v1/tags [get]
func GetTags(c *gin.Context) {
	appG := controller.Gin{C: c}
	name := c.Query("name")
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	state := -1
	if arg := c.Query("state"); arg != "" {
		state, _ = strconv.Atoi(arg)
	}

	tagService := service.Tag{
		Name:     name,
		State:    state,
		PageNum:  util.GetPage(c),
		PageSize: pageSize,
	}
	tags, err := tagService.GetAll()
	if err != nil {
		appG.Failed(controller.Failed, err.Error())
		return
	}

	count, err := tagService.Count()
	if err != nil {
		appG.Failed(controller.Failed, err.Error())
		return
	}
	appG.Success(map[string]interface{}{
		"lists": tags,
		"total": count,
	})
}

type AddTagForm struct {
	Name      string `form:"name" valid:"Required;MaxSize(100)"`
	CreatedBy string `form:"created_by" valid:"Required;MaxSize(100)"`
	State     int    `form:"state" valid:"Range(0,1)"`
}

// @Summary Add article tag
// @Produce  json
// @Param name body string true "Name"
// @Param state body int false "State"
// @Param created_by body int false "CreatedBy"
// @Success 200 {object} controller.Response
// @Failure 500 {object} controller.Response
// @Router /api/v1/tags [post]
func AddTag(c *gin.Context) {
	var (
		appG = controller.Gin{C: c}
		form AddTagForm
	)

	code := appG.BindAndValid(&form)
	if code != controller.Success {
		appG.Failed(code, "")
		return
	}

	tagService := service.Tag{
		Name:      form.Name,
		CreatedBy: form.CreatedBy,
		State:     form.State,
	}
	exists, err := tagService.ExistByName()
	if err != nil {
		appG.Failed(controller.Failed, err.Error())
		return
	}
	if exists {
		appG.Failed(controller.Failed, "已存在")
		return
	}

	err = tagService.Add()
	if err != nil {
		appG.Failed(controller.Failed, err.Error())
		return
	}

	appG.Success("ok")
	return
}

type EditTagForm struct {
	Id         int    `form:"id" valid:"Required;Min(1)"`
	Name       string `form:"name" valid:"Required;MaxSize(100)"`
	ModifiedBy string `form:"modified_by" valid:"Required;MaxSize(100)"`
	State      int    `form:"state" valid:"Range(0,1)"`
}

// @Summary Update article tag
// @Produce  json
// @Param id path int true "Id"
// @Param name body string true "Name"
// @Param state body int false "State"
// @Param modified_by body string true "ModifiedBy"
// @Success 200 {object} controller.Response
// @Failure 500 {object} controller.Response
// @Router /api/v1/tags/{id} [put]
func EditTag(c *gin.Context) {

	appG := controller.Gin{C: c}
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		appG.Failed(controller.Failed, err.Error())
		return
	}

	form := EditTagForm{Id: id}

	code := appG.BindAndValid(&form)
	if code != controller.Success {
		appG.Failed(code, "")
		return
	}

	tagService := service.Tag{
		Id:         form.Id,
		Name:       form.Name,
		ModifiedBy: form.ModifiedBy,
		State:      form.State,
	}

	exists, err := tagService.ExistById()
	if err != nil {
		appG.Failed(controller.Failed, err.Error())
		return
	}

	if !exists {
		appG.Failed(controller.Failed, "tag不存在")
		return
	}

	err = tagService.Edit()
	if err != nil {
		appG.Failed(controller.Failed, err.Error())
		return
	}

	appG.Success("ok")
}

// @Summary Delete article tag
// @Produce  json
// @Param id path int true "Id"
// @Success 200 {object} controller.Response
// @Failure 500 {object} controller.Response
// @Router /api/v1/tags/{id} [delete]
func DeleteTag(c *gin.Context) {
	appG := controller.Gin{C: c}
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		appG.Failed(controller.Failed, err.Error())
		return
	}

	tagService := service.Tag{Id: id}
	exists, err := tagService.ExistById()
	if err != nil {
		appG.Failed(controller.Failed, err.Error())
		return
	}

	if !exists {
		appG.Failed(controller.Failed, "tag不存在")
		return
	}

	if err := tagService.Delete(); err != nil {
		appG.Failed(controller.Failed, err.Error())
		return
	}

	appG.Success("ok")
}

// @Summary get article tag
// @Produce  json
// @Param id path int true "Id"
// @Success 200 {object} controller.Response
// @Failure 500 {object} controller.Response
// @Router /api/v1/tags/{id} [get]
func GetTag(c *gin.Context) {
	appG := controller.Gin{C: c}
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		appG.Failed(controller.Failed, err.Error())
		return
	}

	tagService := service.Tag{Id: id}
	tag, err := tagService.GetTag()
	if err != nil {
		appG.Failed(controller.Failed, err.Error())
		return
	}

	appG.Success(tag)
}
