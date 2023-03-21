package controller

import (
	"apiminiproject/dateset"
	"apiminiproject/model"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type UserController struct{}

func (u UserController) CreateUser(c *gin.Context) {
	var userModel model.UserModel
	var user dateset.User

	e := c.ShouldBind(&user)
	if e != nil {
		// fmt.Println(e)
		panic(e)
	}

	status, err := userModel.CreateUser(
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.Phone)

	if err != nil {
		panic(err)
	}
	c.JSON(200, status)
}

func (u UserController) LoginUser(c *gin.Context) {
	var userModel model.UserModel
	var user dateset.User
	var uservalue dateset.LoginUser
	e := c.ShouldBind(&user)
	if e != nil {
		panic(e)
	}

	statuslogin, id,TypeNumber_User, err := userModel.LoginUser(
		user.Email, user.Password, user.ID)

	uservalue.ID = id
	uservalue.Status = statuslogin
	uservalue.TypeNumber_User = TypeNumber_User
	if err != nil {
		panic(err)
	}
	c.JSON(200, uservalue)
}

func (u UserController) Updateuser(c *gin.Context) {
	var userModel model.UserModel
	var user dateset.UpdateUser

	e := c.ShouldBind(&user)
	if e != nil {
		panic(e)
	}

	statusUpdateuser, err := userModel.Updateuser(
		user.ID,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.Profile_user)

	if err != nil {
		panic(err)
	}
	c.JSON(200, statusUpdateuser)
}
func (u UserController) GetUpdateuser(c *gin.Context) {
	var userModel model.UserModel
	var uservalue dateset.GetUser
	var id dateset.ID

	e := c.ShouldBind(&id)

	if e != nil {
		panic(e)
	}

	FirstName, LastName, Email, Phone, Profile_user := userModel.GetUpdateuser(
		id.ID)
	uservalue.FirstName = FirstName
	uservalue.LastName = LastName
	uservalue.Email = Email
	uservalue.Phone = Phone
	uservalue.Profile_user = Profile_user

	c.JSON(200, uservalue)
}
func (u UserController) AddWorkFreelance(c *gin.Context) {
	var userModel model.UserModel
	var user dateset.AddWorkFreelance

	e := c.ShouldBind(&user)
	if e != nil {
		panic(e)
	}

	status, err := userModel.AddWorkFreelance(
		user.WorkPostID,
		user.TypeWorkNumber,
		user.UserID,
		user.DetailWork,
		user.PricePostWork,
		user.NameWork,
		user.ImageWorkPostFreelance)

	if err != nil {
		panic(err)
	}
	c.JSON(200, status)
}

func (u UserController) UpdateFreelance(c *gin.Context) {
	var userModel model.UserModel
	var user dateset.UpdateFreelance

	e := c.ShouldBind(&user)
	if e != nil {
		panic(e)
	}

	statusUpdateFreelance, err := userModel.UpdateFreelance(
		user.ID,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.Profile_User,
		user.Line,
		user.Facebook,
		user.Instagram)

	if err != nil {
		panic(err)
	}
	c.JSON(200, statusUpdateFreelance)
}
func (u UserController) Updatepostfreelance(c *gin.Context) {
	var userModel model.UserModel
	var user1 dateset.AddWorkFreelance

	e := c.ShouldBind(&user1)
	if e != nil {
		panic(e)
	}

	statusUpdateuser, err := userModel.Updatepostfreelance(
		user1.WorkPostID,
		user1.TypeWorkNumber,
		user1.DetailWork,
		user1.UserID,
		user1.PricePostWork,
		user1.NameWork,
		user1.ImageWorkPostFreelance,
	)

	if err != nil {
		panic(err)
	}
	c.JSON(200, statusUpdateuser)
}

func (u UserController) GetUpdatefreelance(c *gin.Context) {
	var userModel model.UserModel
	var uservalue dateset.GetFreelance
	var id dateset.ID
	e := c.ShouldBind(&id)

	if e != nil {
		panic(e)
	}

	FirstName, LastName, Email, Phone, Profile_user, Line, Facebook, Instagram := userModel.GetUpdatefreelance(
		id.ID)
	uservalue.FirstName = FirstName
	uservalue.LastName = LastName
	uservalue.Email = Email
	uservalue.Phone = Phone
	uservalue.Profile_User = Profile_user
	uservalue.Line = Line
	uservalue.Facebook = Facebook
	uservalue.Instagram = Instagram

	c.JSON(200, uservalue)

}
