package cmntypes

import "github.com/aybjax/nis_lib/pbdto"

type AppQueue interface {
	// Send to queue for transformation and further action
	// sends to itself
	SendToNotify(*pbdto.DiffIds) error
	SendToNotifyListener(func(*pbdto.DiffIds) error)
	// Send to another service about updated embedded info
	Notify(*pbdto.UpdateEmbedded) error
	NotifyListener(func(*pbdto.UpdateEmbedded) error)
}
