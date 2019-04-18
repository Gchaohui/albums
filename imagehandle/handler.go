// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package imagehandle

import (
	"bytes"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type Handler interface {
	// Parameters:
	//  - Image
	Feature(image []byte) (r []float64, err error)
	// Parameters:
	//  - Image
	DeepLearning(image []byte) (r *Result_, err error)
	// Parameters:
	//  - Image
	ObjectDetectionDL(image []byte) (r *Result_, err error)
}

type HandlerClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewHandlerClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *HandlerClient {
	return &HandlerClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewHandlerClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *HandlerClient {
	return &HandlerClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - Image
func (p *HandlerClient) Feature(image []byte) (r []float64, err error) {
	if err = p.sendFeature(image); err != nil {
		return
	}
	return p.recvFeature()
}

func (p *HandlerClient) sendFeature(image []byte) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("Feature", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := FeatureArgs{
		Image: image,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *HandlerClient) recvFeature() (value []float64, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error1 error
		error1, err = error0.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error1
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "Feature failed: out of sequence response")
		return
	}
	result := FeatureResult{}
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
//  - Image
func (p *HandlerClient) DeepLearning(image []byte) (r *Result_, err error) {
	if err = p.sendDeepLearning(image); err != nil {
		return
	}
	return p.recvDeepLearning()
}

func (p *HandlerClient) sendDeepLearning(image []byte) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("DeepLearning", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := DeepLearningArgs{
		Image: image,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *HandlerClient) recvDeepLearning() (value *Result_, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error3 error
		error3, err = error2.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error3
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "DeepLearning failed: out of sequence response")
		return
	}
	result := DeepLearningResult{}
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
//  - Image
func (p *HandlerClient) ObjectDetectionDL(image []byte) (r *Result_, err error) {
	if err = p.sendObjectDetectionDL(image); err != nil {
		return
	}
	return p.recvObjectDetectionDL()
}

func (p *HandlerClient) sendObjectDetectionDL(image []byte) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("ObjectDetectionDL", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := ObjectDetectionDLArgs{
		Image: image,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *HandlerClient) recvObjectDetectionDL() (value *Result_, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error4 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error5 error
		error5, err = error4.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error5
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ObjectDetectionDL failed: out of sequence response")
		return
	}
	result := ObjectDetectionDLResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type HandlerProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      Handler
}

func (p *HandlerProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *HandlerProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *HandlerProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewHandlerProcessor(handler Handler) *HandlerProcessor {

	self6 := &HandlerProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self6.processorMap["Feature"] = &handlerProcessorFeature{handler: handler}
	self6.processorMap["DeepLearning"] = &handlerProcessorDeepLearning{handler: handler}
	self6.processorMap["ObjectDetectionDL"] = &handlerProcessorObjectDetectionDL{handler: handler}
	return self6
}

func (p *HandlerProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x7 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x7.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x7

}

type handlerProcessorFeature struct {
	handler Handler
}

func (p *handlerProcessorFeature) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := FeatureArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("Feature", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := FeatureResult{}
	var retval []float64
	var err2 error
	if retval, err2 = p.handler.Feature(args.Image); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Feature: "+err2.Error())
		oprot.WriteMessageBegin("Feature", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("Feature", thrift.REPLY, seqId); err2 != nil {
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

type handlerProcessorDeepLearning struct {
	handler Handler
}

func (p *handlerProcessorDeepLearning) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := DeepLearningArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("DeepLearning", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := DeepLearningResult{}
	var retval *Result_
	var err2 error
	if retval, err2 = p.handler.DeepLearning(args.Image); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing DeepLearning: "+err2.Error())
		oprot.WriteMessageBegin("DeepLearning", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("DeepLearning", thrift.REPLY, seqId); err2 != nil {
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

type handlerProcessorObjectDetectionDL struct {
	handler Handler
}

func (p *handlerProcessorObjectDetectionDL) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ObjectDetectionDLArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("ObjectDetectionDL", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := ObjectDetectionDLResult{}
	var retval *Result_
	var err2 error
	if retval, err2 = p.handler.ObjectDetectionDL(args.Image); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing ObjectDetectionDL: "+err2.Error())
		oprot.WriteMessageBegin("ObjectDetectionDL", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("ObjectDetectionDL", thrift.REPLY, seqId); err2 != nil {
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

type FeatureArgs struct {
	Image []byte `thrift:"image,1" json:"image"`
}

func NewFeatureArgs() *FeatureArgs {
	return &FeatureArgs{}
}

func (p *FeatureArgs) GetImage() []byte {
	return p.Image
}
func (p *FeatureArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *FeatureArgs) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Image = v
	}
	return nil
}

func (p *FeatureArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Feature_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *FeatureArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("image", thrift.STRING, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:image: %s", p, err)
	}
	if err := oprot.WriteBinary(p.Image); err != nil {
		return fmt.Errorf("%T.image (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:image: %s", p, err)
	}
	return err
}

func (p *FeatureArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FeatureArgs(%+v)", *p)
}

type FeatureResult struct {
	Success []float64 `thrift:"success,0" json:"success"`
}

func NewFeatureResult() *FeatureResult {
	return &FeatureResult{}
}

var FeatureResult_Success_DEFAULT []float64

func (p *FeatureResult) GetSuccess() []float64 {
	return p.Success
}
func (p *FeatureResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FeatureResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *FeatureResult) ReadField0(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list begin: %s", err)
	}
	tSlice := make([]float64, 0, size)
	p.Success = tSlice
	for i := 0; i < size; i++ {
		var _elem8 float64
		if v, err := iprot.ReadDouble(); err != nil {
			return fmt.Errorf("error reading field 0: %s", err)
		} else {
			_elem8 = v
		}
		p.Success = append(p.Success, _elem8)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s", err)
	}
	return nil
}

func (p *FeatureResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Feature_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *FeatureResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.LIST, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := oprot.WriteListBegin(thrift.DOUBLE, len(p.Success)); err != nil {
			return fmt.Errorf("error writing list begin: %s", err)
		}
		for _, v := range p.Success {
			if err := oprot.WriteDouble(float64(v)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p, err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return fmt.Errorf("error writing list end: %s", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *FeatureResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FeatureResult(%+v)", *p)
}

type DeepLearningArgs struct {
	Image []byte `thrift:"image,1" json:"image"`
}

func NewDeepLearningArgs() *DeepLearningArgs {
	return &DeepLearningArgs{}
}

func (p *DeepLearningArgs) GetImage() []byte {
	return p.Image
}
func (p *DeepLearningArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *DeepLearningArgs) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Image = v
	}
	return nil
}

func (p *DeepLearningArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("DeepLearning_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *DeepLearningArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("image", thrift.STRING, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:image: %s", p, err)
	}
	if err := oprot.WriteBinary(p.Image); err != nil {
		return fmt.Errorf("%T.image (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:image: %s", p, err)
	}
	return err
}

func (p *DeepLearningArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DeepLearningArgs(%+v)", *p)
}

type DeepLearningResult struct {
	Success *Result_ `thrift:"success,0" json:"success"`
}

func NewDeepLearningResult() *DeepLearningResult {
	return &DeepLearningResult{}
}

var DeepLearningResult_Success_DEFAULT *Result_

func (p *DeepLearningResult) GetSuccess() *Result_ {
	if !p.IsSetSuccess() {
		return DeepLearningResult_Success_DEFAULT
	}
	return p.Success
}
func (p *DeepLearningResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeepLearningResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *DeepLearningResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &Result_{}
	if err := p.Success.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Success, err)
	}
	return nil
}

func (p *DeepLearningResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("DeepLearning_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *DeepLearningResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Success, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *DeepLearningResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DeepLearningResult(%+v)", *p)
}

type ObjectDetectionDLArgs struct {
	Image []byte `thrift:"image,1" json:"image"`
}

func NewObjectDetectionDLArgs() *ObjectDetectionDLArgs {
	return &ObjectDetectionDLArgs{}
}

func (p *ObjectDetectionDLArgs) GetImage() []byte {
	return p.Image
}
func (p *ObjectDetectionDLArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *ObjectDetectionDLArgs) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Image = v
	}
	return nil
}

func (p *ObjectDetectionDLArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ObjectDetectionDL_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *ObjectDetectionDLArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("image", thrift.STRING, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:image: %s", p, err)
	}
	if err := oprot.WriteBinary(p.Image); err != nil {
		return fmt.Errorf("%T.image (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:image: %s", p, err)
	}
	return err
}

func (p *ObjectDetectionDLArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ObjectDetectionDLArgs(%+v)", *p)
}

type ObjectDetectionDLResult struct {
	Success *Result_ `thrift:"success,0" json:"success"`
}

func NewObjectDetectionDLResult() *ObjectDetectionDLResult {
	return &ObjectDetectionDLResult{}
}

var ObjectDetectionDLResult_Success_DEFAULT *Result_

func (p *ObjectDetectionDLResult) GetSuccess() *Result_ {
	if !p.IsSetSuccess() {
		return ObjectDetectionDLResult_Success_DEFAULT
	}
	return p.Success
}
func (p *ObjectDetectionDLResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ObjectDetectionDLResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *ObjectDetectionDLResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &Result_{}
	if err := p.Success.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Success, err)
	}
	return nil
}

func (p *ObjectDetectionDLResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ObjectDetectionDL_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *ObjectDetectionDLResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Success, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *ObjectDetectionDLResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ObjectDetectionDLResult(%+v)", *p)
}