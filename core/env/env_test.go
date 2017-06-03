/*
Copyright 2016 Medcl (m AT medcl.net)

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

package env

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoad(t *testing.T) {
	env := Environment("../../gopa.full.yml")
	fmt.Println(env)
	fmt.Println(env.SystemConfig)

	assert.Equal(t, "gopa1", env.SystemConfig.ClusterConfig.Name)
	assert.Equal(t, "node1", env.SystemConfig.NodeConfig.Name)

}
