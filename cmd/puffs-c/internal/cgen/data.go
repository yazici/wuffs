// Code generated by running "go generate". DO NOT EDIT.

// Copyright 2017 The Puffs Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cgen

const baseHeader = "" +
	"#ifndef PUFFS_BASE_HEADER_H\n#define PUFFS_BASE_HEADER_H\n\n// Copyright 2017 The Puffs Authors.\n//\n// Licensed under the Apache License, Version 2.0 (the \"License\");\n// you may not use this file except in compliance with the License.\n// You may obtain a copy of the License at\n//\n//    https://www.apache.org/licenses/LICENSE-2.0\n//\n// Unless required by applicable law or agreed to in writing, software\n// distributed under the License is distributed on an \"AS IS\" BASIS,\n// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n// See the License for the specific language governing permissions and\n// limitations under the License.\n\n#include <stdbool.h>\n#include <stdint.h>\n#include <string.h>\n\n// Puffs requires a word size of at least 32 bits because it assumes that\n// converting a u32 to usize will never overflow. For example, the size of a\n// decoded image is often represented, explicitly or implicitly in an image\n// file, as a u32, and it is convenient to compare that to a buffer size.\n//\n// Si" +
	"milarly, the word size is at most 64 bits because it assumes that\n// converting a usize to u64 will never overflow.\n#if __WORDSIZE < 32\n#error \"Puffs requires a word size of at least 32 bits\"\n#elif __WORDSIZE > 64\n#error \"Puffs requires a word size of at most 64 bits\"\n#endif\n\n// PUFFS_VERSION is the major.minor version number as a uint32. The major\n// number is the high 16 bits. The minor number is the low 16 bits.\n//\n// The intention is to bump the version number at least on every API / ABI\n// backwards incompatible change.\n//\n// For now, the API and ABI are simply unstable and can change at any time.\n//\n// TODO: don't hard code this in base-header.h.\n#define PUFFS_VERSION (0x00001)\n\n// puffs_base__slice_u8 is a 1-dimensional buffer (a pointer and length).\n//\n// A value with all fields NULL or zero is a valid, empty slice.\ntypedef struct {\n  uint8_t* ptr;\n  size_t len;\n} puffs_base__slice_u8;\n\n// puffs_base__buf1 is a 1-dimensional buffer (a pointer and length), plus\n// additional indexes into that buffer, p" +
	"lus an opened / closed flag.\n//\n// A value with all fields NULL or zero is a valid, empty buffer.\ntypedef struct {\n  uint8_t* ptr;  // Pointer.\n  size_t len;    // Length.\n  size_t wi;     // Write index. Invariant: wi <= len.\n  size_t ri;     // Read  index. Invariant: ri <= wi.\n  bool closed;   // No further writes are expected.\n} puffs_base__buf1;\n\n// puffs_base__limit1 provides a limited view of a 1-dimensional byte stream:\n// its first N bytes. That N can be greater than a buffer's current read or\n// write capacity. N decreases naturally over time as bytes are read from or\n// written to the stream.\n//\n// A value with all fields NULL or zero is a valid, unlimited view.\ntypedef struct puffs_base__limit1 {\n  uint64_t* ptr_to_len;             // Pointer to N.\n  struct puffs_base__limit1* next;  // Linked list of limits.\n} puffs_base__limit1;\n\ntypedef struct {\n  // TODO: move buf into private_impl? As it is, it looks like users can modify\n  // the buf field to point to a different buffer, which can turn the l" +
	"imit and\n  // mark fields into dangling pointers.\n  puffs_base__buf1* buf;\n  // Do not access the private_impl's fields directly. There is no API/ABI\n  // compatibility or safety guarantee if you do so.\n  struct {\n    puffs_base__limit1 limit;\n    uint8_t* mark;\n  } private_impl;\n} puffs_base__reader1;\n\ntypedef struct {\n  // TODO: move buf into private_impl? As it is, it looks like users can modify\n  // the buf field to point to a different buffer, which can turn the limit and\n  // mark fields into dangling pointers.\n  puffs_base__buf1* buf;\n  // Do not access the private_impl's fields directly. There is no API/ABI\n  // compatibility or safety guarantee if you do so.\n  struct {\n    puffs_base__limit1 limit;\n    uint8_t* mark;\n  } private_impl;\n} puffs_base__writer1;\n\n#endif  // PUFFS_BASE_HEADER_H\n" +
	""

const baseImpl = "" +
	"#ifndef PUFFS_BASE_IMPL_H\n#define PUFFS_BASE_IMPL_H\n\n// Copyright 2017 The Puffs Authors.\n//\n// Licensed under the Apache License, Version 2.0 (the \"License\");\n// you may not use this file except in compliance with the License.\n// You may obtain a copy of the License at\n//\n//    https://www.apache.org/licenses/LICENSE-2.0\n//\n// Unless required by applicable law or agreed to in writing, software\n// distributed under the License is distributed on an \"AS IS\" BASIS,\n// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n// See the License for the specific language governing permissions and\n// limitations under the License.\n\n// puffs_base__empty_struct is used when a Puffs function returns an empty\n// struct. In C, if a function f returns void, you can't say \"x = f()\", but in\n// Puffs, if a function g returns empty, you can say \"y = g()\".\ntypedef struct {\n} puffs_base__empty_struct;\n\n#define PUFFS_BASE__IGNORE_POTENTIALLY_UNUSED_VARIABLE(x) (void)(x)\n\n// PUFFS_BASE__MAGIC is a magic number to " +
	"check that initializers are called.\n// It's not foolproof, given C doesn't automatically zero memory before use,\n// but it should catch 99.99% of cases.\n//\n// Its (non-zero) value is arbitrary, based on md5sum(\"puffs\").\n#define PUFFS_BASE__MAGIC (0xCB3699CCU)\n\n// PUFFS_BASE__ALREADY_ZEROED is passed from a container struct's initializer\n// to a containee struct's initializer when the container has already zeroed\n// the containee's memory.\n//\n// Its (non-zero) value is arbitrary, based on md5sum(\"zeroed\").\n#define PUFFS_BASE__ALREADY_ZEROED (0x68602EF1U)\n\n// Use switch cases for coroutine suspension points, similar to the technique\n// in https://www.chiark.greenend.org.uk/~sgtatham/coroutines.html\n//\n// We use a trivial macro instead of an explicit assignment and case statement\n// so that clang-format doesn't get confused by the unusual \"case\"s.\n#define PUFFS_BASE__COROUTINE_SUSPENSION_POINT(n) \\\n  coro_susp_point = n;                            \\\n  case n:;\n\n#define PUFFS_BASE__COROUTINE_SUSPENSION_POINT_MAYB" +
	"E_SUSPEND(n) \\\n  if (status < 0) {                                             \\\n    goto exit;                                                  \\\n  } else if (status == 0) {                                     \\\n    goto ok;                                                    \\\n  }                                                             \\\n  coro_susp_point = n;                                          \\\n  goto suspend;                                                 \\\n  case n:;\n\n// Clang also defines \"__GNUC__\".\n#if defined(__GNUC__)\n#define PUFFS_BASE__LIKELY(expr) (__builtin_expect(!!(expr), 1))\n#define PUFFS_BASE__UNLIKELY(expr) (__builtin_expect(!!(expr), 0))\n#else\n#define PUFFS_BASE__LIKELY(expr) (expr)\n#define PUFFS_BASE__UNLIKELY(expr) (expr)\n#endif\n\n// Uncomment this #include for printf-debugging.\n// #include <stdio.h>\n\n// ---------------- Static Inline Functions\n//\n// The helpers below are functions, instead of macros, because their arguments\n// can be an expression that we shouldn't evaluate mo" +
	"re than once.\n//\n// They are in base-impl.h and hence copy/pasted into every generated C file,\n// instead of being in some \"base.c\" file, since a design goal is that users of\n// the generated C code can often just #include a single .c file, such as\n// \"gif.c\", without having to additionally include or otherwise build and link\n// a \"base.c\" file.\n//\n// They are static, so that linking multiple puffs .o files won't complain about\n// duplicate function definitions.\n//\n// They are explicitly marked inline, even if modern compilers don't use the\n// inline attribute to guide optimizations such as inlining, to avoid the\n// -Wunused-function warning, and we like to compile with -Wall -Werror.\n\nstatic inline uint16_t puffs_base__load_u16be(uint8_t* p) {\n  return ((uint16_t)(p[0]) << 8) | ((uint16_t)(p[1]) << 0);\n}\n\nstatic inline uint16_t puffs_base__load_u16le(uint8_t* p) {\n  return ((uint16_t)(p[0]) << 0) | ((uint16_t)(p[1]) << 8);\n}\n\nstatic inline uint32_t puffs_base__load_u32be(uint8_t* p) {\n  return ((uint32_t)(p[" +
	"0]) << 24) | ((uint32_t)(p[1]) << 16) |\n         ((uint32_t)(p[2]) << 8) | ((uint32_t)(p[3]) << 0);\n}\n\nstatic inline uint32_t puffs_base__load_u32le(uint8_t* p) {\n  return ((uint32_t)(p[0]) << 0) | ((uint32_t)(p[1]) << 8) |\n         ((uint32_t)(p[2]) << 16) | ((uint32_t)(p[3]) << 24);\n}\n\nstatic inline puffs_base__slice_u8 puffs_base__slice_u8__subslice_i(\n    puffs_base__slice_u8 s,\n    uint64_t i) {\n  if ((i <= SIZE_MAX) && (i <= s.len)) {\n    return ((puffs_base__slice_u8){\n        .ptr = s.ptr + i,\n        .len = s.len - i,\n    });\n  }\n  return ((puffs_base__slice_u8){});\n}\n\nstatic inline puffs_base__slice_u8 puffs_base__slice_u8__subslice_j(\n    puffs_base__slice_u8 s,\n    uint64_t j) {\n  if ((j <= SIZE_MAX) && (j <= s.len)) {\n    return ((puffs_base__slice_u8){.ptr = s.ptr, .len = j});\n  }\n  return ((puffs_base__slice_u8){});\n}\n\nstatic inline puffs_base__slice_u8 puffs_base__slice_u8__subslice_ij(\n    puffs_base__slice_u8 s,\n    uint64_t i,\n    uint64_t j) {\n  if ((i <= j) && (j <= SIZE_MAX) && (j <= s.l" +
	"en)) {\n    return ((puffs_base__slice_u8){\n        .ptr = s.ptr + i,\n        .len = j - i,\n    });\n  }\n  return ((puffs_base__slice_u8){});\n}\n\n// puffs_base__slice_u8__prefix returns up to the first up_to bytes of s.\nstatic inline puffs_base__slice_u8 puffs_base__slice_u8__prefix(\n    puffs_base__slice_u8 s,\n    uint64_t up_to) {\n  if ((uint64_t)(s.len) > up_to) {\n    s.len = up_to;\n  }\n  return s;\n}\n\n// puffs_base__slice_u8__suffix returns up to the last up_to bytes of s.\nstatic inline puffs_base__slice_u8 puffs_base__slice_u8_suffix(\n    puffs_base__slice_u8 s,\n    uint64_t up_to) {\n  if ((uint64_t)(s.len) > up_to) {\n    s.ptr += (uint64_t)(s.len) - up_to;\n    s.len = up_to;\n  }\n  return s;\n}\n\n// puffs_base__slice_u8__copy_from_slice calls memmove(dst.ptr, src.ptr,\n// length) where length is the minimum of dst.len and src.len.\n//\n// Passing a puffs_base__slice_u8 with all fields NULL or zero (a valid, empty\n// slice) is valid and results in a no-op.\nstatic inline uint64_t puffs_base__slice_u8__copy_from_sli" +
	"ce(\n    puffs_base__slice_u8 dst,\n    puffs_base__slice_u8 src) {\n  size_t length = dst.len < src.len ? dst.len : src.len;\n  if (length > 0) {\n    memmove(dst.ptr, src.ptr, length);\n  }\n  return length;\n}\n\nstatic inline uint32_t puffs_base__writer1__copy_from_history32(\n    uint8_t** ptr_ptr,\n    uint8_t* start,  // May be NULL, meaning an unmarked writer1.\n    uint8_t* end,\n    uint32_t distance,\n    uint32_t length) {\n  if (!start) {\n    return 0;\n  }\n  uint8_t* ptr = *ptr_ptr;\n  size_t d = ptr - start;\n  if ((d == 0) || (d < (size_t)(distance))) {\n    return 0;\n  }\n  start = ptr - distance;\n  size_t n = end - ptr;\n  if ((size_t)(length) > n) {\n    length = n;\n  } else {\n    n = length;\n  }\n  // TODO: is manual unrolling actually helpful?\n  for (; n >= 8; n -= 8) {\n    *ptr++ = *start++;\n    *ptr++ = *start++;\n    *ptr++ = *start++;\n    *ptr++ = *start++;\n    *ptr++ = *start++;\n    *ptr++ = *start++;\n    *ptr++ = *start++;\n    *ptr++ = *start++;\n  }\n  for (; n; n--) {\n    *ptr++ = *start++;\n  }\n  *ptr_ptr =" +
	" ptr;\n  return length;\n}\n\nstatic inline uint32_t puffs_base__writer1__copy_from_reader32(\n    uint8_t** ptr_wptr,\n    uint8_t* wend,\n    uint8_t** ptr_rptr,\n    uint8_t* rend,\n    uint32_t length) {\n  uint8_t* wptr = *ptr_wptr;\n  size_t n = length;\n  if (n > wend - wptr) {\n    n = wend - wptr;\n  }\n  uint8_t* rptr = *ptr_rptr;\n  if (n > rend - rptr) {\n    n = rend - rptr;\n  }\n  if (n > 0) {\n    memmove(wptr, rptr, n);\n    *ptr_wptr += n;\n    *ptr_rptr += n;\n  }\n  return n;\n}\n\nstatic inline uint64_t puffs_base__writer1__copy_from_slice(\n    uint8_t** ptr_wptr,\n    uint8_t* wend,\n    puffs_base__slice_u8 src) {\n  uint8_t* wptr = *ptr_wptr;\n  size_t n = src.len;\n  if (n > wend - wptr) {\n    n = wend - wptr;\n  }\n  if (n > 0) {\n    memmove(wptr, src.ptr, n);\n    *ptr_wptr += n;\n  }\n  return n;\n}\n\nstatic inline uint32_t puffs_base__writer1__copy_from_slice32(\n    uint8_t** ptr_wptr,\n    uint8_t* wend,\n    puffs_base__slice_u8 src,\n    uint32_t length) {\n  uint8_t* wptr = *ptr_wptr;\n  size_t n = src.len;\n  if (n > le" +
	"ngth) {\n    n = length;\n  }\n  if (n > wend - wptr) {\n    n = wend - wptr;\n  }\n  if (n > 0) {\n    memmove(wptr, src.ptr, n);\n    *ptr_wptr += n;\n  }\n  return n;\n}\n\n// Note that the *__limit and *__mark methods are private (in base-impl.h) not\n// public (in base-header.h). We assume that, at the boundary between user code\n// and Puffs code, the reader1 and writer1's private_impl fields (including\n// limit and mark) are NULL. Otherwise, some internal assumptions break down.\n// For example, limits could be represented as pointers, even though\n// conceptually they are counts, but that pointer-to-count correspondence\n// becomes invalid if a buffer is re-used (e.g. on resuming a coroutine).\n//\n// Admittedly, some of the Puffs test code calls these methods, but that test\n// code is still Puffs code, not user code. Other Puffs test code modifies\n// private_impl fields directly.\n\nstatic inline puffs_base__reader1 puffs_base__reader1__limit(\n    puffs_base__reader1* o,\n    uint64_t* ptr_to_len) {\n  puffs_base__reader1 r" +
	"et = *o;\n  ret.private_impl.limit.ptr_to_len = ptr_to_len;\n  ret.private_impl.limit.next = &o->private_impl.limit;\n  return ret;\n}\n\nstatic inline puffs_base__empty_struct puffs_base__reader1__mark(\n    puffs_base__reader1* o,\n    uint8_t* mark) {\n  o->private_impl.mark = mark;\n  return ((puffs_base__empty_struct){});\n}\n\n// TODO: static inline puffs_base__writer1 puffs_base__writer1__limit()\n\nstatic inline puffs_base__empty_struct puffs_base__writer1__mark(\n    puffs_base__writer1* o,\n    uint8_t* mark) {\n  o->private_impl.mark = mark;\n  return ((puffs_base__empty_struct){});\n}\n\n#endif  // PUFFS_BASE_IMPL_H\n" +
	""

type template_args_short_read struct {
	PKGPREFIX string
	name      string
}

func template_short_read(b *buffer, args template_args_short_read) error {
	b.printf("short_read_%s:\nif (a_%s.buf && a_%s.buf->closed &&\n!a_%s.private_impl.limit.ptr_to_len) {\nstatus = %sERROR_UNEXPECTED_EOF;\ngoto exit;\n}\nstatus = %sSUSPENSION_SHORT_READ;\ngoto suspend;\n",
		args.name,
		args.name,
		args.name,
		args.name,
		args.PKGPREFIX,
		args.PKGPREFIX,
	)
	return nil
}
