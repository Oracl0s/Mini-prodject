package main

import p "github.com/Oracl0s/Mini-prodject.git/cmd/pass"

func main() {
	p.Users()
	p.Definition()
	p.Ifelse()
	p.Swithes()
	p.Interface()
	p.SayHello(&p.Russian{}, "Russian")
	p.SayHello(&p.America{}, "America")

}
