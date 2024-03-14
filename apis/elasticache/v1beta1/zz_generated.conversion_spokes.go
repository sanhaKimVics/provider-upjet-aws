// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	ujconversion "github.com/crossplane/upjet/pkg/controller/conversion"
	"github.com/crossplane/upjet/pkg/resource"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo converts this ReplicationGroup to the hub type.
func (tr *ReplicationGroup) ConvertTo(dstRaw conversion.Hub) error {
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", tr.GetObjectKind().GroupVersionKind().Version, dstRaw.GetObjectKind().GroupVersionKind().Version)
	}
	return nil
}

// ConvertFrom converts from the hub type to the ReplicationGroup type.
func (tr *ReplicationGroup) ConvertFrom(srcRaw conversion.Hub) error {
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", srcRaw.GetObjectKind().GroupVersionKind().Version, tr.GetObjectKind().GroupVersionKind().Version)
	}
	return nil
}
