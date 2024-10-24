package main

func main() {
	arr := []int{1, 5, 7, 8, 2, 1, 0, -232, -3, -3}
	sorted := quickSort(arr)

}

func quickSort(arr []int) []int {
	pivot := partition(arr, 0, len(arr)-1)
	//TODO: finish
}

func partition(arr []int, low int, high int) interface{} {

}
