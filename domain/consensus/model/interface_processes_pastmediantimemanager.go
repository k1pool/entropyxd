package model

import "github.com/k1pool/entropyxd/domain/consensus/model/externalapi"

// PastMedianTimeManager provides a method to resolve the
// past median time of a block
type PastMedianTimeManager interface {
	PastMedianTime(stagingArea *StagingArea, blockHash *externalapi.DomainHash) (int64, error)
	InvalidateVirtualPastMedianTimeCache()
}
