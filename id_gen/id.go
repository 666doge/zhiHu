package id_gen

import (
	"fmt"
	"github.com/sony/sonyflake"
)

var (
	sf *sonyflake.Sonyflake
	sonyMachineId uint16
)

func getMachineId() (uint16, error) {
	return sonyMachineId, nil
}

func Init(machineId uint16) error {
	sonyMachineId = machineId

	sf = sonyflake.NewSonyflake(sonyflake.Settings{
		MachineID: getMachineId,
	})
	return nil
}

func GetId() (id uint64, err error){
	if sf == nil {
		err = fmt.Errorf("sony flake is not inited")
		return
	}

	id, err = sf.NextID()
	return
}