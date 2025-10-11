package leetcode

func checkTwoChessboards(coordinate1 string, coordinate2 string) bool {
	return (coordinate1[0]-coordinate2[0])%2+(coordinate1[1]-coordinate2[1])%2 == 1
}
