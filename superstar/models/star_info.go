package models

type StarInfo struct {
	Id           int    `xorm:"not null pk autoincr comment('主键ID') int(10)" form:"id"`
	NameZh       string `xorm:"not null comment('中文名') varchar(50)" form:"name_zh"`
	NameEn       string `xorm:"not null comment('英文名') varchar(50)" form:"name_eh"`
	Avatar       string `xorm:"not null comment('头像') int(10)" form:"avatar"`
	Birthday     string `xorm:"not null comment('出生日期') int(10)" form:"birthday"`
	Height       int    `xorm:"not null default 0 comment('身高(cm)') int(10)" form:"height"`
	Weight       int    `xorm:"not null comment('体重(g)') int(10)" form:"weight"`
	Club         string `xorm:"not null comment('俱乐部') varchar(50)" form:"club"`
	Jersy        string `xorm:"not null comment('球衣号码/主打位置') varchar(50)" form:"jersy"`
	Country      string `xorm:"not null comment('国籍') varchar(50)" form:"country"`
	BirthAddress string `xorm:"not null comment('出生地') varchar(255)" form:"birth_address"`
	Feature      string `xorm:"not null comment('个人特点') varchar(255)" form:"feature"`
	MoreInfo     string `xorm:"comment('更多介绍') text" form:"more_info"`
	Status       int    `xorm:"not null default 0 comment('状态 0:正常 1:删除') int(10)" form:"status"`
	Created      int    `xorm:"not null default 0 comment('创建时间') int(10)"`
	Updated      int    `xorm:"not null default 0 comment('更新时间') int(10)"`
}
