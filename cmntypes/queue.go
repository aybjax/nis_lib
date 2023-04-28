package cmntypes

import "github.com/aybjax/nis_lib/pbdto"

type AppQueue interface {
	TransformAndNotify(*pbdto.DiffIds) error
	Notify(*pbdto.UpdateEmbedded) error
	AddListener(func(*pbdto.UpdateEmbedded) error)
}
