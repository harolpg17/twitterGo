package routers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/harolpg17/twitterGo/bd"
	"github.com/harolpg17/twitterGo/jwt"
	"github.com/harolpg17/twitterGo/models"
)

func Login(ctx context.Context) models.RespApi {
	var t models.Usuario
	var r models.RespApi
	r.Status = 400

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = "Error al procesar el body"
		return r
	}

	if len(t.Email) == 0 {
		r.Message = "Email is required"
		return r
	}

	if len(t.Password) == 0 {
		r.Message = "Password is required"
		return r
	}

	userData, existe := bd.IntentoLogin(t.Email, t.Password)
	if !existe {
		r.Message = "Email or password is incorrect"
		return r
	}

	jwtKey, err := jwt.GeneroJWT(ctx, userData)
	if err != nil {
		r.Message = "Error al intentar generar el token correspondiente " + err.Error()
		return r
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	token, err2 := json.Marshal(resp)
	if err2 != nil {
		r.Message = "Error al intentar formatear el token correspondiente a JSON " + err2.Error()
		return r
	}

	cookie := &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: time.Now().Add(24 * time.Hour),
	}

	cookieString := cookie.String()

	res := &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(token),
		Headers: map[string]string{
			"Content-type":               "application/json",
			"Acess-Control-Allow-Origin": "*",
			"Set-Cookie":                 cookieString,
		},
	}

	r.Status = 200
	r.Message = string(token)
	r.CustomResp = res
	return r
}
