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
	studentModel := models.StudentScore{}
	if err := c.Bind(&studentModel); err != nil {
		log.Fatal(err)
	}

	fmt.Println("studentModel: ", studentModel)

	if err := models.SQLiteDB.Create(&studentModel).Error; err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusOK, map[string]interface{}{"success": false, "message": "Create StudentScore error"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true})
}

// GetStudentScore ..
func GetStudentScore(c echo.Context) error {
	studentModel := models.StudentScore{}
	if err := c.Bind(&studentModel); err != nil {
		log.Fatal(err)
	}

	fmt.Println("studentModel: ", studentModel)

	// TODO: GetStudentScore

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true})
}

// GetAllStudentScore ..
func GetAllStudentScore(c echo.Context) error {
	studentModel := models.StudentScore{}
	if err := c.Bind(&studentModel); err != nil {
		log.Fatal(err)
	}

	fmt.Println("studentModel: ", studentModel)

	// TODO: GetAllStudentScore

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true})
}

// UpdateStudentScore ..
func UpdateStudentScore(c echo.Context) error {
	studentModel := models.StudentScore{}
	if err := c.Bind(&studentModel); err != nil {
		log.Fatal(err)
	}

	fmt.Println("studentModel: ", studentModel)

	// TODO: UpdateStudentScore

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true})
}

// DeleteStudentScore ..
func DeleteStudentScore(c echo.Context) error {
	studentModel := models.StudentScore{}
	if err := c.Bind(&studentModel); err != nil {
		log.Fatal(err)
	}

	fmt.Println("studentModel: ", studentModel)

	// TODO: DeleteStudentScore

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true})
}

// RecoverStudentScore ..
func RecoverStudentScore(c echo.Context) error {
	studentModel := models.StudentScore{}
	if err := c.Bind(&studentModel); err != nil {
		log.Fatal(err)
	}

	fmt.Println("studentModel: ", studentModel)

	// TODO: RecoverStudentScore

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true})
}

// GetSubjectGrade ..
func GetSubjectGrade(c echo.Context) error {
	studentModel := models.StudentScore{}
	if err := c.Bind(&studentModel); err != nil {
		log.Fatal(err)
	}

	fmt.Println("studentModel: ", studentModel)

	// TODO: GetSubjectGrade

	// Fix grade: A: 81-100, B: 71-80, C: 61-70, D: 51-60, F: 0-50
	/* expected response:
	[
		{
			"name": "Tong",
			"subject": "eng",
			"grade": "A"
		},
		{
			"name": "Tong",
			"subject": "math",
			"grade": "B"
		},
		{
			"name": "Tong2",
			"subject": "eng",
			"grade": "C"
		}
	]
	*/

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true})
}

// GetSubjectGradeByStudentName ..
func GetSubjectGradeByStudentName(c echo.Context) error {
	studentModel := models.StudentScore{}
	if err := c.Bind(&studentModel); err != nil {
		log.Fatal(err)
	}

	fmt.Println("studentModel: ", studentModel)

	// TODO: GetSubjectGradeByStudentName

	/* expected response:
	{
		"eng": "A",
		"math": "B",
		"social": "C"
	}
	*/

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true})
}
