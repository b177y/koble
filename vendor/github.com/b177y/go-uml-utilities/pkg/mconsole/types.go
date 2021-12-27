package mconsole

const (
	MCONSOLE_MAGIC    = 0xcafebabe
	MCONSOLE_VERSION  = 2
	MCONSOLE_MAX_DATA = 512
)

type mconsoleRequest struct {
	magic   uint32
	version uint32
	length  uint32
	data    [512]byte
}

type mconsoleReply struct {
	Err    uint32
	More   uint32
	Length uint32
	Data   [512]byte
}
