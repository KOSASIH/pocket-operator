/*
Copyright 2023.

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

package pocketset

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	nodesv1alpha1 "github.com/pokt-network/pocket-operator/apis/nodes/v1alpha1"
	"github.com/pokt-network/pocket-operator/apis/nodes/v1alpha1/pocketset/mutate"
)

// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;create;update;patch;delete

// CreateConfigMapParentNameParentNameGenesis creates the ConfigMap resource with name parent.name + -genesis.
func CreateConfigMapParentNameParentNameGenesis(
	parent *nodesv1alpha1.PocketSet,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]interface{}{
				"name":      "" + parent.Name + "-genesis", //  controlled by field:
				"namespace": parent.Name,                   //  controlled by field:
			},
			"data": map[string]interface{}{
				"genesis.json": `{
  "accounts": [
    {
      "address": "6f66574e1f50f0ef72dff748c3f11b9e0e89d32a",
      "amount": "100000000000000"
    },
    {
      "address": "67eb3f0a50ae459fecf666be0e93176e92441317",
      "amount": "100000000000000"
    },
    {
      "address": "3f52e08c4b3b65ab7cf098d77df5bf8cedcf5f99",
      "amount": "100000000000000"
    },
    {
      "address": "113fdb095d42d6e09327ab5b8df13fd8197a1eaf",
      "amount": "100000000000000"
    },
    {
      "address": "43d9ea9d9ad9c58bb96ec41340f83cb2cabb6496",
      "amount": "100000000000000"
    },
    {
      "address": "9ba047197ec043665ad3f81278ab1f5d3eaf6b8b",
      "amount": "100000000000000"
    },
    {
      "address": "88a792b7aca673620132ef01f50e62caa58eca83",
      "amount": "100000000000000"
    }
  ],
  "pools": [
    {
      "address": "DAO",
      "amount": "100000000000000"
    },
    {
      "address": "FeeCollector",
      "amount": "0"
    },
    {
      "address": "AppStakePool",
      "amount": "100000000000000"
    },
    {
      "address": "ValidatorStakePool",
      "amount": "100000000000000"
    },
    {
      "address": "ServiceNodeStakePool",
      "amount": "100000000000000"
    },
    {
      "address": "FishermanStakePool",
      "amount": "100000000000000"
    }
  ],
  "validators": [
    {
      "address": "6f66574e1f50f0ef72dff748c3f11b9e0e89d32a",
      "public_key": "b2eda2232ffb2750bf761141f70f75a03a025f65b2b2b417c7f8b3c9ca91e8e4",
      "chains": null,
      "generic_param": "v1-validator1:8080",
      "staked_amount": "1000000000000",
      "paused_height": -1,
      "unstaking_height": -1,
      "output": "6f66574e1f50f0ef72dff748c3f11b9e0e89d32a",
      "actor_type": 4
    },
    {
      "address": "67eb3f0a50ae459fecf666be0e93176e92441317",
      "public_key": "c16043323c83ffd901a8bf7d73543814b8655aa4695f7bfb49d01926fc161cdb",
      "chains": null,
      "generic_param": "v1-validator2:8080",
      "staked_amount": "1000000000000",
      "paused_height": -1,
      "unstaking_height": -1,
      "output": "67eb3f0a50ae459fecf666be0e93176e92441317",
      "actor_type": 4
    },
    {
      "address": "3f52e08c4b3b65ab7cf098d77df5bf8cedcf5f99",
      "public_key": "a8b6be75d7551da093f788f7286c3a9cb885cfc8e52710eac5f1d5e5b4bf19b2",
      "chains": null,
      "generic_param": "v1-validator3:8080",
      "staked_amount": "1000000000000",
      "paused_height": -1,
      "unstaking_height": -1,
      "output": "3f52e08c4b3b65ab7cf098d77df5bf8cedcf5f99",
      "actor_type": 4
    },
    {
      "address": "113fdb095d42d6e09327ab5b8df13fd8197a1eaf",
      "public_key": "53ee26c82826694ffe1773d7b60d5f20dd9e91bdf8745544711bec5ff9c6fb4a",
      "chains": null,
      "generic_param": "v1-validator4:8080",
      "staked_amount": "1000000000000",
      "paused_height": -1,
      "unstaking_height": -1,
      "output": "113fdb095d42d6e09327ab5b8df13fd8197a1eaf",
      "actor_type": 4
    }
  ],
  "applications": [
    {
      "address": "88a792b7aca673620132ef01f50e62caa58eca83",
      "public_key": "5f78658599943dc3e623539ce0b3c9fe4e192034a1e3fef308bc9f96915754e0",
      "chains": ["0001"],
      "generic_param": "1000000",
      "staked_amount": "1000000000000",
      "paused_height": -1,
      "unstaking_height": -1,
      "output": "88a792b7aca673620132ef01f50e62caa58eca83",
      "actor_type": 1
    }
  ],
  "service_nodes": [
    {
      "address": "43d9ea9d9ad9c58bb96ec41340f83cb2cabb6496",
      "public_key": "16cd0a304c38d76271f74dd3c90325144425d904ef1b9a6fbab9b201d75a998b",
      "chains": ["0001"],
      "generic_param": "v1-validator1:8080",
      "staked_amount": "1000000000000",
      "paused_height": -1,
      "unstaking_height": -1,
      "output": "43d9ea9d9ad9c58bb96ec41340f83cb2cabb6496",
      "actor_type": 2
    }
  ],
  "fishermen": [
    {
      "address": "9ba047197ec043665ad3f81278ab1f5d3eaf6b8b",
      "public_key": "68efd26af01692fcd77dc135ca1de69ede464e8243e6832bd6c37f282db8c9cb",
      "chains": ["0001"],
      "generic_param": "v1-validator1:8080",
      "staked_amount": "1000000000000",
      "paused_height": -1,
      "unstaking_height": -1,
      "output": "9ba047197ec043665ad3f81278ab1f5d3eaf6b8b",
      "actor_type": 3
    }
  ],
  "params": {
    "blocks_per_session": 4,
    "app_minimum_stake": "15000000000",
    "app_max_chains": 15,
    "app_baseline_stake_rate": 100,
    "app_unstaking_blocks": 2016,
    "app_minimum_pause_blocks": 4,
    "app_max_pause_blocks": 672,
    "service_node_minimum_stake": "15000000000",
    "service_node_max_chains": 15,
    "service_node_unstaking_blocks": 2016,
    "service_node_minimum_pause_blocks": 4,
    "service_node_max_pause_blocks": 672,
    "service_nodes_per_session": 24,
    "fisherman_minimum_stake": "15000000000",
    "fisherman_max_chains": 15,
    "fisherman_unstaking_blocks": 2016,
    "fisherman_minimum_pause_blocks": 4,
    "fisherman_max_pause_blocks": 672,
    "validator_minimum_stake": "15000000000",
    "validator_unstaking_blocks": 2016,
    "validator_minimum_pause_blocks": 4,
    "validator_max_pause_blocks": 672,
    "validator_maximum_missed_blocks": 5,
    "validator_max_evidence_age_in_blocks": 8,
    "proposer_percentage_of_fees": 10,
    "missed_blocks_burn_percentage": 1,
    "double_sign_burn_percentage": 5,
    "message_double_sign_fee": "10000",
    "message_send_fee": "10000",
    "message_stake_fisherman_fee": "10000",
    "message_edit_stake_fisherman_fee": "10000",
    "message_unstake_fisherman_fee": "10000",
    "message_pause_fisherman_fee": "10000",
    "message_unpause_fisherman_fee": "10000",
    "message_fisherman_pause_service_node_fee": "10000",
    "message_test_score_fee": "10000",
    "message_prove_test_score_fee": "10000",
    "message_stake_app_fee": "10000",
    "message_edit_stake_app_fee": "10000",
    "message_unstake_app_fee": "10000",
    "message_pause_app_fee": "10000",
    "message_unpause_app_fee": "10000",
    "message_stake_validator_fee": "10000",
    "message_edit_stake_validator_fee": "10000",
    "message_unstake_validator_fee": "10000",
    "message_pause_validator_fee": "10000",
    "message_unpause_validator_fee": "10000",
    "message_stake_service_node_fee": "10000",
    "message_edit_stake_service_node_fee": "10000",
    "message_unstake_service_node_fee": "10000",
    "message_pause_service_node_fee": "10000",
    "message_unpause_service_node_fee": "10000",
    "message_change_parameter_fee": "10000",
    "acl_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "blocks_per_session_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "app_minimum_stake_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "app_max_chains_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "app_baseline_stake_rate_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "app_staking_adjustment_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "app_unstaking_blocks_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "app_minimum_pause_blocks_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "app_max_paused_blocks_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "service_node_minimum_stake_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "service_node_max_chains_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "service_node_unstaking_blocks_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "service_node_minimum_pause_blocks_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "service_node_max_paused_blocks_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "service_nodes_per_session_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "fisherman_minimum_stake_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "fisherman_max_chains_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "fisherman_unstaking_blocks_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "fisherman_minimum_pause_blocks_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "fisherman_max_paused_blocks_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "validator_minimum_stake_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "validator_unstaking_blocks_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "validator_minimum_pause_blocks_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "validator_max_paused_blocks_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "validator_maximum_missed_blocks_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "validator_max_evidence_age_in_blocks_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "proposer_percentage_of_fees_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "missed_blocks_burn_percentage_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "double_sign_burn_percentage_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "message_double_sign_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "message_send_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "message_stake_fisherman_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "message_edit_stake_fisherman_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "message_unstake_fisherman_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "message_pause_fisherman_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "message_unpause_fisherman_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "message_fisherman_pause_service_node_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "message_test_score_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "message_prove_test_score_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "message_stake_app_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "message_edit_stake_app_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "message_unstake_app_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "message_pause_app_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "message_unpause_app_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "message_stake_validator_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "message_edit_stake_validator_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "message_unstake_validator_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "message_pause_validator_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "message_unpause_validator_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "message_stake_service_node_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "message_edit_stake_service_node_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "message_unstake_service_node_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "message_pause_service_node_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "message_unpause_service_node_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45",
    "message_change_parameter_fee_owner": "da034209758b78eaea06dd99c07909ab54c99b45"
  },
  "genesis_time": {
    "seconds": 1663610702,
    "nanos": 405401000
  },
  "chain_id": "testnet",
  "max_block_bytes": 4000000
}
`,
			},
		},
	}

	return mutate.MutateConfigMapParentNameParentNameGenesis(resourceObj, parent, reconciler, req)
}
