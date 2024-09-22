package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/akiradomi/workspace/go-practice/back/model"
	"github.com/akiradomi/workspace/go-practice/back/usecase"
	"github.com/labstack/echo/v4"
)

type UserControllerInterface interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
	CsrfToken(c echo.Context) error
}

type userController struct {
	uu usecase.UserUsecaseInterface
}

// コンストラクタ,main.goから呼び出される
func NewUserController(uu usecase.UserUsecaseInterface) UserControllerInterface {
	return &userController{uu}
}

func (uc *userController) SignUp(c echo.Context) error {
	//ゼロ値のUser構造体を定義
	user := model.User{}
	//routerから渡された値を、ゼロ値のUser構造体にバインド
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	//user_usecaseのSignUpメソッドの呼び出し
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	//返却されたuserResponse構造体をreturn
	return c.JSON(http.StatusCreated, userRes)

}

func (uc *userController) Login(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	//user_usecaseのLoginメソッドの呼び出し、jwtトークンが返却される
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	//クッキーセット
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = "localhost"
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uc *userController) Logout(c echo.Context) error {
	//クッキー初期化
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = "localhost"
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uc *userController) CsrfToken(c echo.Context) error {
	log.Println(c.Get("csrf"))
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}
