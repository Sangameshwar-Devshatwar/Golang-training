package folder1

import (
	"testing"
)

// var u string = "Sangam@123"
// var p string =

func TestLogin(t *testing.T) {
	result := Login("Sangam@123", "123")
	expected := "logged in successfully :-welcome to GO"
	if expected != result {
		t.Errorf("your credentials provided are wrong")//erorf is used to show error where the test case got failed

	}
}
