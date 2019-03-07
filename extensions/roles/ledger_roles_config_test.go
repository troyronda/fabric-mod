/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package roles

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsCommitter(t *testing.T) {
	require := require.New(t)
	require.True(IsCommitter())
}

func TestIsEndorser(t *testing.T) {
	require := require.New(t)
	require.True(IsEndorser())
}

func TestIsValidator(t *testing.T) {
	require := require.New(t)
	require.True(IsValidator())
}

func TestRolesAsString(t *testing.T) {
	require := require.New(t)
	require.Empty(RolesAsString())
}
