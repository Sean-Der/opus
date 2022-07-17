package silk

var (
	// +----------+-----------------------------+
	// | VAD Flag | PDF                         |
	// +----------+-----------------------------+
	// | Inactive | {26, 230, 0, 0, 0, 0}/256   |
	// |          |                             |
	// | Active   | {0, 0, 24, 74, 148, 10}/256 |
	// +----------+-----------------------------+
	//
	// https://datatracker.ietf.org/doc/html/rfc6716#section-4.2.7.3
	icdfFrameTypeVADInactive = []uint{256, 26, 256}
	icdfFrameTypeVADActive   = []uint{256, 24, 98, 246, 256}

	// +-------------+------------------------------------+
	// | Signal Type | PDF                                |
	// +-------------+------------------------------------+
	// | Inactive    | {32, 112, 68, 29, 12, 1, 1, 1}/256 |
	// |             |                                    |
	// | Unvoiced    | {2, 17, 45, 60, 62, 47, 19, 4}/256 |
	// |             |                                    |
	// | Voiced      | {1, 3, 26, 71, 94, 50, 9, 2}/256   |
	// +-------------+------------------------------------+
	//
	// https://datatracker.ietf.org/doc/html/rfc6716#section-4.2.7.4
	icdfIndependentQuantizationGainMSBInactive = []uint{256, 32, 144, 212, 241, 253, 254, 255, 256}
	icdfIndependentQuantizationGainMSBUnvoiced = []uint{256, 2, 19, 64, 124, 186, 233, 252, 256}
	icdfIndependentQuantizationGainMSBVoiced   = []uint{256, 1, 4, 30, 101, 195, 245, 254, 256}

	// +--------------------------------------+
	// | PDF                                  |
	// +--------------------------------------+
	// | {32, 32, 32, 32, 32, 32, 32, 32}/256 |
	// +--------------------------------------+
	//
	// https://datatracker.ietf.org/doc/html/rfc6716#section-4.2.7.4
	icdfIndependentQuantizationGainLSB = []uint{256, 32, 64, 96, 128, 160, 192, 224, 256}

	// +-------------------------------------------------------------------+
	// | PDF                                                               |
	// +-------------------------------------------------------------------+
	// | {6, 5, 11, 31, 132, 21, 8, 4, 3, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, |
	// | 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,       |
	// | 1}/256                                                            |
	// +-------------------------------------------------------------------+
	icdfDeltaQuantizationGain = []uint{
		256, 6, 11, 22, 53, 185, 206, 214, 218, 221, 223, 225, 227, 228,
		229, 230, 231, 232, 233, 234, 235, 236, 237, 238, 239, 240, 241, 242,
		243, 244, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256,
	}
)