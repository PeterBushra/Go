package main

func main() {
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
