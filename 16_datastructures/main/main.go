package main

import (
	"fmt"
)

func main()  {
	//Array, numbered sequence of elements of a single type, with numbers, one dimension,not dynamic
	arrayf()
	//Array2, initialized int values to 0
	arrayf2()
	//Arraybyte
	arraybyte()
	//Arraystring
	arraystring()
	//Slices, it's a reference type
	slice()
	//Slices with make and capacity
	slice2()
	//Slices. running with string
	slide3()
	//Append slice to slice
	appendSliceToSlice()
	//Multidimensional slice
	fmt.Println("********MultiDimensional Slice***********")
	multidimesionalSlice()
	//Slice of slice Int
	sliceOfSliceInt()
	//short slice declaration
	shortSliceDeclaration()
	//increment slice item
	incrementSliceItem()
	//Maps, ks values,dictionary, value uninitialized map is nil, on top hash table
	fmt.Println("********Maps***********")
	mapf()
	//map validation
	mapValidation()

}

func arrayf()  {
	var x[58]string
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
	//65 - 122
	for i:=65; i <= 122; i++ {
		x[i-65] = string(i)
	}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
	fmt.Print("Array ")
}
func arrayf2()  {
	var x [256]int
	fmt.Println(len(x))
	fmt.Println(x[42])
	for i:=0; i < 256; i++ {
		x[i] = i
	}
	for i,v :=range x{
		fmt.Printf("%v - %T - %b\n",v,v,v)
		if i > 50 {
			break
		}
	}
}
func arraybyte()  {
	var x [256]byte
	fmt.Println(len(x))
	fmt.Println(x[0])
	for i:=0; i < 256; i++ {
		x[i] = byte(i)
	}
	for i,v:=range x {
		fmt.Printf("%v - %T - %b\n",v,v,v)
		if i > 50{
			break
		}
	}
}

func arraystring()  {
	var x [256]string
	fmt.Println(len(x))
	fmt.Println(x[0])
	for i :=0; i<256; i++ {
		x[i]=string(i)
	}
	for _,v := range x{
		fmt.Printf("%v - %T - %v\n",v,v,[]byte(v))
	}
}

func slice()  {
	mySlice := []int{1,3,5,7,9,11,}
	fmt.Printf("%T\n",mySlice)
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4]) //slicing a slice
	fmt.Println(mySlice[2]) //index access; accessing by index
	fmt.Println("myString"[2])//index access; accessing by index
}
func slice2()  {
	mySlice := make([]int,0,5)
	fmt.Println("--------------------------------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("--------------------------------------")

	for i:=0; i<80;i++  {
		mySlice = append(mySlice,i) //with append de capacity grows
		fmt.Println("Len:",len(mySlice),"Capacity:",cap(mySlice),"Value:",mySlice[i])
	}
}
func slide3()  {
	greeting := []string{
		"Good morning!",
		"Bonjour",
		"dias!",
		"Bongirono!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}
	for i, currentEntry := range greeting {
		fmt.Println(i,currentEntry)
	}
	for j:=0; j < len(greeting);j++  {
		fmt.Println(greeting[j])
	}
}

func appendSliceToSlice()  {
	mySlide := []int{1,2,3,4,5}
	myOtherSlide := []int{6,7,8,9}
	mySlide = append(mySlide,myOtherSlide...)
	fmt.Println(mySlide)
}
func deleteFromSlice()  {
	mySlice:=[]string{"Monday","Tuesday"}
	myOtherSlice:=[]string{"Wednesday","Thursday","Friday"}
	mySlice = append(mySlice,myOtherSlice...)
	fmt.Println(mySlice)
	mySlice = append(mySlice[:2],mySlice[3:]...)
	fmt.Println(mySlice)
}
func multidimesionalSlice()  {
	record := make([][]string,0)
	fmt.Println(record)
	//student1
	student1 := make([]string,4)
	student1[0] ="Foster"
	student1[1] ="Nathan"
	student1[2] ="100.00"
	student1[3] ="74.00"
	//store the record
	record = append(record,student1)
	//student2
	student2 := make([]string,4)
	student2[0] ="Gomez"
	student2[1] ="Lisa"
	student2[2] ="92.00"
	student2[3] ="96.00"
	record = append(record,student2)
	fmt.Println(record)
}
func sliceOfSliceInt()  {
	transactions := make([][]int,0,3)
	for i:=0; i < 3;i++  {
		transaction := make([]int,0)
		for j:=0; j < 4; j++ {
			transaction = append(transaction,j)
		}
		transactions = append(transactions,transaction)
	}
	fmt.Println(transactions)
}
func shortSliceDeclaration()  {
	student := []string{}
	students := [][]string{}
	//var always initialize variables to default value in this case nil
	var stude []string
	//var studes [][]string
	fmt.Println(stude == nil)
	//student[0] =  "Todd" //error out of range
	student = append(student,"Todd")
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
	//Another way of declaration
	stu1 := make([]string,35)
	stu2 := make([][]string,35)
	fmt.Println(stu1)
	fmt.Println(stu2)
	fmt.Println(stu1 == nil)
}
func incrementSliceItem()  {
	mySlice := make([]int,1)
	fmt.Println(mySlice[0])
	mySlice[0] = 7
	fmt.Println(mySlice[0])
	mySlice[0]++
	fmt.Println(mySlice[0])
}
func mapf()  {
	m := make(map[string]int)
	//m := make(map[int]string)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:",m)
	//Delete an entry
	delete(m,"k2")
	fmt.Println("map_delete:",m)

	v,ok := m["k1"]
	fmt.Println("ok?:",ok,v)

	//Initialize a new map in the same line with this syntax
	n := map[string]int{"foo":1,"bar":2}
	fmt.Println("map:",n)
	//Another way, notice that initialized with var default value of reference type is nil
	var myGreeting = make(map[string]string)
	var myGreeting2 = map[string]string{}
	//var myGreeting map[string]string, never can get away with the entry nil
	myGreeting["Tim"] ="Good morning"
	myGreeting["Jenny"]="Bonjour"
	//Update an entry
	myGreeting["Jenny"]="Bonjourxx"
	myGreeting2["Jenny"]="Bonjourrr"
	fmt.Println(myGreeting)
	fmt.Println("Length of a Map:",len(myGreeting))
	fmt.Println(myGreeting2)
}
func mapValidation()  {
	greeting := map[int]string{
		0:"Good morning!",
		1:"Bonjour!",
		2:"Buenos Dias!",
		3:"Bongiorno!",
	}
	fmt.Println(greeting)
	if v,exists := greeting[2];exists {
		delete(greeting,2)
		fmt.Println("v: ",v)
		fmt.Println("exists: ",exists)
	}else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("v: ",v)
		fmt.Println("exists: ",exists)
	}
	//For
	for k,v:= range greeting{
		fmt.Println(k , " - ",v)
	}
	fmt.Println(greeting)

}