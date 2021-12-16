package day16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	packet1 := ParseData("8A004A801A8002F478")
	t.Log(packet1)
	assert.Equal(t, 16, packet1.GetVersionSum())
	packet2 := ParseData("620080001611562C8802118E34")
	t.Log(packet2)
	assert.Equal(t, 12, packet2.GetVersionSum())
	packet3 := ParseData("C0015000016115A2E0802F182340")
	t.Log(packet3)
	assert.Equal(t, 23, packet3.GetVersionSum())
	packet4 := ParseData("A0016C880162017C3686B18A3D4780")
	t.Log(packet4)
	assert.Equal(t, 31, packet4.GetVersionSum())
}

func TestPart2(t *testing.T) {
	packet1 := ParseData("C200B40A82")
	t.Log(packet1)
	assert.Equal(t, 3, packet1.Evaluate())
	packet2 := ParseData("04005AC33890")
	t.Log(packet2)
	assert.Equal(t, 54, packet2.Evaluate())
	packet3 := ParseData("880086C3E88112")
	t.Log(packet3)
	assert.Equal(t, 7, packet3.Evaluate())
	packet4 := ParseData("CE00C43D881120")
	t.Log(packet4)
	assert.Equal(t, 9, packet4.Evaluate())
	packet5 := ParseData("D8005AC2A8F0")
	t.Log(packet5)
	assert.Equal(t, 1, packet5.Evaluate())
	packet6 := ParseData("F600BC2D8F")
	t.Log(packet6)
	assert.Equal(t, 0, packet6.Evaluate())
	packet7 := ParseData("9C005AC2F8F0")
	t.Log(packet7)
	assert.Equal(t, 0, packet7.Evaluate())
	packet8 := ParseData("9C0141080250320F1802104A08")
	t.Log(packet8)
	assert.Equal(t, 1, packet8.Evaluate())
}

func Test_bitsToInt(t *testing.T) {
	type args struct {
		bits []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{bits: []int{0}}, want: 0},
		{name: "test2", args: args{bits: []int{1, 0}}, want: 2},
		{name: "test3", args: args{bits: []int{1, 0, 1}}, want: 5},
		{name: "test3", args: args{bits: []int{0, 1, 0, 1}}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bitsToInt(tt.args.bits); got != tt.want {
				t.Errorf("bitsToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getBitSliceValue(t *testing.T) {
	type args struct {
		binData []int
		start   int
		end     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{binData: []int{0}, start: 0, end: 0}, want: 0},
		{name: "test2", args: args{binData: []int{1}, start: 0, end: 0}, want: 1},
		{name: "test3", args: args{binData: []int{1, 0, 1}, start: 0, end: 1}, want: 0b10},
		{name: "test4", args: args{binData: []int{0, 1, 0, 1}, start: 0, end: 1}, want: 0b01},
		{name: "test5", args: args{binData: []int{1, 1, 0, 1, 0, 1, 1, 0, 0, 0, 1}, start: 3, end: 6}, want: 0b1011},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBitSliceValue(tt.args.binData, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("getBitSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseBinaryData(t *testing.T) {
	type args struct {
		binData []int
	}
	tests := []struct {
		name         string
		args         args
		version      int
		typeId       int
		value        int
		remainLength int
	}{
		{name: "test1", args: args{binData: []int{1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0}}, remainLength: 3, version: 6, typeId: 4, value: 2021},
		{name: "test2", args: args{binData: []int{0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0}}, remainLength: 7, version: 1, typeId: 6, value: 20 + 10},
		{name: "test3", args: args{binData: []int{1, 1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0}}, remainLength: 5, version: 7, typeId: 3, value: 1 + 2 + 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			packet, remaining := parseBinaryData(tt.args.binData)
			assert.Equal(t, tt.version, packet.getVersion())
			assert.Equal(t, tt.typeId, packet.getTypeId())
			assert.Equal(t, tt.value, packet.getValue())
			assert.Equal(t, tt.remainLength, len(remaining))
		})
	}
}
