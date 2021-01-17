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

//SetName 设置 name
func (b *UserBuilder) SetName(name string) *UserBuilder {
	b.name = name

	return b
}

//SetSex 设置 sex
func (b *UserBuilder) SetSex(sex int) *UserBuilder {
	b.sex = sex

	return b
}
