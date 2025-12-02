package p

import "fmt"

type walker interface {
	walk()
}
type runner interface {
	run()
}
type animal interface {
	makeSound()
}
type greeter interface {
	greet(string) string
}
type Russian struct{}
type America struct{}
type cat struct{}
type dog struct{}

func (b cat) makeSound() {
	fmt.Println("meow !")
}
func (d dog) makeSound() {
	fmt.Println("woof !")
}
func Interface() {
	var b, d animal = &cat{}, &dog{}
	b.makeSound()
	d.makeSound()
}
func (r *Russian) greet(name string) string {
	return fmt.Sprintf("Привет , %s", name)
}
func (a *America) greet(name string) string {
	return fmt.Sprintf("Hello , %s ", name)
}
func SayHello(g greeter, name string) {
	fmt.Println(g.greet(name))
}
func (t *cat) walk() {
	fmt.Println("cat is walking !")
}
func (b *cat) run() {
	fmt.Println("cat is running !")
}
