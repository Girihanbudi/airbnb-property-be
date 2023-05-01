package rest

import (
	"airbnb-property-be/env/appcontext"
	"airbnb-property-be/internal/app/property/preset/request"
	transutil "airbnb-property-be/internal/app/translation/util"
	"airbnb-property-be/internal/pkg/http"
	"airbnb-property-be/internal/pkg/stderror"
	stdresponse "airbnb-property-be/internal/pkg/stdresponse/rest"

	"github.com/gin-gonic/gin"
)

func (h Handler) CreatePropertyType(ctx *gin.Context) {
	clientLocale := appcontext.GetLocale(ctx.Request.Context())

	var req request.CreatePropertyType
	if bindErr := ctx.ShouldBind(&req); bindErr != nil {
		err := transutil.TranslateError(ctx, stderror.DEF_DATA_400, clientLocale)
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}

	bFile, parseErr := http.ParseFileHeaderAsBytes(*req.File)
	if parseErr != nil {
		err := transutil.TranslateError(ctx, stderror.DEF_DATA_400, clientLocale)
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}
	req.BFile = bFile

	h.Property.CreatePropertyType(ctx, req)
}
