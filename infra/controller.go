package infra

type Controller interface {
	Routes()
	Middlewares()
}
