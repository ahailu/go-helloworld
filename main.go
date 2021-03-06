package main

import "fmt"
// import "net/http"

func main() {
    // Hello world
    fmt.Printf("hello, world\n")

	//// Variables
	// var x string = "explicit"
	// y := "inferred"
	// fmt.Printf("explicit x: %s\n", x)
	// fmt.Printf("inferred y: %s\n", y)
	// var (
	//   a = 5
	//   b = 10
	//   c = "hi"
	// )
	// fmt.Printf("inferred a: %d, b: %d, c: %s\n", a, b, c)
	// fmt.Printf("formatting a: %s, b: %s, c: %s\n", a, b, c)

	//// Control structures
	// FOR: 
	// for i := 1; i <= 10; i++ {
	// 	if i % 2 == 0 {
	//     	fmt.Println(i, "even")
	// 	} else {
	// 		fmt.Println(i, "odd")
	// 	}
	// 	// switch i {
	// 	// 	case 0: fmt.Println("Zero")
	// 	// 	case 1: fmt.Println("One")
	// 	// 	case 2: fmt.Println("Two")
	// 	// 	case 3: fmt.Println("Three")
	// 	// 	case 4: fmt.Println("Four")
	// 	// 	case 5: fmt.Println("Five")
	// 	// 	default: fmt.Println("Unknown Number")
	// 	// }
	// }
	//
	// Infinite looP
	// for {
 	//	fmt.Println("infinite")
	// }

	//// Arrays
	// var x [5]int
	// x[4] = 100
 	// fmt.Println(x)

 	//// Slices: A slice is a segment of an array, are indexible and have a length. length is allowed to change
 	// var x []float64


 	//// Maps
 	// mp := make(map[string]int)
	// mp["key"] = 10
	// fmt.Println("added key ", mp)
	// delete(mp, "key")
	// fmt.Println("removed key", mp)
	// fmt.Println("non-existent key fail: ", mp["key"])
	// if value, ok := mp["Un"]; ok {
 	//	fmt.Println("validated key: ", value, ok)
	// }
	
	//Functions
	// m, n := f(7)
	// fmt.Println("multiple val returned", m, n)
	// fmt.Println("variadic 3 args", add(1,2,3))
	// fmt.Println("variadic 4 args", add(1,2,3,4)) // func Println(a ...interface{}) (n int, err error)


	//// Structs + Object composition
	// fmt.Println(person{firstName: "Alice", age: 30})
	// fmt.Println(student{identity: person{firstName: "Dora", lastName: "The Explorer", age: 8}, grade: 3})
	// s := student{identity: person{firstName: "Curious", lastName: "George", age: 7}, grade: 3}
	// fmt.Println(s.identity.name())


    //// Web server: https://golang.org/doc/articles/wiki/
 	// 	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
 	// 	   	fmt.Fprintf(w, "service path %s!", r.URL.Path[0:])
	// })
 	// 	http.ListenAndServe(":8080", nil)
	 
}

// func f(m int) (int, int) {
//   return m*5, m*6
// }

// func add(args ...int) int {
//   total := 0
//   for _, v := range args {
//     total += v
//   }
//   return total
// }


//// Structs + Object composition
// type person struct {
//     firstName string
//     lastName string
//     age  int
// }

// type student struct {
//     identity person 
//     grade int
// }

// func (p *person) name() string {
// 	return p.firstName + " " + p.lastName
// }

