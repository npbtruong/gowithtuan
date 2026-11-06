package dog

//Khi export ra file khác viết hoa chử cái đầu 
type Dog struct{
	Name string
	Owner string
}

func (d Dog) Sound() string {
	return "BARK BARK"
}

func (d Dog) GetName() string{
	return d.Name
}