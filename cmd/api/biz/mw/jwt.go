
package mw

import (
	"context"
	"net/http"
	"time"
<<<<<<< HEAD
	// "encoding/json"
=======
	"encoding/json"
>>>>>>> origin/guo
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/api/biz/model/api"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/api/biz/rpc"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/user"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/consts"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
<<<<<<< HEAD
	// jwtv4 "github.com/golang-jwt/jwt/v4"
=======
	jwtv4 "github.com/golang-jwt/jwt/v4"
>>>>>>> origin/guo
)

var JwtMiddleware *jwt.HertzJWTMiddleware

func InitJWT() {
	JwtMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		Key:           []byte(consts.SecretKey),
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		IdentityKey:   consts.IdentityKey,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
<<<<<<< HEAD
			return &api.User{
				ID: int64(claims[consts.IdentityKey].(float64)),
=======
			
			userid, _ := claims[consts.IdentityKey].(json.Number).Int64()
			// log.Println("^^^^^^^^userid:%v", userid)
			return &api.User{
				ID: userid,
>>>>>>> origin/guo
			}
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					consts.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var err error
			var req api.LoginUserRequest
			if err = c.BindAndValidate(&req); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			if len(req.Username) == 0 || len(req.Password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}
			return rpc.LoginUser(context.Background(), &user.LoginUserRequest{
				Username: req.Username,
				Password: req.Password,
			})
		},
<<<<<<< HEAD
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, utils.H{
				"code":   errno.Success.ErrCode,
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			})
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"code":    errno.AuthorizationFailedErr.ErrCode,
				"message": message,
=======
		LoginResponse: func(ctx context.Context, c *app.RequestContext, status_code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, utils.H{
				"status_code":   errno.Success.ErrCode,
				"status_msg":  errno.Success.ErrMsg,
				"user_id":     0,
				"token":  token,
				// "expire": expire.Format(time.RFC3339),
			})
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, status_code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"status_code":    errno.AuthorizationFailedErr.ErrCode,
				"status_msg":  message,
				
>>>>>>> origin/guo
			})
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			switch t := e.(type) {
			case errno.ErrNo:
				return t.ErrMsg
			default:
				return t.Error()
			}
		},
<<<<<<< HEAD
	})
}
=======
		ParseOptions: []jwtv4.ParserOption{jwtv4.WithJSONNumber()},
	})
}
>>>>>>> origin/guo
