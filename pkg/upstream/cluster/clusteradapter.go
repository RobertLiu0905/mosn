/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cluster

import (
	"fmt"

	"github.com/alipay/sofa-mosn/pkg/api/v2"
)

var clusterMngAdapterInstance *ClusterMngAdapter

func initClusterMngAdapterInstance(clusterMng *clusterManager) {
	clusterMngAdapterInstance = &ClusterMngAdapter{
		clusterMng: clusterMng,
	}
}

func GetClusterMngAdapterInstance() *ClusterMngAdapter {
	return clusterMngAdapterInstance
}

type ClusterMngAdapter struct {
	clusterMng *clusterManager
}

// TriggerClusterAddOrUpdate
// Added or Update Cluster
func (ca *ClusterMngAdapter) TriggerClusterAddOrUpdate(cluster v2.Cluster) error {
	if ca.clusterMng == nil {
		return fmt.Errorf("TriggerClusterAddOrUpdate Error: cluster manager is nil")
	}

	if !ca.clusterMng.AddOrUpdatePrimaryCluster(cluster) {
		return fmt.Errorf("TriggerClusterAddOrUpdate failure, cluster name = %s", cluster.Name)
	}

	return nil
}

// TriggerClusterAddOrUpdate
// Added or Update Cluster and Cluster's hosts
func (ca *ClusterMngAdapter) TriggerClusterAndHostsAddOrUpdate(cluster v2.Cluster, hosts []v2.Host) error {
	if err := ca.TriggerClusterAddOrUpdate(cluster); err != nil {
		return err
	}

	return ca.clusterMng.UpdateClusterHosts(cluster.Name, 0, hosts)
}

// TriggerClusterHostUpdate
// Added or Update Cluster's hosts, return err if cluster not exist
func (ca *ClusterMngAdapter) TriggerClusterHostUpdate(clusterName string, hosts []v2.Host) error {
	if ca.clusterMng == nil {
		return fmt.Errorf("TriggerClusterAddOrUpdate Error: cluster manager is nil")
	}

	return ca.clusterMng.UpdateClusterHosts(clusterName, 0, hosts)
}

// TriggerClusterDel
// used to delete cluster by clusterName
func (ca *ClusterMngAdapter) TriggerClusterDel(clusterName string) error {
	if ca.clusterMng == nil {
		return fmt.Errorf("TriggerClusterAddOrUpdate Error: cluster manager is nil")
	}

	return ca.clusterMng.RemovePrimaryCluster(clusterName)
}
