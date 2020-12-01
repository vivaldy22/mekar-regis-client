package middleware

import (
	"github.com/vivaldy22/mekar-regis-client/tools/consts"
	"github.com/vivaldy22/mekar-regis-client/tools/jwtm"
)

var AdminJWTMiddleware = jwtm.NewJWTMiddleware(consts.HMACADM)