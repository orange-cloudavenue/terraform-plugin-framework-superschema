/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package superschema

import (
	"testing"

	schemaD "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	schemaR "github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// TestComputeIsRequired tests the boolean OR logic for Required across Common and target.
func TestComputeIsRequired(t *testing.T) {
	tests := []struct {
		name     string
		common   schemaR.StringAttribute
		target   schemaR.StringAttribute
		expected bool
	}{
		{
			name:     "both false",
			common:   schemaR.StringAttribute{Required: false},
			target:   schemaR.StringAttribute{Required: false},
			expected: false,
		},
		{
			name:     "common true, target false",
			common:   schemaR.StringAttribute{Required: true},
			target:   schemaR.StringAttribute{Required: false},
			expected: true,
		},
		{
			name:     "common false, target true",
			common:   schemaR.StringAttribute{Required: false},
			target:   schemaR.StringAttribute{Required: true},
			expected: true,
		},
		{
			name:     "both true",
			common:   schemaR.StringAttribute{Required: true},
			target:   schemaR.StringAttribute{Required: true},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := computeIsRequired(tt.common, tt.target)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestComputeIsOptional tests the boolean OR logic for Optional.
func TestComputeIsOptional(t *testing.T) {
	tests := []struct {
		name     string
		common   schemaR.StringAttribute
		target   schemaR.StringAttribute
		expected bool
	}{
		{
			name:     "both false",
			common:   schemaR.StringAttribute{Optional: false},
			target:   schemaR.StringAttribute{Optional: false},
			expected: false,
		},
		{
			name:     "common true, target false",
			common:   schemaR.StringAttribute{Optional: true},
			target:   schemaR.StringAttribute{Optional: false},
			expected: true,
		},
		{
			name:     "common false, target true",
			common:   schemaR.StringAttribute{Optional: false},
			target:   schemaR.StringAttribute{Optional: true},
			expected: true,
		},
		{
			name:     "both true",
			common:   schemaR.StringAttribute{Optional: true},
			target:   schemaR.StringAttribute{Optional: true},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := computeIsOptional(tt.common, tt.target)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestComputeIsComputed tests the boolean OR logic for Computed.
func TestComputeIsComputed(t *testing.T) {
	tests := []struct {
		name     string
		common   schemaR.StringAttribute
		target   schemaR.StringAttribute
		expected bool
	}{
		{
			name:     "both false",
			common:   schemaR.StringAttribute{Computed: false},
			target:   schemaR.StringAttribute{Computed: false},
			expected: false,
		},
		{
			name:     "common true, target false",
			common:   schemaR.StringAttribute{Computed: true},
			target:   schemaR.StringAttribute{Computed: false},
			expected: true,
		},
		{
			name:     "common false, target true",
			common:   schemaR.StringAttribute{Computed: false},
			target:   schemaR.StringAttribute{Computed: true},
			expected: true,
		},
		{
			name:     "both true",
			common:   schemaR.StringAttribute{Computed: true},
			target:   schemaR.StringAttribute{Computed: true},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := computeIsComputed(tt.common, tt.target)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestComputeIsSensitive tests the boolean OR logic for Sensitive.
func TestComputeIsSensitive(t *testing.T) {
	tests := []struct {
		name     string
		common   schemaR.StringAttribute
		target   schemaR.StringAttribute
		expected bool
	}{
		{
			name:     "both false",
			common:   schemaR.StringAttribute{Sensitive: false},
			target:   schemaR.StringAttribute{Sensitive: false},
			expected: false,
		},
		{
			name:     "common true, target false",
			common:   schemaR.StringAttribute{Sensitive: true},
			target:   schemaR.StringAttribute{Sensitive: false},
			expected: true,
		},
		{
			name:     "common false, target true",
			common:   schemaR.StringAttribute{Sensitive: false},
			target:   schemaR.StringAttribute{Sensitive: true},
			expected: true,
		},
		{
			name:     "both true",
			common:   schemaR.StringAttribute{Sensitive: true},
			target:   schemaR.StringAttribute{Sensitive: true},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := computeIsSensitive(tt.common, tt.target)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestComputeIsWriteOnly tests the boolean OR logic for WriteOnly.
func TestComputeIsWriteOnly(t *testing.T) {
	tests := []struct {
		name     string
		common   schemaR.StringAttribute
		target   schemaR.StringAttribute
		expected bool
	}{
		{
			name:     "both false",
			common:   schemaR.StringAttribute{WriteOnly: false},
			target:   schemaR.StringAttribute{WriteOnly: false},
			expected: false,
		},
		{
			name:     "common true, target false",
			common:   schemaR.StringAttribute{WriteOnly: true},
			target:   schemaR.StringAttribute{WriteOnly: false},
			expected: true,
		},
		{
			name:     "common false, target true",
			common:   schemaR.StringAttribute{WriteOnly: false},
			target:   schemaR.StringAttribute{WriteOnly: true},
			expected: true,
		},
		{
			name:     "both true",
			common:   schemaR.StringAttribute{WriteOnly: true},
			target:   schemaR.StringAttribute{WriteOnly: true},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := computeIsWriteOnly(tt.common, tt.target)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestComputeMarkdownDescription tests string concatenation for MarkdownDescription.
func TestComputeMarkdownDescription(t *testing.T) {
	tests := []struct {
		name     string
		common   schemaR.StringAttribute
		target   schemaR.StringAttribute
		expected string
	}{
		{
			name:     "both empty",
			common:   schemaR.StringAttribute{MarkdownDescription: ""},
			target:   schemaR.StringAttribute{MarkdownDescription: ""},
			expected: "",
		},
		{
			name:     "only common",
			common:   schemaR.StringAttribute{MarkdownDescription: "common description"},
			target:   schemaR.StringAttribute{MarkdownDescription: ""},
			expected: "common description",
		},
		{
			name:     "only target",
			common:   schemaR.StringAttribute{MarkdownDescription: ""},
			target:   schemaR.StringAttribute{MarkdownDescription: "target description"},
			expected: "target description",
		},
		{
			name:     "both present",
			common:   schemaR.StringAttribute{MarkdownDescription: "common"},
			target:   schemaR.StringAttribute{MarkdownDescription: "target"},
			expected: "common target",
		},
		{
			name:     "with extra spaces",
			common:   schemaR.StringAttribute{MarkdownDescription: "  common  "},
			target:   schemaR.StringAttribute{MarkdownDescription: "  target  "},
			expected: "  common target  ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := computeMarkdownDescription(tt.common, tt.target)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}

// TestComputeDescription tests string concatenation for Description.
func TestComputeDescription(t *testing.T) {
	tests := []struct {
		name     string
		common   schemaR.StringAttribute
		target   schemaR.StringAttribute
		expected string
	}{
		{
			name:     "both empty",
			common:   schemaR.StringAttribute{Description: ""},
			target:   schemaR.StringAttribute{Description: ""},
			expected: "",
		},
		{
			name:     "only common",
			common:   schemaR.StringAttribute{Description: "common"},
			target:   schemaR.StringAttribute{Description: ""},
			expected: "common",
		},
		{
			name:     "only target",
			common:   schemaR.StringAttribute{Description: ""},
			target:   schemaR.StringAttribute{Description: "target"},
			expected: "target",
		},
		{
			name:     "both present",
			common:   schemaR.StringAttribute{Description: "common"},
			target:   schemaR.StringAttribute{Description: "target"},
			expected: "common target",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := computeDescription(tt.common, tt.target)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}

// TestComputeDeprecationMessage tests string concatenation for DeprecationMessage.
func TestComputeDeprecationMessage(t *testing.T) {
	tests := []struct {
		name     string
		common   schemaR.StringAttribute
		target   schemaR.StringAttribute
		expected string
	}{
		{
			name:     "both empty",
			common:   schemaR.StringAttribute{DeprecationMessage: ""},
			target:   schemaR.StringAttribute{DeprecationMessage: ""},
			expected: "",
		},
		{
			name:     "only common",
			common:   schemaR.StringAttribute{DeprecationMessage: "use new field"},
			target:   schemaR.StringAttribute{DeprecationMessage: ""},
			expected: "use new field",
		},
		{
			name:     "only target",
			common:   schemaR.StringAttribute{DeprecationMessage: ""},
			target:   schemaR.StringAttribute{DeprecationMessage: "use new field"},
			expected: "use new field",
		},
		{
			name:     "both present",
			common:   schemaR.StringAttribute{DeprecationMessage: "common reason"},
			target:   schemaR.StringAttribute{DeprecationMessage: "target reason"},
			expected: "common reason target reason",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := computeDeprecationMessage(tt.common, tt.target)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}

// TestDataSourceAttributeType verifies datasource variant merge logic.
func TestComputeDescriptionWithDataSourceAttribute(t *testing.T) {
	tests := []struct {
		name     string
		common   schemaR.StringAttribute
		target   schemaD.StringAttribute
		expected string
	}{
		{
			name:     "both empty",
			common:   schemaR.StringAttribute{Description: ""},
			target:   schemaD.StringAttribute{Description: ""},
			expected: "",
		},
		{
			name:     "common only",
			common:   schemaR.StringAttribute{Description: "common desc"},
			target:   schemaD.StringAttribute{Description: ""},
			expected: "common desc",
		},
		{
			name:     "target only",
			common:   schemaR.StringAttribute{Description: ""},
			target:   schemaD.StringAttribute{Description: "datasource desc"},
			expected: "datasource desc",
		},
		{
			name:     "both present",
			common:   schemaR.StringAttribute{Description: "common"},
			target:   schemaD.StringAttribute{Description: "datasource"},
			expected: "common datasource",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := computeDescription(tt.common, tt.target)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}
