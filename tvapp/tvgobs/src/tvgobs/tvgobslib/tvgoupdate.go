package tvgobslib

import "fmt"

var finished bool = false
// UpdateMain needs to be exported
func UpdateMain() (finished bool) {
	SetUp()
	finished = true
	fmt.Println("Update Complete")
	return
}
