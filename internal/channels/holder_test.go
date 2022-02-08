/*
 * Copyright 2022-present Open Networking Foundation
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */

package channels

import (
	"testing"
)

func TestAddChannel(t *testing.T) {
	topic := "TestChannel"

	Init()
	AddChannel(topic)
	chanRef := GetChannel(topic)
	if chanRef == nil {
		t.Error("failed to create and retrive channel ref")
	}
}
