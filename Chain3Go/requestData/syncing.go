// requestData project syncing.go
package requestData

import "library/Lchain3/types"

type SyncingResponse struct {
	StartingBlock types.ComplexIntResponse `json:"startingBlock"`
	CurrentBlock  types.ComplexIntResponse `json:"currentBlock"`
	HighestBlock  types.ComplexIntResponse `json:"highestBlock"`
}
