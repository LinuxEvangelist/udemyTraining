package main

import "fmt"
// same package (go run main.og type.go)
func main()  {
	switch "Medhi" {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Medhi":
		fmt.Println("Wassup Medhi")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("Have you no friends?")
	}
	//fallthrough
	/*switch "Marcus" {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Medhi":
		fmt.Println("Wassup Medhi")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	case "Marcus":
		fmt.Println("Wassup Marcus")
		fallthrough
	case "Jenny1":
		fmt.Println("Wassup Jenny1")
		fallthrough
	case "Jennyy":
		fmt.Println("Wassup Jennyy")
	default:
		fmt.Println("Have you no friends?")
	}*/
	//Multiple Cases
	/*switch "Jenny" {
	case "Tim","Jenny":
		fmt.Println("Wassup Tim , or ,err, Jenny")
	case "Marcus","Medhi":
		fmt.Println("Both of your names start with M")
	case "Julian","Sushant":
		fmt.Println("Wassup Julian / Sushant")*/

	//no expression
	/*myFriendsName := "A!"
	switch {
		case len(myFriendsName) == 2:
			fmt.Println("Wassup my friend with name of length 2")
		case myFriendsName == "Tim":
			fmt.Println("Wassup Tim")
		case myFriendsName == "Jenny":
			fmt.Println("Wassup Jenny")
		case myFriendsName =="Marcus", myFriendsName=="Medhi":
			fmt.Println("Your name is either Marcus or Medhi")
		case myFriendsName == "Julian":
			fmt.Println("Wassup Julian")
		case myFriendsName =="Sushant":
			fmt.Println("Wassup Sushant")
		default:
			fmt.Println("nothing matched; this is the default")
	}*/


	SwitchOnType(7)


}