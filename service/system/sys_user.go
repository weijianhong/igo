package system

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gofrs/uuid/v5"
	"github.com/weijianhong/igo/global"
	"github.com/weijianhong/igo/model/system"
	"github.com/weijianhong/igo/utils"
	"go.uber.org/zap"
)

type UserService struct{}

func (userService *UserService) Register(u system.SysUser) (userInter system.SysUser, err error) {
	var user system.SysUser
	err = global.DB.QueryRowContext(context.Background(),
		`SELECT password FROM sys_users WHERE username = ?`, u.Username).
		Scan(&user.Password)
	if err == nil {
		if errors.Is(err, sql.ErrNoRows) {
			return system.SysUser{}, errors.New("用户名已注册")
		}
		return user, err
	}
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		global.LOG.Error("", zap.Error(err))

		return system.SysUser{}, err
	}

	// 否则 附加uuid 密码hash加密 注册
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.Must(uuid.NewV4())
	_, err = global.DB.Exec("insert into sys_users (`username`,`password`,`uuid`) values (?,?)",
		u.Username, u.Password, u.UUID)
	if err != nil {
		return system.SysUser{}, err
	}
	return u, err
}

func (userService *UserService) Login(u *system.SysUser) (userInter *system.SysUser, err error) {
	if nil == global.DB {
		return nil, fmt.Errorf("db not init")
	}

	var user system.SysUser
	err = global.DB.QueryRowContext(context.Background(),
		`SELECT password,enable FROM sys_users WHERE username = ?`, u.Username).
		Scan(&user.Password, &user.Enable)
	if err != nil {
		return nil, err
	}

	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return nil, errors.New("密码错误")
	}

	//MenuServiceApp.UserAuthorityDefaultRouter(&user)
	return &user, err
}

//@function: ChangePassword
//@description: 修改用户密码
//@param: u *model.SysUser, newPassword string
//@return: userInter *model.SysUser,err error

//func (userService *UserService) ChangePassword(u *system.SysUser, newPassword string) (userInter *system.SysUser, err error) {
//	var user system.SysUser
//	if err = global.GVA_DB.Where("id = ?", u.ID).First(&user).Error; err != nil {
//		return nil, err
//	}
//	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
//		return nil, errors.New("原密码错误")
//	}
//	user.Password = utils.BcryptHash(newPassword)
//	err = global.GVA_DB.Save(&user).Error
//	return &user, err
//
//}

//@function: GetUserInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

//func (userService *UserService) GetUserInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
//	limit := info.PageSize
//	offset := info.PageSize * (info.Page - 1)
//	db := global.GVA_DB.Model(&system.SysUser{})
//	var userList []system.SysUser
//	err = db.Count(&total).Error
//	if err != nil {
//		return
//	}
//	err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
//	return userList, total, err
//}

//@function: SetUserAuthority
//@description: 设置一个用户的权限
//@param: uuid uuid.UUID, authorityId string
//@return: err error

//func (userService *UserService) SetUserAuthority(id uint, authorityId uint) (err error) {
//	assignErr := global.GVA_DB.Where("sys_user_id = ? AND sys_authority_authority_id = ?", id, authorityId).First(&system.SysUserAuthority{}).Error
//	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
//		return errors.New("该用户无此角色")
//	}
//	err = global.GVA_DB.Where("id = ?", id).First(&system.SysUser{}).Update("authority_id", authorityId).Error
//	return err
//}

//@function: SetUserAuthorities
//@description: 设置一个用户的权限
//@param: id uint, authorityIds []string
//@return: err error

//func (userService *UserService) SetUserAuthorities(id uint, authorityIds []uint) (err error) {
//	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
//		TxErr := tx.Delete(&[]system.SysUserAuthority{}, "sys_user_id = ?", id).Error
//		if TxErr != nil {
//			return TxErr
//		}
//		var useAuthority []system.SysUserAuthority
//		for _, v := range authorityIds {
//			useAuthority = append(useAuthority, system.SysUserAuthority{
//				SysUserId: id, SysAuthorityAuthorityId: v,
//			})
//		}
//		TxErr = tx.Create(&useAuthority).Error
//		if TxErr != nil {
//			return TxErr
//		}
//		TxErr = tx.Where("id = ?", id).First(&system.SysUser{}).Update("authority_id", authorityIds[0]).Error
//		if TxErr != nil {
//			return TxErr
//		}
//		// 返回 nil 提交事务
//		return nil
//	})
//}

//@function: DeleteUser
//@description: 删除用户
//@param: id float64
//@return: err error

//func (userService *UserService) DeleteUser(id int) (err error) {
//	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
//		if err := tx.Where("id = ?", id).Delete(&system.SysUser{}).Error; err != nil {
//			return err
//		}
//		if err := tx.Delete(&[]system.SysUserAuthority{}, "sys_user_id = ?", id).Error; err != nil {
//			return err
//		}
//		return nil
//	})
//}

//@function: SetUserInfo
//@description: 设置用户信息
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser
//
//func (userService *UserService) SetUserInfo(req system.SysUser) error {
//	return global.GVA_DB.Model(&system.SysUser{}).
//		Select("updated_at", "nick_name", "header_img", "phone", "email", "sideMode", "enable").
//		Where("id=?", req.ID).
//		Updates(map[string]interface{}{
//			"updated_at": time.Now(),
//			"nick_name":  req.NickName,
//			"header_img": req.HeaderImg,
//			"phone":      req.Phone,
//			"email":      req.Email,
//			"side_mode":  req.SideMode,
//			"enable":     req.Enable,
//		}).Error
//}

//@function: SetUserInfo
//@description: 设置用户信息
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser

//func (userService *UserService) SetSelfInfo(req system.SysUser) error {
//	return global.GVA_DB.Model(&system.SysUser{}).
//		Where("id=?", req.ID).
//		Updates(req).Error
//}

//@function: GetUserInfo
//@description: 获取用户信息
//@param: uuid uuid.UUID
//@return: err error, user system.SysUser

//func (userService *UserService) GetUserInfo(uuid uuid.UUID) (user system.SysUser, err error) {
//	var reqUser system.SysUser
//	err = global.GVA_DB.Preload("Authorities").Preload("Authority").First(&reqUser, "uuid = ?", uuid).Error
//	if err != nil {
//		return reqUser, err
//	}
//	MenuServiceApp.UserAuthorityDefaultRouter(&reqUser)
//	return reqUser, err
//}

//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.SysUser

//func (userService *UserService) FindUserById(id int) (user *system.SysUser, err error) {
//	var u system.SysUser
//	err = global.GVA_DB.Where("id = ?", id).First(&u).Error
//	return &u, err
//}

//@function: FindUserByUuid
//@description: 通过uuid获取用户信息
//@param: uuid string
//@return: err error, user *model.SysUser

//func (userService *UserService) FindUserByUuid(uuid string) (user *system.SysUser, err error) {
//	var u system.SysUser
//	if err = global.GVA_DB.Where("uuid = ?", uuid).First(&u).Error; err != nil {
//		return &u, errors.New("用户不存在")
//	}
//	return &u, nil
//}

//@function: resetPassword
//@description: 修改用户密码
//@param: ID uint
//@return: err error

//func (userService *UserService) ResetPassword(ID uint) (err error) {
//	err = global.GVA_DB.Model(&system.SysUser{}).Where("id = ?", ID).Update("password", utils.BcryptHash("123456")).Error
//	return err
//}
