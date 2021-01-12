package main

import "fmt"

type contactInfo struct{
	email string
	zipCode int
}

type person struct{
	firstName string
	lastName string
	//contact contactInfo
	contactInfo

}
func main(){

	/**alex := person{ "Alex" , "Anderson"}
	fmt.Println(alex)

	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Printf("%+v",alex)**/


	jim := person{
		firstName: "Jim",
		lastName: "Party",
		//contact
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94000,

		},
	}

	/*jimPointer := &jim
	jimPointer.updateName("jimmy")*/
	jim.updateName("jimmy")
	jim.print()


}

func (pointerToPerson *person) updateName(newFirstName string){
	(*pointerToPerson).firstName = newFirstName
}

//Receiver func.
func (p person) print(){
	fmt.Printf("%+v",p)

}
