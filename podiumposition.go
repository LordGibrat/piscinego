package piscine

func PodiumPosition(podium [][]string) [][]string {
	for left, right := 0, len(podium)-1; left < right; left, right = left+1, right-1 {
		podium[left], podium[right] = podium[right], podium[left]
	}
	return podium
}
