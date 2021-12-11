/*
 * CoreDB API
 *
 * An interface to distributed nodes for personal storage.
 *
 * API version: 1.0.4
 * Contact: info@coredb.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package coredb

type SubjectAsset struct {

	AssetId string `json:"assetId"`

	OriginalId string `json:"originalId,omitempty"`

	Transform string `json:"transform,omitempty"`

	State string `json:"state,omitempty"`

	Size int64 `json:"size,omitempty"`

	Hash string `json:"hash,omitempty"`

	Created int64 `json:"created,omitempty"`
}