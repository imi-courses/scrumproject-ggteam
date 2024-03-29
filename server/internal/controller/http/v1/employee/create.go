package employee

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sethvargo/go-password/password"

	"github.com/imi-courses/scrumproject-ggteam/server/internal/controller/http/v1/exception"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/dto"
	"github.com/imi-courses/scrumproject-ggteam/server/internal/entity"
)

type createRequest dto.CreateEmployeeWithoutPassword

type createResponse struct {
	*entity.Employee
	Password string `json:"password"`
}

func (r *route) create(c *gin.Context) {
	var body createRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		exception.BadRequest(c, err.Error())
		return
	}

	_, err := r.ue.FindOne(c, entity.Employee{Email: body.Email})
	if err == nil {
		exception.BadRequest(
			c,
			fmt.Sprintf("A employee with this email %s already exist", body.Email),
		)
		return
	}

	password, err := password.Generate(8, 8, 0, false, false)
	if err != nil {
		exception.InternalServerError(c, err.Error())
		return
	}

	hashedPassword, err := r.uh.HashPassword(password)
	if err != nil {
		exception.InternalServerError(c, err.Error())
		return
	}

	createEmployee := dto.CreateEmployee{
		Firstname:  body.Firstname,
		Middlename: body.Middlename,
		Surname:    body.Surname,
		Email:      body.Email,
		Password:   hashedPassword,
	}

	employee, err := r.ue.Create(c.Request.Context(), createEmployee)
	if err != nil {
		r.l.Error(err.Error())
		exception.InternalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, createResponse{
		Employee: employee,
		Password: password,
	})
}
