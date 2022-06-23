package main

// #cgo LDFLAGS: -L./opus-celt-decoder -locd
// #include <stdio.h>
// #include <stdlib.h>
// #include <./opus-celt-decoder/opus.h>
//
//void DecodeAndCompare(void *input_buffer, int input_buffer_len, void *output_buffer) {
//      OpusDecoder *dec = NULL;
//      int err, output_samples;
//
//      if ((dec = opus_decoder_create(48000, 2, &err)) && err != OPUS_OK) {
//            fprintf(stderr, "Cannot create decoder: %s\n", opus_strerror(err));
//            return;
//      }
//
//      if ((output_samples = opus_decode(dec, input_buffer, input_buffer_len, output_buffer, 1920, 0)) < 0) {
//            fprintf(stderr, "error decoding frame: %s\n", opus_strerror(output_samples));
//						return;
//      }
//}
import "C"
import (
	"fmt"
	"os"

	"github.com/pion/opus"
)

func convertUsingC(in []byte) []byte {
	decodeAndCompareOutput := C.CBytes(make([]byte, 3840))
	defer C.free(decodeAndCompareOutput)

	decodeAndCompareInput := C.CBytes(in)
	defer C.free(decodeAndCompareInput)
	C.DecodeAndCompare(decodeAndCompareInput, C.int(len(in)), decodeAndCompareOutput)

	return C.GoBytes(decodeAndCompareOutput, 3840)
}

func main() {
	decoder := &opus.Decoder{}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	file, err := os.Open(homeDir + "/opus/single_buffer")
	if err != nil {
		panic(err)
	}

	fileinfo, err := file.Stat()
	if err != nil {
		panic(err)
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	if _, err = file.Read(buffer); err != nil {
		panic(err)
	}

	outputBuffer := convertUsingC(buffer)
	fmt.Printf("% 02x \n", outputBuffer)

	bandwidth, isStereo, frames, err := decoder.Decode(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("bandwidth(%s) isStereo(%t) framesCount(%d)\n", bandwidth.String(), isStereo, len(frames))
}
