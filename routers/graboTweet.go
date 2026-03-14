package routers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/harolpg17/twitterGo/bd"
	"github.com/harolpg17/twitterGo/models"
)

func GraboTweet(ctx context.Context, claim models.Claim) models.RespApi {
	var mensaje models.Tweet
	var r models.RespApi
	r.Status = 400
	IDUsuario := claim.ID.Hex()

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &mensaje)
	if err != nil {
		r.Message = "Error en el formato del mensaje"
		return r
	}

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		r.Message = "Error al insertar el tweet"
		return r
	}

	if !status {
		r.Message = "No se ha logrado insertar el tweet"
		return r
	}

	r.Status = 200
	r.Message = "Tweet insertado correctamente"
	return r

}
