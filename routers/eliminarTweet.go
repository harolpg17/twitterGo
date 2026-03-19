package routers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/harolpg17/twitterGo/bd"
	"github.com/harolpg17/twitterGo/models"
)

func EliminarTweet(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400

	IDUsuario := claim.ID.Hex()
	IDTweet := request.QueryStringParameters["id"]

	if len(IDTweet) == 0 {
		r.Message = "ID of the tweet is required"
		return r
	}

	err := bd.BorroTweet(IDTweet, IDUsuario)
	if err != nil {
		r.Message = "Error deleting the tweet"
		return r
	}

	r.Message = "Tweet deleted successfully"
	r.Status = 200
	return r
}
