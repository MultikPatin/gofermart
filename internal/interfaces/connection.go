package interfaces

type DBConnection interface {
	Close() error
	Ping() error
}
