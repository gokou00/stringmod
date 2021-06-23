package strings


func internalFunction(){
	// test test
}

func CountOddEven(s string)(odds,evens int){
	odds ,evens = 0,0

	for _, c := range s{
		if int(c) % 2 == 0{
			evens++
		}else{
			odds++
		}
	}
	return	
}