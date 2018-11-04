// Copyright © 2018 The IPFN Developers. All Rights Reserved.
// Copyright © 2016-2018 IBM Corp. All Rights Reserved.
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

package sw

import (
	"errors"
	"reflect"
	"testing"

	mocks2 "github.com/ipfn/ipfn/src/go/crypto/bccsp/mocks"
	"github.com/ipfn/ipfn/src/go/crypto/bccsp/sw/mocks"
	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	t.Parallel()

	expectedKey := &mocks2.MockKey{}
	expectedPlaintext := []byte{1, 2, 3, 4}
	expectedOpts := &mocks2.EncrypterOpts{}
	expectedCiphertext := []byte{0, 1, 2, 3, 4}
	expectedErr := errors.New("no error")

	encryptors := make(map[reflect.Type]Encryptor)
	encryptors[reflect.TypeOf(&mocks2.MockKey{})] = &mocks.Encryptor{
		KeyArg:       expectedKey,
		PlaintextArg: expectedPlaintext,
		OptsArg:      expectedOpts,
		EncValue:     expectedCiphertext,
		EncErr:       expectedErr,
	}

	csp := CSP{encryptors: encryptors}

	ct, err := csp.Encrypt(expectedKey, expectedPlaintext, expectedOpts)
	assert.Equal(t, expectedCiphertext, ct)
	assert.Equal(t, expectedErr, err)
}