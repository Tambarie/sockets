package main
import(
	"fmt"
	"net"
	"os"
)

//// IPMask To handle masking operations, there is the type
//type IPMask []byte
//func IPv4Mask(a,b,c,d byte) IPMask{
//
//}
//
//// DefaultMask Function to create a mask from a 4-byte IPv4 address
//func (ip IP) DefaultMask() IPMask {
//
//}
//
//// Mask Method which returns the default mask
//func (ip IP) Mask(mask IPMask)  IP{
//
//}


func main()  {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s dotted-ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	dotAddr := os.Args[1]

	addr := net.ParseIP(dotAddr)
	if addr == nil{
		fmt.Println("Invalid address")
		os.Exit(1)
	}
	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, bits := mask.Size()
	fmt.Println("Address is ", addr.String(),
		"Default mask length is ", bits,
		"Leading ones count is ",ones,
		"Mask is (hex) ", mask.String(),
		"Network is ", network.String())

	os.Exit(0)

}