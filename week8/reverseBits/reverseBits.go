package main

import "github.com/sh1yu/assertion"

func reverseBits(num uint32) uint32 {
	// 奇偶位交换
	num = (num&0x55555555)<<1 | (num&0xaaaaaaaa)>>1
	// 每2位交换
	num = (num&0x33333333)<<2 | (num&0xcccccccc)>>2
	// 每4位交换
	num = (num&0x0f0f0f0f)<<4 | (num&0xf0f0f0f0)>>4
	// 每8位交换
	num = (num&0x00ff00ff)<<8 | (num&0xff00ff00)>>8
	// 每16位交换
	num = (num&0x0000ffff)<<16 | (num&0xffff0000)>>16

	return num
}

func main() {
	assertion.Equals(uint32(3221225471), reverseBits(0b11111111111111111111111111111101))
}
