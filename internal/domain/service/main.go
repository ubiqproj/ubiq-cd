package service

type Service interface {
	Path() string
	Run() error
	Stop() error
}
