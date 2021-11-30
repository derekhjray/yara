// Copyright @ 2021-2022 Derek Ray <derek9ray@gmail.com>
// Copyright © 2015-2020 Hilko Bengen <bengen@hilluzination.de>
// All rights reserved.
//
// Use of this source code is governed by the license that can be
// found in the LICENSE file.

package yara

// #cgo !user_libs,darwin,arm64 LDFLAGS: -L${SRCDIR}/libs/darwin/arm64 -lyara -lm -lmagic -ljansson -lz -lbz2 -llzma -lcrypto
// #cgo !user_libs,darwin,arm64 CFLAGS: -I${SRCDIR}/libs/darwin/arm64/include
// #cgo !user_libs,linux,amd64 LDFLAGS: -L${SRCDIR}/libs/linux/amd64 -lyara -lm -lmagic -ljansson -lz -lbz2 -llzma -lcrypto
// #cgo !user_libs,linux,amd64 CFLAGS: -I${SRCDIR}/libs/linux/amd64/include
// #cgo user_libs LDFLAGS: -lyara
/*
#include <yara.h>
#if YR_VERSION_HEX < 0x040100
#error YARA version 4.1 required
#endif
*/
import "C"
