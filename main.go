package main

import "fmt"

//Animal is the type for our abstract factory
type Animal interface{
	Says()
	LikesMilk() bool
}

//Dog is the concrete factory for dogs
type Dog struct{}

//implement abstract factory for dogs

func(d *Dog) Says(){
	fmt.Println("Woof")
}

func(d *Dog) LikesMilk() bool{
	return false
}


//Cat is the concrete factory for cats
type Cat struct{}

//implement abstract factory for cats

func(c *Cat) Says(){
	fmt.Println("Meow")
}

func(c *Cat) LikesMilk() bool{
	return true
}


type AnimalFactory interface{
	New() Animal
}

type DogFactory struct{}

func (df *DogFactory) New() Animal{
	return &Dog{} 
}

type CatFactory struct{}

func (cf *CatFactory) New() Animal{
	return &Cat{}
}

func main(){
	//create one each of a dog and car facgory
	dogFactory := DogFactory{}
	catFactory := CatFactory{}

	//call new method to create dog and cat
	dog := dogFactory.New()
	cat := catFactory.New()

	dog.Says()
	fmt.Println("A dog likes milk ", dog.LikesMilk())
	cat.Says()
	fmt.Println("A cat likes milk ", cat.LikesMilk())
}