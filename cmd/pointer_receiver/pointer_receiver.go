package main

type Getter interface {
	Get() string
}

type A struct {
}

func (a A) Get() string {
	return ""
}

func main() {
	var a *A
	a.Get()

	var _ Getter = a
}
