package sample

import (
	"github.com/golang/protobuf/ptypes"
	"gitlab.com/techschool/pcbook/pb"
)

func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:   randomKeyboardLayout(),
		Blacklit: randomBool(),
	}
	return keyboard
}

func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)
	numberCores := randomInt(2, 8)
	numberThread := randomInt(2, 8)
	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.5)
	cpu := &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberThread),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}

	return cpu
}

func NewGPU() *pb.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)
	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.8)
	gpu := &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: &pb.Memory{
			Value: uint64(randomInt(2, 6)),
			Unit:  pb.Memory_GIGABYTE,
		},
	}
	return gpu
}

func NewRam() *pb.Memory {
	ram := &pb.Memory{
		Value: uint64(randomInt(2, 6)),
		Unit:  pb.Memory_GIGABYTE,
	}
	return ram
}

func NewStorage(storageType pb.Storage_Driver) *pb.Storage {
	storage := &pb.Storage{
		Driver: storageType,
		Memory: &pb.Memory{
			Value: uint64(randomInt(1, 6)),
			Unit:  pb.Memory_TERABYTE,
		},
	}
	return storage
}

func NewScreen() *pb.Screen {
	height := uint32(randomInt(1880, 4320))
	screen := &pb.Screen{
		Multitouch: randomBool(),
		Panel:      randomScreenPanel(),
		Resolution: &pb.Screen_Resolution{
			Height: height,
			Width:  height * 16 / 9,
		},
		SizeInch: randomFloat32(13, 17),
	}
	return screen
}

func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()
	cpu := NewCPU()
	laptop := &pb.Laptop{
		Id: randomID(),
		Brand: brand,
		Name: randomLaptopName(brand),
		Cpu: cpu,
		Ram: NewRam(),
		Gpus: []*pb.GPU{NewGPU()},
		Storages: []*pb.Storage{NewStorage(pb.Storage_SSD), NewStorage(pb.Storage_HDD)},
		Screen: NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pb.Laptop_WeightKg{randomFloat64(1.8, 3.0)},
		PriceUsd: randomFloat64(1500, 3000),
		ReleaseYear: uint32(randomInt(2015, 2020)),
		UpdateAt: ptypes.TimestampNow(),
	}
	return laptop
}