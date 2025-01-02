```go
- game
	- cmd // Для хранения пакетов main
		-life
			- main.go // Точка входа в программу
	- http
		- server // HTTP-сервер
			server.go // Код сервера
			- handler  // Регистрация функций обработчиков
				handler.go
	- internal
		- application  // Конфигурация и код вызова приложения
			application.go
		- service // Сервис, который инициализирует и хранит состояния игры
			service.go
	- pkg // Для хранения пакетов
		- life // Сама логика игры
			- life.go
```