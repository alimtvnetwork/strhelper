package attrmeta

import (
	"gitlab.com/evatix-go/core/coreinterface"
	"gitlab.com/evatix-go/core/coreinterface/errcoreinf"
	"gitlab.com/evatix-go/core/coreinterface/loggerinf"
)

type MetaAttributesCollector interface {
	loggerinf.MetaAttributesStacker
	coreinterface.LengthGetter
	coreinterface.BasicSlicerContractsBinder

	Clone() MetaAttributesCollector
	errcoreinf.CompiledVoidLogger

	AsCollection() *Collection
}
