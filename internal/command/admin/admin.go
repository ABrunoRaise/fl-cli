// Copyright 2022 Giuseppe De Palma, Matteo Trentin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package admin

type Admin struct {
	Dev          dev   `cmd:"" help:"deploy locally with 1 core and 1 worker docker containers"`
	Reset        reset `cmd:"" help:"removes the deployment of local containers"`
	Kubernetes   k8s   `cmd:"" name:"kubernetes" aliases:"k,k8s" help:"deploy on an existing kubernetes cluster"`
	KubernetesRM k8sRm `cmd:"" name:"kubernetes-rm" aliases:"krm, k8srm" help:"remove the deployment on an existing kubernetes cluster"`
}
