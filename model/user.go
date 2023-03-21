package model

import (
	"apiminiproject/database"
	"fmt"
	"strings"
)

type UserModel struct{}

func (u UserModel) CreateUser(firstname string,
	lastname string,
	email string,
	password string,
	phone string) (string, error) {

	db := database.Connectdata()

	stmt, err := db.Prepare("insert into user (FirstName, LastName, Email, Password, Phone) values(?,?,?,?,?)")
	if err != nil {
		fmt.Print(err)
	}

	_, err = stmt.Exec(
		firstname,
		lastname,
		email,
		password,
		phone)
	check := "Complete!"

	if err != nil {
		check = "Fail"
	}
	return check, nil
}

func (u UserModel) LoginUser(email string,
	password string, id int) (string, int, int, error) {

	db := database.Connectdata()

	findemail := "@"
	statuslogin := "ชื่อผู้ใช้งานหรือรหัสผ่านไม่ถูกต้องอีเมล"
	var iddb int
	var TypeNumber_User int
	if strings.Contains(email, findemail) {
		var passworddb string
		err := db.QueryRow("select Password, ID, TypeNumber_User from user WHERE Email=? union select Password, ID, TypeNumber_User from company WHERE CompanyEmail=?", email, email).Scan(&passworddb, &iddb, &TypeNumber_User)

		if err != nil {
			fmt.Print(err)
		}

		if passworddb == password {
			statuslogin = "ถูกต้อง"
		}

	} else {
		var passworddb string
		err := db.QueryRow("select Password, ID, TypeNumber_User from user WHERE Phone=? union select Password, ID, TypeNumber_User from company WHERE CompanyPhone=?", email, email).Scan(&passworddb, &iddb, &TypeNumber_User)

		if err != nil {
			fmt.Print(err)
		}

		fmt.Print(passworddb + " pass")
		if passworddb == password {
			statuslogin = "ถูกต้อง"
		}
	}
	return statuslogin, iddb, TypeNumber_User, nil
}

func (u UserModel) Updateuser(idint int,
	firstname string,
	lastname string,
	email string,
	phone string,
	profile_user string,
) (string, error) {

	db := database.Connectdata()
	_, err := db.Exec("UPDATE user SET FirstName = ?, LastName = ?, Email = ?, Phone = ?, Profile_user = ? WHERE id = ?", firstname, lastname, email, phone, profile_user, idint)

	if err != nil {
		fmt.Print(err)
	}
	statusUpdateuser := "สำเร็จ"
	return statusUpdateuser, nil
}
func (u UserModel) GetUpdateuser(id int) (string, string, string, string, string) {
	db := database.Connectdata()
	var FirstName string
	var LastName string
	var Email string
	var Phone string
	var Profile_user string
	err := db.QueryRow("select FirstName,LastName,Email,Phone,Profile_user from user WHERE id = ?", id).Scan(&FirstName, &LastName, &Email, &Phone, &Profile_user)
	if err != nil {
		fmt.Print(err)
	}
	return FirstName, LastName, Email, Phone, Profile_user
}
func (u UserModel) AddWorkFreelance(WorkPostID int,
	TypeWorkNumber int,
	UserID int,
	DetailWork string,
	PricePostWork string,
	NameWork string,
	ImageWorkPostFreelance string) (string, error) {

	db := database.Connectdata()

	stmt, err := db.Prepare("insert into work_post_freelance (Work_Post_ID, Type_Work_Number, User_ID, Detail_Work, Price_Post_Work,Name_Work,Image_Work_Post_Freelance) values(?,?,?,?,?,?,?)")
	if err != nil {
		fmt.Print(err)
	}

	_, err = stmt.Exec(
		WorkPostID,
		TypeWorkNumber,
		UserID,
		DetailWork,
		PricePostWork,
		NameWork,
		ImageWorkPostFreelance)
	check := "Complete!"

	if err != nil {
		check = "Fail"
	}
	return check, nil
}

func (u UserModel) Updatepostfreelance(workpostid int,
	typeWorknumber int,
	detailwork string,
	userid int,
	pricepostwork string,
	namework string,
	imageworkpostfreelance string,
) (string, error) {

	db := database.Connectdata()

	_, err := db.Exec("UPDATE work_post_freelance SET Type_Work_Number = ?, Detail_Work = ?, Price_Post_Work = ?, Name_Work = ?, Image_Work_Post_Freelance = ? WHERE Work_Post_ID = ?", typeWorknumber, detailwork, pricepostwork, namework, imageworkpostfreelance, workpostid)

	if err != nil {
		fmt.Print(err)
	}
	statusUpdateuser := "สำเร็จ"
	return statusUpdateuser, nil
}

func (u UserModel) UpdateFreelance(
	id int,
	firstname string,
	lastname string,
	email string,
	phone string,
	profileuser string,
	line string,
	facebook string,
	instagram string) (string, error) {

	db := database.Connectdata()
	TypeNumber_User := 1
	_, err := db.Exec("UPDATE user SET FirstName = ?, LastName = ?, Email = ?, Phone = ?, Profile_User = ?, Line = ?, Facebook = ?, Instagram = ? , TypeNumber_User = ? WHERE id = ?", firstname, lastname, email, phone, profileuser, line, facebook, instagram, TypeNumber_User, id)

	if err != nil {
		fmt.Print(err)
	}
	statusUpdateFreelance := "สำเร็จ"
	return statusUpdateFreelance, nil
}

func (u UserModel) GetUpdatefreelance(id int) (string, string, string, string, string, string, string, string) {
	fmt.Print(id)
	db := database.Connectdata()
	var FirstName string
	var LastName string
	var Email string
	var Phone string
	var Profile_user string
	var Line string
	var Facebook string
	var Instagram string
	err := db.QueryRow("select FirstName,LastName,Email,Phone,Profile_user,Line,Facebook,Instagram from user WHERE id = ?", id).Scan(&FirstName, &LastName, &Email, &Phone, &Profile_user, &Line, &Facebook, &Instagram)
	if err != nil {
		fmt.Print(err)
	}
	return FirstName, LastName, Email, Phone, Profile_user, Line, Facebook, Instagram
}
