// Code generated by protoapi:go; DO NOT EDIT.

package todolistsvr

import (
	"github.com/labstack/echo"
	"github.com/yoozoo/protoapi/protoapigo"
)

// TodolistService is the interface contains all the controllers
type TodolistService interface {
	Add(c echo.Context, req *AddReq) (resp *AddResp, bizError *AddError, err error)

	List(c echo.Context, req *Empty) (resp *ListResp, err error)
}

func _add_Handler(srv TodolistService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		req := new(AddReq)

		if err = c.Bind(req); err != nil {
			resp := &CommonError{BindError: &BindError{err.Error()}}
			return c.JSON(420, resp)
		}
		/*

			if valErr := req.Validate(); valErr != nil {
				resp := &CommonError{ValidateError: valErr}
				return c.JSON(420, resp)
			}

		*/
		resp, bizError, err := srv.Add(c, req)
		if err != nil {
			// e:= err.(*CommonError) will panic if assertion fail, which is not what we want
			if e, ok := err.(*CommonError); ok {
				return c.JSON(420, e)
			}
			return c.String(500, err.Error())
		}
		if bizError != nil {
			return c.JSON(400, bizError)
		}

		return c.JSON(200, resp)
	}
}
func _list_Handler(srv TodolistService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		req := new(Empty)

		if err = c.Bind(req); err != nil {
			resp := &CommonError{BindError: &BindError{err.Error()}}
			return c.JSON(420, resp)
		}
		/*

			if valErr := req.Validate(); valErr != nil {
				resp := &CommonError{ValidateError: valErr}
				return c.JSON(420, resp)
			}

		*/
		resp, err := srv.List(c, req)
		if err != nil {
			// e:= err.(*CommonError) will panic if assertion fail, which is not what we want
			if e, ok := err.(*CommonError); ok {
				return c.JSON(420, e)
			}
			return c.String(500, err.Error())
		}

		return c.JSON(200, resp)
	}
}

// RegisterTodolistService is used to bind routers
func RegisterTodolistService(e *echo.Echo, srv TodolistService) {
	RegisterTodolistServiceWithPrefix(e, srv, "")
}

// RegisterTodolistServiceWithPrefix is used to bind routers with custom prefix
func RegisterTodolistServiceWithPrefix(e *echo.Echo, srv TodolistService, prefix string) {
	// switch to strict JSONAPIBinder, if using echo's DefaultBinder
	if _, ok := e.Binder.(*echo.DefaultBinder); ok {
		e.Binder = new(protoapigo.JSONAPIBinder)
	}
	e.POST(prefix+"/TodolistService.add", _add_Handler(srv))
	e.POST(prefix+"/TodolistService.list", _list_Handler(srv))
}
