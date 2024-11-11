package provider

type IF interface {
	Get() <-chan string
}
