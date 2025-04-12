package interfaces

type Database interface {
	Close() error
	Ping() error
}
