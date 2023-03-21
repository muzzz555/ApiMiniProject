package controller

import (
	"apiminiproject/dateset"
	"apiminiproject/model"
	"github.com/gin-gonic/gin"
)

type systemController struct{}




func (u UserController) GetAllWork(c *gin.Context) {
	systemModel := model.SystemModel{}
	var typeworkid dateset.Typeworkid

	e := c.ShouldBind(&typeworkid)

	if e != nil {
		panic(e)
	}

	if typeworkid.Select_ID == 1 {
		var allwork dateset.AllWorkFreelance
		works,err := systemModel.GetWorkAllfreelance(typeworkid)

		if err != nil {
			panic(err)
		}
		allwork.Allwork = works;
		c.JSON(200,allwork)

	} else {
		var allwork dateset.AllWorkCompany
		works,err := systemModel.GetWorkAllCompany(typeworkid)

		if err != nil {
			panic(err)
		}
		allwork.Allwork = works;
		c.JSON(200, allwork)
	}
}

func (u UserController) GetWorkfreelance(c *gin.Context) {
	systemModel := model.SystemModel{}
	var workpostid dateset.Workpostid

	e := c.ShouldBind(&workpostid)

	if e != nil {
		panic(e)
	}

		work,err := systemModel.GetWorkfreelance(workpostid)

		if err != nil {
			panic(err)
		}
		c.JSON(200, work)
}

func (u UserController) GetWorkcompany(c *gin.Context) {
	systemModel := model.SystemModel{}
	var workpostid dateset.Work_Post_ID

	e := c.ShouldBind(&workpostid)

	if e != nil {
		panic(e)
	}

	work, err := systemModel.GetWorkcompany(workpostid)

	if err != nil {
		panic(err)
	}
	c.JSON(200, work)
}

