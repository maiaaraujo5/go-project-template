package routes

func Add(params Params) {

	for _, handler := range params.Handlers {
		handler.AddRoute(params.Echo)
	}
}
