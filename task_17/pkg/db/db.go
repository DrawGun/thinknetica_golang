package db

// User учетные данные пользователя
type User struct {
	Login    string
	Password string
}

// UserPasswords аналог таблицы пользователей
var UserPasswords = map[string]string{
	"admin": "$2a$10$w.fe2uqr0UOaLvo83r.jOe0tRZn8AVQFHKQSpXCw8k7XIs8kaJzte",
	"guest": "$2a$10$ouLmtFssm.KrfL0yDwWxi.cowOLfOCZ2QS6K/6CIg/2q26T4bqlwC",
}

// AccessRights аналог таблицы с правами пользоавтелей
var AccessRights = map[string][]string{
	"admin": []string{"all"},
	"guest": []string{"read"},
}
