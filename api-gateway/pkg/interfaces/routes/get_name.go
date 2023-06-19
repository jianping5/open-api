package routes

import (
	// "context"
	"api-gateway/pkg/interfaces/pb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetNameRequestBody struct {
	Id int64 `json:"id"`
	
}

func GetName(ctx *gin.Context, c pb.InterfacesServiceClient) {
	body := GetNameRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// userId, _ := ctx.Get("userId")

	res, err := c.GetName(context.Background(), &pb.GetNameRequest{
		Id: body.Id,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

