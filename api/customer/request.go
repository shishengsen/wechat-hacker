package customer

type Request struct {

}

type LoginInput struct {
	Username string `form:"username" json:"username" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

type RegisterInput struct {
	UserName   string `form:"UserName" json:"userName" validate:"required"`
	Phone      string `form:"Phone" json:"phone" validate:"required"`
	Password   string `form:"Password" json:"password" validate:"required"`
	Role       int    `form:"Role" json:"role" validate:"required"`
	CustomerId int    `form:"CustomerId" json:"customerId" validate:"required"`
	CompanyIds string `form:"CompanyIds" json:"companyIds" validate:"required"`
	ApiToken   string `form:"ApiToken" json:"apiToken" validate:"required"`
}

