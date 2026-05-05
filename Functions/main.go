//package main

//func add(x int, y int) int {
//	return x + y
//}

// The above add function can be rewritten as the below

//func add(x, y int) int {
//	return x + y
//}
//
//func main() {
//	result := add(2, 4)
//	fmt.Printf("Result : %v\n", result)
//}

//Assignment
//We often will need to manipulate strings in our messaging app.
//For example, adding some personalization by using a customer's name within a template.
//The concat function should take two strings and smash them together.
//hello + world = helloworld

//package main
//
//import "fmt"
//
//func concat(s1 string, s2 string) string {
//	return s1 + s2
//}
//
//// don't touch below this line
//
//func main() {
//	test("Lane,", " happy birthday!")
//	test("Zuck,", " hope that Metaverse thing works out")
//	test("Go", " is fantastic")
//}
//
//func test(s1 string, s2 string) {
//	fmt.Println(concat(s1, s2))
//}

//package main
//
//import "fmt"
//
//// Pass by value - The functions in Go usually receive the copy of that variable so that the caller's original data is not mutated.
//
//func main() {
//	x := 5
//	increment(x)
//	fmt.Println(x) // This will still print 5
//}
//
//func increment(x int) {
//	x++
//}

// Assignment
// monthlyBillIncrease: Should return the increase in the bill from the previous to the current month. If the bill decreased, return a negative number.
// getBillForMonth: Should return the total cost for the number of messages sent.
// Fix the bugs in the monthlyBillIncrease and getBillForMonth functions.
// Looks like whoever wrote the functions didn't know the getBillForMonth
// function's bill parameter would be passed by value.
// It's not actually updating the lastMonthBill and thisMonthBill variables as intended so monthlyBillIncrease isn't returning the right result.
//
// Change getBillForMonth so it only take 2 parameters, get rid of bill.
// Instead, simply return the total cost of the messages.
// monthlyBillIncrease should use the result of calling getBillForMonth to calculate the increase between months.
package main

func monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) int {
	var lastMonthBill int = getBillForMonth(costPerSend, numLastMonth)
	var thisMonthBill int = getBillForMonth(costPerSend, numThisMonth)
	return thisMonthBill - lastMonthBill
}

func getBillForMonth(costPerSend, messagesSent int) int {
	bill := costPerSend * messagesSent
	return bill
}
