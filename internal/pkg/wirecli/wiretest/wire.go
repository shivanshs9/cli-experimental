//+build wireinject

/*
Copyright 2019 The Kubernetes Authors.
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

package wiretest

import (
	"io"

	"github.com/google/wire"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"sigs.k8s.io/cli-experimental/internal/pkg/apply"
	"sigs.k8s.io/cli-experimental/internal/pkg/clik8s"
	"sigs.k8s.io/cli-experimental/internal/pkg/dy"
	"sigs.k8s.io/cli-experimental/internal/pkg/dy/dispatch"
	"sigs.k8s.io/cli-experimental/internal/pkg/dy/list"
	"sigs.k8s.io/cli-experimental/internal/pkg/status"
)

func InitializeStatus(clik8s.ResourceConfigs, *object.Commit, io.Writer) (*status.Status, func(), error) {
	panic(wire.Build(ProviderSet))
}

func InitializeApply(clik8s.ResourceConfigs, *object.Commit, io.Writer) (*apply.Apply, func(), error) {
	panic(wire.Build(ProviderSet))
}

func InitializeCommandBuilder(io.Writer) (*dy.CommandBuilder, func(), error) {
	panic(wire.Build(ProviderSet))
}

func InitializeDispatcher(io.Writer) (*dispatch.Dispatcher, func(), error) {
	panic(wire.Build(ProviderSet))
}

func InitializeLister(io.Writer) (*list.CommandLister, func(), error) {
	panic(wire.Build(ProviderSet))
}
