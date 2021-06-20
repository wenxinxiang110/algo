package array

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	// is find the mountain's top
	findFirstTop := false

	if arr[0] > arr[1] {
		return false
	}

	for i := 1; i < len(arr); i++ {
		if !findFirstTop {
			if arr[i] > arr[i-1] {
				continue
			}
			findFirstTop = true
		}

		if findFirstTop {
			if arr[i] >= arr[i-1] {
				return false
			}
		}

	}

	if arr[len(arr)-1] >= arr[len(arr)-2] {
		return false
	}

	return true
}
