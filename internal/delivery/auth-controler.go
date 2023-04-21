package delivery

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/muhangga/internal/entity"
	"github.com/muhangga/internal/service"
	"github.com/muhangga/internal/utils"
)

type authDelivery struct {
	authService service.AuthService
	jwtService  service.JwtService
}

func NewAuthDelivery(authService service.AuthService, jwtService service.JwtService) *authDelivery {
	return &authDelivery{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (d *authDelivery) Login(c *gin.Context) {
	var loginDTO entity.LoginDTO
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		resp := utils.JsonError(http.StatusInternalServerError, "invalid json", err.Error(), utils.EmptyObject{})
		c.JSON(resp.Status, resp)
		return
	}

	validate := validator.New()
	if err := validate.Struct(loginDTO); err != nil {
		resp := utils.ValidationResponse(err)
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	user, err := d.authService.Login(loginDTO)
	if err != nil {
		resp := utils.JsonError(http.StatusInternalServerError, "Failed login", err.Error(), utils.EmptyObject{})
		c.JSON(resp.Status, resp)
		return
	}

	resp := utils.JsonSuccess(http.StatusOK, "login success", user)
	c.JSON(resp.Status, resp)
}

func (d *authDelivery) Register(c *gin.Context) {
	var registerDTO entity.RegisterDTO

	if err := c.ShouldBindJSON(&registerDTO); err != nil {
		resp := utils.JsonError(http.StatusInternalServerError, "invalid json", err.Error(), utils.EmptyObject{})
		c.JSON(resp.Status, resp)
		return
	}

	validate := validator.New()
	if err := validate.Struct(registerDTO); err != nil {
		resp := utils.ValidationResponse(err)
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	if !d.authService.IsEmailAvailable(registerDTO.Email) {
		resp := utils.JsonError(http.StatusBadRequest, "email already taken", errors.New("email already taken").Error(), utils.EmptyObject{})
		c.JSON(resp.Status, resp)
		return
	}

	user, err := d.authService.Register(registerDTO)
	if err != nil {
		resp := utils.JsonError(http.StatusInternalServerError, "invalid json", err.Error(), utils.EmptyObject{})
		c.JSON(resp.Status, resp)
		return
	}

	resp := utils.JsonSuccess(http.StatusCreated, "register success", user)
	c.JSON(resp.Status, resp)
}
