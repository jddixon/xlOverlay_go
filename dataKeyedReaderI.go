package overlay

// xlOverlay_go/dataKeyedReader.go

import (
	xi "github.com/jddixon/xlNodeID_go"
)

type DataKeyedReaderI interface {

	/**
	 * Retrieve data by content key (content hash).  This interface
	 * blocks.
	 */
	Get(nodeID *xi.NodeID) error

	/**
	 * Retrieve a serialized SignedList, given its key, calculated
	 * from the RSA public key and title of the list.  This is NOT
	 * a content key.  The call blocks.
	 */
	GetSigned(nodeID *xi.NodeID) error
}
