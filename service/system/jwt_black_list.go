package system

import (
	"context"
	"github.com/weijianhong/igo/global"
	"github.com/weijianhong/igo/model/system"
	"github.com/weijianhong/igo/utils"
)

type JwtService struct{}

//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error

func (jwtService *JwtService) JsonInBlacklist(jwtList system.JwtBlacklist) (err error) {

	_, err = global.DB.Exec("insert into ? (`jwt`) values (?)", jwtList.TableName(), jwtList.Jwt)
	if err != nil {
		return err
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool

func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	_, ok := global.BlackCache.Get(jwt)
	return ok
}

//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: userName string
//@return: redisJWT string, err error

func (jwtService *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = global.REDIS.Get(context.Background(), userName).Result()
	return redisJWT, err
}

//@function: SetRedisJWT
//@description: jwt存入redis并设置过期时间
//@param: jwt string, userName string
//@return: err error

func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	dr, err := utils.ParseDuration(global.CONFIG.JWT.ExpiresTime)
	if err != nil {
		return err
	}
	timer := dr
	err = global.REDIS.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

func LoadAll() {
	var data []string

	var jwtList system.JwtBlacklist

	rows, err := global.DB.Query("SELECT jwt FROM ?", jwtList.TableName())
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var jwt string
		if err = rows.Scan(&jwt); err != nil {
			return
		}
		data = append(data, jwt)
	}

	if errRow := rows.Err(); errRow != nil {
		return
	}

	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	}

	// jwt黑名单 加入 BlackCache 中
}
