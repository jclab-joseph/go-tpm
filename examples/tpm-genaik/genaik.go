//go:build !windows

// Copyright (c) 2014, Google LLC All rights reserved.
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

package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"os"

	"github.com/jclab-joseph/go-tpm/tpm"
)

var (
	ownerAuthEnvVar = "TPM_OWNER_AUTH"
	srkAuthEnvVar   = "TPM_SRK_AUTH"
	aikAuthEnvVar   = "TPM_AIK_AUTH"
)

func main() {
	var blobname = flag.String("blob", "aikblob", "The name of the file to create")
	var tpmname = flag.String("tpm", "/dev/tpm0", "The path to the TPM device to use")
	flag.Parse()

	rwc, err := tpm.OpenTPM(*tpmname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't open the TPM file %s: %s\n", *tpmname, err)
		return
	}

	// Compute the auth values as needed.
	var ownerAuth [20]byte
	ownerInput := os.Getenv(ownerAuthEnvVar)
	if ownerInput != "" {
		oa := sha1.Sum([]byte(ownerInput))
		copy(ownerAuth[:], oa[:])
	}

	var srkAuth [20]byte
	srkInput := os.Getenv(srkAuthEnvVar)
	if srkInput != "" {
		sa := sha1.Sum([]byte(srkInput))
		copy(srkAuth[:], sa[:])
	}

	var aikAuth [20]byte
	aikInput := os.Getenv(aikAuthEnvVar)
	if aikInput != "" {
		aa := sha1.Sum([]byte(aikInput))
		copy(aikAuth[:], aa[:])
	}

	// TODO(tmroeder): add support for Privacy CAs.
	blob, err := tpm.MakeIdentity(rwc, srkAuth[:], ownerAuth[:], aikAuth[:], nil, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't make an new AIK: %s\n", err)
		return
	}

	if err := os.WriteFile(*blobname, blob, 0600); err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't write to file %s: %s\n", *blobname, err)
		return
	}

	return
}
