/*
Copyright 2019 The Vitess Authors.

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
package tabletmanager

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/liquidata-inc/vitess/go/test/endtoend/cluster"
)

func TestFallbackSecurityPolicy(t *testing.T) {
	ctx := context.Background()
	mTablet := clusterInstance.GetVttabletInstance(masterUID)

	//Init Tablets
	err := clusterInstance.VtctlclientProcess.InitTablet(mTablet, cell, keyspaceName, hostname, shardName)
	assert.Nil(t, err)

	// Start Mysql Processes
	err = cluster.StartMySQL(ctx, mTablet, username, clusterInstance.TmpDirectory)
	assert.Nil(t, err)

	// Requesting an unregistered security_policy should fallback to deny-all.
	clusterInstance.VtTabletExtraArgs = []string{"-security_policy", "bogus"}
	err = clusterInstance.StartVttablet(mTablet, "NOT_SERVING", false, cell, keyspaceName, hostname, shardName)
	assert.Nil(t, err)

	// It should deny ADMIN role.
	url := fmt.Sprintf("http://localhost:%d/streamqueryz/terminate", mTablet.HTTPPort)
	assertNotAllowedURLTest(t, url)

	// It should deny MONITORING role.
	url = fmt.Sprintf("http://localhost:%d/debug/health", mTablet.HTTPPort)
	assertNotAllowedURLTest(t, url)

	// It should deny DEBUGGING role.
	url = fmt.Sprintf("http://localhost:%d/queryz", mTablet.HTTPPort)
	assertNotAllowedURLTest(t, url)

	// Reset the VtTabletExtraArgs
	clusterInstance.VtTabletExtraArgs = []string{}
	// Tear down custom processes
	killTablets(t, mTablet)
}

func assertNotAllowedURLTest(t *testing.T, url string) {
	resp, err := http.Get(url)
	assert.Nil(t, err)

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	assert.True(t, resp.StatusCode > 400)
	assert.Contains(t, string(body), "Access denied: not allowed")
}

func assertAllowedURLTest(t *testing.T, url string) {
	resp, err := http.Get(url)
	assert.Nil(t, err)

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	assert.NotContains(t, string(body), "Access denied: not allowed")
}

func TestDenyAllSecurityPolicy(t *testing.T) {
	ctx := context.Background()
	mTablet := clusterInstance.GetVttabletInstance(masterUID)

	//Init Tablets
	err := clusterInstance.VtctlclientProcess.InitTablet(mTablet, cell, keyspaceName, hostname, shardName)
	assert.Nil(t, err)

	// Start Mysql Processes
	err = cluster.StartMySQL(ctx, mTablet, username, clusterInstance.TmpDirectory)
	assert.Nil(t, err)

	// Requesting a deny-all security_policy.
	clusterInstance.VtTabletExtraArgs = []string{"-security_policy", "deny-all"}
	err = clusterInstance.StartVttablet(mTablet, "NOT_SERVING", false, cell, keyspaceName, hostname, shardName)
	assert.Nil(t, err)

	// It should deny ADMIN role.
	url := fmt.Sprintf("http://localhost:%d/streamqueryz/terminate", mTablet.HTTPPort)
	assertNotAllowedURLTest(t, url)

	// It should deny MONITORING role.
	url = fmt.Sprintf("http://localhost:%d/debug/health", mTablet.HTTPPort)
	assertNotAllowedURLTest(t, url)

	// It should deny DEBUGGING role.
	url = fmt.Sprintf("http://localhost:%d/queryz", mTablet.HTTPPort)
	assertNotAllowedURLTest(t, url)

	// Reset the VtTabletExtraArgs
	clusterInstance.VtTabletExtraArgs = []string{}
	// Tear down custom processes
	killTablets(t, mTablet)
}

func TestReadOnlySecurityPolicy(t *testing.T) {
	ctx := context.Background()
	mTablet := clusterInstance.GetVttabletInstance(0)

	//Init Tablets
	err := clusterInstance.VtctlclientProcess.InitTablet(mTablet, cell, keyspaceName, hostname, shardName)
	assert.Nil(t, err)

	// Start Mysql Processes
	err = cluster.StartMySQL(ctx, mTablet, username, clusterInstance.TmpDirectory)
	assert.Nil(t, err)

	// Requesting a read-only security_policy.
	clusterInstance.VtTabletExtraArgs = []string{"-security_policy", "read-only"}
	err = clusterInstance.StartVttablet(mTablet, "NOT_SERVING", false, cell, keyspaceName, hostname, shardName)
	assert.Nil(t, err)

	// It should deny ADMIN role.
	url := fmt.Sprintf("http://localhost:%d/streamqueryz/terminate", mTablet.HTTPPort)
	assertNotAllowedURLTest(t, url)

	// It should deny MONITORING role.
	url = fmt.Sprintf("http://localhost:%d/debug/health", mTablet.HTTPPort)
	assertAllowedURLTest(t, url)

	// It should deny DEBUGGING role.
	url = fmt.Sprintf("http://localhost:%d/queryz", mTablet.HTTPPort)
	assertAllowedURLTest(t, url)

	// Reset the VtTabletExtraArgs
	clusterInstance.VtTabletExtraArgs = []string{}
	// Tear down custom processes
	killTablets(t, mTablet)
}
