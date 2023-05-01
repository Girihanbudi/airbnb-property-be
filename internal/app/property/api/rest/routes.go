package rest

import authmid "airbnb-property-be/internal/app/middleware/auth"

func (h Handler) RegisterApi() {
	sessions := h.Router.Group("/sessions")
	{
		sessions.GET("/google", authmid.GinValidateNoJwtTokenFound, h.CreatePropertyType)
	}
}
