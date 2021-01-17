//go:generate ghf
//+build ghf

package example

//User 用户
type User struct {
}

//UserBuilder 用户
//@build
type UserBuilder struct {
	name string
	sex  int
}

//Build 建造者方法
func (u *UserBuilder) Build() *User {
	return nil
}
