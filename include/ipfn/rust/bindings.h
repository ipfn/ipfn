/**
 * Copyright © 2018 The IPFN Developers. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/* Generated with cbindgen:0.6.6 */

/* Warning, this file is autogenerated by cbindgen. Don't modify this manually. */

#include <cstdint>
#include <cstdlib>

namespace ipfn {
namespace ffi {

// Normal example docs.
struct Normal {
  int32_t x;
  int32_t y;
};

extern "C" {

// Addition example docs.
extern int32_t addition(int32_t a, int32_t b);

// Normalize example docs.
extern int32_t normalize(Normal n);

extern void theme_song_free(char *s);

extern char *theme_song_generate(const char *name);

} // extern "C"

} // namespace ffi
} // namespace ipfn