package main

func main (){  // STRING, NUMBER , BOOLEAN
	// STRING 

	// o'zg - nomi - tipi - VALUE - qiymat
	var message string = "Hello Student"
	println(message)

	var messagel string         // DECLARE --> tanitdim
	messagel = "Hello World"    // ASSIGNMEND --> tenglashtirdim
	println(messagel)

	var box1, box2 string = "first", "second"
	println(box1,box2)

	//NOTE !

	var box = "This is variable"  // TIPI berilmagan ammo bu ham ishlaydi !
	println(box)                  // tipi berilmasa o'zi avtomat topib sekin ishlaydi
	                         //  agar tipi berilsa kod tezlik jihatidan yaxshi ishlaydi
//______________________________________________________________________________________

// NUMBER
// INT
	var a, b, c string    // BUTUN SON  Ex: 7, -5, 11, 23, -94, 768, 811
	println(a,b,c)     // --> 0 0 0 

	var n1, n2, n3 int = 11, 22, 33
	println(n1,n2,n3)

	var num1, num2, num3 int = 11, 22, 33
	println(num1 + num2 + num3 )
//____________________________________________________________________________________

}