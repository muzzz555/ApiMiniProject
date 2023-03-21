package model

import (
	"fmt"
	"apiminiproject/database"
	"apiminiproject/dateset"
)

type SystemModel struct{}

func (u SystemModel) GetWorkAllfreelance(Type_Work_ID dateset.Typeworkid ) ([]dateset.Getworkpagehomefreelance,error)  {
	var works_freelance []dateset.Getworkpagehomefreelance
	// var allwork dateset.Allwork 
	db := database.Connectdata()
		rows, err := db.Query("select Work_Post_ID,Type_Work_Name, Price_Post_Work, Name_Work, FirstName, LastName, Image_Work_Post_Freelance from user,work_post_freelance,type_work WHERE User_ID=ID AND Type_Work_ID=Type_Work_Number AND Type_Work_ID = ?", Type_Work_ID.Type_Work_ID)
		if err != nil{
			return works_freelance,err
		}
		fmt.Print(rows)

		for rows.Next() {
			var work dateset.Getworkpagehomefreelance

			if err := rows.Scan(&work.Work_Post_ID,&work.Type_Work_Name, &work.Price_Post_Work, &work.Name_Work, &work.FirstName, &work.LastName, &work.Image_Work_Post_Freelance); 
			
			err != nil{
					return works_freelance,err
			}
			fmt.Print(work.Image_Work_Post_Freelance)
			works_freelance = append(works_freelance, work)
		}
		
		return works_freelance,nil
	} 

func (u SystemModel) GetWorkAllCompany(Type_Work_ID dateset.Typeworkid ) ([]dateset.Getworkpagehomecompany,error)  {
	var works_company []dateset.Getworkpagehomecompany

	db := database.Connectdata()

		rows, err := db.Query("select Work_Post_ID,Type_Work_Name, Price_Work_Min, Name_Work, Image_Work_Post_Company, CompanyName from company,work_post_company,type_work WHERE Company_ID=ID AND Type_Work_ID=Type_Work_Number AND Type_Work_ID = ?", Type_Work_ID.Type_Work_ID)
	
		if err != nil{
			return works_company,err
		}

		for rows.Next() {
			var work dateset.Getworkpagehomecompany

			if err := rows.Scan(&work.Work_Post_ID,&work.Type_Work_Name, &work.Price_Work_Min, &work.Name_Work, &work.Image_Work_Post_Company ,&work.CompanyName); 
			
			err != nil{
					return works_company,err
			}
			works_company = append(works_company, work)			
		}
		return works_company,nil
}

func (u SystemModel) GetWorkfreelance(Type_Work_ID dateset.Workpostid ) ([]dateset.Getworkfreelance,error)  {
	var work_freelance []dateset.Getworkfreelance

	db := database.Connectdata()

		rows, err := db.Query("select Price_Post_Work, Name_Work, Detail_Work, Type_Work_Name, FirstName, LastName, Line, Facebook, Instagram, Image_Work_Post_Freelance, Email, Phone from user,work_post_freelance,type_work WHERE User_ID=ID AND Type_Work_ID=Type_Work_Number AND Work_Post_ID = ?", Type_Work_ID.Work_Post_Id)
	
		if err != nil{
			return work_freelance,err
		}

		for rows.Next() {
			var work dateset.Getworkfreelance

			if err := rows.Scan(&work.Price_Post_Work, &work.Name_Work, &work.Detail_Work, &work.Type_Work_Name , &work.FirstName, &work.LastName, &work.Line, &work.Facebook, &work.Instagram, &work.Image_Work_Post_Freelance, &work.Email, &work.Phone ); 
			
			err != nil{
					return work_freelance,err
			}
			work_freelance = append(work_freelance, work)			
		}
		return work_freelance,nil
}

func (u SystemModel) GetWorkcompany(Type_Work_ID dateset.Work_Post_ID) ([]dateset.GetWorkCompany,error) {
	var work_company []dateset.GetWorkCompany

	db := database.Connectdata()

		rows, err := db.Query("select CompanyName, CompanyEmail, CompanyPhone, Type_Work_Name, Position, Num_Person, Price_Work_Min, Price_Work_Max, Education, Detail_Work, Image_Work_Post_Company from work_post_company, type_work, company where Company_ID = ID and Type_Work_ID = Type_Work_Number and Work_Post_ID = ?",Type_Work_ID.Work_Post_ID)

		if err != nil {
			return work_company, err
		}

		for rows.Next() {
			var work dateset.GetWorkCompany

			if err := rows.Scan(&work.CompanyName, &work.CompanyEmail, &work.CompanyPhone, &work.Type_Work_Name, &work.Position, &work.Num_Person, &work.Price_Work_Min, &work.Price_Work_Max, &work.Education, &work.Detail_Work, &work.Image_Work_Post_Company, );

			err != nil {
				return work_company, err
			}
			work_company = append(work_company, work)
		}
		return work_company, nil
}