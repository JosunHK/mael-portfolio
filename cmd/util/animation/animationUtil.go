package animationUtil

func isPortrait(h, w int32) bool {
	if h > w {
		return true
	}
	return false
}
