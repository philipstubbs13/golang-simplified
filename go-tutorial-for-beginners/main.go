package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func sayGreeting(n string) {
// 	fmt.Printf("Good morning %v \n", n)
// }
// func sayBye(n string) {
// 	fmt.Printf("Goodbe %v \n", n)
// }

// func cycleNames(n []string, f func(string)) {
// 	for _, v := range n {
// 		f(v)
// 	}
// }

// func circleArea(r float64) float64 {
// 	return math.Pi * r * r
// }

// func getInitials(n string) (string, string) {
// 	s := strings.ToUpper(n)
// 	names := strings.Split(s, " ")

// 	var initials []string
// 	for _, v := range names {
// 		initials = append(initials, v[:1])
// 	}

// 	if len(initials) > 1 {
// 		return initials[0], initials[1]
// 	}

// 	return initials[0], ""
// }

// var score = 99.5

// func updateName(x *string) {
// 	*x = "wedge"
// }

// func updateMenu(y map[string]float64) {
// 	y["coffee"] = 2.99
// }

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
	
	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if  err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)

		fmt.Println("item added - ", name, price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if  err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)

		fmt.Println("tip added - ", tip)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("you saved the file - ", b.name)
	default:
		fmt.Println("that was not a valid option...")
		promptOptions(b)
	}

}

func main() {
	mybill := createBill()
	promptOptions(mybill)

	// name := "tifa"

	// updateName(name)

	// fmt.Println("memory address of name is: ", &name);

	// m := &name
	// fmt.Println("memory address:", m)
	// fmt.Println("value at memory address:", *m)

	// fmt.Println(name)
	// updateName(m)
	// fmt.Println(name);
	
	// menu := map[string]float64 {
	// 	"pie": 5.95,
	// 	"ice cream": 3.99,
	// }

	// updateMenu(menu)
	// fmt.Println(menu)

	// menu := map[string]float64{
	// 	"soup": 4.99,
	// 	"pie": 7.99,
	// 	"salad": 6.99,
	// 	"toffee pudding": 3.55,
	// }

	// fmt.Println(menu)
	// fmt.Println(menu["pie"])

	// Looping maps
	// for k, v := range menu {
	// 	fmt.Println(k, "-", v)
	// }

	// ints as key type
	// phonebook := map[int]string{
	// 	267584967: "mario",
	// 	985494849: "luigi",
	// 	583948949: "peach",
	// }

	// fmt.Println(phonebook)
	// fmt.Println(phonebook[267584967])

	// phonebook[985494849] = "bowser"
	// fmt.Println(phonebook)

	// phonebook[583948949] = "yoshi"
	// fmt.Println(phonebook)


	// sayHello("mario")

	// for _, v := range points {
	// 	fmt.Println(v)
	// }

	// showScore()

	// fn1, sn1 := getInitials("tifa lockhart")
	// fmt.Println(fn1, sn1);

	// fn2, sn2 := getInitials("cloud strife")
	// fmt.Println(fn2, sn2);

	// fn3, sn3 := getInitials("barret")
	// fmt.Println(fn3, sn3);

	// sayGreeting("mario")
	// sayGreeting("luigi")
	// sayBye("mario")

	// cycleNames([]string{"cloud", "tifa", "barret"}, sayGreeting)

	// a1 := circleArea(10.5)
	// a2 := circleArea(15)

	// fmt.Println(a1, a2)
	// fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f", a1, a2)

	// age := 45

	// fmt.Println(age <= 50)
	// fmt.Println(age >= 50)
	// fmt.Println(age == 45)
	// fmt.Println(age != 50)

	// if age < 30 {
	// 	fmt.Println("age is less than 30")
	// } else if age < 40 {
	// 	fmt.Println("age is less than 40")
	// } else {
	// 	fmt.Println("age is not less than 45")
	// }

	// names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}

	// for index, value := range names {
	// 	if index == 1 {
	// 		fmt.Println("continuing at pos\n", index)
	// 		continue
	// 	}
	// 	if index > 2 {
	// 		fmt.Println("breaking at pos", index)
	// 		break
	// 	}

	// 	fmt.Printf("the value at pos %v is %v\n", index, value)
	// }
	
	// x := 0
	// for x < 5{
	// 	fmt.Println("value of x is:", x)
	// 	x++
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("value of i is:", i)
	// }

	// names := []string{"mario", "luigi", "yoshi", "peach"}
	
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// for index, value := range names {
	// 	fmt.Printf("the value at index %v is %v\n", index, value)
	// }

	// for _, value := range names {
	// 	fmt.Printf("the value is %v\n", value)
	// 	value = "new string"
	// }

	// fmt.Println(names)

	// greeting := "hello there friends!"

	// fmt.Println(strings.Contains(greeting, "hello!"))
	// fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "th"))
	// fmt.Println(strings.Split(greeting, " "))

	// the original value is unchanged
	// fmt.Println("original string value =", greeting)

	// ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	// sort.Ints(ages);
	// fmt.Println(ages);

	// index := sort.SearchInts(ages, 30);
	// fmt.Println(index)

	// names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	// sort.Strings(names)
	// fmt.Println(names)

	// fmt.Println(sort.SearchStrings(names, "bowser"))


	// var ages [3]int = [3]int{20, 25, 30}
	// var ages = [3]int{20, 25, 30}
	
	// names := [4]string{"yoshi", "mario", "peach", "bowser"}
	// names[1] = "luigi"

	// fmt.Println(ages, len(ages))
	// fmt.Println(names, len(names))

	// slices (use arrays under the hood)
	// var scores = []int{100, 50, 60}
	// scores[2] = 25
	// scores = append(scores, 85)

	// fmt.Println(scores, len(scores))

	// slice ranges
	// rangeOne := names[1:3]
	// rangeTwo := names[2:]
	// rangeThree := names[:3]

	// fmt.Println(rangeOne, rangeTwo, rangeThree)

	// rangeOne = append(rangeOne, "koopa")
	// fmt.Println(rangeOne)

	// age := 35
	// name := "shaun"

	// Print
	// fmt.Print("hello, ")
	// fmt.Print("world! \n")
	// fmt.Print("new line! \n")

	// Println
	// fmt.Println("hello ninjas!")
	// fmt.Println("goodbye ninjas!")
	// fmt.Println("my age is", age, "and my name is", name)

	// Printf (formatted strings) %_ = format specifier
	// fmt.Printf("my age is %v and my name is %v \n", age, name)
	// fmt.Printf("my age is %q and my name is %q \n", age, name)
	// fmt.Printf("age is of type %T\n", age)
	// fmt.Printf("you scored %0.1f points! \n", 225.55)

	// Sprintf (save formatted strings)
	// var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	// fmt.Println("the saved string is:", str)
	
	// strings
	// var nameOne string = "mario"
	// var nameTwo = "luigi"
	// var nameThree string

	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameOne = "peach"
	// nameThree = "bowser"

	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameFour := "yoshi"
	// fmt.Println(nameFour)

	// ints
	// var ageOne int = 20
	// var ageTwo = 30
	// ageThree := 40

	// fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	// var numOne int8 = 25
	// var numTwo int8 = -128
	// var numThree uint16 = 255

	// var scoreOne float32 = 25.98
	// var scoreTwo float64 = 399488457839384754839585.7
	// scoreThree := 1.5
}