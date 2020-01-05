package models

type UserInfo struct {
	Id         int    `xorm:"not null pk autoincr comment('主键ID') INT(10)"`
	Name       string `xorm:"not null default '' comment('中文名') VARCHAR(50)"`
	SysCreated int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
	SysUpdated int    `xorm:"not null default 0 comment('最后修改时间') INT(10)"`
}
