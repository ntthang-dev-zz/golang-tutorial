package booleansBasic

import "fmt"

func booleansBasic() {
	var isTrue bool = true
	var isFalse bool = false
	// AND
	if isTrue && isFalse {
		fmt.Println("Both Conditions need to be True")
	}
	// OR - not exclusive
	if isTrue || isFalse {
		fmt.Println("Only one condition needs to be True")
	}
}
