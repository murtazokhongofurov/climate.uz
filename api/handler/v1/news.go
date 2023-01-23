package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/climate.uz/api/models"
	"gitlab.com/climate.uz/internal/controller/storage/repo"
	"gitlab.com/climate.uz/pkg/utils"
)

// New news product
// @Summary Add news product
// @Description This method add news products
// @Tags 		News
// @Security	BearerAuth
// @Accept 		json
// @Produce 	json
// @Param 		news 	body 	 models.NewsProductReq true "NewsProductRequest"
// @Success 	201 	{object} repo.NewsProductResponse
// @Failure 	500 	{object} models.FailureInfo
// @Router  	/news 	[post]
func (h *handlerV1) CreateNewsProduct(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			Code:    http.StatusInternalServerError,
			Message: "Invalid access token",
			Error:   err,
		})
		h.log.Error("Error while getting claims of access token ", err)
		return
	}
	var (
		body models.NewsProductReq
	)
	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			Code:    http.StatusBadRequest,
			Message: "Please enter right info",
		})
		h.log.Error("Error while binding request", err)
		return
	}

	res, err := h.storage.NewProduct().CreateNews(&repo.NewsProductRequest{
		CategoryId:  body.CategoryId,
		Title:       body.Title,
		MediaLink:   body.MediaLink,
		Description: body.Description,
		Price:       body.Price,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			Code:    http.StatusInternalServerError,
			Message: "Something went wrong",
		})
		h.log.Error("error while create news product", err)
		return
	}
	c.JSON(http.StatusCreated, res)
}

// Get news products
// @Summary 	get news products
// @Description This method get news product info by id
// @Tags 		News
// @Accept 		json
// @Produce 	json
// @Param 		id 		path 	 int true "newsproductID"
// @Success 	200 	{object} repo.NewsProductResponse
// @Failure 	500 	{object} models.FailureInfo
// @Router 		/news/{id} 	[get]
func (h *handlerV1) GetNewsProduct(c *gin.Context) {
	id := c.Param("id")
	news_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			Code:    http.StatusBadRequest,
			Message: "Please enter right info",
		})
		h.log.Error("error while parsing news product id")
		return
	}
	res, err := h.storage.NewProduct().GetNewsById(news_id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			Code:    http.StatusInternalServerError,
			Message: "Something went wrong",
		})
		h.log.Error("Error while getting news product by id", err)
		return
	}
	c.JSON(http.StatusOK, res)
}

// Get all news products
// @Summary 	get all news products in category
// @Description This method get all news products by category id
// @Tags 		News
// @Accept 		json
// @Produce 	json
// @Param 		id 		path 			int 	true "categoryId"
// @Success 	200 	{object}		models.AllNewsProducts
// @Failure 	500 	{object}		models.FailureInfo
// @Router 		/news/incategory/{id} 	[get]
func (h *handlerV1) GetNewsByCategoryId(c *gin.Context) {
	id := c.Param("id")
	categoryId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			Code:    http.StatusBadRequest,
			Message: "please enter right info",
		})
		h.log.Error("Error while parsing category id", err)
		return
	}
	res, err := h.storage.NewProduct().GetNewsByCategoryId(categoryId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			Code:    http.StatusInternalServerError,
			Message: "Something went wrong",
		})
		h.log.Error("error while getting news products by category Id", err)
		return
	}

	c.JSON(http.StatusOK, res)

}

// Get All news products
// @Summary 	get all news products
// @Description This method get all news product infos
// @Tags 		News
// @Accept 		json
// @Produce 	json
// @Param 		page 	query 		int 	false "page"
// @Param 		limit 	query 		int 	false "limit"
// @Param 		search 	query 		string 	false "search"
// @Success 	200 	{object}	models.AllNewsProducts
// @Failure 	500 	{object}	models.FailureInfo
// @Router 		/allnews 	[get]
func (h *handlerV1) GetAllNewsProducts(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			Code:    http.StatusBadRequest,
			Message: "please enter right info",
		})
		h.log.Error("error while parsequeryparams", errStr)
		return
	}

	res, errs := h.storage.NewProduct().GetAllNews(&repo.AllNewsProductParams{
		Page:   params.Page,
		Limit:  params.Limit,
		Search: params.Search,
	})
	if errs != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			Code:    http.StatusBadRequest,
			Message: "Something went wrong",
		})
		h.log.Error("Error while getting all news product", errs)
		return
	}
	c.JSON(http.StatusOK, res)
}

// Edit News product
// @Summary 	update 	news product
// @Description this 	method update news product
// @Tags 		News
// @Security	BearerAuth
// @Accept 		json
// @Produce 	json
// @Param 		news 	body 	 models.UpdateNewsProductReq true "UpdateRequest"
// @Success 	200 	{object} models.NewsProductRes
// @Failure 	500 	{object} models.FailureInfo
// @Router 		/news 	[patch]
func (h *handlerV1) UpdateNewsProduct(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			Code:    http.StatusInternalServerError,
			Message: "Invalid access token",
			Error:   err,
		})
		h.log.Error("Error while getting claims of access token ", err)
		return
	}
	var (
		body models.UpdateNewsProductReq
	)
	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			Code:    http.StatusBadRequest,
			Message: "please enter right info",
		})
		h.log.Error("Error while binding request", err)
		return
	}
	res, err := h.storage.NewProduct().UpdateNews(&repo.NewsProductUpdateReq{
		Id:          body.Id,
		Title:       body.Title,
		MediaLink:   body.MediaLink,
		Description: body.Description,
		Price:       body.Price,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			Code:    http.StatusInternalServerError,
			Message: "Something went wrong",
		})
		h.log.Error("error while updating news product", err)
		return
	}
	c.JSON(http.StatusOK, res)
}

// Delete News product
// @Summary 	delete news product info
// @Description this 	method delete news product by id
// @Tags 		News
// @Security	BearerAuth
// @Accept 		json
// @Produce 	json
// @Param 		id 	    	path 	int 	true "id"
// @Success 	200 		{object} models.SuccessInfo
// @Failure 	500 		{object} models.FailureInfo
// @Router 		/news/{id}	[delete]
func (h *handlerV1) DeleteNewsProduct(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			Code:    http.StatusInternalServerError,
			Message: "Invalid access token",
			Error:   err,
		})
		h.log.Error("Error while getting claims of access token ", err)
		return
	}

	id := c.Param("id")

	newId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			Code:    http.StatusBadRequest,
			Message: "please enter right info",
		})
		h.log.Error("error while parsing id", err)
		return
	}
	_, err = h.storage.NewProduct().DeleteNewsById(newId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			Code:    http.StatusInternalServerError,
			Message: "Something went wrong",
		})
		h.log.Error("Error while deleting news product", err)
		return
	}
	c.JSON(http.StatusOK, models.SuccessInfo{
		StatusCode: http.StatusOK,
		Message:    "Successfully deleted news product info",
	})
}

// Delete News product
// @Summary 	delete 	news product
// @Description this 	method delete news product
// @Tags 		News
// @Security	BearerAuth
// @Accept 		json
// @Produce 	json
// @Param 		id 	    	path 	int 	true "category id"
// @Success 	200 		{object} models.SuccessInfo
// @Failure 	500 		{object} models.FailureInfo
// @Router 		/news/category/{id} 	[delete]
func (h *handlerV1) DeleteNewsByCategoryId(c *gin.Context) {
	_, err := GetClaims(*h, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			Code:    http.StatusInternalServerError,
			Message: "Invalid access token",
			Error:   err,
		})
		h.log.Error("Error while getting claims of access token ", err)
		return
	}
	id := c.Param("id")

	categoryId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.FailureInfo{
			Code:    http.StatusBadRequest,
			Message: "please enter right info",
		})
		h.log.Error("error while parsing id", err)
		return
	}
	_, err = h.storage.NewProduct().DeleteNewsByCategoryId(categoryId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			Code:    http.StatusInternalServerError,
			Message: "Something went wrong",
		})
		h.log.Error("Error while deleting news product", err)
		return
	}
	c.JSON(http.StatusOK, models.SuccessInfo{
		StatusCode: http.StatusOK,
		Message:    "Successfully deleted all news products in this category have been",
	})
}
