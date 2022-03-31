package service

import (
	"log"

	"github.com/eze-insurance/database"
	"github.com/eze-insurance/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	mail "github.com/xhit/go-simple-mail/v2"
)

const (
	Philip = "Philip"
)

var htmlBody = `
		<html>
			<head>
				<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
				<title>Hello, World</title>
			</head>
			<body>
				<p>This is an email using Go</p>
			</body>
		`

func CreateShuttlersPolicy(c *fiber.Ctx) error {
	shuttlers := new(model.ShuttlersRequest)

	err := c.BodyParser(shuttlers)

	if err != nil {
		log.Fatal(err)
	}

	err = validate.Struct(shuttlers)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {

			errormsg := err.StructNamespace() + " " + err.Tag()
			errors = append(errors, errormsg)
		}

		return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"error":  errors,
			"status": false})
	}

	shuttlerStruct := model.Shuttlers{
		FullName:         shuttlers.FullName,
		Address:          shuttlers.Address,
		Business:         shuttlers.Business,
		MaritalStatus:    shuttlers.MaritalStatus,
		DateOfBirth:      shuttlers.DateOfBirth,
		PhoneNumber:      shuttlers.PhoneNumber,
		Email:            shuttlers.Email,
		NextOfKin:        shuttlers.NextOfKin,
		NextOfKinAddress: shuttlers.NextOfKin,
		NextOfKinPhone:   shuttlers.NextOfKinPhone,
	}

	db := database.DbConnect()

	result := db.Create(&shuttlerStruct)

	if result.Error != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error":  result.Error,
			"status": false})
	}

	// Create email
	email := mail.NewMSG()
	email.SetFrom("From Me <me@host.com>")
	email.AddTo("you@example.com")
	//email.AddCc("another_you@example.com")
	email.SetSubject("New Go Email")

	email.SetBody(mail.TextHTML, htmlBody)
	email.AddAttachment("super_cool_file.png")

	// Send email
	err = email.Send(MailService())

	if err != nil {
		log.Fatal(err)
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "Created New Shuttlers Policy",
		"policy":  shuttlerStruct,
		"status":  true})
}
