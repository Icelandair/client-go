/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package install installs the experimental API group, making it available as
// an option to all of the API encoding/decoding machinery.
package install

import (
	"github.com/Icelandair/client-go/1.5/pkg/apimachinery/announced"
	"github.com/Icelandair/client-go/1.5/pkg/apis/storage"
	"github.com/Icelandair/client-go/1.5/pkg/apis/storage/v1beta1"
	"github.com/Icelandair/client-go/1.5/pkg/util/sets"
)

func init() {
	if err := announced.NewGroupMetaFactory(
		&announced.GroupMetaFactoryArgs{
			GroupName:                  storage.GroupName,
			VersionPreferenceOrder:     []string{v1beta1.SchemeGroupVersion.Version},
			ImportPrefix:               "github.com/Icelandair/client-go/1.5/pkg/apis/storage",
			RootScopedKinds:            sets.NewString("StorageClass"),
			AddInternalObjectsToScheme: storage.AddToScheme,
		},
		announced.VersionToSchemeFunc{
			v1beta1.SchemeGroupVersion.Version: v1beta1.AddToScheme,
		},
	).Announce().RegisterAndEnable(); err != nil {
		panic(err)
	}
}
