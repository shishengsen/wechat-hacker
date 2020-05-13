package customer

type Request struct {

}

type LoginInput struct {
	Username string `form:"username" json:"username" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

type RegisterInput struct {
	UserName   string `form:"userName" json:"userName" validate:"required"`
	Phone      string `form:"phone" json:"phone" validate:"required"`
	Password   string `form:"password" json:"password" validate:"required"`
	Role       int    `form:"role" json:"role" validate:"required"`
	CustomerId int    `form:"customerId" json:"customerId" validate:"required"`
	CompanyIds string `form:"companyIds" json:"companyIds" validate:"required"`
	ApiToken   string `form:"apiToken" json:"apiToken" validate:"required"`
}

type ModifyPasswordInput struct {
	Username string `form:"username" json:"username" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
	NewPassword string `form:"newPassword" json:"newPassword" validate:"required"`
}