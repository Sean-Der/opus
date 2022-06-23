#include <stdio.h>
#include <stdlib.h>
#include "opus.h"
#include "/Users/duboisea/opus/single_buffer.h"

void print_buffer_int16(int16_t *buff, int len) {
      for (int i = 0; i < len; i++) {
            printf("%hd, ", buff[i]);
            if ((i % 25) == 0 && i != 0) {
                  printf("\n");
            }
      }
}

void compare_buffer(int16_t *a, int len) {
      int16_t *b = (int16_t*)calloc(sizeof(int16_t), 1920);
      fread(b, sizeof(int16_t), 1920, fopen("/Users/duboisea/opus/decoded_single_buffer" , "r"));

      for (int i = 0; i < len; i++) {
            if (a[i] != b[i]) {
                  fprintf(stderr, "buffer mismatch at %d", i);
            } else {
                  fprintf(stderr, "buffer match at %d", i);
            }
      }
}

int main() {
      OpusDecoder *dec = NULL;
      int err, output_samples;
      int16_t *out = (int16_t*)calloc(sizeof(int16_t), 1920);

      if ((dec = opus_decoder_create(48000, 2, &err)) && err != OPUS_OK) {
            fprintf(stderr, "Cannot create decoder: %s\n", opus_strerror(err));
            return EXIT_FAILURE;
      }

      if ((output_samples = opus_decode(dec, single_buffer, single_buffer_len, out, 1920, 0)) < 0) {
            fprintf(stderr, "error decoding frame: %s\n", opus_strerror(output_samples));
      }

      compare_buffer(out, 1920);
}
