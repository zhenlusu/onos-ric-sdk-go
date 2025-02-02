// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package encoding

// Type encoding type (asn.1 or protobuf)
type Type int32

const (
	// ASN1 :
	ASN1 Type = iota

	// PROTO :
	PROTO
)

// String :
func (t Type) String() string {
	return [...]string{"ASN.1", "PROTO"}[t]
}
