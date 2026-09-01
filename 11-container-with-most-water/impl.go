package containerwater

// MaxArea finds two lines that together with the x-axis form a container that contains the most water.
// Time Complexity: O(N), Space Complexity: O(1)
func MaxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxWater := 0

	for left < right {
		hLeft := height[left]
		hRight := height[right]

		// Height of the water is bounded by the shorter line
		currentHeight := hLeft
		if hRight < currentHeight {
			currentHeight = hRight
		}

		currentWidth := right - left
		currentArea := currentHeight * currentWidth

		if currentArea > maxWater {
			maxWater = currentArea
		}

		// Move the pointer of the shorter line inward to search for higher bounds
		if hLeft < hRight {
			left++
		} else {
			right--
		}
	}

	return maxWater
}
