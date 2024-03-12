package printgovers

import "runtime"

// my function to retrieve the version of golang in use
func PrintVers() string {
	//  using the stdlib runtime() to return the version of go
	return runtime.Version()
}
