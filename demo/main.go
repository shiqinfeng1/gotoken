package main

import (
	"fmt"
	"gotoken"
	"gotoken/jwt"
)

func main() {
	tkSrv := gotoken.New(jwt.NewJwtBuilder(),
		gotoken.WithExpired(3),
		gotoken.WithSecret([]byte("iam secret")),
		gotoken.WithIssuer("issu1"),
		gotoken.WithSubject("sub1"),
		gotoken.WithID("uuid"),
		gotoken.WithAudience([]string{"audi1", "audi2"}),
		gotoken.WithSignMethod(gotoken.JwtSignHS256),
	)
	token, err := tkSrv.Generate()
	if err != nil {
		panic(err)
	}
	fmt.Println("get new jwt token:", token)
}
