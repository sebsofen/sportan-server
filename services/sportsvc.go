// Autogenerated by Thrift Compiler (1.0.0-dev)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package services

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type SportSvc interface {
	// Parameters:
	//  - Token
	//  - Sport
	CreateSport(token string, sport *Sport) (r *Sport, err error)
	// Parameters:
	//  - Bla
	GetAllSports(bla string) (r []*Sport, err error)
	// Parameters:
	//  - Token
	//  - Sportid
	GetSportById(token string, sportid string) (r *Sport, err error)
}

type SportSvcClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewSportSvcClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *SportSvcClient {
	return &SportSvcClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewSportSvcClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *SportSvcClient {
	return &SportSvcClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - Token
//  - Sport
func (p *SportSvcClient) CreateSport(token string, sport *Sport) (r *Sport, err error) {
	if err = p.sendCreateSport(token, sport); err != nil {
		return
	}
	return p.recvCreateSport()
}

func (p *SportSvcClient) sendCreateSport(token string, sport *Sport) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("createSport", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := SportSvcCreateSportArgs{
		Token: token,
		Sport: sport,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *SportSvcClient) recvCreateSport() (value *Sport, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "createSport" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "createSport failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "createSport failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error31 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error32 error
		error32, err = error31.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error32
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "createSport failed: invalid message type")
		return
	}
	result := SportSvcCreateSportResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - Bla
func (p *SportSvcClient) GetAllSports(bla string) (r []*Sport, err error) {
	if err = p.sendGetAllSports(bla); err != nil {
		return
	}
	return p.recvGetAllSports()
}

func (p *SportSvcClient) sendGetAllSports(bla string) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("getAllSports", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := SportSvcGetAllSportsArgs{
		Bla: bla,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *SportSvcClient) recvGetAllSports() (value []*Sport, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "getAllSports" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "getAllSports failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getAllSports failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error33 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error34 error
		error34, err = error33.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error34
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "getAllSports failed: invalid message type")
		return
	}
	result := SportSvcGetAllSportsResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - Token
//  - Sportid
func (p *SportSvcClient) GetSportById(token string, sportid string) (r *Sport, err error) {
	if err = p.sendGetSportById(token, sportid); err != nil {
		return
	}
	return p.recvGetSportById()
}

func (p *SportSvcClient) sendGetSportById(token string, sportid string) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("getSportById", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := SportSvcGetSportByIdArgs{
		Token:   token,
		Sportid: sportid,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *SportSvcClient) recvGetSportById() (value *Sport, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "getSportById" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "getSportById failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getSportById failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error35 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error36 error
		error36, err = error35.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error36
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "getSportById failed: invalid message type")
		return
	}
	result := SportSvcGetSportByIdResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type SportSvcProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      SportSvc
}

func (p *SportSvcProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *SportSvcProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *SportSvcProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewSportSvcProcessor(handler SportSvc) *SportSvcProcessor {

	self37 := &SportSvcProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self37.processorMap["createSport"] = &sportSvcProcessorCreateSport{handler: handler}
	self37.processorMap["getAllSports"] = &sportSvcProcessorGetAllSports{handler: handler}
	self37.processorMap["getSportById"] = &sportSvcProcessorGetSportById{handler: handler}
	return self37
}

func (p *SportSvcProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x38 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x38.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x38

}

type sportSvcProcessorCreateSport struct {
	handler SportSvc
}

func (p *sportSvcProcessorCreateSport) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := SportSvcCreateSportArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("createSport", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := SportSvcCreateSportResult{}
	var retval *Sport
	var err2 error
	if retval, err2 = p.handler.CreateSport(args.Token, args.Sport); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing createSport: "+err2.Error())
		oprot.WriteMessageBegin("createSport", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("createSport", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type sportSvcProcessorGetAllSports struct {
	handler SportSvc
}

func (p *sportSvcProcessorGetAllSports) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := SportSvcGetAllSportsArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getAllSports", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := SportSvcGetAllSportsResult{}
	var retval []*Sport
	var err2 error
	if retval, err2 = p.handler.GetAllSports(args.Bla); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getAllSports: "+err2.Error())
		oprot.WriteMessageBegin("getAllSports", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getAllSports", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type sportSvcProcessorGetSportById struct {
	handler SportSvc
}

func (p *sportSvcProcessorGetSportById) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := SportSvcGetSportByIdArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getSportById", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := SportSvcGetSportByIdResult{}
	var retval *Sport
	var err2 error
	if retval, err2 = p.handler.GetSportById(args.Token, args.Sportid); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getSportById: "+err2.Error())
		oprot.WriteMessageBegin("getSportById", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("getSportById", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Token
//  - Sport
type SportSvcCreateSportArgs struct {
	Token string `thrift:"token,1" db:"token" json:"token"`
	Sport *Sport `thrift:"sport,2" db:"sport" json:"sport"`
}

func NewSportSvcCreateSportArgs() *SportSvcCreateSportArgs {
	return &SportSvcCreateSportArgs{}
}

func (p *SportSvcCreateSportArgs) GetToken() string {
	return p.Token
}

var SportSvcCreateSportArgs_Sport_DEFAULT *Sport

func (p *SportSvcCreateSportArgs) GetSport() *Sport {
	if !p.IsSetSport() {
		return SportSvcCreateSportArgs_Sport_DEFAULT
	}
	return p.Sport
}
func (p *SportSvcCreateSportArgs) IsSetSport() bool {
	return p.Sport != nil
}

func (p *SportSvcCreateSportArgs) Read(iprot thrift.TProtocol) error {
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
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
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

func (p *SportSvcCreateSportArgs) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Token = v
	}
	return nil
}

func (p *SportSvcCreateSportArgs) ReadField2(iprot thrift.TProtocol) error {
	p.Sport = &Sport{}
	if err := p.Sport.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Sport), err)
	}
	return nil
}

func (p *SportSvcCreateSportArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("createSport_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
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

func (p *SportSvcCreateSportArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("token", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:token: ", p), err)
	}
	if err := oprot.WriteString(string(p.Token)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.token (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:token: ", p), err)
	}
	return err
}

func (p *SportSvcCreateSportArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("sport", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:sport: ", p), err)
	}
	if err := p.Sport.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Sport), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:sport: ", p), err)
	}
	return err
}

func (p *SportSvcCreateSportArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SportSvcCreateSportArgs(%+v)", *p)
}

// Attributes:
//  - Success
type SportSvcCreateSportResult struct {
	Success *Sport `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewSportSvcCreateSportResult() *SportSvcCreateSportResult {
	return &SportSvcCreateSportResult{}
}

var SportSvcCreateSportResult_Success_DEFAULT *Sport

func (p *SportSvcCreateSportResult) GetSuccess() *Sport {
	if !p.IsSetSuccess() {
		return SportSvcCreateSportResult_Success_DEFAULT
	}
	return p.Success
}
func (p *SportSvcCreateSportResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SportSvcCreateSportResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.ReadField0(iprot); err != nil {
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

func (p *SportSvcCreateSportResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &Sport{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *SportSvcCreateSportResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("createSport_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
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

func (p *SportSvcCreateSportResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *SportSvcCreateSportResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SportSvcCreateSportResult(%+v)", *p)
}

// Attributes:
//  - Bla
type SportSvcGetAllSportsArgs struct {
	Bla string `thrift:"bla,1" db:"bla" json:"bla"`
}

func NewSportSvcGetAllSportsArgs() *SportSvcGetAllSportsArgs {
	return &SportSvcGetAllSportsArgs{}
}

func (p *SportSvcGetAllSportsArgs) GetBla() string {
	return p.Bla
}
func (p *SportSvcGetAllSportsArgs) Read(iprot thrift.TProtocol) error {
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
			if err := p.ReadField1(iprot); err != nil {
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

func (p *SportSvcGetAllSportsArgs) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Bla = v
	}
	return nil
}

func (p *SportSvcGetAllSportsArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getAllSports_args"); err != nil {
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

func (p *SportSvcGetAllSportsArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("bla", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:bla: ", p), err)
	}
	if err := oprot.WriteString(string(p.Bla)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.bla (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:bla: ", p), err)
	}
	return err
}

func (p *SportSvcGetAllSportsArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SportSvcGetAllSportsArgs(%+v)", *p)
}

// Attributes:
//  - Success
type SportSvcGetAllSportsResult struct {
	Success []*Sport `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewSportSvcGetAllSportsResult() *SportSvcGetAllSportsResult {
	return &SportSvcGetAllSportsResult{}
}

var SportSvcGetAllSportsResult_Success_DEFAULT []*Sport

func (p *SportSvcGetAllSportsResult) GetSuccess() []*Sport {
	return p.Success
}
func (p *SportSvcGetAllSportsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SportSvcGetAllSportsResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.ReadField0(iprot); err != nil {
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

func (p *SportSvcGetAllSportsResult) ReadField0(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*Sport, 0, size)
	p.Success = tSlice
	for i := 0; i < size; i++ {
		_elem39 := &Sport{}
		if err := _elem39.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem39), err)
		}
		p.Success = append(p.Success, _elem39)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *SportSvcGetAllSportsResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getAllSports_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
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

func (p *SportSvcGetAllSportsResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.LIST, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Success)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Success {
			if err := v.Write(oprot); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *SportSvcGetAllSportsResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SportSvcGetAllSportsResult(%+v)", *p)
}

// Attributes:
//  - Token
//  - Sportid
type SportSvcGetSportByIdArgs struct {
	Token   string `thrift:"token,1" db:"token" json:"token"`
	Sportid string `thrift:"sportid,2" db:"sportid" json:"sportid"`
}

func NewSportSvcGetSportByIdArgs() *SportSvcGetSportByIdArgs {
	return &SportSvcGetSportByIdArgs{}
}

func (p *SportSvcGetSportByIdArgs) GetToken() string {
	return p.Token
}

func (p *SportSvcGetSportByIdArgs) GetSportid() string {
	return p.Sportid
}
func (p *SportSvcGetSportByIdArgs) Read(iprot thrift.TProtocol) error {
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
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
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

func (p *SportSvcGetSportByIdArgs) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Token = v
	}
	return nil
}

func (p *SportSvcGetSportByIdArgs) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Sportid = v
	}
	return nil
}

func (p *SportSvcGetSportByIdArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getSportById_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
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

func (p *SportSvcGetSportByIdArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("token", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:token: ", p), err)
	}
	if err := oprot.WriteString(string(p.Token)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.token (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:token: ", p), err)
	}
	return err
}

func (p *SportSvcGetSportByIdArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("sportid", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:sportid: ", p), err)
	}
	if err := oprot.WriteString(string(p.Sportid)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.sportid (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:sportid: ", p), err)
	}
	return err
}

func (p *SportSvcGetSportByIdArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SportSvcGetSportByIdArgs(%+v)", *p)
}

// Attributes:
//  - Success
type SportSvcGetSportByIdResult struct {
	Success *Sport `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewSportSvcGetSportByIdResult() *SportSvcGetSportByIdResult {
	return &SportSvcGetSportByIdResult{}
}

var SportSvcGetSportByIdResult_Success_DEFAULT *Sport

func (p *SportSvcGetSportByIdResult) GetSuccess() *Sport {
	if !p.IsSetSuccess() {
		return SportSvcGetSportByIdResult_Success_DEFAULT
	}
	return p.Success
}
func (p *SportSvcGetSportByIdResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SportSvcGetSportByIdResult) Read(iprot thrift.TProtocol) error {
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
		case 0:
			if err := p.ReadField0(iprot); err != nil {
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

func (p *SportSvcGetSportByIdResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &Sport{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *SportSvcGetSportByIdResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getSportById_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
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

func (p *SportSvcGetSportByIdResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *SportSvcGetSportByIdResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SportSvcGetSportByIdResult(%+v)", *p)
}
