package handlers

import "github.com/hertz-contrib/jwt"

var jwtAuthMiddleware *jwt.HertzJWTMiddleware

func SetAuthMiddleware(h *jwt.HertzJWTMiddleware) {
	jwtAuthMiddleware = h
}
