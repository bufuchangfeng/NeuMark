package models

import "time"

//用户
//学号 姓名 性别 年级 院系 专业 校区 班级 头像
//用户状态		 用法：0 / 1      禁止 / 允许
//手机号 微信号 QQ 用法：联系方式
//最后登录时间     用法：xx时间前在线
//创建时间        用法：已加入 NeuMark xx 天

type User struct {
	ID int `gorm:"primary_key"`
	Sid string
	Name string
	Sex string
	Grade string
	Institute string
	Major string
	Campus string
	Class string

	// 头像
	// 用户状态

	Phone string
	WeChat string
	QQ	string

	LoginAt time.Time
	CreatedAt time.Time

	Goods_ []Goods
	Comments []Comment
	Buys []Buy
}

