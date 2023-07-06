int min(int x, int y){
	if (x < y){
		return x;
	}
	return y;
}

int max(int x, int y){
	if (x > y){
		return x;
	}
	return y;
}


int kItemsWithMaximumSum(int numOnes, int numZeros, int numNegOnes, int k){
	return min(k, numOnes) - max(k-numOnes-numZeros, 0);
}