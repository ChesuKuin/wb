package main

import "fmt"

// Human — родительская структура с произвольным набором полей и методов.
type Human struct {
	Name string
	Age  int
}

func (h Human) SayHello() string {
	return fmt.Sprintf("Привет, меня зовут %s", h.Name)
}

func (h *Human) Birthday() {
	h.Age++
	fmt.Printf("%s теперь %d лет\n", h.Name, h.Age)
}

// Action встраивает Human (embedded struct) — это аналог наследования
// через композицию: Action автоматически получает все поля и методы Human.
type Action struct {
	Human
	ActionType string
}

func (a Action) Describe() string {
	return fmt.Sprintf("%s выполняет действие: %s", a.Name, a.ActionType)
}

// SayHello переопределяет метод Human — при этом можно явно вызвать
// "родительскую" версию через a.Human.SayHello().
func (a Action) SayHello() string {
	return fmt.Sprintf("Action: %s (%s)", a.Human.SayHello(), a.ActionType)
}

func main() {
	a := Action{
		Human:      Human{Name: "Рири", Age: 20},
		ActionType: "бежит",
	}

	// методы Human доступны напрямую на Action благодаря embedding
	fmt.Println(a.SayHello())
	fmt.Println(a.Describe())
	a.Birthday()

	// поля Human тоже доступны напрямую
	fmt.Println(a.Name, a.Age)
}
