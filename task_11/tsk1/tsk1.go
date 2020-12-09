// Package tsk1 реализует решение для задания 1
package tsk1

// Employee описывает сотрудника организации
type Employee struct {
	age int
}

// Customer описывает клиента организации
type Customer struct {
	age int
}

// Age возвращает возраст сотрудника
func (e Employee) Age() int {
	return e.age
}

// Age возвращает возраст клиента
func (c Customer) Age() int {
	return c.age
}

type user interface {
	Age() int
}

// MaxAge возвращает возраст самого старшего пользователя
func MaxAge(users ...user) int {
	var age int

	for _, u := range users {
		if age < u.Age() {
			age = u.Age()
		}
	}
	return age
}
