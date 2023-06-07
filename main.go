package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/public", "public")

	e.GET("/hello", helloword)
	e.GET("/", home)
	e.GET("/addProject", addProject)
	e.GET("/projeect-detail", projectDetail)
	e.GET("/contactMe", contactMe)
	e.POST("/addFormProject", addFormProject)

	e.Logger.Fatal(e.Start("localhost:8000"))
}

func helloword(c echo.Context) error {
	return c.String(http.StatusOK, "helloworld")
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

func addProject(c echo.Context) error {
	var template, err = template.ParseFiles("views/addProject.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return template.Execute(c.Response(), nil)
}

func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"id":           id,
		"projectName":  "Peserta Dumbways Indonesia",
		"durasi":       "1 bulan",
		"technologies": []string{"Golang", "Node JS", "React JS", "Pyhton"},
		"description":  "Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.",
	}
	var template, err = template.ParseFiles("views/add-project-detail.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return template.Execute(c.Response(), data)
}

func contactMe(c echo.Context) error {
	var template, err = template.ParseFiles("views/contact-me.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return template.Execute(c.Response(), nil)
}

func addFormProject(c echo.Context) error {
	projectName := c.FormValue("projectName")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	nodeJS := c.FormValue("tech-one")
	reactJS := c.FormValue("tech-two")
	pyhton := c.FormValue("tech-three")
	Golang := c.FormValue("tech-four")
	description := c.FormValue("desc")

	println("Project name : " + projectName)
	println("start date : " + startDate)
	println("end date : " + endDate)
	println("Technologies : " + nodeJS + reactJS + pyhton + Golang)
	println("description : " + description)

	return c.Redirect(http.StatusMovedPermanently, "/")
}
