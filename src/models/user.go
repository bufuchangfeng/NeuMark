package models

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

	// for shop
	Password string
	CategoryID int
	// 头像
	// 用户状态

	Phone string
	WeChat string
	QQ	string

	Sells []Sell	`gorm:"FOREIGNKEY:UserID;ASSOCIATION_FOREIGNKEY:ID"`
	Comments []Comment `gorm:"FOREIGNKEY:UserID;ASSOCIATION_FOREIGNKEY:ID"`
	Buys []Buy		`gorm:"FOREIGNKEY:UserID;ASSOCIATION_FOREIGNKEY:ID"`
	Collects []Collect `gorm:"FOREIGNKEY:UserID;ASSOCIATION_FOREIGNKEY:ID"`

	Usertype string

}

