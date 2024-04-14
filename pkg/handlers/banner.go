package handler

import (
	"banner_service/pkg/model"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllBanners(c *gin.Context) {
	var form model.AdminGet
	var banners []model.Banner
	var err error
	if err = c.ShouldBindQuery(&form); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if form.Offset == 0 {
		form.Offset = -1
	}
	if form.Limit == 0 {
		form.Limit = -1
	}

	banners, err = h.services.GetAllBanners(form)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, banners)
}

func (h *Handler) createBanner(c *gin.Context) {
	var input model.Banner
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	sort.SliceStable(input.TagIds, func(i int, j int) bool {
		return input.TagIds[i] < input.TagIds[j]
	})
	id, err := h.services.Banner.CreateBanner(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"banner_id": id,
	})
}

func (h *Handler) modifyBanner(c *gin.Context) {
	var err error
	var id int
	var input model.Banner

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	idStr := c.Param("id")
	id, err = strconv.Atoi(idStr)
	if err != nil {
		newErrorResponse(c, 400, err.Error())
		return
	}

	err = h.services.FindId(id)
	if err != nil {
		newErrorResponse(c, 404, err.Error())
	}

	if input.TagIds != nil {
		sort.SliceStable(input.TagIds, func(i int, j int) bool {
			return input.TagIds[i] < input.TagIds[j]
		})
	}
	input.Id = id
	input.UpdatedAt = time.Now()
	err = h.services.UpdateBanner(input)
	if err != nil {
		newErrorResponse(c, 500, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) getUserBanner(c *gin.Context) {
	var form model.UserGet
	var content model.JSONB
	var err error
	form.UseLastRevision = false

	if err = c.ShouldBindQuery(&form); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	content, err = h.services.Banner.GetUserBanner(form)
	if err != nil {
		newErrorResponse(c, 404, err.Error())
	}

	c.JSON(http.StatusOK, content)
}

func (h *Handler) deleteBanner(c *gin.Context) {
	var err error
	var id int
	idStr := c.Param("id")
	id, err = strconv.Atoi(idStr)

	if err != nil {
		newErrorResponse(c, 400, err.Error())
	}

	err = h.services.FindId(id)
	if err != nil {
		newErrorResponse(c, 404, err.Error())
	}

	err = h.services.Banner.DeleteBanner(id)
	if err != nil {
		newErrorResponse(c, 500, err.Error())
	}
	c.JSON(204, nil)
}
