package middleware

import (
	"context"
	"github.com/ZhaoJun-hz/go-web-base/api"
	"github.com/ZhaoJun-hz/go-web-base/global"
	"github.com/ZhaoJun-hz/go-web-base/global/constants"
	"github.com/ZhaoJun-hz/go-web-base/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

const (
	ERROR_CODE_INVALID_TOKEN = 10401
	TOKEN_NAME               = "Authorization"
	TOKEN_PREFIX             = "Bearer "
)

func tokenErr(c *gin.Context) {
	api.Fail(c, api.Response{
		Status: http.StatusUnauthorized,
		Code:   ERROR_CODE_INVALID_TOKEN,
		Msg:    "token invalid",
	})
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(TOKEN_NAME)
		// token  不存在就直接返回
		if token == "" || !strings.HasPrefix(token, TOKEN_PREFIX) {
			tokenErr(c)
			return
		}
		// token 无法解析，直接返回
		token = token[len(TOKEN_PREFIX):]
		jwtCustomClaims, err := utils.ParseToken(token)
		userId := jwtCustomClaims.ID
		if err != nil || userId == 0 {
			tokenErr(c)
			return
		}

		// Token 与访问者登录对应的Token不一致，直接返回
		userIdStr := strconv.Itoa(int(userId))
		//strings.ReplaceAll(constants.LOGIN_USER_TOKEN_REDIS_KEY, "{ID}", strconv.Itoa(int(user.ID)))

		tokenRedis := global.RedisClient.Get(context.Background(),
			strings.ReplaceAll(constants.LOGIN_USER_TOKEN_REDIS_KEY, "{ID}", userIdStr))
		if tokenRedis.Err() != nil && token != tokenRedis.Val() {
			tokenErr(c)
			return
		}

		// 过期？？
		c.Set("LOGIN_USER_ID", userId)
		c.Next()
	}

}
