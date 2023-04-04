// Copyright (c) 2018, Google LLC All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tpm2

import (
	"bytes"
	"testing"

	. "github.com/jclab-joseph/go-tpm/tpm2"
)

func TestKDFa(t *testing.T) {
	tcs := []struct {
		hashAlg  Algorithm
		key      []byte
		contextU []byte
		contextV []byte
		label    string
		bits     int
		expected []byte
	}{
		{
			hashAlg:  AlgSHA256,
			key:      []byte{'y', 'o', 'l', 'o', 0},
			contextU: []byte{'k', 'e', 'k', 0},
			contextV: []byte{'y', 'o', 'y', 'o', 0},
			label:    "IDENTITY",
			bits:     128,
			expected: []byte{0xd2, 0xd7, 0x2c, 0xc7, 0xa8, 0xa5, 0xeb, 0x09, 0xe8, 0xc7, 0x90, 0x12, 0xe2, 0xda, 0x9f, 0x22},
		},
		{
			hashAlg:  AlgSHA256,
			key:      []byte{'c', 'a', 0},
			contextU: []byte{'a', 'b', 'c', 0},
			label:    "IDENTITY",
			bits:     1024,
			expected: []byte{0x1a, 0xae, 0x71, 0x51, 0xac, 0x1a, 0x56, 0x90, 0xed, 0xa7, 0xdc, 0xab, 0xd5, 0x68, 0x00, 0xc1, 0x1c, 0x56, 0xa3, 0x81, 0x0b, 0xa0, 0x59, 0x82, 0x6f, 0xe4, 0x77, 0x63, 0x48, 0xd6, 0xae, 0x8e, 0x5d, 0x5d, 0x18, 0xc7, 0xcc, 0xf8, 0x37, 0x3f, 0x7b, 0x94, 0x2a, 0xda, 0x8b, 0x91, 0x2b, 0x12, 0xda, 0x56, 0xfb, 0x37, 0xf6, 0x4b, 0x93, 0x58, 0x72, 0x84, 0x1e, 0xc0, 0x7d, 0x38, 0xe1, 0xfb, 0x8e, 0x7e, 0xc8, 0x6e, 0xfc, 0xbf, 0xb4, 0x44, 0x75, 0x6b, 0xc8, 0x86, 0x3f, 0x85, 0x8d, 0x26, 0x90, 0xa6, 0x21, 0xc9, 0xaf, 0xb9, 0x83, 0xcd, 0x77, 0xe7, 0xa1, 0x04, 0x8a, 0xe1, 0xa7, 0x59, 0x8a, 0xc8, 0x95, 0x32, 0x3d, 0x44, 0xc1, 0x02, 0x27, 0xaf, 0x0a, 0x00, 0x14, 0x4c, 0xab, 0x55, 0x11, 0x10, 0x75, 0xdc, 0x6b, 0x72, 0xad, 0x6e, 0xb1, 0x63, 0xc7, 0x45, 0x8b, 0x87, 0x8e, 0x8c},
		},
		{
			hashAlg:  AlgSHA1,
			key:      []byte{'c', 'a', 0},
			contextU: []byte{'a', 'b', 'c', 0},
			label:    "IDENTITY",
			bits:     256,
			expected: []byte{0x83, 0xf3, 0x54, 0xaf, 0xcf, 0x92, 0x3d, 0xe2, 0x11, 0x2e, 0x08, 0x91, 0x43, 0x4c, 0xd0, 0xbd, 0xc8, 0xac, 0xbf, 0x01, 0xb8, 0x11, 0xc0, 0xe8, 0xcd, 0x06, 0x2d, 0xed, 0x39, 0xe3, 0x1f, 0x7f},
		},
		{
			// Based on an example on Page 44. Note that the upper 7 bits of the zeroth byte are clear.
			hashAlg:  AlgSHA1,
			key:      []byte{0x27, 0x1f, 0xa0, 0x8b, 0xbd, 0xc5, 0x6, 0xe, 0xc3, 0xdf, 0xa9, 0x28, 0xff, 0x9b, 0x73, 0x12, 0x3a, 0x12, 0xda, 0xc},
			contextU: []byte{0xce, 0x24, 0x4f, 0x39, 0x5d, 0xca, 0x73, 0x91},
			contextV: []byte{0xda, 0x50, 0x40, 0x31, 0xdd, 0xf1, 0x2e, 0x83},
			label:    "KDFSELFTESTLABEL",
			bits:     521,
			expected: []byte{0x1, 0x71, 0x4e, 0x86, 0x50, 0x72, 0xc1, 0xe4, 0xc4, 0xb, 0x36, 0x70, 0x28, 0x11, 0x12, 0x63, 0x2f, 0xc9, 0xec, 0xba, 0x25, 0xe6, 0x6a, 0xf6, 0x8d, 0x18, 0xa2, 0xb6, 0x2d, 0xe9, 0xcb, 0xb4, 0x45, 0x21, 0xda, 0x2b, 0xc9, 0xa4, 0x96, 0x86, 0x2e, 0xb3, 0xf3, 0x6, 0x94, 0x9f, 0x9e, 0xfe, 0x8a, 0x9e, 0x1c, 0xcb, 0xce, 0x3b, 0x4d, 0x66, 0x8f, 0xfd, 0x75, 0xc9, 0x39, 0x4b, 0xa5, 0x94, 0x58, 0xfe},
		},
	}

	for _, tc := range tcs {
		o, err := KDFa(tc.hashAlg, tc.key, tc.label, tc.contextU, tc.contextV, tc.bits)
		if err != nil {
			t.Fatalf("KDFa(%v, %v, %q, %v, %v, %v) returned error: %v", tc.hashAlg, tc.key, tc.label, tc.contextU, tc.contextV, tc.bits, err)
		}
		if !bytes.Equal(tc.expected, o) {
			t.Errorf("Test with KDFa(%v, %v, %q, %v, %v, %v) returned incorrect result", tc.hashAlg, tc.key, tc.label, tc.contextU, tc.contextV, tc.bits)
			t.Logf("  Got:  %v", o)
			t.Logf("  Want: %v", tc.expected)
		}
	}
}

func TestKDFe(t *testing.T) {
	label := "DUPLICATE"

	// Test vectors taken from running the microsoft TPM simulator, ms-tpm-20-ref.
	tests := []struct {
		name       string
		bits       int
		hashAlg    Algorithm
		z          []byte
		partyUInfo []byte
		partyVInfo []byte
		expected   []byte
	}{
		{
			name:       "NIST-P224-SHA256",
			bits:       256,
			hashAlg:    AlgSHA256,
			z:          []byte{0x6e, 0xd1, 0x5c, 0x60, 0xfd, 0x43, 0x3f, 0x5d, 0xdb, 0x28, 0xd, 0x7b, 0xe4, 0x3f, 0x8a, 0xc5, 0xa4, 0x52, 0x4c, 0x13, 0xb9, 0x2f, 0xf2, 0x93, 0x61, 0x94, 0x29, 0xef},
			partyUInfo: []byte{0xaf, 0x6, 0xe1, 0xa4, 0x22, 0xed, 0xbe, 0x6f, 0x41, 0xe1, 0xf8, 0xb3, 0xce, 0xa, 0xf2, 0x1f, 0xc8, 0xb1, 0x1, 0x3c, 0x1f, 0xc8, 0xd5, 0x50, 0xcc, 0xae, 0xe6, 0x6d},
			partyVInfo: []byte{0xa0, 0x2e, 0x47, 0x5e, 0xc7, 0x53, 0x44, 0x4d, 0x1b, 0xc1, 0xad, 0x10, 0xbc, 0xa3, 0xa7, 0xda, 0x72, 0xee, 0x65, 0x29, 0x7b, 0x4, 0xd5, 0xf4, 0x2a, 0xa8, 0x81, 0x2c},
			expected:   []byte{0x33, 0xa1, 0x99, 0xf3, 0x24, 0xbe, 0x56, 0x22, 0x49, 0x4e, 0x7c, 0x79, 0xcb, 0xc, 0x8c, 0x84, 0x22, 0x73, 0xd7, 0x68, 0x8d, 0x3a, 0x64, 0xda, 0x97, 0xfb, 0x48, 0xea, 0xea, 0x44, 0xf0, 0xa3},
		},
		{
			name:       "NIST-P224-SHA1",
			bits:       160,
			hashAlg:    AlgSHA1,
			z:          []byte{0x15, 0xf6, 0xca, 0x83, 0x6a, 0x67, 0x92, 0xec, 0x4a, 0xb2, 0xc9, 0x89, 0x84, 0xd9, 0x8c, 0xd4, 0xb8, 0xf, 0xe9, 0x6c, 0xf, 0x40, 0x26, 0xd5, 0xe7, 0x5c, 0x6f, 0x4f},
			partyUInfo: []byte{0x12, 0x2f, 0xdd, 0x75, 0x5, 0xe1, 0xde, 0xd5, 0x81, 0x1e, 0xac, 0x38, 0x35, 0x11, 0x6c, 0x9c, 0x1c, 0xd5, 0xbf, 0xc6, 0x85, 0x74, 0x84, 0xba, 0x31, 0x53, 0xf, 0xb8},
			partyVInfo: []byte{0x58, 0x27, 0xb4, 0x21, 0x43, 0x0, 0x63, 0x6e, 0xd6, 0x15, 0xe7, 0x56, 0xd, 0xd7, 0xaa, 0xc4, 0xbb, 0x94, 0xc7, 0xc2, 0x73, 0xdf, 0xf1, 0x10, 0xd2, 0xc3, 0x80, 0x8a},
			expected:   []byte{0x3c, 0x1e, 0x8e, 0x76, 0xde, 0x44, 0xaf, 0x9e, 0xfe, 0xea, 0x6e, 0xa7, 0xce, 0x6b, 0x43, 0x39, 0x63, 0x59, 0xe1, 0xc2},
		},
		{
			name:       "NIST-P256-SHA256",
			bits:       256,
			hashAlg:    AlgSHA256,
			z:          []byte{0x69, 0xfd, 0x74, 0xd0, 0x52, 0xa5, 0xcd, 0x9d, 0x6a, 0xdd, 0xf4, 0xea, 0x79, 0x1f, 0x47, 0x37, 0xeb, 0x29, 0x72, 0x46, 0xdc, 0xe6, 0x4, 0xf6, 0x38, 0x5c, 0x13, 0x92, 0xe0, 0xfb, 0x56, 0xb9},
			partyUInfo: []byte{0xc2, 0x9f, 0x8c, 0xff, 0x1c, 0xe0, 0x54, 0xed, 0x59, 0x45, 0xab, 0xd3, 0x93, 0xee, 0x2e, 0xed, 0x3f, 0x67, 0x43, 0x7d, 0x9d, 0xb0, 0x7b, 0x93, 0x21, 0x56, 0xc6, 0x5d, 0x32, 0x92, 0x13, 0x73},
			partyVInfo: []byte{0x10, 0x79, 0x37, 0xdc, 0x44, 0xe5, 0xbb, 0x50, 0x94, 0xd3, 0xd3, 0xa, 0x55, 0xef, 0xac, 0x77, 0xb6, 0xdb, 0x32, 0x53, 0xba, 0x42, 0xda, 0xbc, 0x80, 0x44, 0x46, 0xb9, 0x38, 0x64, 0xe0, 0x55},
			expected:   []byte{0xb1, 0x52, 0x67, 0xe3, 0xf9, 0x60, 0xf8, 0xf8, 0xf3, 0xb9, 0x4f, 0xf3, 0xfc, 0x64, 0x11, 0x9f, 0x68, 0x76, 0x97, 0x4f, 0x98, 0x4f, 0x82, 0xd5, 0xb, 0xd5, 0x4a, 0xbc, 0x7a, 0x64, 0x6, 0xe8},
		},
		{
			name:       "NIST-P256-SHA1",
			bits:       160,
			hashAlg:    AlgSHA1,
			z:          []byte{0x82, 0x1f, 0xd7, 0x93, 0x2a, 0x44, 0xcf, 0x38, 0x33, 0x73, 0x68, 0x5b, 0x71, 0x92, 0xd4, 0x6c, 0xe9, 0xc9, 0xc4, 0x7a, 0x4b, 0x99, 0x76, 0x1e, 0x88, 0x78, 0xbd, 0xbf, 0x25, 0xc8, 0xb8, 0x68},
			partyUInfo: []byte{0x68, 0xd1, 0x23, 0x7b, 0x42, 0xf4, 0xf8, 0x2d, 0xde, 0x83, 0x9f, 0xd2, 0x77, 0x87, 0xa, 0x4f, 0x21, 0xc0, 0x2e, 0x6b, 0x2d, 0x19, 0xe2, 0xf4, 0x25, 0x1d, 0x7b, 0x66, 0x87, 0xe5, 0x1d, 0x6a},
			partyVInfo: []byte{0xfb, 0xf1, 0xc3, 0xa7, 0x29, 0xcb, 0xd4, 0xe0, 0xbd, 0xcf, 0x8d, 0xbf, 0xc9, 0x47, 0xf4, 0xbf, 0x95, 0x88, 0xb4, 0x2, 0x9c, 0xbb, 0x3d, 0x65, 0x8f, 0xd5, 0x7e, 0xda, 0x36, 0xb1, 0x41, 0xbd},
			expected:   []byte{0x78, 0x9b, 0x33, 0x6d, 0x41, 0x19, 0x8f, 0x37, 0x33, 0xd2, 0xbe, 0x8f, 0x5c, 0xb3, 0x5c, 0x98, 0xc4, 0x2b, 0x51, 0x9e},
		},
		{
			name:       "NIST-P384-SHA256",
			bits:       256,
			hashAlg:    AlgSHA256,
			z:          []byte{0xbb, 0x3a, 0x3f, 0x26, 0xd6, 0x81, 0xb, 0xff, 0x29, 0x67, 0x63, 0x49, 0x95, 0xfd, 0x57, 0x9f, 0x80, 0x6f, 0xa6, 0x8f, 0xd, 0x9d, 0xb5, 0xbd, 0x62, 0xd1, 0xc, 0x30, 0x5c, 0xea, 0xa9, 0xda, 0x8e, 0xec, 0x67, 0x3e, 0xe0, 0x89, 0x51, 0xbc, 0x63, 0xae, 0x9c, 0x6f, 0xb5, 0xfe, 0x98, 0xb1},
			partyUInfo: []byte{0x20, 0xf0, 0xf4, 0xc, 0xf1, 0x85, 0xca, 0x7, 0xa4, 0x56, 0x99, 0xf7, 0x4a, 0x7, 0x19, 0xc6, 0x21, 0x84, 0x3e, 0x9d, 0x7, 0x1e, 0x4f, 0x6b, 0xe8, 0xc3, 0x7a, 0xbd, 0xa8, 0x38, 0xee, 0x1a, 0x2a, 0x9, 0x82, 0x4d, 0x22, 0x6b, 0x36, 0xa7, 0xae, 0x50, 0x18, 0xaa, 0x9, 0x97, 0xd6, 0x17},
			partyVInfo: []byte{0xcb, 0x67, 0x7d, 0x12, 0x74, 0x25, 0x9f, 0x55, 0x2a, 0xa4, 0x52, 0x60, 0xd9, 0x65, 0x46, 0xc5, 0x3f, 0xe1, 0xa0, 0x4a, 0xf7, 0x8a, 0xc, 0xc, 0xbb, 0xb4, 0xf4, 0x3b, 0x32, 0x71, 0x63, 0xc5, 0xd, 0x76, 0x8, 0xbc, 0x2c, 0xf6, 0x4c, 0x1f, 0x2c, 0xbe, 0x79, 0x9a, 0xaa, 0x9f, 0xbe, 0x6d},
			expected:   []byte{0x94, 0x54, 0xb7, 0x2e, 0xb5, 0xf, 0xb, 0x25, 0x66, 0x24, 0x69, 0x76, 0xa7, 0x44, 0x5d, 0xa1, 0xa1, 0xeb, 0xfa, 0x82, 0xc9, 0x14, 0xc4, 0x65, 0xac, 0x75, 0x4e, 0x94, 0x1d, 0xe0, 0xfa, 0xaa},
		},
		{
			name:       "NIST-P384-SHA1",
			bits:       160,
			hashAlg:    AlgSHA1,
			z:          []byte{0x6a, 0x21, 0xc6, 0x9, 0x50, 0x1e, 0x63, 0x6e, 0x65, 0x20, 0x49, 0xd1, 0xac, 0xc9, 0xd9, 0xea, 0x23, 0xc2, 0x7f, 0xf8, 0xc5, 0x8e, 0xf4, 0xe6, 0xb2, 0x91, 0x41, 0xa7, 0x63, 0xda, 0xf7, 0x94, 0x40, 0x16, 0xcd, 0x4, 0x4f, 0xa5, 0x8b, 0xf8, 0x5a, 0x59, 0x82, 0xfc, 0x11, 0x8d, 0xb3, 0x9e},
			partyUInfo: []byte{0x9f, 0xc, 0x9f, 0xcf, 0xd8, 0xd5, 0xa6, 0x21, 0x88, 0x78, 0xe7, 0xfc, 0xb5, 0x13, 0x9b, 0xd5, 0xe8, 0xc5, 0x8d, 0xcb, 0x9e, 0x5f, 0x16, 0xd2, 0xa7, 0x6b, 0xe5, 0xc5, 0x6f, 0x9b, 0x1e, 0x13, 0x6e, 0x99, 0x78, 0x4b, 0xe3, 0xf1, 0x98, 0x1f, 0xe, 0x74, 0xa5, 0xc9, 0x5d, 0x1, 0xce, 0x88},
			partyVInfo: []byte{0x2, 0x31, 0x9a, 0x6a, 0xb9, 0xc7, 0x3e, 0x59, 0xd9, 0x93, 0x3e, 0x3f, 0x70, 0x2b, 0xa4, 0x33, 0xea, 0x8a, 0xc8, 0x55, 0x52, 0xbe, 0x9, 0x85, 0x5e, 0x7a, 0xe2, 0xe1, 0x4e, 0xce, 0x3a, 0x97, 0xf7, 0xe9, 0xd4, 0x8a, 0x93, 0xd5, 0x25, 0x7e, 0x62, 0x76, 0x73, 0x79, 0x85, 0xa1, 0x39, 0xa5},
			expected:   []byte{0x2, 0x7e, 0x4b, 0xb0, 0x36, 0x25, 0xd6, 0x19, 0x2, 0xe8, 0x6d, 0x80, 0xf7, 0xe2, 0xe2, 0x8d, 0xfa, 0x8f, 0x1a, 0xf2},
		},
		{
			name:       "NIST-P521-SHA256",
			bits:       256,
			hashAlg:    AlgSHA256,
			z:          []byte{0x1, 0xaf, 0x5f, 0x2a, 0x95, 0x8d, 0xee, 0x7f, 0x66, 0x17, 0xa5, 0x52, 0x2c, 0xaa, 0x91, 0xe4, 0x72, 0x21, 0xe2, 0xf2, 0xf6, 0x39, 0xb, 0xec, 0x39, 0x5d, 0xc6, 0xf0, 0x34, 0x86, 0x48, 0xc9, 0x73, 0xf1, 0x99, 0x68, 0xe4, 0x8f, 0x1d, 0x89, 0x43, 0xc7, 0xf5, 0xcd, 0x6a, 0x92, 0x88, 0xe9, 0x78, 0x3b, 0x73, 0x3a, 0xd2, 0x44, 0x3b, 0xab, 0x63, 0xef, 0x28, 0x5d, 0x26, 0x10, 0x21, 0x7a, 0x8, 0xeb},
			partyUInfo: []byte{0xf, 0xcb, 0xe5, 0xe4, 0xca, 0xf1, 0xe1, 0x1f, 0x46, 0xfe, 0xeb, 0x72, 0x62, 0xd4, 0xb4, 0x27, 0x6c, 0x6, 0x50, 0x9c, 0x96, 0xc7, 0x43, 0xdb, 0x23, 0x8e, 0x48, 0xa0, 0x90, 0xf1, 0x24, 0x3f, 0x6d, 0x33, 0xbe, 0xff, 0x9, 0x40, 0xf0, 0xa6, 0xfb, 0xcc, 0x61, 0x63, 0x72, 0x92, 0x3a, 0x3a, 0xde, 0x83, 0xe5, 0xa8, 0xa6, 0x6c, 0x37, 0xb5, 0x53, 0x0, 0x4, 0x2f, 0x10, 0xcb, 0x39, 0x2, 0x13},
			partyVInfo: []byte{0x0, 0xee, 0xd7, 0x6b, 0x36, 0x5c, 0xba, 0xeb, 0xdd, 0xf5, 0x5a, 0xd2, 0xa0, 0xb6, 0xcb, 0x95, 0xb2, 0x85, 0xce, 0xd2, 0x64, 0xe2, 0x39, 0x5, 0x74, 0x7b, 0x8d, 0x4, 0x1a, 0xa9, 0x68, 0xbd, 0x66, 0xec, 0x69, 0x3d, 0xd9, 0xf8, 0x7f, 0xd4, 0xad, 0x43, 0x3a, 0x23, 0xe8, 0xcc, 0xa6, 0x4d, 0xa1, 0x65, 0xd, 0xac, 0xf0, 0x84, 0x7a, 0x4d, 0x4a, 0x8a, 0xf4, 0x2e, 0x73, 0xb2, 0x37, 0x85, 0x67, 0x8b},
			expected:   []byte{0xbc, 0x34, 0x57, 0x10, 0xa2, 0x4, 0x54, 0xf, 0xc5, 0x9a, 0xe8, 0x5e, 0x23, 0x7, 0x1d, 0x0, 0xfe, 0xfb, 0x28, 0xc, 0xe6, 0xa1, 0xd0, 0xbb, 0x52, 0x73, 0x9f, 0x8c, 0xf9, 0x5, 0xda, 0x30},
		},
		{
			name:       "NIST-P521-SHA1",
			bits:       160,
			hashAlg:    AlgSHA1,
			z:          []byte{0x1, 0x3d, 0x67, 0x71, 0x7e, 0x61, 0xf0, 0x2b, 0xdd, 0x24, 0xb8, 0xd7, 0x2b, 0x45, 0x5, 0x89, 0xad, 0x52, 0x69, 0x40, 0xc, 0x86, 0x98, 0x90, 0xa2, 0xf4, 0xaf, 0xad, 0xda, 0xb8, 0x4d, 0x95, 0x84, 0x38, 0xc9, 0x24, 0x80, 0x16, 0x88, 0xec, 0x9, 0x85, 0x2e, 0x47, 0xf9, 0x25, 0xb5, 0xe3, 0xf, 0xc1, 0x78, 0xaa, 0xa5, 0x84, 0x83, 0x9a, 0x1d, 0x4c, 0x15, 0x97, 0x3f, 0xec, 0x59, 0x24, 0xe1, 0xb},
			partyUInfo: []byte{0x0, 0x8a, 0xe6, 0xa1, 0xad, 0x3e, 0xef, 0xa0, 0xaf, 0x49, 0x41, 0x85, 0x33, 0xe9, 0x18, 0x8a, 0xa4, 0x36, 0x95, 0x4c, 0xb9, 0x41, 0x42, 0xd6, 0xe4, 0x66, 0xa1, 0x48, 0xb1, 0x65, 0x8d, 0xdc, 0xba, 0x9c, 0x25, 0x72, 0x3, 0xf6, 0xde, 0x1c, 0x4, 0x66, 0x7d, 0x3b, 0x7d, 0x5a, 0x8c, 0x82, 0x8d, 0x1f, 0x9f, 0x46, 0xc4, 0x0, 0x99, 0xef, 0xc3, 0xa1, 0xca, 0xe, 0x98, 0xa0, 0x23, 0x8b, 0x6b, 0x5e},
			partyVInfo: []byte{0x1, 0x56, 0xf, 0x5e, 0xcf, 0xef, 0xb4, 0x4, 0x94, 0x32, 0xb1, 0x7b, 0x24, 0x86, 0x92, 0x41, 0xba, 0x2d, 0xfc, 0xb2, 0x3a, 0xce, 0x96, 0xee, 0x52, 0x23, 0x51, 0xcf, 0x4a, 0x3e, 0x1d, 0xf4, 0xb8, 0x27, 0xe6, 0xa7, 0xb5, 0xb9, 0xce, 0x37, 0x4d, 0xa, 0xc4, 0xb9, 0x20, 0x2d, 0x3f, 0xa, 0x14, 0x47, 0x7f, 0x51, 0x29, 0x2c, 0x3e, 0x41, 0x7a, 0x1, 0x11, 0x23, 0x9c, 0x8, 0x23, 0xfe, 0xef, 0xa3},
			expected:   []byte{0xee, 0x5, 0xca, 0xe6, 0x4c, 0xa5, 0x6, 0xfe, 0x89, 0x51, 0x7d, 0x7b, 0x8e, 0xf9, 0x73, 0xf1, 0xb, 0x39, 0xd4, 0x5c},
		},
		{
			name:       "SHORTER-THAN-HASH",
			bits:       100,
			hashAlg:    AlgSHA256,
			z:          []byte{0x9e, 0xc0, 0x69, 0x1d, 0x3c, 0x5d, 0x35, 0xfb, 0xb2, 0xdc, 0xf0, 0x82, 0xb0, 0xbb, 0x4d, 0x1d, 0xf0, 0x7e, 0xe0, 0xc5, 0xbf, 0x27, 0x25, 0x1f, 0xff, 0xa9, 0x73, 0xee, 0x33, 0x9e, 0x5e, 0x62},
			partyUInfo: []byte{0x68, 0xd6, 0x2d, 0x34, 0x49, 0xe0, 0xdb, 0x8c, 0x1a, 0xf5, 0x7a, 0xdb, 0xa1, 0x53, 0x43, 0xbc, 0x34, 0xf2, 0xa6, 0xe7, 0x6a, 0x97, 0x50, 0xf4, 0x76, 0xe6, 0x18, 0x62, 0xdb, 0x8f, 0xb, 0xec},
			partyVInfo: []byte{0xaa, 0xb7, 0xca, 0xb3, 0x61, 0xd8, 0xf3, 0x1d, 0xf, 0x6a, 0x98, 0xcc, 0x3c, 0x11, 0xbb, 0xe9, 0x98, 0x3b, 0xf9, 0x1f, 0xc4, 0xc5, 0x3e, 0x90, 0xd5, 0xcb, 0x10, 0xeb, 0x74, 0xfd, 0x5c, 0x3a},
			expected:   []byte{0x5, 0x0, 0x54, 0xd3, 0xd7, 0x78, 0x5, 0x91, 0x20, 0xaf, 0x6d, 0x3, 0x13},
		},
		{
			name:       "LONGER-THAN-HASH",
			bits:       1600,
			hashAlg:    AlgSHA256,
			z:          []byte{0x1e, 0xb2, 0xa, 0x75, 0x56, 0xd7, 0xb7, 0xc9, 0x1d, 0x7d, 0x69, 0x74, 0xc9, 0x4a, 0x9b, 0xfb, 0xfa, 0xf1, 0x25, 0x63, 0xcb, 0xde, 0x93, 0xb4, 0x5, 0x6b, 0xb9, 0xb6, 0x2b, 0x2b, 0x23, 0x82},
			partyUInfo: []byte{0x57, 0xea, 0xfe, 0xe9, 0xcb, 0x12, 0xae, 0xb7, 0x43, 0x7d, 0xbc, 0xe8, 0x40, 0xaa, 0x9d, 0xbd, 0x59, 0x6f, 0x3b, 0x10, 0x32, 0xa, 0xb0, 0x65, 0x75, 0xb2, 0x37, 0xf3, 0x38, 0xfa, 0xe6, 0xc2},
			partyVInfo: []byte{0x30, 0x99, 0xaa, 0x56, 0xbc, 0x2b, 0xe9, 0xbc, 0xca, 0x32, 0x4f, 0x3, 0x10, 0x3a, 0x56, 0xac, 0x3, 0x31, 0x4d, 0x10, 0x5f, 0x4d, 0xbc, 0x90, 0x32, 0xfe, 0x87, 0x9f, 0xe9, 0xc0, 0xdf, 0x68},
			expected:   []byte{0xe0, 0xd2, 0x56, 0xf0, 0x5f, 0xce, 0xda, 0x6c, 0x5f, 0xe6, 0xb0, 0x9f, 0xb5, 0xaa, 0xbf, 0xab, 0xa5, 0xe4, 0xca, 0x7c, 0x43, 0x12, 0xec, 0xbe, 0x89, 0xf, 0x7c, 0x57, 0x47, 0xf5, 0xca, 0xad, 0xa0, 0x4e, 0xbc, 0x13, 0xff, 0x15, 0x7f, 0xdb, 0x73, 0x76, 0xdc, 0xdc, 0xf5, 0x61, 0x59, 0x35, 0xd, 0x7c, 0xfb, 0x1c, 0x2, 0xba, 0xe0, 0x18, 0x28, 0x7c, 0xc3, 0x4b, 0x67, 0xf3, 0x2b, 0xf0, 0xd8, 0x9d, 0x7f, 0x36, 0xe8, 0x3f, 0x5b, 0xcf, 0x76, 0xe7, 0x2, 0x34, 0xac, 0xda, 0x4a, 0xe5, 0x9d, 0xa6, 0x1, 0x93, 0x28, 0x17, 0x80, 0xa, 0xf3, 0x4b, 0xd7, 0x54, 0x36, 0xc6, 0x59, 0x1e, 0xbd, 0xbb, 0x97, 0x88, 0x66, 0x4, 0x14, 0x4, 0x4b, 0xe, 0x26, 0xf0, 0x6d, 0xeb, 0x8a, 0x34, 0xca, 0xa4, 0xe0, 0xa4, 0x90, 0xae, 0x3, 0xdd, 0x11, 0x80, 0xac, 0x17, 0x50, 0xa0, 0x1b, 0x8d, 0xeb, 0x7a, 0x4a, 0x79, 0x96, 0x91, 0x64, 0x17, 0xc4, 0x21, 0xe0, 0xf6, 0x5b, 0x57, 0x1, 0xb5, 0xec, 0x14, 0xd7, 0xb1, 0x19, 0x5d, 0x6e, 0xc5, 0x33, 0x7f, 0x3, 0xaa, 0x41, 0x9d, 0x72, 0x7d, 0x20, 0xa5, 0x75, 0xde, 0xfd, 0xe2, 0x2e, 0xe4, 0x54, 0x2b, 0xbf, 0x9e, 0xad, 0x57, 0xd0, 0x3e, 0x2e, 0x76, 0x28, 0xd8, 0x58, 0x80, 0xaf, 0x53, 0x1b, 0x2d, 0x3, 0xc3, 0xd7, 0xe7, 0x57, 0x9c, 0xcf, 0xf6, 0xf9, 0x7e, 0x7f, 0xa8, 0x2e, 0x54},
		},
		{
			// Based on an example on Page 44. Note that the upper 7 bits of the zeroth byte are clear.
			name:       "NOT-MULTIPLE",
			bits:       521,
			hashAlg:    AlgSHA256,
			z:          []byte{0x94, 0xb2, 0x61, 0x44, 0xbf, 0xa0, 0x6f, 0x0, 0x94, 0x10, 0x67, 0x54, 0xb7, 0x9c, 0x24, 0xc3, 0xb9, 0xcb, 0x20, 0x1, 0xf2, 0xba, 0xd, 0x28, 0x42, 0x51, 0xc1, 0x70, 0xdc, 0xae, 0x2f, 0x34},
			partyUInfo: []byte{0x4d, 0x2f, 0xe0, 0x86, 0x78, 0x22, 0xb1, 0x5d, 0xbe, 0x6a, 0x14, 0x7c, 0xd5, 0x19, 0xf7, 0x68, 0x41, 0x25, 0x6e, 0xb, 0xe3, 0x2b, 0x2d, 0x53, 0x13, 0x31, 0xc8, 0xaa, 0xc2, 0xf9, 0x71, 0x25},
			partyVInfo: []byte{0xbc, 0xca, 0x60, 0x47, 0xb4, 0xe0, 0x7d, 0x5f, 0x92, 0x88, 0x6a, 0xc3, 0x6a, 0x3d, 0x48, 0xaf, 0x5d, 0x6f, 0x34, 0x10, 0xc7, 0x7c, 0x7d, 0x4f, 0xa6, 0xe0, 0xd3, 0x4d, 0x6a, 0x8f, 0xbd, 0xdc},
			expected:   []byte{0x1, 0x24, 0xfc, 0x70, 0xf4, 0xf0, 0x26, 0xcd, 0x92, 0x45, 0x14, 0x2a, 0xdc, 0xa8, 0x3a, 0x0, 0x77, 0x32, 0x16, 0x3, 0x8, 0x92, 0xaf, 0x57, 0x92, 0xae, 0x48, 0xf5, 0xd0, 0x40, 0x74, 0x6d, 0x64, 0xe0, 0x68, 0x44, 0x2d, 0x49, 0x9, 0x99, 0x29, 0xfa, 0x2d, 0x8a, 0x4d, 0xd5, 0x6f, 0x46, 0xbf, 0x48, 0x3d, 0x1f, 0xc6, 0x58, 0x90, 0x3f, 0xd6, 0x2d, 0xc, 0x32, 0x6, 0x4b, 0xca, 0xa1, 0xa1, 0x3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := KDFe(test.hashAlg, test.z, label, test.partyUInfo, test.partyVInfo, test.bits)
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(result, test.expected) {
				t.Errorf("got: %v\n expected: %v", result, test.expected)
			}
		})
	}
}
