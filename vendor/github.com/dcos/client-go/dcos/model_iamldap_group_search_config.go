/*
 * DC/OS
 *
 * DC/OS API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dcos

type IamldapGroupSearchConfig struct {
	SearchBase           string   `json:"search-base"`
	SearchFilterTemplate string   `json:"search-filter-template"`
	MemberAttributes     []string `json:"member-attributes,omitempty"`
}
