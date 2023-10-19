package progressBar

import "github.com/cheggaaa/pb/v3"

type ManagePool struct {
	Pool *pb.Pool
}

func NewManagePool() *ManagePool {
	return &ManagePool{Pool: &pb.Pool{}}
}

func (receiver ManagePool) AddPool(managePool ...*pb.ProgressBar) {
	receiver.Pool.Add(managePool...)
}

func (receiver ManagePool) Start() error {
	return receiver.Pool.Start()
}
