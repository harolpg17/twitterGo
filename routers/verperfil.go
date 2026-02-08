package routers

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/harolpg17/twitterGo/bd"
	"github.com/harolpg17/twitterGo/models"
)

func VerPerfil(request events.APIGatewayProxyRequest) models.RespApi {
	var r models.RespApi
	r.Status = 400

	fmt.Println("Entre en VerPerfil")
	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "El id es obligatorio"
		return r
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		r.Message = "Error al buscar el perfil" + err.Error()
		return r
	}

	respJson, err := json.Marshal(perfil)
	if err != nil {
		r.Status = 500
		r.Message = "Error al convertir el perfil a json" + err.Error()
		return r
	}

	r.Status = 200
	r.Message = string(respJson)
	return r
}
