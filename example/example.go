package main

import "fmt"

type DomainError struct {
	reason string
}

func (e *DomainError) Error() string {
	if len(e.reason) > 0 {
		return e.reason
	}
	return "DomainError!"
}

type scope int

const (
	ScopeInfo scope = iota
	ScopeWarn
	ScopeError
)

type ApplicationError struct {
	scope scope
}

func (e *ApplicationError) Error() string {
	return "ApplicationError!"
}

func handleError(err error) {
	switch e := err.(type) {
	case *DomainError:
		fmt.Println("Domain:", e)
	case *ApplicationError:
		fmt.Println("Application:", e, "Scope level:", e.scope)
	}

}

func main() {
	de := &DomainError{}
	handleError(de)

	de = &DomainError{
		reason: "not a valid num",
	}
	handleError(de)

	ae := &ApplicationError{
		scope: ScopeInfo,
	}
	handleError(ae)
}
