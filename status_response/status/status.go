package status

// Status represents Status of StatusResponse
type Status int

// Status list
const (
	Processed  Status = 2
	Pending    Status = 0
	Cancelled  Status = -1
	Failed     Status = -2
	Chargeback Status = -3
)

func (status Status) String() string {
	switch status {
	case Processed:
		return "processed"
	case Pending:
		return "pending"
	case Cancelled:
		return "cancelled"
	case Failed:
		return "failed"
	case Chargeback:
		return "chargeback"
	default:
		return "unkown"
	}
}
