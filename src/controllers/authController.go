package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Login struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type Loginsucess struct {
    Token string `json:"token"`
}



func AuthLogin(c *gin.Context) {
	//var loginData Login

    //user, err := models.VerifyUser(loginData.Username, loginData.Password)
	/*
    if err != nil {

        c.JSON(http.StatusUnauthorized, utils.TypeErrorResponse{
            Error: "Credenciais inválidas",
        })
        return
    }
    if user == nil {
        c.JSON(http.StatusInternalServerError, utils.TypeErrorResponse{
            Error: "Usuário não encontrado",
        })
        return
    }

    token, err := utils.GenerateJWT(user.ID)
    if err != nil {

        c.JSON(http.StatusInternalServerError, utils.TypeErrorResponse{
            Error: "Falha ao gerar o token",
        })
        return
    }

    c.JSON(http.StatusOK, Loginsucess{
        Token: token,
    })*/

	c.JSON(http.StatusOK, "ok")

}
