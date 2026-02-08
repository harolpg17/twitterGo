package routers

import (
	"context"
	"encoding/json"

	"github.com/harolpg17/twitterGo/bd"
	"github.com/harolpg17/twitterGo/models"
)

func ModificarPerfil(ctx context.Context, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400

	var t models.Usuario
	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = "Error al procesar el body" + err.Error()
		return r
	}

	status, err := bd.ModificoRegistro(t, claim.ID.Hex())
	if err != nil {
		r.Message = "Error al modificar el registro " + err.Error()
		return r
	}

	if !status {
		r.Message = "No se pudo modificar el registro"
		return r
	}

	r.Status = 200
	r.Message = "Registro modificado correctamente"
	return r
}
