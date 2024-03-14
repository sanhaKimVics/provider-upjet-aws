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

// ConvertTo converts this HoursOfOperation to the hub type.
func (tr *HoursOfOperation) ConvertTo(dstRaw conversion.Hub) error {
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", tr.GetObjectKind().GroupVersionKind().Version, dstRaw.GetObjectKind().GroupVersionKind().Version)
	}
	return nil
}

// ConvertFrom converts from the hub type to the HoursOfOperation type.
func (tr *HoursOfOperation) ConvertFrom(srcRaw conversion.Hub) error {
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", srcRaw.GetObjectKind().GroupVersionKind().Version, tr.GetObjectKind().GroupVersionKind().Version)
	}
	return nil
}

// ConvertTo converts this Queue to the hub type.
func (tr *Queue) ConvertTo(dstRaw conversion.Hub) error {
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", tr.GetObjectKind().GroupVersionKind().Version, dstRaw.GetObjectKind().GroupVersionKind().Version)
	}
	return nil
}

// ConvertFrom converts from the hub type to the Queue type.
func (tr *Queue) ConvertFrom(srcRaw conversion.Hub) error {
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", srcRaw.GetObjectKind().GroupVersionKind().Version, tr.GetObjectKind().GroupVersionKind().Version)
	}
	return nil
}

// ConvertTo converts this RoutingProfile to the hub type.
func (tr *RoutingProfile) ConvertTo(dstRaw conversion.Hub) error {
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", tr.GetObjectKind().GroupVersionKind().Version, dstRaw.GetObjectKind().GroupVersionKind().Version)
	}
	return nil
}

// ConvertFrom converts from the hub type to the RoutingProfile type.
func (tr *RoutingProfile) ConvertFrom(srcRaw conversion.Hub) error {
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", srcRaw.GetObjectKind().GroupVersionKind().Version, tr.GetObjectKind().GroupVersionKind().Version)
	}
	return nil
}
