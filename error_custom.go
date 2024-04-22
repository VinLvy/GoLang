package main

import "fmt"

type validationerror struct {
	Message string
}

func (v *validationerror) Error() string {
	return v.Message
}

type notfounderror struct {
	Message string
}

func (n *notfounderror) Error() string {
	return n.Message
}

func savedata(id string, data any) error {
	if id == "" {
		return &validationerror{"validation error"}
	}
	if id != "daoa" {
		return &notfounderror{"data not found"}
	}
	// ok
	return nil
}

func ercu() {
	err := savedata("d", nil)
	if err != nil {
		// terjadi error
		//if validationerr, ok := err.(*validationerror); ok {
		//	fmt.Println("validation error:", validationerr.Error())
		//} else if notfounderr, ok := err.(*notfounderror); ok {
		//	fmt.Println("not found error:", notfounderr.Error())
		//} else {
		//	fmt.Println("unknown error:", err.Error())
		//}

		switch finalerror := err.(type) {
		case *validationerror:
			fmt.Println("validation error:", finalerror.Error())
		case *notfounderror:
			fmt.Println("not found error:", finalerror.Error())
		default:
			fmt.Println("unknown error:", finalerror.Error())
		}
	} else {
		// sukses
		fmt.Println("sukses")
	}
}
