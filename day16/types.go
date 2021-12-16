package day16

type Packet interface {
	getVersion() int
	getTypeId() int
	getValue() int
	GetVersionSum() int
	Evaluate() int
}

type litteralPacket struct {
	version int
	typeId  int
	value   int
}

func (p litteralPacket) getVersion() int {
	return p.version
}

func (p litteralPacket) getTypeId() int {
	return p.typeId
}

func (p litteralPacket) getValue() int {
	return p.value
}

func (p litteralPacket) Evaluate() int {
	return p.value
}

func (p litteralPacket) GetVersionSum() int {
	return p.version
}

type operatorPacket struct {
	version int
	typeId  int
	packets []Packet
}

func (p operatorPacket) getVersion() int {
	return p.version
}

func (p operatorPacket) getTypeId() int {
	return p.typeId
}

func (p operatorPacket) getValue() int {
	value := 0
	for _, packet := range p.packets {
		value += packet.getValue()
	}
	return value
}

func (p operatorPacket) GetVersionSum() int {
	value := p.version
	for _, packet := range p.packets {
		value += packet.GetVersionSum()
	}
	return value
}

func (p operatorPacket) Evaluate() int {
	value := 0
	// +
	if p.typeId == 0 {
		for _, packet := range p.packets {
			value += packet.Evaluate()
		}
		return value

	}
	// *
	if p.typeId == 1 {
		value := 1
		for _, packet := range p.packets {
			value *= packet.Evaluate()
		}
		return value
	}
	// min
	if p.typeId == 2 {
		value := p.packets[0].Evaluate()
		for _, packet := range p.packets {
			newVal := packet.Evaluate()
			if newVal < value {
				value = newVal
			}
		}
		return value
	}
	// max
	if p.typeId == 3 {
		value := p.packets[0].Evaluate()
		for _, packet := range p.packets {
			newVal := packet.Evaluate()
			if newVal > value {
				value = newVal
			}
		}
		return value
	}
	// greater
	if p.typeId == 5 {
		if p.packets[0].Evaluate() > p.packets[1].Evaluate() {
			return 1
		}
		return 0
	}
	// less
	if p.typeId == 6 {
		if p.packets[0].Evaluate() < p.packets[1].Evaluate() {
			return 1
		}
		return 0
	}
	// equal
	if p.typeId == 7 {
		if p.packets[0].Evaluate() == p.packets[1].Evaluate() {
			return 1
		}
		return 0
	}
	panic("Unknown typeId")
}
