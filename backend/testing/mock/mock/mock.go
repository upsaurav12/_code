package mock

type Greeter interface {
	Greet(name string) string
}

type RealGreeter struct{}

func (r RealGreeter) Greet(name string) string {
	return "Hello, " + name
}

func SayHello(g Greeter, name string) string {
	return g.Greet(name) // Can use any Greeter, not just RealGreeter
}
