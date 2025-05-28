package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"

	"alaricode/go-fiber/pkg/tadapter"
	"alaricode/go-fiber/views"
)

type UserHandler struct {
	repo  *UserRepository
	store *session.Store
}

func NewUserHandler(router fiber.Router, repo *UserRepository, store *session.Store) *UserHandler {
	h := &UserHandler{repo: repo, store: store}
	router.Post("/register", h.Register)
	router.Post("/login", h.Login)
	router.Get("/logout", h.Logout)
	return h
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "" || password == "" {
		c.Status(fiber.StatusBadRequest)
		return tadapter.Render(c, views.RegisterError("Email и пароль обязательны"))
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return tadapter.Render(c, views.RegisterError("Ошибка при шифровании пароля"))
	}

	err = h.repo.CreateUser(name, email, string(hash))
	if err != nil {
		c.Status(fiber.StatusConflict)
		return tadapter.Render(c, views.RegisterError("Пользователь с таким email уже существует"))
	}

	c.Set("HX-Redirect", "/login")
	return c.SendStatus(fiber.StatusNoContent)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	user, err := h.repo.FindByEmail(email)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) != nil {
		c.Status(fiber.StatusUnauthorized)
		return tadapter.Render(c, views.LoginError("Неверный email или пароль"))
	}

	sess, _ := h.store.Get(c)
	sess.Set("userID", user.ID)
	sess.Set("name", user.Name)
	sess.Save()

	c.Set("HX-Redirect", "/")
	return c.SendStatus(fiber.StatusNoContent)
}

func (h *UserHandler) Logout(c *fiber.Ctx) error {
	sess, _ := h.store.Get(c)
	sess.Destroy()
	return c.Redirect("/")
}
