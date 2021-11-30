// Copyright © 2015-2020 Hilko Bengen <bengen@hilluzination.de>
// All rights reserved.
//
// Use of this source code is governed by the license that can be
// found in the LICENSE file.

package yara

// #cgo darwin,arm64 LDFLAGS: -L${SRCDIR}/libs/darwin/arm64 -lyara -lm -lmagic -ljansson -lz -lbz2 -llzma -lcrypto
// #cgo darwin,arm64 CFLAGS: -I${SRCDIR}/libs/darwin/arm64/include
// #cgo !yara_no_pkg_config,yara_static   pkg-config: --static yara
// #cgo yara_no_pkg_config                LDFLAGS:    -lyara
/*
#include <yara.h>
#if YR_VERSION_HEX < 0x040100
#error YARA version 4.1 required
#endif
*/
import "C"
