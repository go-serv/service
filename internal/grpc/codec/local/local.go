package net

import (
	"github.com/go-serv/service/internal/ancillary/net"
	cc "github.com/go-serv/service/internal/grpc/codec"
	"github.com/go-serv/service/pkg/z"
)

const (
	SharedMem_IPC z.HeaderFlagsType = 1 << iota
)

type codec struct {
	z.CodecInterface
}

type dataFrame struct {
	z.DataFrameInterface
	shmObjName   string
	shmBlockSize uint32
	shmDataSize  uint32
}

func (c *codec) NewDataFrame() z.DataFrameInterface {
	df := new(dataFrame)
	df.DataFrameInterface = cc.NewDataFrame()
	return df
}

func (df *dataFrame) Parse(b []byte, hook func(*net.NetReader) error) error {
	return df.DataFrameInterface.Parse(b, func(netr *net.NetReader) (err error) {
		// Do nothing unless the shared memory IPC flag is set
		if df.HeaderFlags().Has(SharedMem_IPC) {
			df.shmObjName, err = netr.ReadString()
			if err != nil {
				return
			}
			df.shmBlockSize, err = net.GenericNetReader[uint32](netr)
			if err != nil {
				return
			}
			df.shmDataSize, err = net.GenericNetReader[uint32](netr)
			if err != nil {
				return
			}
		}
		return
	})
}

func (df *dataFrame) Compose([]byte) (out []byte, err error) {
	var header []byte
	if df.HeaderFlags().Has(SharedMem_IPC) {
		netw := net.NewWriter()
		netw.WriteString(df.shmObjName)
		err = net.GenericNetWriter[uint32](netw, df.shmBlockSize)
		if err != nil {
			return
		}
		err = net.GenericNetWriter[uint32](netw, df.shmDataSize)
		if err != nil {
			return
		}
		header = netw.Bytes()
	}
	out, err = df.DataFrameInterface.Compose(header)
	return
}

func (df *dataFrame) SharedMemObjectName() string {
	return df.shmObjName
}

func (df *dataFrame) WithSharedMemObjectName(name string) {
	df.shmObjName = name
}

func (df *dataFrame) SharedMemBlockSize() uint32 {
	return df.shmDataSize
}

func (df *dataFrame) WithSharedMemBlockSize(size uint32) {
	df.shmDataSize = size
}

func (df *dataFrame) SharedMemDataSize() uint32 {
	return df.shmDataSize
}

func (df *dataFrame) WithSharedMemDataSize(size uint32) {
	df.shmDataSize = size
}
