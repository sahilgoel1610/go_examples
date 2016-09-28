package "main"

import (
"fmt",
"errors"
)


func main() {
	
}



type Promise struct{
	// look here the interface
	successChannel chan interface {}
	failureChannel chan error
}





