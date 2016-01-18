package slacker

// Publishable is an interface that allows types to declare they are
// "publishable". Meaning that you can send the RTM API the data returned by
// the method Publishable
type Publishable interface {
	Publishable() ([]byte, error)
}
