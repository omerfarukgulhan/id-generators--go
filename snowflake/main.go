package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	timestampBits   = 41
	datacenterBits  = 5
	machineBits     = 5
	sequenceBits    = 12
	maxDatacenterID = (1 << datacenterBits) - 1 // 31
	maxMachineID    = (1 << machineBits) - 1    // 31
	maxSequence     = (1 << sequenceBits) - 1   // 4095
	epoch           = int64(1288834974657)
	timestampShift  = sequenceBits + machineBits + datacenterBits
	datacenterShift = sequenceBits + machineBits
	machineShift    = sequenceBits
)

type Snowflake struct {
	sync.Mutex
	timestamp    int64
	datacenterID int64
	machineID    int64
	sequence     int64
}

func NewSnowflake(datacenterID, machineID int64) (*Snowflake, error) {
	if datacenterID > maxDatacenterID || datacenterID < 0 {
		return nil, fmt.Errorf("datacenter ID must be between 0 and %d", maxDatacenterID)
	}

	if machineID > maxMachineID || machineID < 0 {
		return nil, fmt.Errorf("machine ID must be between 0 and %d", maxMachineID)
	}

	return &Snowflake{
		timestamp:    0,
		datacenterID: datacenterID,
		machineID:    machineID,
		sequence:     0,
	}, nil
}

func (s *Snowflake) GenerateID() int64 {
	s.Lock()
	defer s.Unlock()
	now := time.Now().UnixMilli()
	if now == s.timestamp {
		s.sequence = (s.sequence + 1) & maxSequence
		if s.sequence == 0 {
			for now <= s.timestamp {
				now = time.Now().UnixMilli()
			}
		}
	} else {
		s.sequence = 0
	}

	s.timestamp = now
	id := ((now - epoch) << timestampShift) |
		(s.datacenterID << datacenterShift) |
		(s.machineID << machineShift) |
		s.sequence

	return id
}

func main() {
	snowflake, err := NewSnowflake(1, 1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i := 0; i < 5; i++ {
		id := snowflake.GenerateID()
		println(id)
	}
}
