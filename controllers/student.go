package controllers

import (
	"echo-lab-go/models"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CreateStudentScore ..
func CreateStudentScore(c echo.Context) error {
	studentModel := models.StudentScoreName{}
	if err := c.Bind(&studentModel); err != nil {
		log.Fatal(err)
	}

	reqMap := map[string]interface{}{}
	if err := c.Bind(&reqMap); err != nil {
		log.Fatal(err)
	}

	fmt.Println("studentModel: ", studentModel)
	fmt.Println("reqMap: ", reqMap)

	if err := models.DB.Create(&studentModel).Error; err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK, map[string]interface{}{"success": false, "message": "Create StudentScore error"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true})
}

// GetStudentScore ..
func GetStudentScore(c echo.Context) error {
	studentModel := models.StudentScoreName{}
	if err := c.Bind(&studentModel); err != nil {
		log.Fatal(err)
	}

	reqMap := map[string]interface{}{}
	if err := c.Bind(&reqMap); err != nil {
		log.Fatal(err)
	}

	fmt.Println("studentModel: ", studentModel)
	fmt.Println("reqMap: ", reqMap)

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true})
}

// GetAllStudentScore ..
func GetAllStudentScore(c echo.Context) error {
	studentModel := models.StudentScoreName{}
	if err := c.Bind(&studentModel); err != nil {
		log.Fatal(err)
	}

	reqMap := map[string]interface{}{}
	if err := c.Bind(&reqMap); err != nil {
		log.Fatal(err)
	}

	fmt.Println("studentModel: ", studentModel)
	fmt.Println("reqMap: ", reqMap)

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true})
}

// UpdateStudentScore ..
func UpdateStudentScore(c echo.Context) error {
	studentModel := models.StudentScoreName{}
	if err := c.Bind(&studentModel); err != nil {
		log.Fatal(err)
	}

	reqMap := map[string]interface{}{}
	if err := c.Bind(&reqMap); err != nil {
		log.Fatal(err)
	}

	fmt.Println("studentModel: ", studentModel)
	fmt.Println("reqMap: ", reqMap)

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true})
}

// DeleteStudentScore ..
func DeleteStudentScore(c echo.Context) error {
	studentModel := models.StudentScoreName{}
	if err := c.Bind(&studentModel); err != nil {
		log.Fatal(err)
	}

	reqMap := map[string]interface{}{}
	if err := c.Bind(&reqMap); err != nil {
		log.Fatal(err)
	}

	fmt.Println("studentModel: ", studentModel)
	fmt.Println("reqMap: ", reqMap)

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true})
}

// RecoverStudentScore ..
func RecoverStudentScore(c echo.Context) error {
	studentModel := models.StudentScoreName{}
	if err := c.Bind(&studentModel); err != nil {
		log.Fatal(err)
	}

	reqMap := map[string]interface{}{}
	if err := c.Bind(&reqMap); err != nil {
		log.Fatal(err)
	}

	fmt.Println("studentModel: ", studentModel)
	fmt.Println("reqMap: ", reqMap)

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true})
}

// GetSubjectGrade ..
func GetSubjectGrade(c echo.Context) error {
	studentModel := models.StudentScoreName{}
	if err := c.Bind(&studentModel); err != nil {
		log.Fatal(err)
	}

	reqMap := map[string]interface{}{}
	if err := c.Bind(&reqMap); err != nil {
		log.Fatal(err)
	}

	fmt.Println("studentModel: ", studentModel)
	fmt.Println("reqMap: ", reqMap)

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true})
}

// GetSubjectGradeByStudentName ..
func GetSubjectGradeByStudentName(c echo.Context) error {
	studentModel := models.StudentScoreName{}
	if err := c.Bind(&studentModel); err != nil {
		log.Fatal(err)
	}

	reqMap := map[string]interface{}{}
	if err := c.Bind(&reqMap); err != nil {
		log.Fatal(err)
	}

	fmt.Println("studentModel: ", studentModel)
	fmt.Println("reqMap: ", reqMap)

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true})
}
