package vacancy

import (
	"alaricode/go-fiber/pkg/tadapter"
	"alaricode/go-fiber/pkg/validator"
	"alaricode/go-fiber/views/components"

	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog"
)

type VacancyHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
	repository   *VacancyRepository
	store        *session.Store
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger, repository *VacancyRepository, store *session.Store) {
	h := &VacancyHandler{
		router:       router,
		customLogger: customLogger,
		repository:   repository,
		store:        store,
	}
	vacancyGroup := h.router.Group("/vacancy")
	vacancyGroup.Post("/", h.createVacancy)
}

func (h *VacancyHandler) createVacancy(c *fiber.Ctx) error {
	sess, err := h.store.Get(c)
	if err != nil || sess.Get("userID") == nil {
		component := components.Notification("Вы должны быть авторизованы для публикации вакансии", components.NotificationFail)
		return tadapter.Render(c, component)
	}

	form := VacancyCreateForm{
		Email:    c.FormValue("email"),
		Location: c.FormValue("location"),
		Type:     c.FormValue("type"),
		Company:  c.FormValue("company"),
		Role:     c.FormValue("role"),
		Salary:   c.FormValue("salary"),
	}
	errors := validate.Validate(
		&validators.EmailIsPresent{Name: "Email", Field: form.Email, Message: "Email не задан или неверный"},
		&validators.StringIsPresent{Name: "Location", Field: form.Location, Message: "Расположение не задано"},
		&validators.StringIsPresent{Name: "Type", Field: form.Type, Message: "Сфера компании не задана"},
		&validators.StringIsPresent{Name: "Company", Field: form.Company, Message: "Название компании не задано"},
		&validators.StringIsPresent{Name: "Role", Field: form.Role, Message: "Должность не задана"},
		&validators.StringIsPresent{Name: "Salary", Field: form.Salary, Message: "Зарплата не задана"},
	)

	if len(errors.Errors) > 0 {
		component := components.Notification(validator.FormatErrors(errors), components.NotificationFail)
		return tadapter.Render(c, component)
	}

	err = h.repository.AddVacancy(form)
	if err != nil {
		component := components.Notification(err.Error(), components.NotificationFail)
		return tadapter.Render(c, component)
	}

	component := components.Notification("Вакансия успешно создана", components.NotificationSuccess)
	return tadapter.Render(c, component)
}
