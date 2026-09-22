package ICMP
import ( "encoding/binary"
		 "fmt")
type ICMPPacket struct{
	Type uint8
	Code uint8
	Checksum uint16
	Identifier uint16
	Sequence uint16
	Payload []byte
}
func Marshal( p ICMPPacket) []byte{
	packet := make([]byte, 8+len(p.Payload))
	packet[0]=p.Type
	packet[1]=p.Code
	
	binary.BigEndian.PutUint16(packet[4:6],p.Identifier)
	binary.BigEndian.PutUint16(packet[6:8],p.Sequence)
	copy(packet[8:],p.Payload)
	sum:=Checksum(packet)
	binary.BigEndian.PutUint16(packet[2:4],sum)
	return packet
}
func Checksum(data[]byte) uint16{
	var sum uint32
	for i:=0; i+1<len(data); i+=2{
		word:=uint32(data[i]<<8 | data[i+1])
		sum+=word
	}
	if len(data)%2!=0 {
		sum+=uint32(data[len(data)-1]<<8)
	}
	if sum>>16!=0 {
		sum = (sum & 0xffff) + sum>>16
	}
	return ^uint16(sum)
}
func Parse(data[] byte) (ICMPPacket){
	packet:=ICMPPacket{
		Type: data[0],
		Code: data[1],
		Checksum: binary.BigEndian.Uint16(data[2:4]),
		Identifier: binary.BigEndian.Uint16(data[4:6]),
		Sequence: binary.BigEndian.Uint16(data[6:8]),
		Payload: data[8:],
		
	}
	return packet
}
func main(){
	packet:=ICMPPacket{
		Type:8,
		Code:0,
		Identifier:2807,
		Sequence:1,
		Payload:[]byte("hello"),
	}
	data:=Marshal(packet)
	
}
