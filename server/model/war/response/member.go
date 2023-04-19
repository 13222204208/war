package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/war"

type LoginResponse struct {
	User      war.Member `json:"user"`
	Token     string     `json:"token"`
	ExpiresAt int64      `json:"expiresAt"`
}
