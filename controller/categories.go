package controller

import (
	"github.com/xguox/coconut/model"

	"github.com/xguox/coconut/util"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateCategory godoc
// @Summary 添加新的商品分类
// @Description 添加新的商品分类
// @Accept json
// @Produce json
// @Tags category
// @Param Authorization header string true "认证 Token"
// @Param body body model.CategoryValidator true "创建分类请求参数"
// @Success 200 {string} json "{mark: 0, msg: "", data: {id: 1, name: "category_name"}}"
// @Failure 422 {string} json "{mark: 1, msg: "err msg", data: {}}"
// @Router /admin/categories [post]
func CreateCategory(c *gin.Context) {
	resp := util.NewResp()
	v := model.CategoryValidator{}
	if err := v.Bind(c); err != nil {
		resp.Mark = util.MarkWarn
		resp.Msg = "参数有误" // TODO: 拼接 util.NewValidatorError(err)
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	if err := model.SaveData(&v.CategoryModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   v.CategoryModel.ID,
		"name": v.CategoryModel.Name,
	})
}

func FetchCategories(c *gin.Context) {
	categories := model.GetCategories()
	if len(categories) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no categories found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": categories})
}

func UpdateCategory(c *gin.Context) {
	id := c.Params.ByName("id")
	category, err := model.GetCategoryByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Category not found"})
		return
	}

	v := model.CategoryValidator{CategoryModel: category}

	if err := v.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, util.NewValidatorError(err))
		return
	}

	if err := model.SaveData(&v.CategoryModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": http.StatusUnprocessableEntity, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   v.CategoryModel.ID,
		"name": v.CategoryModel.Name,
	})
}

func DestroyCategory(c *gin.Context) {
	id := c.Params.ByName("id")
	category, err := model.GetCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "category not found"})
	} else {
		category.SetDeletedAt(time.Now())
		c.JSON(http.StatusOK, gin.H{})
	}
}
