package controller

import (
	"apiminiproject/dateset"
	"apiminiproject/model"
	"github.com/gin-gonic/gin"
)

type companyController struct{}

func (u UserController) CreateAccountCompany(c *gin.Context) {
	var companyModel model.CompanyModel
	var company dateset.Company

	e := c.ShouldBind(&company)
	if e != nil {
		// fmt.Println(e)
		panic(e)
	}

	status, err := companyModel.CreateAccountCompany(
		company.CompanyName,
		company.CompanyEmail,
		company.CompanyPhone,
		company.Address,
		company.Subdistrict,
		company.District,
		company.Province,
		company.Postcode,
		company.Password)

	if err != nil {
		panic(err)
	}
	c.JSON(200, status)
}

func (u UserController) AddWorkCompany(c *gin.Context) {
	var companyModel model.CompanyModel
	var company dateset.AddWorkCompany

	e := c.ShouldBind(&company)
	if e != nil {
		// fmt.Println(e)w
		panic(e)
	}

	status, err := companyModel.AddWorkCompany(
		company.CompanyID,
		company.TypeWorkNumber,
		company.NameWork,
		company.DetailWork,
		company.Position,
		company.NumPerson,
		company.PriceWorkMin,
		company.PriceWorkMax,
		company.Education,
		company.ImageWorkPostCompany)

	if err != nil {
		panic(err)
	}
	c.JSON(200, status)
}

func (u UserController) Updatepostcompany(c *gin.Context) {
	var companyModel model.CompanyModel
	var company dateset.AddWorkCompany

	e := c.ShouldBind(&company)
	if e != nil {
		panic(e)
	}

	statusUpdatecompany, err := companyModel.Updatepostcompany(
		company.WorkPostID,
		company.TypeWorkNumber,
		company.NameWork,
		company.DetailWork,
		company.Position,
		company.NumPerson,
		company.PriceWorkMin,
		company.PriceWorkMax,
		company.Education,
		company.ImageWorkPostCompany,
	)

	if err != nil {
		panic(err)
	}
	c.JSON(200, statusUpdatecompany)
}
func (u UserController) Updatecompany(c *gin.Context) {
	var userModel model.UserModel
	var user dateset.UpdateCompany

	e := c.ShouldBind(&user)
	if e != nil {
		panic(e)
	}

	statusUpdateuser, err := userModel.Updatecompany(
		user.ID,
		user.CompanyName,
		user.CompanyEmail,
		user.CompanyPhone,
		user.Address,
		user.Subdistrict,
		user.District,
		user.Province,
		user.Postcode,
		user.ProfileCompany)

	if err != nil {
		panic(err)
	}
	c.JSON(200, statusUpdateuser)
}
func (u UserController) GetUpdatecompany(c *gin.Context) {
	var companyModel model.CompanyModel
	var company dateset.GetCompany
	var id dateset.ID

	e := c.ShouldBind(&id)

	if e != nil {
		panic(e)
	}

	CompanyName, CompanyEmail, CompanyPhone, Address, Subdistrict, District, Province, Postcode, ProfileCompany := companyModel.GetUpdatecompany(
		id.ID)
	company.CompanyName = CompanyName
	company.CompanyEmail = CompanyEmail
	company.CompanyPhone = CompanyPhone
	company.Address = Address
	company.Subdistrict = Subdistrict
	company.District = District
	company.Province = Province
	company.Postcode = Postcode
	company.ProfileCompany = ProfileCompany

	c.JSON(200, company)
}
