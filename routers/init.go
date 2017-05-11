package routers

import (
	"../auth"
)

var (
	SecretKey = auth.GetSecretKey()
)