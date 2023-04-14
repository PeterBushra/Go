package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func BaseDT() {
	// ****** Strings
	var interString string = "first inter sentence in \n GO"
	println(interString)
	rawString := `first raw sentence in \n GO`
	println(rawString)

	// ***** Numbers
	// Go doesn't support implicit case between diff types
	// int16 != int32 [error]

	// bool

	// error

	/// ********* Multiple Variables to be initialized at once
	a, b := 10, 5
	println(a)
	println(b)

	// ***** GO Support Constant
	const constVal = 3.14
	const constStrVal string = "Peter"
	const c = constVal
	const (
		pi      = 3.14
		name    = "peter"
		company = "Canonical"
	)
	println(company)

	// all const must be compiler you can't assign a function return to a const

	// sepcial keyword called iota takes the value as its related position in const exp gorup
	const (
		IOTAa = iota
		IOTAb = iota
	)
	println(IOTAa)
	println(IOTAb)

	const (
		x1 = 2 * 5
		x2 // 2*5
		x3 = iota
		x4 // iota +1
	)
	println(x1, x2, x3, x4)

	// ********* pointers
	aR := 52
	pB := &aR
	println(aR, pB, *pB)
	aR = 412
	println(aR, pB, *pB)
	*pB = 111
	println(aR, pB, *pB)
}

func main() {
	//BaseDT()
	//arr()
	//slice()
	//maps()
	//structs()
	//loops()
	//conditions()
	panicAndRecover()

}

func panicAndRecover() {
	fmt.Println("panicAndRecover 1")
	func1()
	fmt.Println("panicAndRecover 2")

}

func func1() {
	defer func() {
		fmt.Println(recover())
	}()
	fmt.Println("func1")
	panic("Ops!")
	fmt.Println("func2")
}

func conditions() {
	// if , switch [nothing new]

	/// ************** Defer
	fmt.Println("print 1")

	defer fmt.Println("defer print 1")

	fmt.Println("print 2")

	defer fmt.Println("defer print 2")

	/*USE CASE*/ // Last in First out
	/*
		db, _ := sql.Open("", "connstr")
		defer db.Close()

		rows, _ := db.Query("select ...")
		rows.Close()
	*/

}

func loops() {
	//infinite loop
	/*
		x := 1
		for {
			fmt.Println(x)
			x += 1
		}
	*/
	/*
		x := 1
		for x < 5 {
			fmt.Println(x)
			x += 1
		}

		for t := 124; t < 127; t++ {
			fmt.Println(t)
		}
	*/
	// Loop Thorugh Collection
	arr := [4]int{1, 2, 3, 4}
	for key, val := range arr {
		fmt.Println(key, val)
	}
}

func structs() {
	// structs are comparable and copied by value
	var s struct {
		id   int
		name string
	}
	fmt.Println(s)

	s.name = "Peter"
	s.id = 1
	fmt.Println(s)

	type myStruct struct {
		id   int
		name string
	}

	var s1 myStruct
	s1.id = 2
	s1.name = "Boshra"
	fmt.Println(s1)

	var s2 myStruct
	s2.id = 3
	s2.name = "Aziz"
	fmt.Println(s2)

	s4 := s2
	fmt.Println(s4)
	fmt.Println(s4 == s2)

}

func maps() {
	// maps copying as reference
	var m map[string]int
	fmt.Println(m)
	m = map[string]int{"peter": 1, "Boshra": 2}
	fmt.Println(m)
	fmt.Println(m["peter"])
	fmt.Println(m["xx"])

	delete(m, "peter")
	fmt.Println(m["peter"])
	fmt.Println(m)
	v, ok := m["peter"]
	fmt.Println(v, ok)

	m2 := m
	fmt.Println("m2: ", m2)
	m["XX"] = 123
	fmt.Println("m: ", m)
	fmt.Println("m2: ", m2)
	//m == m2 // not applicable
}

func slice() {
	// not fixed as array
	var s []int
	fmt.Println(s)
	s = []int{1, 2, 3}
	fmt.Println(s[0])
	s[1] = 132
	fmt.Println(s[1])

	s = append(s, 4, 5, 6, 7)
	fmt.Println(s)

	slices.Delete(s, 1, 4)
	fmt.Println(s)

}
func arr() {
	// Arrays in go are copying the value on equal not copy the address

	var array [3]int
	fmt.Println(array)
	array = [3]int{1, 2, 3}
	fmt.Println(array)
	array2 := array
	println(array2 == array)
	array[0] = 99
	fmt.Println(array)
	fmt.Println(array2)
	println(array2 == array)
}
