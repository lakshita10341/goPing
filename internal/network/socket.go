package network
import "golang.org/x/sys/unix"
type RawSocket struct {
	fd int
}
func newRawSocket() (*RawSocket, error){
	num, err:= unix.Socket(
		unix.AF_INET,
		unix.SOCK_RAW,
		unix.IPPROTO_ICMP,
	)

	if err!=nil {
		return nil, err
	}
	return &RawSocket{
		fd: num
	}, nil
}
func (s *RawSocket) Close() error{
	return unix.Close(s.fd)
}