package person_handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"rinha-de-backend/src/dtos/person_dto"
	"rinha-de-backend/src/interfaces"
)

type PersonHandler struct {
	service interfaces.PersonService
}

func New(service interfaces.PersonService) *PersonHandler {
	return &PersonHandler{
		service: service,
	}
}

func (ph PersonHandler) Routes(app *fiber.App) {
	routerGeneral := app.Group("/contagem-pessoas")
	routerGeneral.Get("/", ph.count)

	router := app.Group("/pessoas")

	router.Get("/", ph.listPerson)
	router.Get("/:id", ph.GetPerson)
	router.Post("/", ph.CreatePerson)
}

func (ph PersonHandler) count(c *fiber.Ctx) error {
	ctx := c.UserContext()

	p, err := ph.service.List(ctx, person_dto.ListRequestParams{})
	if err != nil {
		return c.SendStatus(http.StatusUnprocessableEntity)
	}

	return c.SendString(fmt.Sprintf("%d", len(p)))
}

func (ph PersonHandler) listPerson(c *fiber.Ctx) error {
	ctx := c.UserContext()
	params := c.Query("t", "")

	if params == "" {
		return c.SendStatus(http.StatusBadRequest)
	}

	p, err := ph.service.List(ctx, person_dto.ListRequestParams{
		Params: params,
		Size:   50,
	})
	if err != nil {
		return c.SendStatus(http.StatusUnprocessableEntity)
	}

	return c.JSON(p)
}

func (ph PersonHandler) GetPerson(c *fiber.Ctx) error {
	ctx := c.UserContext()
	ID := c.Params("id")

	personID, err := uuid.Parse(ID)
	if err != nil {
		return c.SendStatus(http.StatusNotFound)
	}

	p, err := ph.service.GetByID(ctx, personID)
	if err != nil {
		return c.SendStatus(http.StatusNotFound)
	}

	return c.JSON(p)
}

func (ph PersonHandler) CreatePerson(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var payload person_dto.CreatePayload
	err := c.BodyParser(&payload)
	if err != nil {
		return c.SendStatus(http.StatusUnprocessableEntity)
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterValidation("birth", ValidateMyVal)

	err = validate.Struct(payload)
	if err != nil {
		return c.SendStatus(http.StatusUnprocessableEntity)
	}

	p, err := ph.service.Create(ctx, payload)
	if err != nil {
		return c.SendStatus(http.StatusUnprocessableEntity)
	}

	c.Response().Header.Add("Location", fmt.Sprintf("/pessoas/%s", p.ID.String()))
	c.Response().Header.Add("Content-Type", "application/json")

	return c.SendStatus(http.StatusCreated)
}

func ValidateMyVal(fl validator.FieldLevel) bool {
	_, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}

	return true
}
