package main

import _case "go_project/interface/case"

func main() {
	//cat := _case.NewCat()
	//animalLife(cat)
	//dog := _case.NewDog()
	//animalLife(dog)
	dove := _case.NewDove()
	animalLife(dove)
}

func animalLife(a _case.AnimalI) {
	a.Each()
	a.Drink()
	a.Sleep()
	a.Run()
}
