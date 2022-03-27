package attrmeta

import (
	"fmt"

	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreinterface"
	"gitlab.com/evatix-go/core/coreinterface/entityinf"
	"gitlab.com/evatix-go/core/coreinterface/enuminf"
	"gitlab.com/evatix-go/core/coreinterface/errcoreinf"
	"gitlab.com/evatix-go/core/coreinterface/loggerinf"
	"gitlab.com/evatix-go/core/coreinterface/serializerinf"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/enum/logtype"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

type nopAttr struct {
	another *Collection
}

func (it *nopAttr) IsSilent() bool {
	return true
}

func (it *nopAttr) Items() map[string]interface{} {
	return nil
}

func (it *nopAttr) GetAsStrings() []string {
	return nil
}

func (it *nopAttr) GetVal(keyName string) (val interface{}) {
	return nil
}

func (it *nopAttr) RemoveAt(index int) (isSuccess bool) {
	panic(errcore.NotSupportedType.ErrorNoRefs("not supported for map type"))
}

func (it *nopAttr) ListStrings() []string {
	return it.GetAsStrings()
}

func (it *nopAttr) JsonString() string {
	return ""
}

func (it *nopAttr) JsonStringMust() string {
	jsonResult := it.JsonPtr()
	jsonResult.MustBeSafe()

	return jsonResult.JsonString()
}

func (it nopAttr) Json() corejson.Result {
	return corejson.Empty.Result()
}

func (it nopAttr) JsonPtr() *corejson.Result {
	return corejson.Empty.ResultPtr()
}

func (it *nopAttr) JsonParseSelfInject(jsonResult *corejson.Result) error {
	return errcore.
		NotSupportedType.
		ErrorNoRefs("for nop type")
}

func (it *nopAttr) LoggerTyper() enuminf.LoggerTyper {
	return logtype.Silent.AsLoggerTyper()
}

func (it *nopAttr) On(
	isLog bool,
) loggerinf.MetaAttributesStacker {
	if isLog {
		return it.another
	}

	return it
}

// OnTitle
//
//  not impl todo
func (it *nopAttr) OnTitle(
	isLog bool,
	title string,
) loggerinf.MetaAttributesWithoutTileStacker {
	panic("not implemented and supported")
}

func (it *nopAttr) Msg(
	message string,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) Title(
	title string,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) TitleAttr(
	title, attr string,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) Str(
	title, val string,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) NullKey(
	title string,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) StandardSlicer(
	title string,
	standardSlice coreinterface.StandardSlicer,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) Stringers(
	title string,
	stringers ...fmt.Stringer,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) Hex(
	title string, hexValues []byte,
) loggerinf.MetaAttributesStacker {
	return it.RawJson(
		title, hexValues)
}

func (it *nopAttr) RawJson(
	title string,
	rawJsonBytes []byte,
) loggerinf.MetaAttributesStacker {
	return it
}

// Error
//
//  skip on error
func (it *nopAttr) Error(
	title string,
	err error,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) SimpleBytesResulter(
	title string,
	result serializerinf.SimpleBytesResulter,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) BaseJsonResulter(
	title string,
	result serializerinf.BaseJsonResulter,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) BasicJsonResulter(
	title string,
	result serializerinf.BasicJsonResulter,
) loggerinf.MetaAttributesStacker {
	return it.BaseJsonResulter(
		title,
		result)
}

func (it *nopAttr) JsonResulter(
	title string,
	result serializerinf.JsonResulter,
) loggerinf.MetaAttributesStacker {
	return it.BaseJsonResulter(
		title,
		result)
}

func (it *nopAttr) MapIntegerAny(
	title string,
	mapAny map[int]interface{},
) loggerinf.MetaAttributesStacker {
	return it.JsonResulter(
		title,
		corejson.NewPtr(mapAny))
}

func (it *nopAttr) Meta(
	title string,
	metaAttr loggerinf.MetaAttributesCompiler,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) MapBool(
	title string,
	mapInt map[string]bool,
) loggerinf.MetaAttributesStacker {
	return it.JsonResulter(
		title,
		corejson.NewPtr(mapInt))
}

func (it *nopAttr) MapInt(
	title string,
	mapInt map[string]int,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) MapAnyAny(
	title string,
	mapAny map[interface{}]interface{},
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) MapAny(
	title string,
	mapAny map[string]interface{},
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) MapIntAny(
	title string, mapAny map[int]interface{},
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) MapIntString(
	title string, mapAny map[int]string,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) MapJsonResult(
	title string, mapAny map[string]corejson.Result,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) JsonResult(
	title string, json *corejson.Result,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) JsonResultItems(
	title string, jsons ...*corejson.Result,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) DefaultStackTraces() loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) ErrWithTypeTraces(
	title string,
	errType errcoreinf.BasicErrorTyper,
	err error,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) ErrorsWithTypeTraces(
	title string,
	errType errcoreinf.BasicErrorTyper,
	errorItems ...error,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnErrStackTraces(err error) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnErrWrapperOrCollectionStackTraces(
	errWrapperOrCollection errcoreinf.BaseErrorOrCollectionWrapper,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) FullStringer(title string, fullStringer errcoreinf.FullStringer) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyFullStringer(fullStringer errcoreinf.FullStringer) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) FullTraceAsAttr(
	title string, attrFullStringWithTraces errcoreinf.FullStringWithTracesGetter,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) BasicErrWrapper(
	errWrapperOrCollection errcoreinf.BasicErrWrapper,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) BaseRawErrCollectionDefiner(
	errWrapperOrCollection errcoreinf.BaseRawErrCollectionDefiner,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) BaseErrorWrapperCollectionDefiner(
	errWrapperCollection errcoreinf.BaseErrorWrapperCollectionDefiner,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) ErrWrapperOrCollection(
	errWrapperOrCollection errcoreinf.BaseErrorOrCollectionWrapper,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) CompiledBasicErrWrapper(
	compiler errcoreinf.CompiledBasicErrWrapper,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) Namer(title string, namer enuminf.Namer) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyNamer(namer enuminf.Namer) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) EnumTitleEnum(
	title enuminf.SimpleEnumer, enum enuminf.BasicEnumer,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) SimpleEnumTitleEnum(
	title enuminf.SimpleEnumer,
	enum enuminf.SimpleEnumer,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) Enum(
	title string,
	enum enuminf.BasicEnumer,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) Enums(key string, enums ...enuminf.BasicEnumer) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyEnum(enum enuminf.BasicEnumer) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyEnums(enums ...enuminf.BasicEnumer) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyString(value string) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyStrings(values ...string) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyStringer(stringer fmt.Stringer) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyStringers(stringers ...fmt.Stringer) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyIntegers(values ...int) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyBooleans(values ...bool) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyBytes(rawBytes []byte) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyRawJson(rawBytes []byte) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyBytesErr(rawBytes []byte, err error) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlySimpleBytesResulter(
	result serializerinf.SimpleBytesResulter,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyBaseJsonResulter(
	result serializerinf.BaseJsonResulter,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyBasicJsonResulter(result serializerinf.BasicJsonResulter) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyJsonResulter(result serializerinf.JsonResulter) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyAny(anyItem interface{}) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyAnyItems(values ...interface{}) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyMetaAttr(
	metaAttr loggerinf.MetaAttributesCompiler,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyAnyIf(isLog bool, anyItem interface{}) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyAnyItemsIf(
	isLog bool, anyItems ...interface{},
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyMapBool(mapItems map[string]bool) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyMapInt(mapInt map[string]int) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyMapAny(mapAny map[string]interface{}) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyMapIntAny(mapAny map[int]interface{}) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyMapIntString(mapAny map[int]string) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyMapJsonResult(mapAny map[string]corejson.Result) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyJson(json *corejson.Result) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyJsons(jsons ...*corejson.Result) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) Booleans(title string, isResults ...bool) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) Jsoner(
	jsoner corejson.Jsoner,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) Jsoners(
	jsoners ...corejson.Jsoner,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) JsonerTitle(
	title string,
	jsoner corejson.Jsoner,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) JsonerIf(
	isLog bool,
	jsoner corejson.Jsoner,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) JsonersIf(
	isLog bool,
	jsoners ...corejson.Jsoner,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) Serializer(
	serializer loggerinf.Serializer,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) Serializers(
	serializers ...loggerinf.Serializer,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) SerializerFunc(
	serializerFunc func() ([]byte, error),
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) SerializerFunctions(
	serializerFunctions ...func() ([]byte, error),
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) StandardTaskEntityDefiner(
	entity entityinf.StandardTaskEntityDefiner,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) TaskEntityDefiner(
	entity entityinf.TaskEntityDefiner,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) StandardTaskEntityDefinerTitle(
	title string,
	entity entityinf.StandardTaskEntityDefiner,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) TaskEntityDefinerTitle(
	title string,
	entity entityinf.TaskEntityDefiner,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) LoggerModel(
	loggerModel loggerinf.SingleLogModeler,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) LoggerModelTitle(
	title string,
	loggerModel loggerinf.SingleLogModeler,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) Integers(
	key string,
	integerItems ...int,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) Fmt(
	title,
	format string,
	v ...interface{},
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) FmtIf(
	isLog bool,
	title,
	format string,
	v ...interface{},
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyFmt(
	format string,
	v ...interface{},
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) OnlyFmtIf(
	isLog bool,
	format string,
	v ...interface{},
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) RawPayloadsGetter(
	payloadsGetter loggerinf.RawPayloadsGetter,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) RawPayloadsGetterTitle(
	title string,
	payloadsGetter loggerinf.RawPayloadsGetter,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) RawPayloadsGetterIf(
	isLog bool, payloadsGetter loggerinf.RawPayloadsGetter,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) Inject(
	others ...loggerinf.MetaAttributesStacker,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) InjectMap(
	compiledMap map[string]interface{},
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) ConcatNew(
	others ...loggerinf.MetaAttributesStacker,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) Dispose() {}

func (it *nopAttr) Finalize() string {
	return ""
}

func (it *nopAttr) CompileIf(isCompile bool) string {
	return ""
}

func (it *nopAttr) CompileToJsonResult() *corejson.Result {
	return nil
}

func (it *nopAttr) Compile() string {
	return ""
}

func (it *nopAttr) CompileFmt(
	formatter string, v ...interface{},
) string {
	return ""
}

func (it *nopAttr) Commit() {
	it.Finalize()
}

func (it *nopAttr) CompileAnyTo(toPointer interface{}) error {
	return nil
}

func (it *nopAttr) CompileAny() interface{} {
	return nil
}

func (it *nopAttr) CompileStacks() []string {
	return nil
}

func (it *nopAttr) ReflectSetTo(toPointer interface{}) error {
	return nil
}

func (it *nopAttr) CompileMap() map[string]interface{} {
	return nil
}

func (it *nopAttr) CompileBytes() ([]byte, error) {
	return nil, nil
}

func (it *nopAttr) CompileBytesIf(isCompile bool) ([]byte, error) {
	return nil, nil
}

func (it *nopAttr) CompileBytesMust() []byte {
	return nil
}

func (it *nopAttr) Id(
	value string,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) IdInt(
	value int,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) IdAny(
	value interface{},
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) StrNonEmpty(
	name, value string,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) Stringer(
	name string,
	stringer fmt.Stringer,
) loggerinf.MetaAttributesStacker {
	return it.Str(name, stringer.String())
}

func (it *nopAttr) Strings(
	name string,
	items ...string,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) StringArray(
	name string,
	items []string,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) Int(
	name string, value int,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) IntArray(
	name string, value []int,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) Float32(
	name string, value float32,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) Float32Array(
	name string, value []float32,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) Float64(
	name string, value float64,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) Float64Array(
	name string, value []float64,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) UInt(
	name string, value uint,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) UIntArray(
	name string, value []uint,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) Bool(
	name string, isTrue bool,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) BoolArray(
	name string, rawBooleans []bool,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) Byte(
	name string, value byte,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) Bytes(
	name string, rawBytes []byte,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) DefaultBytesWithErr(
	rawBytes []byte,
	err error,
) loggerinf.MetaAttributesStacker {
	return it.BytesWithErr(
		"BytesWithErr",
		rawBytes,
		err)
}

func (it *nopAttr) BytesWithErr(
	name string,
	rawBytes []byte,
	err error,
) loggerinf.MetaAttributesStacker {
	it.Bytes(name, rawBytes)
	it.NameDotErr(name, err)

	return it
}

func (it *nopAttr) Any(
	name string,
	value interface{},
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) AnyItems(
	name string,
	anyItems ...interface{},
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) MarshallJson(
	name string,
	anyItem interface{},
) *corejson.Result {
	return nil
}

func (it *nopAttr) Marshall(
	name string,
	anyItem interface{},
) []byte {
	return nil
}

func (it *nopAttr) CoreJson(
	name string,
	json *corejson.Result,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) NameDotErr(
	name string,
	err error,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) AnErr(
	name string,
	err error,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) Err(
	err error,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) Errs(
	name string,
	errs ...error,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) RawErrCollection(
	key string,
	rawErrCollection errcoreinf.BaseRawErrCollectionDefiner,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) ErrorWrapper(
	name string,
	errWp *errorwrapper.Wrapper,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) ErrorWrapperCollection(
	name string,
	errCollection *errwrappers.Collection,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) StackTraces(
	stackSkipIndex int,
	title string,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) StackTracesDefaultCount(
	stackSkipIndex int,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) StackTracesDefault() loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) Length() int {
	return 0
}

func (it *nopAttr) Count() int {
	return it.Length()
}

func (it *nopAttr) IsEmpty() bool {
	return it.Length() == 0
}

func (it *nopAttr) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *nopAttr) LastIndex() int {
	return it.Length() - 1
}

func (it *nopAttr) HasIndex(index int) bool {
	return index <= it.LastIndex()
}

func (it *nopAttr) HasKey(key string) bool {
	return false
}

func (it *nopAttr) IsKeyMissing(key string) bool {
	return true
}

func (it *nopAttr) DeleteKeys(
	keys ...string,
) loggerinf.MetaAttributesStacker {
	return it
}

func (it *nopAttr) Clear() {}

func (it *nopAttr) CompileStringNoMessage() string {
	return ""
}

func (it *nopAttr) String() string {
	return it.CompileStringNoMessage()
}

func (it *nopAttr) CompileJson(
	message string,
) corejson.Result {
	return corejson.Empty.Result()
}

func (it *nopAttr) CompileString(message string) string {
	json := it.CompileJson(message)

	return json.JsonString()
}

func (it *nopAttr) CompiledAsBasicErr(
	basicErrTyper errcoreinf.BasicErrorTyper,
) errcoreinf.BasicErrWrapper {
	return nil
}

func (it *nopAttr) CompiledAsErrorWrapper(
	variant errtype.Variation,
	finalMessage string,
) *errorwrapper.Wrapper {
	return nil
}

func (it *nopAttr) CompiledInjectToErrorCollection(
	errCollection *errwrappers.Collection,
	variant errtype.Variation,
	finalMessage string,
) *errwrappers.Collection {
	return errCollection
}

func (it nopAttr) AsMetaAttributesStacker() loggerinf.MetaAttributesStacker {
	return &it
}

func (it *nopAttr) Clone() *nopAttr {
	return it
}

func (it nopAttr) AsBasicSliceContractsBinder() coreinterface.BasicSlicerContractsBinder {
	return &it
}

func (it nopAttr) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return &it
}

func (it nopAttr) AsStandardSlicerContractsBinder() coreinterface.StandardSlicerContractsBinder {
	return &it
}
