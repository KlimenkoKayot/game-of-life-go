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
	- pkg // (ОК) Для хранения пакетов
		- life // (ОК) Сама логика игры
			- life.go
        - support // (ОК) Вспомогательные функции
            - support.go
```

***

Хз надо или нет
```go
		- service // Сервис, который инициализирует и хранит состояния игры
			service.go
```