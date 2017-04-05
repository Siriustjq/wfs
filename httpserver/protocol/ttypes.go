// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package protocol

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

// 返回对象
//
// Attributes:
//  - Status: 状态
type WfsAck struct {
	Status *int32 `thrift:"status,1" json:"status,omitempty"`
}

func NewWfsAck() *WfsAck {
	return &WfsAck{}
}

var WfsAck_Status_DEFAULT int32

func (p *WfsAck) GetStatus() int32 {
	if !p.IsSetStatus() {
		return WfsAck_Status_DEFAULT
	}
	return *p.Status
}
func (p *WfsAck) IsSetStatus() bool {
	return p.Status != nil
}

func (p *WfsAck) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *WfsAck) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Status = &v
	}
	return nil
}

func (p *WfsAck) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("WfsAck"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *WfsAck) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetStatus() {
		if err := oprot.WriteFieldBegin("status", thrift.I32, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:status: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.Status)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.status (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:status: ", p), err)
		}
	}
	return err
}

func (p *WfsAck) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("WfsAck(%+v)", *p)
}

// 返回对象
//
// Attributes:
//  - Status: 状态
type Wfs struct {
	Status *int32 `thrift:"status,1" json:"status,omitempty"`
}

func NewWfs() *Wfs {
	return &Wfs{}
}

var Wfs_Status_DEFAULT int32

func (p *Wfs) GetStatus() int32 {
	if !p.IsSetStatus() {
		return Wfs_Status_DEFAULT
	}
	return *p.Status
}
func (p *Wfs) IsSetStatus() bool {
	return p.Status != nil
}

func (p *Wfs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *Wfs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Status = &v
	}
	return nil
}

func (p *Wfs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Wfs"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Wfs) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetStatus() {
		if err := oprot.WriteFieldBegin("status", thrift.I32, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:status: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.Status)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.status (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:status: ", p), err)
		}
	}
	return err
}

func (p *Wfs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Wfs(%+v)", *p)
}

// 文件对象
//
// Attributes:
//  - Name: 名称
//  - FileBody: 对象
//  - FileType: 类型
type WfsFile struct {
	Name     *string `thrift:"name,1" json:"name,omitempty"`
	FileBody []byte  `thrift:"fileBody,2" json:"fileBody,omitempty"`
	FileType *string `thrift:"fileType,3" json:"fileType,omitempty"`
}

func NewWfsFile() *WfsFile {
	return &WfsFile{}
}

var WfsFile_Name_DEFAULT string

func (p *WfsFile) GetName() string {
	if !p.IsSetName() {
		return WfsFile_Name_DEFAULT
	}
	return *p.Name
}

var WfsFile_FileBody_DEFAULT []byte

func (p *WfsFile) GetFileBody() []byte {
	return p.FileBody
}

var WfsFile_FileType_DEFAULT string

func (p *WfsFile) GetFileType() string {
	if !p.IsSetFileType() {
		return WfsFile_FileType_DEFAULT
	}
	return *p.FileType
}
func (p *WfsFile) IsSetName() bool {
	return p.Name != nil
}

func (p *WfsFile) IsSetFileBody() bool {
	return p.FileBody != nil
}

func (p *WfsFile) IsSetFileType() bool {
	return p.FileType != nil
}

func (p *WfsFile) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *WfsFile) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Name = &v
	}
	return nil
}

func (p *WfsFile) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.FileBody = v
	}
	return nil
}

func (p *WfsFile) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.FileType = &v
	}
	return nil
}

func (p *WfsFile) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("WfsFile"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *WfsFile) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetName() {
		if err := oprot.WriteFieldBegin("name", thrift.STRING, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:name: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Name)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.name (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:name: ", p), err)
		}
	}
	return err
}

func (p *WfsFile) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetFileBody() {
		if err := oprot.WriteFieldBegin("fileBody", thrift.STRING, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:fileBody: ", p), err)
		}
		if err := oprot.WriteBinary(p.FileBody); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.fileBody (2) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:fileBody: ", p), err)
		}
	}
	return err
}

func (p *WfsFile) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetFileType() {
		if err := oprot.WriteFieldBegin("fileType", thrift.STRING, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:fileType: ", p), err)
		}
		if err := oprot.WriteString(string(*p.FileType)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.fileType (3) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:fileType: ", p), err)
		}
	}
	return err
}

func (p *WfsFile) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("WfsFile(%+v)", *p)
}