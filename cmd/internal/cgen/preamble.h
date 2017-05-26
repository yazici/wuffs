// After editing this file, run "go generate" in this directory.

// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

#include <stdbool.h>
#include <stdint.h>

#if __WORDSIZE < 32
#error "Puffs requires a word size of at least 32 bits"
#endif
