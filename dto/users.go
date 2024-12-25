package dto

type ReqUserChangeEmail struct{}
type ResUserChangeEmail struct{}
type ReqUserChangePassword struct{}
type ResUserChangePassword struct{}
type ReqUserDelete struct{}
type ResUserDelete struct{}
type ReqUserDetail struct{}
type ResUserDetail struct{}
type ReqUserDisable struct{}
type ResUserDisable struct{}
type ReqUserList struct{}
type ResUserList struct{}
type ReqUserLogin struct{}
type ResUserLogin struct{}
type ReqUserRegister struct {
	Email       string `validate:"required,email,max=128" json:"email"`
	Password    string `validate:"required,min=8" json:"password"`
	DisplayName string `validate:"required,max=128" json:"displayName"`
}

type ResUserRegister struct {
	ID string `json:"id"`
}
type ReqUserUpdate struct{}
type ResUserUpdate struct{}
