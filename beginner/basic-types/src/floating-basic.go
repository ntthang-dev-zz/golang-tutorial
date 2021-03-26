package floatingBasic

import "fmt"

func floatingBasic() {
	var f1 float32
	var f2 float64

	var maxFloat32 float32
	maxFloat32 = 16777216
	fmt.Println(maxFloat32 == maxFloat32+10) // you would typically expect this to return false
	// it returns true
	fmt.Println(maxFloat32 + 10)      // 16777216
	fmt.Println(maxFloat32 + 2000000) // 16777216

	// converting from int to float
	var myint int
	myfloat := float64(myint)

	// converting from float to int
	var myfloat2 float64
	myint2 := int(myfloat2)
}
