package service


type Animal interface {
	Sound() string
	GetName() string
	//chỉ chứa method ko chứa thuộc tính
}

type AnimalExtension interface{
	Animal
	EatSound() string
}

type AnimalAction interface{
	Bite() string
}