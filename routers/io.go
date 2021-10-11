package routers

import (
	"net/http"
	"wallet/controller"
)

var IO = []Router{
	{
		URI:                    "/all",
		Method:                 http.MethodGet,
		Function:               controller.GetAllIO,
		AuthenticationRequired: false,
	},
}
