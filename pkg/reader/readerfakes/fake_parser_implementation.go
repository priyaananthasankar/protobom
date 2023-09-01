// Code generated by counterfeiter. DO NOT EDIT.
package readerfakes

import (
	"io"
	"os"
	"sync"

	"github.com/bom-squad/protobom/pkg/formats"
	"github.com/bom-squad/protobom/pkg/native"
	"github.com/bom-squad/protobom/pkg/reader/options"
	"github.com/bom-squad/protobom/pkg/sbom"
)

type FakeParserImplementation struct {
	DetectFormatStub        func(*options.Options, io.ReadSeeker) (formats.Format, error)
	detectFormatMutex       sync.RWMutex
	detectFormatArgsForCall []struct {
		arg1 *options.Options
		arg2 io.ReadSeeker
	}
	detectFormatReturns struct {
		result1 formats.Format
		result2 error
	}
	detectFormatReturnsOnCall map[int]struct {
		result1 formats.Format
		result2 error
	}
	GetUnserializerStub        func(*options.Options, formats.Format) (native.Unserializer, error)
	getUnserializerMutex       sync.RWMutex
	getUnserializerArgsForCall []struct {
		arg1 *options.Options
		arg2 formats.Format
	}
	getUnserializerReturns struct {
		result1 native.Unserializer
		result2 error
	}
	getUnserializerReturnsOnCall map[int]struct {
		result1 native.Unserializer
		result2 error
	}
	OpenDocumentFileStub        func(string) (*os.File, error)
	openDocumentFileMutex       sync.RWMutex
	openDocumentFileArgsForCall []struct {
		arg1 string
	}
	openDocumentFileReturns struct {
		result1 *os.File
		result2 error
	}
	openDocumentFileReturnsOnCall map[int]struct {
		result1 *os.File
		result2 error
	}
	ParseStreamStub        func(native.Unserializer, *options.Options, io.Reader) (*sbom.Document, error)
	parseStreamMutex       sync.RWMutex
	parseStreamArgsForCall []struct {
		arg1 native.Unserializer
		arg2 *options.Options
		arg3 io.Reader
	}
	parseStreamReturns struct {
		result1 *sbom.Document
		result2 error
	}
	parseStreamReturnsOnCall map[int]struct {
		result1 *sbom.Document
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeParserImplementation) DetectFormat(arg1 *options.Options, arg2 io.ReadSeeker) (formats.Format, error) {
	fake.detectFormatMutex.Lock()
	ret, specificReturn := fake.detectFormatReturnsOnCall[len(fake.detectFormatArgsForCall)]
	fake.detectFormatArgsForCall = append(fake.detectFormatArgsForCall, struct {
		arg1 *options.Options
		arg2 io.ReadSeeker
	}{arg1, arg2})
	stub := fake.DetectFormatStub
	fakeReturns := fake.detectFormatReturns
	fake.recordInvocation("DetectFormat", []interface{}{arg1, arg2})
	fake.detectFormatMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeParserImplementation) DetectFormatCallCount() int {
	fake.detectFormatMutex.RLock()
	defer fake.detectFormatMutex.RUnlock()
	return len(fake.detectFormatArgsForCall)
}

func (fake *FakeParserImplementation) DetectFormatCalls(stub func(*options.Options, io.ReadSeeker) (formats.Format, error)) {
	fake.detectFormatMutex.Lock()
	defer fake.detectFormatMutex.Unlock()
	fake.DetectFormatStub = stub
}

func (fake *FakeParserImplementation) DetectFormatArgsForCall(i int) (*options.Options, io.ReadSeeker) {
	fake.detectFormatMutex.RLock()
	defer fake.detectFormatMutex.RUnlock()
	argsForCall := fake.detectFormatArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeParserImplementation) DetectFormatReturns(result1 formats.Format, result2 error) {
	fake.detectFormatMutex.Lock()
	defer fake.detectFormatMutex.Unlock()
	fake.DetectFormatStub = nil
	fake.detectFormatReturns = struct {
		result1 formats.Format
		result2 error
	}{result1, result2}
}

func (fake *FakeParserImplementation) DetectFormatReturnsOnCall(i int, result1 formats.Format, result2 error) {
	fake.detectFormatMutex.Lock()
	defer fake.detectFormatMutex.Unlock()
	fake.DetectFormatStub = nil
	if fake.detectFormatReturnsOnCall == nil {
		fake.detectFormatReturnsOnCall = make(map[int]struct {
			result1 formats.Format
			result2 error
		})
	}
	fake.detectFormatReturnsOnCall[i] = struct {
		result1 formats.Format
		result2 error
	}{result1, result2}
}

func (fake *FakeParserImplementation) GetUnserializer(arg1 *options.Options, arg2 formats.Format) (native.Unserializer, error) {
	fake.getUnserializerMutex.Lock()
	ret, specificReturn := fake.getUnserializerReturnsOnCall[len(fake.getUnserializerArgsForCall)]
	fake.getUnserializerArgsForCall = append(fake.getUnserializerArgsForCall, struct {
		arg1 *options.Options
		arg2 formats.Format
	}{arg1, arg2})
	stub := fake.GetUnserializerStub
	fakeReturns := fake.getUnserializerReturns
	fake.recordInvocation("GetUnserializer", []interface{}{arg1, arg2})
	fake.getUnserializerMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeParserImplementation) GetUnserializerCallCount() int {
	fake.getUnserializerMutex.RLock()
	defer fake.getUnserializerMutex.RUnlock()
	return len(fake.getUnserializerArgsForCall)
}

func (fake *FakeParserImplementation) GetUnserializerCalls(stub func(*options.Options, formats.Format) (native.Unserializer, error)) {
	fake.getUnserializerMutex.Lock()
	defer fake.getUnserializerMutex.Unlock()
	fake.GetUnserializerStub = stub
}

func (fake *FakeParserImplementation) GetUnserializerArgsForCall(i int) (*options.Options, formats.Format) {
	fake.getUnserializerMutex.RLock()
	defer fake.getUnserializerMutex.RUnlock()
	argsForCall := fake.getUnserializerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeParserImplementation) GetUnserializerReturns(result1 native.Unserializer, result2 error) {
	fake.getUnserializerMutex.Lock()
	defer fake.getUnserializerMutex.Unlock()
	fake.GetUnserializerStub = nil
	fake.getUnserializerReturns = struct {
		result1 native.Unserializer
		result2 error
	}{result1, result2}
}

func (fake *FakeParserImplementation) GetUnserializerReturnsOnCall(i int, result1 native.Unserializer, result2 error) {
	fake.getUnserializerMutex.Lock()
	defer fake.getUnserializerMutex.Unlock()
	fake.GetUnserializerStub = nil
	if fake.getUnserializerReturnsOnCall == nil {
		fake.getUnserializerReturnsOnCall = make(map[int]struct {
			result1 native.Unserializer
			result2 error
		})
	}
	fake.getUnserializerReturnsOnCall[i] = struct {
		result1 native.Unserializer
		result2 error
	}{result1, result2}
}

func (fake *FakeParserImplementation) OpenDocumentFile(arg1 string) (*os.File, error) {
	fake.openDocumentFileMutex.Lock()
	ret, specificReturn := fake.openDocumentFileReturnsOnCall[len(fake.openDocumentFileArgsForCall)]
	fake.openDocumentFileArgsForCall = append(fake.openDocumentFileArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.OpenDocumentFileStub
	fakeReturns := fake.openDocumentFileReturns
	fake.recordInvocation("OpenDocumentFile", []interface{}{arg1})
	fake.openDocumentFileMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeParserImplementation) OpenDocumentFileCallCount() int {
	fake.openDocumentFileMutex.RLock()
	defer fake.openDocumentFileMutex.RUnlock()
	return len(fake.openDocumentFileArgsForCall)
}

func (fake *FakeParserImplementation) OpenDocumentFileCalls(stub func(string) (*os.File, error)) {
	fake.openDocumentFileMutex.Lock()
	defer fake.openDocumentFileMutex.Unlock()
	fake.OpenDocumentFileStub = stub
}

func (fake *FakeParserImplementation) OpenDocumentFileArgsForCall(i int) string {
	fake.openDocumentFileMutex.RLock()
	defer fake.openDocumentFileMutex.RUnlock()
	argsForCall := fake.openDocumentFileArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeParserImplementation) OpenDocumentFileReturns(result1 *os.File, result2 error) {
	fake.openDocumentFileMutex.Lock()
	defer fake.openDocumentFileMutex.Unlock()
	fake.OpenDocumentFileStub = nil
	fake.openDocumentFileReturns = struct {
		result1 *os.File
		result2 error
	}{result1, result2}
}

func (fake *FakeParserImplementation) OpenDocumentFileReturnsOnCall(i int, result1 *os.File, result2 error) {
	fake.openDocumentFileMutex.Lock()
	defer fake.openDocumentFileMutex.Unlock()
	fake.OpenDocumentFileStub = nil
	if fake.openDocumentFileReturnsOnCall == nil {
		fake.openDocumentFileReturnsOnCall = make(map[int]struct {
			result1 *os.File
			result2 error
		})
	}
	fake.openDocumentFileReturnsOnCall[i] = struct {
		result1 *os.File
		result2 error
	}{result1, result2}
}

func (fake *FakeParserImplementation) ParseStream(arg1 native.Unserializer, arg2 *options.Options, arg3 io.Reader) (*sbom.Document, error) {
	fake.parseStreamMutex.Lock()
	ret, specificReturn := fake.parseStreamReturnsOnCall[len(fake.parseStreamArgsForCall)]
	fake.parseStreamArgsForCall = append(fake.parseStreamArgsForCall, struct {
		arg1 native.Unserializer
		arg2 *options.Options
		arg3 io.Reader
	}{arg1, arg2, arg3})
	stub := fake.ParseStreamStub
	fakeReturns := fake.parseStreamReturns
	fake.recordInvocation("ParseStream", []interface{}{arg1, arg2, arg3})
	fake.parseStreamMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeParserImplementation) ParseStreamCallCount() int {
	fake.parseStreamMutex.RLock()
	defer fake.parseStreamMutex.RUnlock()
	return len(fake.parseStreamArgsForCall)
}

func (fake *FakeParserImplementation) ParseStreamCalls(stub func(native.Unserializer, *options.Options, io.Reader) (*sbom.Document, error)) {
	fake.parseStreamMutex.Lock()
	defer fake.parseStreamMutex.Unlock()
	fake.ParseStreamStub = stub
}

func (fake *FakeParserImplementation) ParseStreamArgsForCall(i int) (native.Unserializer, *options.Options, io.Reader) {
	fake.parseStreamMutex.RLock()
	defer fake.parseStreamMutex.RUnlock()
	argsForCall := fake.parseStreamArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeParserImplementation) ParseStreamReturns(result1 *sbom.Document, result2 error) {
	fake.parseStreamMutex.Lock()
	defer fake.parseStreamMutex.Unlock()
	fake.ParseStreamStub = nil
	fake.parseStreamReturns = struct {
		result1 *sbom.Document
		result2 error
	}{result1, result2}
}

func (fake *FakeParserImplementation) ParseStreamReturnsOnCall(i int, result1 *sbom.Document, result2 error) {
	fake.parseStreamMutex.Lock()
	defer fake.parseStreamMutex.Unlock()
	fake.ParseStreamStub = nil
	if fake.parseStreamReturnsOnCall == nil {
		fake.parseStreamReturnsOnCall = make(map[int]struct {
			result1 *sbom.Document
			result2 error
		})
	}
	fake.parseStreamReturnsOnCall[i] = struct {
		result1 *sbom.Document
		result2 error
	}{result1, result2}
}

func (fake *FakeParserImplementation) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.detectFormatMutex.RLock()
	defer fake.detectFormatMutex.RUnlock()
	fake.getUnserializerMutex.RLock()
	defer fake.getUnserializerMutex.RUnlock()
	fake.openDocumentFileMutex.RLock()
	defer fake.openDocumentFileMutex.RUnlock()
	fake.parseStreamMutex.RLock()
	defer fake.parseStreamMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeParserImplementation) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
