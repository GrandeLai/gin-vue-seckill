package controller

import (
	"context"
	"fmt"
	utils2 "gin-vue-seckill/jdaw_user_srv/common/utils"
	"gin-vue-seckill/jdaw_user_srv/common/utils/snowflake"
	dao "gin-vue-seckill/jdaw_user_srv/dao/user"
	"gin-vue-seckill/jdaw_user_srv/models"
	"gin-vue-seckill/jdaw_user_srv/proto/front_user"
	"github.com/patrickmn/go-cache"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Front_user struct{}

// Return a new handler
//func New() *Front_user {
//	return &Front_user{}
//}

var c = cache.New(60*time.Second, 10*time.Second)

//用户注册
func (fu *Front_user) FrontUserRegister(ctx context.Context, in *front_user.FrontUserRequest, out *front_user.FrontUserResponse) error {
	//校验验证码是否正确

	code, is_ok := c.Get(in.Email)
	if is_ok {
		if code != in.Code {
			out.Code = 500
			out.Msg = "邮箱验证码错误"
		} else {
			id, err := snowflake.GetID()
			if err != nil {
				return err
			}
			user := models.FrontUser{
				UserId:     int64(id),
				Email:      in.Email,
				Password:   in.Password,
				Status:     1,
				CreateTime: time.Now(),
			}
			fmt.Println(user)
			if err := dao.CreateFrontUser(&user); err != nil {
				return err
			} else {
				out.Code = 200
				out.Msg = "注册成功，跳转至登录页面"
			}
			//data_source.Db.Create(&user)
			//fmt.Println(err)
			//out.Code = 200
			//out.Msg = "注册成功，跳转至登录页面"
		}
	} else {
		out.Code = 500
		out.Msg = "注册失败，请重新尝试"
	}

	return nil
}

//发送邮件
func (fu *Front_user) FrontUserSendEmail(ctx context.Context, in *front_user.FrontUserMailRequest, out *front_user.FrontUserResponse) error {
	email := in.Mail
	user, _ := dao.GetFrontUser(email)
	if user != nil {
		out.Code = 500
		out.Msg = "当前邮箱已存在，请选择其他邮箱"
		return nil
	}
	randnum := utils2.GenRandNum(6)
	utils2.SendEmail(email, randnum)
	//生成的验证码放在session中
	//初始化
	c.Set(email, randnum, cache.DefaultExpiration)
	out.Code = 200
	out.Msg = "发送成功"
	return nil
}

func (fu *Front_user) FrontUserLogin(ctx context.Context, in *front_user.FrontUserRequest, out *front_user.FrontUserResponse) error {
	email := in.Email
	password := in.Password
	user, err := dao.GetFrontUser(email)
	if user == nil {
		out.Code = 500
		out.Msg = "当前用户不存在"
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		out.Code = 500
		out.Msg = "密码错误"
		return err
	}
	out.Code = 200
	out.Msg = "登录成功"
	out.UserName = email
	return nil
}
