package routes

import (
	// "context"
	"api-gateway/pkg/interfaces/pb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddUserInterfaceInfoRequest struct {
	UserId string `json:"userId"`
	InterfaceInfoId string `json:"interfaceInfoId"`
}

func GetName(ctx *gin.Context, c pb.InterfacesServiceClient) {
	body := AddUserInterfaceInfoRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// userId, _ := ctx.Get("userId")

	res, err := c.GetName(context.Background(), &pb.AddUserInterfaceInfoRequest{
		UserId: body.UserId,
		InterfaceInfoId: body.InterfaceInfoId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

