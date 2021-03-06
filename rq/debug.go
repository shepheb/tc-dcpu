package rq

import "strings"

var regNames = map[string]uint16{
	"r0": 0,
	"r1": 1,
	"r2": 2,
	"r3": 3,
	"r4": 4,
	"r5": 5,
	"r6": 6,
	"r7": 7,
}

var registers = []string{"r0", "r1", "r2", "r3", "r4", "r5", "r6", "r7", "pc", "sp", "lr", "cpsr", "spsr"}

func (c *rq) Registers() []string {
	return registers
}

func (c *rq) RegByName(name string) (uint32, string, bool) {
	if r, ok := regNames[name]; ok {
		return uint32(c.regs[r]), strings.ToLower(name), true
	}

	if name == "pc" || name == "PC" {
		return uint32(c.pc), "pc", true
	} else if name == "sp" || name == "SP" {
		return uint32(c.sp), "sp", true
	} else if name == "lr" || name == "LR" {
		return uint32(c.lr), "lr", true
	} else if name == "cpsr" || name == "CPSR" {
		return uint32(c.cpsr), "cpsr", true
	} else if name == "spsr" || name == "SPSR" {
		return uint32(c.spsr), "spsr", true
	} else {
		return 0, "", false
	}
}

func (*rq) RegisterWidth(name string) int {
	return 16
}
