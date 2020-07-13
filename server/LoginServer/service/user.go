package service

import (
	utils2 "KServer/library/utils"
	"KServer/manage"
	"KServer/server/utils"
	"KServer/server/utils/pd"
	"fmt"
	"github.com/kataras/iris/v12"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Account *pd.Account
	Manage  manage.IManage
	Encrypt *utils2.Encrypt
}

func NewUser(manage manage.IManage) *User {
	u := User{}
	u.Manage = manage
	return &u
}
func (u *User) PreHandler(ctx iris.Context) {
	ctx.Next()
}

func (u *User) AccountRegister(ctx iris.Context) {
	if err := ctx.ReadJSON(&u.Account); err != nil {
		_, _ = ctx.JSON(iris.Map{"state": "fail", "msg": "系统错误"})
		return
	}
	if len(u.Account.Account) < 1 || len(u.Account.Account) < 1 {
		_, _ = ctx.JSON(iris.Map{"state": "fail", "msg": "账号或密码不能为空"})
		return
	}
	coll := u.Manage.DB().Mongo().GetCollection("user_account")
	err := coll.Find(bson.M{"account": u.Account.Account}).One(&u.Account)
	if err == nil {
		_, _ = ctx.JSON(iris.Map{"state": "fail", "msg": "账号已存在"})
		return
	}
	u.Account.UUID = u.Encrypt.NewUuid()
	u.Account.Token = u.Encrypt.NewToken()
	err = coll.Insert(&u.Account)
	if err != nil {
		_, _ = ctx.JSON(iris.Map{"state": "fail", "msg": err.Error()})
		return
	}
	key := utils.ClientLoginInfoKey + u.Account.UUID
	_, _ = u.Manage.DB().Redis().GetMasterConn().Set(key).ProtoBuf(u.Account)
	//_, _ = u.Redis.SetValueByProto(key, &u.Account)
	u.Account.PassWord = "******" //  返回隐藏密码
	_, _ = ctx.JSON(iris.Map{"state": "success", "msg": "注册成功", "result": u.Account})

}
func (u *User) AccountLogin(ctx iris.Context) {
	if err := ctx.ReadJSON(&u.Account); err != nil {
		_, _ = ctx.JSON(iris.Map{"state": "fail", "msg": "系统错误"})
		return
	}
	if len(u.Account.Account) < 1 || len(u.Account.Account) < 1 {
		_, _ = ctx.JSON(iris.Map{"state": "fail", "msg": "账号或密码不能为空"})
		return
	}
	//key := utils.ClientLoginInfoKey + u.Account.UUID
	//fmt.Println(u.Account.UUID)

	//dbUser := &pd.Account{}

	/*
		if u.Manage.DB().Redis().GetSlaveConn().Get(key).ProtoBuf(dbUser) == nil {
			if dbUser.Account == u.Account.Account && dbUser.PassWord == u.Account.PassWord {
				if dbUser.Online == utils.ClientOnline {
					//dbUser.Token = u.Encrypt.NewToken()
					_, _ = ctx.JSON(iris.Map{"state": "fail", "msg": "账号已在线"})
					return
				}
				dbUser.Token = u.Encrypt.NewToken()
				_, err := u.Manage.DB().Redis().GetMasterConn().Set(utils.ClientLoginInfoKey + u.Account.UUID).ProtoBuf(dbUser)
				if err != nil {
					fmt.Println("err1=", err)
				}
				//_, _ = u.Redis.SetValueByProto(key, dbUser)
				dbUser.PassWord = "******" //  返回隐藏密码
				fmt.Println(u.Account.UUID)

				_, _ = ctx.JSON(iris.Map{"state": "success", "msg": "登陆成功", "result": dbUser})
			} else {
				_, _ = ctx.JSON(iris.Map{"state": "fail", "msg": "密码错误"})
			}
		} else {

	*/
	coll := u.Manage.DB().Mongo().GetCollection("user_account")
	loginPassWord := u.Account.PassWord
	err := coll.Find(bson.M{"account": u.Account.Account}).One(&u.Account)
	if err != nil {
		_, _ = ctx.JSON(iris.Map{"state": "fail", "msg": "账号不存在"})
		return
	}
	if loginPassWord != u.Account.PassWord {
		_, _ = ctx.JSON(iris.Map{"state": "fail", "msg": "密码错误"})
		return
	}
	if u.Account.Online == utils.ClientOnline {
		_, _ = ctx.JSON(iris.Map{"state": "fail", "msg": "账号已在线"})
		return
	}
	u.Account.Token = u.Encrypt.NewToken()
	_ = coll.UpData(bson.M{"account": u.Account.Account, "PassWord": u.Account.PassWord}, &u.Account)
	_, err = u.Manage.DB().Redis().GetMasterConn().Set(utils.ClientLoginInfoKey + u.Account.UUID).ProtoBuf(u.Account)
	if err != nil {
		fmt.Println("err2=", err)
	}
	//_, _ = u.Redis.SetValueByProto(key, &u.Account)
	u.Account.PassWord = "******" //  返回隐藏密码
	fmt.Println(u.Account.UUID)

	_, _ = ctx.JSON(iris.Map{"state": "success", "msg": "登陆成功", "result": u.Account})

}
