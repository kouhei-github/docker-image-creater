package services

import "github.com/cheggaaa/pb/v3"

type PoolImp interface {
	AddPool(managePool ...*pb.ProgressBar)
	Start() error
}

type SteBarImp interface {
	Progress()
}
