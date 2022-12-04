package contract

type Handler interface {
	Execute() error
}
