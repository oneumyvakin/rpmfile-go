# rpmfile-go
Working with rpm files with Golang

# Usage

- go get github.com/oneumyvakin/rpmfile-go

```
package main

import (
  "github.com/oneumyvakin/rpmfile-go"
  "fmt"
)

func main() {

	input_file := "C:\\Users\\Acha\\l10n.rpm"
	var rpm rpmfile.Rpm_file
	//rpm.set_debug()
	rpm.Open(input_file)

	fmt.Printf("Signature: %#v\n", rpm.Signature.Headersignatures)
	fmt.Printf("SHA1: %s\n", rpm.Signature.Sha1)

	fmt.Printf("Version: %s\n", rpm.Header.Version)
	fmt.Printf("Name: %s\n", rpm.Header.Name)
	fmt.Printf("Buildhost: %s\n", rpm.Header.Buildhost)
	fmt.Printf("OS: %s\n", rpm.Header.Os)
	fmt.Printf("Buildhost: %s\n", rpm.Header.Arch)

}
```
