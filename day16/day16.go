package day16

import (
	"adventofcode2021/common"
	"fmt"
	"strconv"
)

func ParseData(input string) Packet {
	bits := make([]int, 0)
	for _, char := range input {
		value, _ := strconv.ParseUint(string(char), 16, 64)
		valueAsBinary := fmt.Sprintf("%04b", value)
		for _, bit := range valueAsBinary {
			bitValue, _ := strconv.ParseUint(string(bit), 2, 64)
			bits = append(bits, int(bitValue))
		}
	}

	packet, _ := parseBinaryData(bits)
	return packet
}

func parseBinaryData(bits []int) (Packet, []int) {
	var packet Packet
	var subpackets []Packet

	length := len(bits)
	remainingBits := bits

	version := getBitSliceValue(bits, 0, 2)
	typeId := getBitSliceValue(bits, 3, 5)

	// litteral
	if typeId == 4 {
		litteralBits := make([]int, 0)
		index := 6
		for {
			litteralBits = append(litteralBits, bits[index+1])
			litteralBits = append(litteralBits, bits[index+2])
			litteralBits = append(litteralBits, bits[index+3])
			litteralBits = append(litteralBits, bits[index+4])
			breakBit := getBitSliceValue(bits, index, index)
			if breakBit == 0 {
				break
			}
			index += 5
		}
		remainingBits = getSubBits(bits, index+5, length-1)

		litteralValue := getBitSliceValue(litteralBits, 0, len(litteralBits)-1)
		packet = litteralPacket{
			version: version,
			typeId:  typeId,
			value:   litteralValue,
		}
	} else {
		lengthTypeId := getBitSliceValue(bits, 6, 6)
		if lengthTypeId == 0 {
			subPacketsLength := int(getBitSliceValue(bits, 7, 21))
			lastBitIndex := 7 + 15 + subPacketsLength
			subPacketsBits := getSubBits(bits, 22, lastBitIndex-1)
			remainingBits = getSubBits(bits, lastBitIndex, length-1)

			for {
				nextSubPacket, nextSubRemainingBits := parseBinaryData(subPacketsBits)
				subpackets = append(subpackets, nextSubPacket)
				if len(nextSubRemainingBits) == 0 {
					break
				}
				subPacketsBits = nextSubRemainingBits
			}

		} else {
			packetCount := int(getBitSliceValue(bits, 7, 17))

			lastBitIndex := 7 + 11
			remainingBits = getSubBits(bits, lastBitIndex, length-1)
			for i := 1; i <= packetCount; i++ {
				nextSubPacket, nextRemainingBits := parseBinaryData(remainingBits)
				subpackets = append(subpackets, nextSubPacket)
				remainingBits = nextRemainingBits
			}
		}
		packet = operatorPacket{
			version: version,
			typeId:  typeId,
			packets: subpackets,
		}

	}

	return packet, remainingBits
}

func bitsToInt(bits []int) int {
	str := ""
	if len(bits) == 0 {
		return 0
	}
	for _, bit := range bits {
		str += fmt.Sprint(bit)
	}
	value, err := strconv.ParseInt(str, 2, 64)
	common.Check(err)
	return int(value)
}

func getBitSliceValue(bitsData []int, start int, end int) int {
	return bitsToInt(getSubBits(bitsData, start, end))
}

func getSubBits(bitsData []int, start int, end int) []int {
	bits := make([]int, 0)
	for x := start; x <= end; x++ {
		bits = append(bits, bitsData[x])
	}
	return bits
}

func Step1(data []string) int {
	return len(data)
}

func Step2(data []string) int {
	return len(data) * 2
}
