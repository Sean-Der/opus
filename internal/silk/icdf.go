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
	// PDF for Delta Quantization Gain Coding
	icdfDeltaQuantizationGain = []uint{
		256, 6, 11, 22, 53, 185, 206, 214, 218, 221, 223, 225, 227, 228,
		229, 230, 231, 232, 233, 234, 235, 236, 237, 238, 239, 240, 241, 242,
		243, 244, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256,
	}

	// +-----------+----------+--------------------------------------------+
	// | Audio     | Signal   | PDF                                        |
	// | Bandwidth | Type     |                                            |
	// +-----------+----------+--------------------------------------------+
	// | NB or MB  | Inactive | {44, 34, 30, 19, 21, 12, 11, 3, 3, 2, 16,  |
	// |           | or       | 2, 2, 1, 5, 2, 1, 3, 3, 1, 1, 2, 2, 2, 3,  |
	// |           | unvoiced | 1, 9, 9, 2, 7, 2, 1}/256                   |
	// |           |          |                                            |
	// | NB or MB  | Voiced   | {1, 10, 1, 8, 3, 8, 8, 14, 13, 14, 1, 14,  |
	// |           |          | 12, 13, 11, 11, 12, 11, 10, 10, 11, 8, 9,  |
	// |           |          | 8, 7, 8, 1, 1, 6, 1, 6, 5}/256             |
	// |           |          |                                            |
	// | WB        | Inactive | {31, 21, 3, 17, 1, 8, 17, 4, 1, 18, 16, 4, |
	// |           | or       | 2, 3, 1, 10, 1, 3, 16, 11, 16, 2, 2, 3, 2, |
	// |           | unvoiced | 11, 1, 4, 9, 8, 7, 3}/256                  |
	// |           |          |                                            |
	// | WB        | Voiced   | {1, 4, 16, 5, 18, 11, 5, 14, 15, 1, 3, 12, |
	// |           |          | 13, 14, 14, 6, 14, 12, 2, 6, 1, 12, 12,    |
	// |           |          | 11, 10, 3, 10, 5, 1, 1, 1, 3}/256          |
	// +-----------+----------+--------------------------------------------+
	// PDFs for Normalized LSF Stage-1 Index Decoding
	icdfNormalizedLSFStageOneIndexNarrowbandOrMediumbandUnvoiced = []uint{
		256, 44, 78, 108, 127, 148, 160, 171, 174, 177, 179, 195, 197, 199, 200,
		205, 207, 208, 211, 214, 215, 216, 218, 220, 222, 225, 226, 235, 244, 246,
		253, 255, 256,
	}
	icdfNormalizedLSFStageOneIndexNarrowbandOrMediumbandVoiced = []uint{
		256, 1, 11, 12, 20, 23, 31, 39, 53, 66, 80, 81, 95, 107, 120, 131, 142, 154,
		165, 175, 185, 196, 204, 213, 221, 228, 236, 237, 238, 244, 245, 251, 256,
	}
	icdfNormalizedLSFStageOneIndexWidebandUnvoiced = []uint{
		256, 31, 52, 55, 72, 73, 81, 98, 102, 103, 121, 137, 141, 143, 146, 147, 157,
		158, 161, 177, 188, 204, 206, 208, 211, 213, 224, 225, 229, 238, 246, 253, 256,
	}
	icdfNormalizedLSFStageOneIndexWidebandVoiced = []uint{
		256, 1, 5, 21, 26, 44, 55, 60, 74, 89, 90, 93, 105, 118, 132, 146, 152, 166, 178,
		180, 186, 187, 199, 211, 222, 232, 235, 245, 250, 251, 252, 253, 256,
	}

	// +-------------------------------+
	// | PDF                           |
	// +-------------------------------+
	// | {156, 60, 24, 9, 4, 2, 1}/256 |
	// +-------------------------------+
	//
	// Normalized LSF Index Extension Decoding
	icdfNormalizedLSFStageTwoIndexExtension = []uint{256, 156, 216, 240, 249, 253, 255, 256}

	// +----------+--------------------------------------+
	// | Codebook | PDF                                  |
	// +----------+--------------------------------------+
	// | a        | {1, 1, 1, 15, 224, 11, 1, 1, 1}/256  |
	// |          |                                      |
	// | b        | {1, 1, 2, 34, 183, 32, 1, 1, 1}/256  |
	// |          |                                      |
	// | c        | {1, 1, 4, 42, 149, 55, 2, 1, 1}/256  |
	// |          |                                      |
	// | d        | {1, 1, 8, 52, 123, 61, 8, 1, 1}/256  |
	// |          |                                      |
	// | e        | {1, 3, 16, 53, 101, 74, 6, 1, 1}/256 |
	// |          |                                      |
	// | f        | {1, 3, 17, 55, 90, 73, 15, 1, 1}/256 |
	// |          |                                      |
	// | g        | {1, 7, 24, 53, 74, 67, 26, 3, 1}/256 |
	// |          |                                      |
	// | h        | {1, 1, 18, 63, 78, 58, 30, 6, 1}/256 |
	// +----------+--------------------------------------+
	//  PDFs for NB/MB Normalized LSF Stage-2 Index Decoding
	//
	// +----------+---------------------------------------+
	// | Codebook | PDF                                   |
	// +----------+---------------------------------------+
	// | i        | {1, 1, 1, 9, 232, 9, 1, 1, 1}/256     |
	// |          |                                       |
	// | j        | {1, 1, 2, 28, 186, 35, 1, 1, 1}/256   |
	// |          |                                       |
	// | k        | {1, 1, 3, 42, 152, 53, 2, 1, 1}/256   |
	// |          |                                       |
	// | l        | {1, 1, 10, 49, 126, 65, 2, 1, 1}/256  |
	// |          |                                       |
	// | m        | {1, 4, 19, 48, 100, 77, 5, 1, 1}/256  |
	// |          |                                       |
	// | n        | {1, 1, 14, 54, 100, 72, 12, 1, 1}/256 |
	// |          |                                       |
	// | o        | {1, 1, 15, 61, 87, 61, 25, 4, 1}/256  |
	// |          |                                       |
	// | p        | {1, 7, 21, 50, 77, 81, 17, 1, 1}/256  |
	// +----------+---------------------------------------+
	// PDFs for WB Normalized LSF Stage-2 Index Decoding
	//
	// NB/MD and WD ICDF are combined because the codebooks
	// do not overlap
	//
	icdfNormalizedLSFStageTwoIndex = [][]uint{
		// Narrowband and Mediumband
		{256, 1, 2, 3, 18, 242, 253, 254, 255, 256},
		{256, 1, 2, 4, 38, 221, 253, 254, 255, 256},
		{256, 1, 2, 6, 48, 197, 252, 254, 255, 256},
		{256, 1, 2, 10, 62, 185, 246, 254, 255, 256},
		{256, 1, 4, 20, 73, 174, 248, 254, 255, 256},
		{256, 1, 4, 21, 76, 166, 239, 254, 255, 256},
		{256, 1, 8, 32, 85, 159, 226, 252, 255, 256},
		{256, 1, 2, 20, 83, 161, 219, 249, 255, 256},

		// Wideband
		{256, 1, 2, 3, 12, 244, 253, 254, 255, 256},
		{256, 1, 2, 4, 32, 218, 253, 254, 255, 256},
		{256, 1, 2, 5, 47, 199, 252, 254, 255, 256},
		{256, 1, 2, 12, 61, 187, 252, 254, 255, 256},
		{256, 1, 5, 24, 72, 172, 249, 254, 255, 256},
		{256, 1, 2, 16, 70, 170, 242, 254, 255, 256},
		{256, 1, 2, 17, 78, 165, 226, 251, 255, 256},
		{256, 1, 8, 29, 79, 156, 237, 254, 255, 256},
	}
)
