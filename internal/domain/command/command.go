package command

type Type string

const (
	SET     Type = "SET"
	GET     Type = "GET"
	DEL     Type = "DEL"
	EXPIRE  Type = "EXPIRE"
	TTL     Type = "TTL"
	PERSIST Type = "PERSIST"
	QUIT    Type = "QUIT"

	KEYS   Type = "KEYS"
	EXISTS Type = "EXISTS"
	SIZE   Type = "SIZE"
	PING   Type = "PING"
	INFO   Type = "INFO"
)

func (t Type) String() string {
	return string(t)
}

func (t Type) IsValid() bool {
	switch t {
	case SET, GET, DEL, EXPIRE, TTL, PERSIST, QUIT, KEYS, EXISTS, SIZE, PING, INFO:
		return true
	default:
		return false
	}
}

func (t Type) IsWriteCommand(line string) bool {
	switch t {
	case SET, DEL, EXPIRE, PERSIST:
		return true
	default:
		return false
	}
}
