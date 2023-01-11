package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		"enode://dae27f8361eb42cfdff3fcc5711fa6ac98db57ae1dd884a54e0da31f2089ad80d7dba60d1351a4715504ac8011c25f6398ba3258ffff7f99fc1030f823bca340@185.25.48.50:15060",
		"enode://eff4046530334404b485e1ac4a91e5631ddaf9e4574719bf2133fd5f4ae071cb576843cace4934986a11a596b49f9e93a0f01947d01000560e1838c23a9b055b@185.25.48.213:15060",
		"enode://2c4055b73a7cebce2bc657b7a5174c94dcf0de4ae5b28c68e10d472614cfc549b0ff01b3d683a9f4a49b02744a9f5b8ea645fc422766db083c7d21c873155ecc@185.25.48.217:15060",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}
