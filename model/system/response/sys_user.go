package response

import "github.com/weijianhong/igo/model/system"

type SysUserResponse struct {
	User system.SysUser `json:"user"`
	//system.System
}

type LoginResponse struct {
	User      system.SysUser `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}
