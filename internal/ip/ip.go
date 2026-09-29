package ip
import "fmt"

type IPv4Header struct {
	Version    uint8
	HeaderLen  uint8
	Protocol   uint8
}

func ParseIPv4 ( data []byte) (IPv4Header, error){
	if len(data) < 20 {
		return IPv4Header{}, fmt.Errorf("packet too short")
	}

	version := data[0] >> 4
	ihl := data[0] & 0x0F

	headerLen := ihl * 4

	if len(data) < int(headerLen) {
		return IPv4Header{}, fmt.Errorf("invalid IPv4 header length")
	}

	return IPv4Header{
		Version:   version,
		HeaderLen: headerLen,
		Protocol:  data[9],
	}, nil
}