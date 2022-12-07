package api

import (
	"unicode/utf8"

	"github.com/enzanumo/ky-theater-web/internal/model"
	"github.com/enzanumo/ky-theater-web/internal/service"
	"github.com/enzanumo/ky-theater-web/pkg/app"
	"github.com/enzanumo/ky-theater-web/pkg/errcode"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Login 用户登录
func Login(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	user, err := service.DoLogin(c, &param)
	if err != nil {
		logrus.Errorf("service.DoLogin err: %v", err)
		response.ToErrorResponse(err.(*errcode.Error))
		return
	}

	token, err := app.GenerateToken(user)
	if err != nil {
		logrus.Errorf("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})
}

// Register 用户注册
func Register(c *gin.Context) {

	param := service.RegisterRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 用户名检查
	err := service.ValidUsername(param.Username)
	if err != nil {
		logrus.Errorf("service.Register err: %v", err)
		response.ToErrorResponse(err.(*errcode.Error))
		return
	}

	// 密码检查
	err = service.CheckPassword(param.Password)
	if err != nil {
		logrus.Errorf("service.Register err: %v", err)
		response.ToErrorResponse(err.(*errcode.Error))
		return
	}

	user, err := service.Register(
		param.Username,
		param.Password,
	)

	if err != nil {
		logrus.Errorf("service.Register err: %v", err)
		response.ToErrorResponse(errcode.UserRegisterFailed)
		return
	}

	response.ToResponse(gin.H{
		"id":       user.ID,
		"username": user.Username,
	})
}

// 获取用户基本信息
func GetUserInfo(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)

	if username, exists := c.Get("USERNAME"); exists {
		param.Username = username.(string)
	}

	user, err := service.GetUserInfo(&param)

	if err != nil {
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	phone := ""
	if user.Phone != "" && len(user.Phone) == 11 {
		phone = user.Phone[0:3] + "****" + user.Phone[7:]
	}

	response.ToResponse(gin.H{
		"id":       user.ID,
		"nickname": user.Nickname,
		"username": user.Username,
		"status":   user.Status,
		"balance":  user.Balance,
		"phone":    phone,
		"is_admin": user.IsAdmin,
	})
}

// 修改密码
func ChangeUserPassword(c *gin.Context) {
	param := service.ChangePasswordReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	user := &model.User{}
	if u, exists := c.Get("USER"); exists {
		user = u.(*model.User)
	}

	// 密码检查
	err := service.CheckPassword(param.Password)
	if err != nil {
		logrus.Errorf("service.Register err: %v", err)
		response.ToErrorResponse(err.(*errcode.Error))
		return
	}

	// 旧密码校验
	if !service.ValidPassword(user.Password, param.OldPassword, user.Salt) {
		response.ToErrorResponse(errcode.ErrorOldPassword)
		return
	}

	// 更新入库
	user.Password, user.Salt = service.EncryptPasswordAndSalt(param.Password)
	service.UpdateUserInfo(user)

	response.ToResponse(nil)
}

// 修改昵称
func ChangeNickname(c *gin.Context) {
	param := service.ChangeNicknameReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	user := &model.User{}
	if u, exists := c.Get("USER"); exists {
		user = u.(*model.User)
	}

	if utf8.RuneCountInString(param.Nickname) < 2 || utf8.RuneCountInString(param.Nickname) > 12 {
		response.ToErrorResponse(errcode.NicknameLengthLimit)
		return
	}

	// 执行绑定
	user.Nickname = param.Nickname
	service.UpdateUserInfo(user)

	response.ToResponse(nil)
}

// 用户绑定手机号
func BindUserPhone(c *gin.Context) {
	param := service.UserPhoneBindReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	user := &model.User{}
	if u, exists := c.Get("USER"); exists {
		user = u.(*model.User)
	}

	// 手机重复性检查
	if service.CheckPhoneExist(user.ID, param.Phone) {
		response.ToErrorResponse(errcode.ExistedUserPhone)
		return
	}

	if err := service.CheckPhoneCaptcha(param.Phone, param.Captcha); err != nil {
		logrus.Errorf("service.CheckPhoneCaptcha err: %v\n", err)
		response.ToErrorResponse(err)
		return
	}

	// 执行绑定
	user.Phone = param.Phone
	service.UpdateUserInfo(user)

	response.ToResponse(nil)
}

// 修改用户状态
func ChangeUserStatus(c *gin.Context) {
	param := service.ChangeUserStatusReq{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		logrus.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	if param.Status != model.UserStatusNormal && param.Status != model.UserStatusClosed {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	user, err := service.GetUserByID(param.ID)
	if err != nil {
		logrus.Errorf("service.GetUserByID err: %v\n", err)
		response.ToErrorResponse(errcode.NoExistUsername)
		return
	}

	// 执行更新
	user.Status = param.Status
	service.UpdateUserInfo(user)

	response.ToResponse(nil)
}

func GetUserProfile(c *gin.Context) {
	response := app.NewResponse(c)
	username := c.Query("username")

	user, err := service.GetUserByUsername(username)
	if err != nil {
		logrus.Errorf("service.GetUserByUsername err: %v\n", err)
		response.ToErrorResponse(errcode.NoExistUsername)
		return
	}

	response.ToResponse(gin.H{
		"id":       user.ID,
		"nickname": user.Nickname,
		"username": user.Username,
		"status":   user.Status,
		"is_admin": user.IsAdmin,
	})
}

func GetUserWalletBills(c *gin.Context) {
	response := app.NewResponse(c)

	userID, _ := c.Get("UID")
	bills, totalRows, err := service.GetUserWalletBills(userID.(int64), (app.GetPage(c)-1)*app.GetPageSize(c), app.GetPageSize(c))

	if err != nil {
		logrus.Errorf("service.GetUserWalletBills err: %v\n", err)
		response.ToErrorResponse(errcode.GetCollectionsFailed)
		return
	}

	response.ToResponseList(bills, totalRows)
}

func userFrom(c *gin.Context) (*model.User, bool) {
	if u, exists := c.Get("USER"); exists {
		user, ok := u.(*model.User)
		return user, ok
	}
	return nil, false
}

func userIdFrom(c *gin.Context) (int64, bool) {
	if u, exists := c.Get("UID"); exists {
		uid, ok := u.(int64)
		return uid, ok
	}
	return -1, false
}
