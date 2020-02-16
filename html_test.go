package html

import (
	"fmt"
	"testing"
)

func TestBase_Get(t *testing.T) {
	fmt.Println(Body("23423424", M{
		"font-size": "16px",
	}))
}
