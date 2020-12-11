// Package tsk2 реализует решение для задания 2
package tsk2

import "fmt"

// Employee описывает сотрудника организации
type Employee struct {
	age int
}

// Customer описывает клиента организации
type Customer struct {
	age int
}

// MaxAge возвращает самого старшего пользователя
func MaxAge(users ...interface{}) (interface{}, error) {
	var oldest interface{}
	var maxAge int

	for _, u := range users {
		var age int

		switch t := u.(type) {
		case Employee:
			age = t.age
		case Customer:
			age = t.age
		default:
			return nil, fmt.Errorf("wrong age attribute of %T", t)
		}

		if age > maxAge {
			maxAge = age
			oldest = u
		}
	}

	return oldest, nil
}
