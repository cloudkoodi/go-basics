package aulas

// var taskDone bool - bigger names is OK in package scope

func NameConvention() {
	var mv int // max value
	_ = mv

	// var max_value int  - NOT OK

	// var packetReceived int - NOT OK, too long

	// var b int - OK

	// const MAX_VALUE = 100 - NOT OK

	// const N = 100 Better

	// var Max = 100 - can be exported

	// maxValue := 1000 recommended
	// max_value := 1000 not recommended

	// writeToDB := true - ok, idiomatic
	//writeToDb := false - not OK
}
