package html

import (
	"fmt"
	"testing"
)

func TestBase_Get(t *testing.T) {
	fmt.Println(Div(
		A("asfasdfa",
			M{"color": "#676565"},
			M{"class": "dropdown-toggle", "href": "#", "data-toggle": "dropdown"},
		)+Ul("23234234",
			M{"min-width": "20px !important", "left": "-32px", "overflow": "hidden"},
			M{"class": "dropdown-menu", "role": "menu", "aria-labelledby": "dLabel"}),
		M{"text-align": "center"}, M{"class": "dropdown"}))
}
