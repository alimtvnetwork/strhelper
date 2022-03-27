package attrmeta

import (
	"gitlab.com/evatix-go/core/coreinterface"
	"gitlab.com/evatix-go/core/coreinterface/errcoreinf"
	"gitlab.com/evatix-go/core/coreinterface/loggerinf"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
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
