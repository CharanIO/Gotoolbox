//every code bit in go belongs to the package and package belogs to module
//step 1:Enable dependency tracking by go mod init modulename
//step 2:declare the package
//step 3:write the code

package main

import "fmt"

func main() {
	fmt.Printf("hello world")
}
