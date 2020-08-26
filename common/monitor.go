package common

type Monitor interface {
	GetApiName() string
	Update()
}
