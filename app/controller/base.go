package controller

type Controller interface {
	Routes()
	Middlewares()
}
