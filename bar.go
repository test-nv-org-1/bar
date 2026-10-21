package bar

import (
	"fmt"
)

const version = "v0.1.0"

func Bar() string {	
	return fmt.Sprintf("bar %s", version)
}
