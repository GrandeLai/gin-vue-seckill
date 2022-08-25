package controller

import (
	"context"
	dao "gin-vue-seckill/jdaw_user_srv/dao/user"
	"gin-vue-seckill/jdaw_user_srv/proto/admin_user"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type Admin_user struct{}

// Return a new handler
//func New() *Admin_user {
//	return &Admin_user{}
//}

func (au *Admin_user) AdminUserLogin(ctx context.Context, in *admin_user.AdminUserRequest, out *admin_user.AdminUserResponse) error {
	username := in.Username
	password := in.Password
	user, err := dao.GetAdminUser(username)
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
	out.UserName = username
	return nil
}

var timeLayoutStr = "2006-01-02 15:04:05"

func (au *Admin_user) GetFrontUserList(ctx context.Context, in *admin_user.FrontUserListRequest, out *admin_user.FrontUserListResponse) error {
	cp := in.CurrentPage
	ps := in.PageSize
	currentPage := int(cp)
	pageSize := int(ps)
	frontuserlist, err, count := dao.GetFrontUserList(currentPage, pageSize)
	if err != nil {
		out.Code = 200
		out.Msg = "当前无用户"
		return err
	} else {
		out.Code = 200
		out.Msg = "已查到用户"
		frontList := make([]*admin_user.FrontUserDetail, 0, len(frontuserlist))
		for _, frontuser := range frontuserlist {
			frontuserre := &admin_user.FrontUserDetail{
				Email:      frontuser.Email,
				Status:     strconv.Itoa(frontuser.Status),
				Desc:       frontuser.Desc,
				CreateTime: frontuser.CreateTime.Format(timeLayoutStr),
			}
			frontList = append(frontList, frontuserre)
		}
		out.Frontuserlist = frontList
		out.PageSize = int32(pageSize)
		out.Current = int32(currentPage)
		out.Total = int32(count)
		return nil
	}
}
