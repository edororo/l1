package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h Human) Introduce() {
	fmt.Printf("Меня зовут %s, мне %v лет.\n", h.name, h.age)
}

func (h Human) SayGoodbye() {
	fmt.Printf("Было приятно представиться, прощайте!\n")
}

type Action struct {
	Human
	work string
}

func (a Action) DoAction() {
	defer a.SayGoodbye()
	a.Introduce()
	fmt.Printf("Моя профессия %s.\n", a.work)
}

func main() {
	person := Action{Human{"Вася", 27}, "учитель"}
	person.DoAction()
}
