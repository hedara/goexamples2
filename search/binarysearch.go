package search

func FindSearchIterations(lower int,upper int, item int) int {
	counter := 1
	pick := lower + (upper - lower)/2
	if item > pick {
		return counter + FindSearchIterations(pick+1, upper,item)
	} else if item < pick {
		return counter + FindSearchIterations(lower,pick, item)
	}else {
		return 0
	}
}
