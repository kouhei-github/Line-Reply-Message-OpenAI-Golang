package utils

//type error interface {
//	Error() string
//}

type MyError struct {
	Message string
}

func (myErr MyError) Error() string {
	msg := "[code]\n" + myErr.Message + "[/code]\n"
	return msg
}
