// This file was auto-generated.
// DO NOT EDIT MANUALLY!!!
package api

import (
	"github.com/labstack/echo"
	"github.com/midnight-trigger/todo/api/controller"
	"github.com/midnight-trigger/todo/third_party/jwt"
)

func RegisterRoutes(e *echo.Echo) {
	PostSigninUser(e, &controller.User{})
	PostUser(e, &controller.User{})
}
func RegisterAuthRoutes(e *echo.Group) {
	PostTodo(e, &controller.Todo{})
}
func PostSigninUser(
	e *echo.Echo,
	inter *controller.User,
) {
	e.POST("api/v1/users/signin", func(c echo.Context) error {
		res := inter.PostSigninUser(c)
		return c.JSON(res.Meta.Code, res)
	})
}
func PostUser(
	e *echo.Echo,
	inter *controller.User,
) {
	e.POST("api/v1/users", func(c echo.Context) error {
		res := inter.PostUser(c)
		return c.JSON(res.Meta.Code, res)
	})
}
func PostTodo(
	e *echo.Group,
	inter *controller.Todo,
) {
	e.POST("api/v1/todos", func(c echo.Context) error {
		claims, r := jwt.GetJWTClaims(c)
		if claims == nil {
			return c.JSON(r.Code, r)
		}
		res := inter.PostTodo(c, claims)
		return c.JSON(res.Meta.Code, res)
	})
}
