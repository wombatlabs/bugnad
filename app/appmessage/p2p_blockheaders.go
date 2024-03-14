package appmessage

// BlockHeadersMessage represents a bugna BlockHeaders message
type BlockHeadersMessage struct {
	baseMessage
	BlockHeaders []*MsgBlockHeader
}

// Command returns the protocol command string for the message
func (msg *BlockHeadersMessage) Command() MessageCommand {
	return CmdBlockHeaders
}

// NewBlockHeadersMessage returns a new bugna BlockHeaders message
func NewBlockHeadersMessage(blockHeaders []*MsgBlockHeader) *BlockHeadersMessage {
	return &BlockHeadersMessage{
		BlockHeaders: blockHeaders,
	}
}
