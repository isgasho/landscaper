// Copyright 2020 Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DataObjectContext defines the context of a data object.
type DataObjectContext string

const (
	// ExportDataObjectContext is the data object type of a exported object.
	ExportDataObjectContext DataObjectContext = "export"
	// ExportDataObjectContext is the data object type of a imported object.
	ImportDataObjectContext DataObjectContext = "import"
)

// DataObjectTypeAnnotation defines the name of the annotation that specifies the type of the dataobject.
const DataObjectTypeAnnotation = "dataobjects.landscaper.gardener.cloud/type"

// DataObjectContextLabel defines the name of the label that specifies the context of the dataobject.
const DataObjectContextLabel = "dataobjects.landscaper.gardener.cloud/context"

// DataObjectKeyLabel defines the name of the label that specifies the export or imported key of the dataobject.
const DataObjectKeyLabel = "dataobjects.landscaper.gardener.cloud/key"

// DataObjectSourceLabel defines the name of the label that specifies the source of the dataobject.
const DataObjectSourceLabel = "dataobjects.landscaper.gardener.cloud/source"

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DataObjectList contains a list of DataObject
type DataObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataType `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DataObject are resources that can hold any kind json or yaml data.
// +kubebuilder:resource:path="dataobjects",scope="Cluster",shortName={"do","dobj"},singular="dataobject"
type DataObject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Data contains the data of the object as string.
	Data json.RawMessage `json:"data"`
}
