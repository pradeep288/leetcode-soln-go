// Reference: https://www.youtube.com/watch?v=5LMkddl2NCk
func countPrimes(n int) int {
    
    if n == 0 || n== 1 {
		return 0
	}
	
    compositeArray := make([]bool, n)

	compositeArray[0], compositeArray[1] = true, true

	limit:= int(math.Sqrt(float64(n)))
	
	for i:=2; i<=limit; i++{
		if ! compositeArray[i] {
			for j := i * i; j < n; j = j + i {
				compositeArray[j] = true
			}
		}
	}
	
    count:= 0
	for i:= range compositeArray{
		if compositeArray[i] == false {
			count += 1
		}
	}

    return count 
    
}
