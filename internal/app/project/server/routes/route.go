package routes

import echoSwagger "github.com/swaggo/echo-swagger"

func Add(params Params) {
	for _, handler := range params.Handlers {
		handler.AddRoute(params.Echo)
	}

	params.Echo.GET("v1/swagger/*", echoSwagger.WrapHandler)
}
