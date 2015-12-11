#include <cilk/cilk.h>
#include "foo.h"

int sum(int *arr, int i0, int i1) {
  if ((i1-i0) > 1) {
    int im = (i1+i0) / 2;
    int a = cilk_spawn sum(arr, i0, im);
    int b = cilk_spawn sum(arr, im, i1);
    return a+b;
  } else {
    return arr[i0];
  }
}
