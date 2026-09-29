package ping
import (
		"fmt"
		"time"
		"net"
		"goPing/internal/icmp"
		"goPing/internal/ip"
)
func Ping(address string) error{
	conn, err :=net.Dial("ip4:icmp",address)
	if err!=nil{
		return err
	}
	
	defer conn.Close()
	packet:=icmp.ICMPPacket{
		Type:       8,
        Code:       0,
        Identifier: 2807,
        Sequence:   1,
        Payload:    []byte("hello"),
	}
	data:=icmp.Marshal(packet)
	start:=time.Now()
	_,err=conn.Write(data)
	if err!=nil{
		return err
	}
	buffer:=make([]byte,1500);
	conn.SetReadDeadline(time.Now().Add(2*time.Second))
	n,err:=conn.Read(buffer)
	if err!=nil{
		return err
	}
	rtt:=time.Since(start)
	ipHeader, err := ip.ParseIPv4(buffer[:n])
	if err != nil {
		return err
	}

	fmt.Printf("IP version: %d\n", ipHeader.Version)
	fmt.Printf("IP header length: %d bytes\n", ipHeader.HeaderLen)
	fmt.Printf("Protocol: %d\n", ipHeader.Protocol)
	icmpData := buffer[ipHeader.HeaderLen:n]
	if icmp.Checksum(icmpData) != 0 {
		return fmt.Errorf("invalid ICMP checksum")
	}
	reply, err := icmp.Parse(icmpData)
	if err != nil {
		return err
	}

	fmt.Printf("Type       : %d\n", reply.Type)
	fmt.Printf("Code       : %d\n", reply.Code)
	fmt.Printf("Checksum   : %04x\n", reply.Checksum)
	fmt.Printf("Identifier : %d\n", reply.Identifier)
	fmt.Printf("Sequence   : %d\n", reply.Sequence)
	fmt.Printf("RTT        : %v\n", rtt)
	if reply.Type!=0 {
		return fmt.Errorf("not an Echo Reply")
	}
	if reply.Identifier!=packet.Identifier{
		return fmt.Errorf("identifier mismatch")
	}
	if reply.Sequence!=packet.Sequence {
		return fmt.Errorf("sequence mismatch")
	}

	fmt.Printf("received %d bytes\n", n)
	fmt.Printf("% x\n", buffer[:n])
	if icmp.Checksum(icmpData)!=0 {
		return fmt.Errorf("invalid ICMP checksum")
	}
	fmt.Printf("Reply recieved: %d bytes, RTT: %n bytes",n,rtt)



	return nil;

}