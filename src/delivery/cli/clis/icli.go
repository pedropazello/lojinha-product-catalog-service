package clis

type ICLI interface {
	GetStdInput() []string
	PutStdOutput() string
}
