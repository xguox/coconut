package controller

import (
	"coconut/model"
	. "coconut/serializer"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProductVariants(c *gin.Context) {
	id := c.Params.ByName("id")
	_product, err := model.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "no product found"})
		return
	}
	var variants = _product.Variants
	if len(_product.Variants) == 0 {
		variants = append(variants, *_product.GetDefaultVariant())
	}
	s := VariantsSerializer{variants}
	c.JSON(http.StatusOK, gin.H{"data": s.Response()})
}
