/*******************************************************************************
 * Copyright 2021 Intel Corporation
 * Copyright 2019 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *
 * @author: Daniel Harms, Dell
 * @author: Tingyu Zeng, Dell
 *******************************************************************************/

package secretstore

const (
	// this is a feature key to indicate whether to enable Registry Consul's ACL or not
	RegistryACLFeatureFlag = "ENABLE_REGISTRY_ACL"

	VaultToken             = "X-Vault-Token"
	TokenCreatorPolicyName = "privileged-token-creator"

	// This is an admin token policy that allow for creation of
	// per-service tokens and policies
	TokenCreatorPolicy = `
path "auth/token/create" {
  capabilities = ["create", "update", "sudo"]
}

path "auth/token/create-orphan" {
  capabilities = ["create", "update", "sudo"]
}

path "auth/token/create/*" {
  capabilities = ["create", "update", "sudo"]
}

path "sys/policies/acl/edgex-service-*"
{
  capabilities = ["create", "read", "update", "delete" ]
}

path "sys/policies/acl"
{
  capabilities = ["list"]
}
`
)
