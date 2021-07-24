package dto

// Input interface represent input of use cases
type Input interface {
	String() string
	Validate() error
}

// Output interface represent output of use cases
type Output interface {
	String() string
}
