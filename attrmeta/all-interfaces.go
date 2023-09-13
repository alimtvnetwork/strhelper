package attrmeta

import (
	"gitlab.com/auk-go/core/coreinterface"
	"gitlab.com/auk-go/core/coreinterface/errcoreinf"
	"gitlab.com/auk-go/core/coreinterface/loggerinf"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/errorwrapper/errwrappers"
)

type MetaAttributesCollector interface {
	loggerinf.MetaAttributesStacker
	coreinterface.LengthGetter
	coreinterface.BasicSlicerContractsBinder

	Clone() MetaAttributesCollector
	errcoreinf.CompiledVoidLogger

	CompiledAsErrorWrapper(
		variant errtype.Variation,
		finalMessage string,
	) *errorwrapper.Wrapper

	CompiledInjectToErrorCollection(
		errCollection *errwrappers.Collection,
		variant errtype.Variation,
		finalMessage string,
	) *errwrappers.Collection

	AsCollection() *Collection
}
