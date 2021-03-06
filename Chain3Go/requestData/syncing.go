// requestData project syncing.go
package requestData

import "serverLayer/Chain3Go/types"

type SyncingResponse struct {
	StartingBlock types.ComplexIntResponse `json:"startingBlock"`
	CurrentBlock  types.ComplexIntResponse `json:"currentBlock"`
	HighestBlock  types.ComplexIntResponse `json:"highestBlock"`
}
