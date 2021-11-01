package quest

func reorderedPowerOf2(n int) bool {

	//2的幂 那在二进制下只有一个1

	tmp:=n
	valid:=true
	for tmp!=1{
		if tmp&1==1{
			valid=false
			break
		}
		tmp=tmp>>1
	}
	if valid{
		return true
	}
}
