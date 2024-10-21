package main

import "fmt"

func main(){
	fmt.Println("Learn Array ");

	var friendsNames [4]string

	friendsNames[0] = "Apple";
	friendsNames[1] = "Apple";
	friendsNames[2] = "Apple";
	friendsNames[3] = "Apple";

	fmt.Println(friendsNames)
	fmt.Println(len(friendsNames))

	var vegList = [3]string{"potato","tomato","lotamo"}

	fmt.Println(vegList)
}