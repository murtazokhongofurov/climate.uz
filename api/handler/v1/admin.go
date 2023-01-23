package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/climate.uz/api/models"
)

// admin register
// @Summary 	Login admin
// @Description this method admin login
// @Tags 		Admin
// @Accept 		json
// @Produce 	json
// @Param  		username path 	  string true "username"
// @Param   	password path 	  string true "password"
// @Success 	201 	 {object} models.SuccessInfo
// @Failure 	500 	 {object} models.FailureInfo
// @Failure 	400 	 {object} models.FailureInfo
// @Failure 	409 	 {object} models.FailureInfo
// @Router 		/admin/{username}/{password} 	[get]
func (h *handlerV1) AdminLogin(c *gin.Context) {
	var (
		username = c.Param("username")
		password = c.Param("password")
	)

	res, err := h.storage.Admin().GetAdmin(username)
	if err != nil {
		c.JSON(http.StatusNotFound, models.FailureInfo{
			Code:    http.StatusNotFound,
			Message: "please enter right info",
		})
		h.log.Error("error while getting admin info", err)
		return
	}
	if password != res.Password {
		c.JSON(http.StatusConflict, gin.H{
			"message": "invalid login or password",
		})
	}
	// if !etc.CheckPasswordHash(password, res.Password) {
	// 	c.JSON(http.StatusConflict, models.FailureInfo{
	// 		Code:    http.StatusConflict,
	// 		Message: "username or password error",
	// 	})
	// 	h.log.Error("error checking password", err)
	// 	return
	// }
	h.jwthandler.Role = "admin"
	h.jwthandler.Sub = res.Id
	h.jwthandler.Aud = []string{"cliamte.uz"}
	h.jwthandler.SigninKey = h.cfg.SigninKey
	h.jwthandler.Log = h.log
	tokens, err := h.jwthandler.GenerateAuthJWT()
	accessToken := tokens[0]

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.FailureInfo{
			Code:    http.StatusInternalServerError,
			Message: "something went wrong, please try again",
		})
		h.log.Error("error occured while generating tokens ", err)
		return
	}
	res.AccessToken = accessToken
	res.Password = ""
	c.JSON(http.StatusOK, res)
}
