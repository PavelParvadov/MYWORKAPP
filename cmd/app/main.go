package main

import (
	"MYWORKAPP/config"
	"MYWORKAPP/internal/home"
	"MYWORKAPP/pkg/logger"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	recoverMiddleware "github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Инициализация конфига
	config.Init()
	logCnf := config.NewLogConfig()
	CustomLogger := logger.NewLogger(logCnf)

	// Настройка логгера

	log.SetLevel(log.Level(logCnf.Level))

	app := fiber.New()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: CustomLogger,
	}))
	app.Use(recoverMiddleware.New())

	// Регистрация обработчиков
	home.NewHomeHandler(app, CustomLogger)

	// Запуск сервера
	app.Listen(":3000")

}
