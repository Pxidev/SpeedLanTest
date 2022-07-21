package vars

var (
	Version = "0.0.1"
	Port    = ""
)

func SetPort(portFunc string) {
	Port = portFunc
}
