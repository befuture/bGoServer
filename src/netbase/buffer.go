package netbase
/*
// Website: http://www.befuture.net/
// Project: http://www.github.com/befuture/GoServer/
//
// Copyright 2013 CrBeE(Jacky Lin). All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
*/

import (
	"bytes"
	"encoding/binary"
)

type Buffer struct {
	Buffer *bytes.Buffer
}

func CreateBuffer() *Buffer {
	var buffer *Buffer
	buffer = new(Buffer)
	return buffer
}

func (self *Buffer) Init() {
	self.Buffer = new(bytes.Buffer)
}

func (self *Buffer) PushData(buff []byte) {
	binary.Write(self.Buffer, binary.LittleEndian, buff)
}

/*
	function: return the Bytes in the buffer, no cutting of orginal one.
	return: []byte <- the bytes in buffer with the type []byte.
			error <- always nil
*/
func (self *Buffer) Bytes() ([]byte, error) {
	return self.Buffer.Bytes(), nil
}

func (self *Buffer) ReadBytes(length int) ([]byte, error) {
	result := make([]byte, length)
	_, err := self.Buffer.Read(result)
	return result, err
}

func (self *Buffer) ReadFloat() (float32, error) {
	var result float32
	err := binary.Read(self.Buffer, binary.LittleEndian, &result)
	return result, err
}

func (self *Buffer) ReadDouble() (float64, error) {
	var result float64
	err := binary.Read(self.Buffer, binary.LittleEndian, &result)
	return result, err
}

func (self *Buffer) ReadInt32() (int32, error) {
	var result int32
	err := binary.Read(self.Buffer, binary.LittleEndian, &result)
	return result, err
}

func (self *Buffer) ReadInt16() (int16, error) {
	var result int16
	err := binary.Read(self.Buffer, binary.LittleEndian, &result)
	return result, err
}

func (self *Buffer) ReadDword() (uint32, error) {
	var result uint32
	err := binary.Read(self.Buffer, binary.LittleEndian, &result)
	return result, err
}

func (self *Buffer) ReadWord() (uint16, error) {
	var result uint16
	err := binary.Read(self.Buffer, binary.LittleEndian, &result)
	return result, err
}

func (self *Buffer) ReadByte() (byte, error) {
	var result byte
	err := binary.Read(self.Buffer, binary.LittleEndian, &result)
	return result, err
}

/*
	Not available 
*/
func (self *Buffer) ReadAnsiString(length int) (string, error) {
	return "", nil
}

/*
	Not available 
*/
func (self *Buffer) ReadUTF8String(length int) (string, error) {
	return "", nil
}