package silk

type (
	// Bandwidth for Silk can be NB (narrowband) MB (medium-band) or WB (wideband)
	Bandwidth byte

	frameSignalType             byte
	frameQuantizationOffsetType byte
)

const (
	nanoseconds20Ms = 20000000

	frameSignalTypeInactive frameSignalType = iota + 1
	frameSignalTypeUnvoiced
	frameSignalTypeVoiced

	frameQuantizationOffsetTypeLow frameQuantizationOffsetType = iota + 1
	frameQuantizationOffsetTypeHigh
)

// Bandwidth constants
const (
	BandwidthNarrowband Bandwidth = iota + 1
	BandwidthMediumband
	BandwidthWideband
)

func maxUint32(a, b uint32) uint32 {
	if a > b {
		return a
	}
	return b
}

func maxInt32(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func clamp(low, in, high int32) int32 {
	if in > high {
		return high
	} else if in < low {
		return low
	}

	return in
}

// The sign of x, i.e.,
//            ( -1,  x < 0
//  sign(x) = <  0,  x == 0
//            (  1,  x > 0
// https://datatracker.ietf.org/doc/html/rfc6716#section-1.1.4
func sign(x int) int {
	switch {
	case x < 0:
		return -1
	case x == 0:
		return 0
	default:
		return 1
	}
}
