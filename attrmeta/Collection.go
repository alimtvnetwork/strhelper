package attrmeta

import (
	"fmt"
	"log"

	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/coredynamic"
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/coreinterface"
	"gitlab.com/evatix-go/core/coreinterface/entityinf"
	"gitlab.com/evatix-go/core/coreinterface/enuminf"
	"gitlab.com/evatix-go/core/coreinterface/errcoreinf"
	"gitlab.com/evatix-go/core/coreinterface/loggerinf"
	"gitlab.com/evatix-go/core/coreinterface/serializerinf"
	"gitlab.com/evatix-go/core/coretaskinfo"
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/enum/logtype"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/errorwrapper/refs"
	"gitlab.com/evatix-go/strhelper/strs"
)

type Collection struct {
	logType logtype.Variant
	items   map[string]interface{}
}

func NewCollectionCap(
	capacity int,
) *Collection {
	return &Collection{
		items: make(map[string]interface{}, capacity),
	}
}

func NewCollectionDefault() *Collection {
	return &Collection{
		items: make(
			map[string]interface{},
			constants.Capacity12),
	}
}

func (it Collection) AsCollection() *Collection {
	return &it
}

func (it *Collection) Log() {
	if it.IsEmpty() {
		return
	}

	fmt.Println(it.Compile())
}

func (it *Collection) LogPretty() {
	if it.IsEmpty() {
		return
	}

	fmt.Println(
		it.CompileToJsonResult().
			PrettyJsonString())
}

func (it *Collection) LogWithTraces() {
	if it.IsEmpty() {
		return
	}

	fmt.Println(
		it.Compile() +
			constants.DefaultLine +
			codestack.StacksStringDefault())
}

func (it *Collection) LogFatal() {
	if it.IsEmpty() {
		return
	}

	log.Fatalln(it.Compile())
}

func (it *Collection) LogFatalWithTraces() {
	if it.IsEmpty() {
		return
	}

	log.Fatalln(it.Compile() + constants.DefaultLine + codestack.StacksStringDefault())
}

func (it *Collection) LogIf(isLog bool) {
	if !isLog {
		return
	}

	it.Log()
}

func (it *Collection) IsSilent() bool {
	if it == nil {
		return true
	}

	return it.logType.IsSilent()
}

func (it *Collection) Items() map[string]interface{} {
	return it.items
}

func (it *Collection) GetAsStrings() []string {
	if it.IsEmpty() {
		return []string{}
	}

	slice := make(
		[]string,
		it.Length())

	index := 0
	for key, valueAny := range it.items {
		slice[index] = fmt.Sprintf(
			constants.KeyValQuotationWrapJsonFormat,
			key,
			valueAny)
		index++
	}

	return slice
}

func (it *Collection) GetVal(keyName string) (val interface{}) {
	return it.items[keyName]
}

func (it *Collection) ListStrings() []string {
	return it.GetAsStrings()
}

func (it *Collection) JsonString() string {
	return it.JsonPtr().JsonString()
}

func (it *Collection) JsonStringMust() string {
	jsonResult := it.JsonPtr()
	jsonResult.MustBeSafe()

	return jsonResult.JsonString()
}

func (it *Collection) Json() corejson.Result {
	return corejson.New(it.items)
}

func (it *Collection) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it.items)
}

func (it *Collection) JsonParseSelfInject(jsonResult *corejson.Result) error {
	if it == nil {
		return errcore.CannotBeNilType.ErrorNoRefs("collection cannot be nil")
	}

	return jsonResult.Deserialize(it.items)
}

func (it *Collection) LoggerTyper() enuminf.LoggerTyper {
	return it.logType.AsLoggerTyper()
}

func (it *Collection) On(
	isLog bool,
) loggerinf.MetaAttributesStacker {
	if isLog {
		return it
	}

	// no log
	return &nopAttr{
		another: it,
	}
}

func (it *Collection) Msg(
	message string,
) loggerinf.MetaAttributesStacker {
	it.items["message"] = message

	return it
}

func (it *Collection) Title(
	title string,
) loggerinf.MetaAttributesStacker {
	it.items["title"] = title

	return it
}

func (it *Collection) TitleAttr(
	title, attr string,
) loggerinf.MetaAttributesStacker {
	it.items[title] = attr

	return it
}

func (it *Collection) Info(
	info *coretaskinfo.Info,
) loggerinf.MetaAttributesStacker {
	it.items[info.Name()] = info.
		PrettyJsonString()

	return it
}

func (it *Collection) InfoWithPayload(
	info *coretaskinfo.Info,
	payloads []byte,
) loggerinf.MetaAttributesStacker {
	it.items[info.Name()] = info.
		PrettyJsonStringWithPayloads(payloads)

	return it
}

func (it *Collection) FriendlyErr(
	name string,
	frdErrWrap *errorwrapper.FriendlyError,
	isIncludeTraces bool,
) loggerinf.MetaAttributesStacker {
	if frdErrWrap.IsEmpty() {
		return it
	}

	it.items[name+".FriendlyMessage"] = frdErrWrap.FriendlyMsg()
	it.items[name] = frdErrWrap.FullStringWithTracesIf(isIncludeTraces)

	return it
}

func (it *Collection) Str(
	title, val string,
) loggerinf.MetaAttributesStacker {
	it.items[title] = val

	return it
}

func (it *Collection) NullKey(
	title string,
) loggerinf.MetaAttributesStacker {
	it.items[title] = constants.NilAngelBracket

	return it
}

func (it *Collection) StandardSlicer(
	title string,
	standardSlice coreinterface.StandardSlicer,
) loggerinf.MetaAttributesStacker {
	if standardSlice == nil {
		it.items[title] = constants.NilAngelBracket

		return it
	}

	it.items[title] = standardSlice.ListStrings()

	return it
}

func (it *Collection) Stringers(
	title string,
	stringers ...fmt.Stringer,
) loggerinf.MetaAttributesStacker {
	slice := strs.StringersToSlice(
		stringers...)
	it.items[title] = slice

	return it
}

func (it *Collection) Hex(
	title string, hexValues []byte,
) loggerinf.MetaAttributesStacker {
	return it.RawJson(
		title, hexValues)
}

func (it *Collection) RawJson(
	title string,
	rawJsonBytes []byte,
) loggerinf.MetaAttributesStacker {
	var toString string

	if len(rawJsonBytes) > 0 {
		toString = string(rawJsonBytes)
	}

	it.items[title] = toString

	return it
}

// Error
//
//  skip on error
func (it *Collection) Error(
	title string,
	err error,
) loggerinf.MetaAttributesStacker {
	if err == nil {
		return it
	}

	it.items[title] = errcore.CompiledError(
		err, title)

	return it
}

func (it *Collection) SimpleBytesResulter(
	title string,
	result serializerinf.SimpleBytesResulter,
) loggerinf.MetaAttributesStacker {
	if result == nil {
		it.items[title] = constants.NilAngelBracket

		return it
	}

	if result.HasError() {
		it.items[title] = fmt.Sprintf(
			"Value : %s, Error: %s",
			string(result.SafeValues()),
			errcore.ToString(result.MeaningfulError()))

		return it
	}

	it.items[title] = string(result.SafeValues())

	return it
}

func (it *Collection) BaseJsonResulter(
	title string,
	result serializerinf.BaseJsonResulter,
) loggerinf.MetaAttributesStacker {
	it.items[title] = jsonResultToDisplayString(
		result)

	return it
}

func (it *Collection) BasicJsonResulter(
	title string,
	result serializerinf.BasicJsonResulter,
) loggerinf.MetaAttributesStacker {
	return it.BaseJsonResulter(
		title,
		result)
}

func (it *Collection) JsonResulter(
	title string,
	result serializerinf.JsonResulter,
) loggerinf.MetaAttributesStacker {
	return it.BaseJsonResulter(
		title,
		result)
}

func (it *Collection) MapIntegerAny(
	title string,
	mapAny map[int]interface{},
) loggerinf.MetaAttributesStacker {
	return it.JsonResulter(
		title,
		corejson.NewPtr(mapAny))
}

func (it *Collection) Meta(
	title string,
	metaAttr loggerinf.MetaAttributesCompiler,
) loggerinf.MetaAttributesStacker {
	if metaAttr == nil {
		it.items[title] = constants.NilAngelBracket

		return it
	}

	it.items[title] = metaAttr.Compile()

	return it
}

func (it *Collection) MapBool(
	title string,
	mapInt map[string]bool,
) loggerinf.MetaAttributesStacker {
	return it.JsonResulter(
		title,
		corejson.NewPtr(mapInt))
}

func (it *Collection) MapInt(
	title string,
	mapInt map[string]int,
) loggerinf.MetaAttributesStacker {
	return it.JsonResulter(
		title,
		corejson.NewPtr(mapInt))
}

func (it *Collection) MapAnyAny(
	title string,
	mapAny map[interface{}]interface{},
) loggerinf.MetaAttributesStacker {
	return it.JsonResulter(
		title,
		corejson.NewPtr(mapAny))
}

func (it *Collection) MapAny(
	title string,
	mapAny map[string]interface{},
) loggerinf.MetaAttributesStacker {
	return it.JsonResulter(
		title,
		corejson.NewPtr(mapAny))
}

func (it *Collection) MapIntAny(
	title string, mapAny map[int]interface{},
) loggerinf.MetaAttributesStacker {
	return it.JsonResulter(
		title,
		corejson.NewPtr(mapAny))
}

func (it *Collection) MapIntString(
	title string, mapAny map[int]string,
) loggerinf.MetaAttributesStacker {
	return it.JsonResulter(
		title,
		corejson.NewPtr(mapAny))
}

func (it *Collection) MapJsonResult(
	title string, mapAny map[string]corejson.Result,
) loggerinf.MetaAttributesStacker {
	return it.JsonResulter(
		title,
		corejson.NewPtr(mapAny))
}

func (it *Collection) JsonResult(title string, json *corejson.Result) loggerinf.MetaAttributesStacker {
	return it.JsonResulter(
		title,
		json)
}

func (it *Collection) JsonResultItems(title string, jsons ...*corejson.Result) loggerinf.MetaAttributesStacker {
	compiledString := jsonResultsDisplayString(
		jsons...)

	return it.Str(
		title,
		compiledString)
}

func (it *Collection) DefaultStackTraces() loggerinf.MetaAttributesStacker {
	it.items["StackTraces"] = codestack.StacksString(
		codestack.Skip1)

	return it
}

func (it *Collection) ErrWithTypeTraces(
	title string,
	errType errcoreinf.BasicErrorTyper,
	err error,
) loggerinf.MetaAttributesStacker {
	if err == nil {
		return it
	}

	it.items[title] =
		errnew.Error.Default(
			errtype.NewUsingTyper(errType),
			err).
			FullStringWithTraces()

	return it
}

func (it *Collection) ErrorsWithTypeTraces(
	title string,
	errType errcoreinf.BasicErrorTyper,
	errorItems ...error,
) loggerinf.MetaAttributesStacker {
	if len(errorItems) == 0 {
		return it
	}

	it.items[title] = errnew.
		Error.Default(
		errtype.NewUsingTyper(errType),
		errcore.MergeErrors(errorItems...)).
		FullStringWithTraces()

	return it
}

func (it *Collection) OnErrStackTraces(err error) loggerinf.MetaAttributesStacker {
	if err == nil {
		return it
	}

	it.items["err-with-traces"] = errnew.
		Error.Default(
		errtype.Unknown,
		err).
		FullStringWithTraces()

	return it
}

func (it *Collection) OnErrWrapperOrCollectionStackTraces(
	errWrapperOrCollection errcoreinf.BaseErrorOrCollectionWrapper,
) loggerinf.MetaAttributesStacker {
	if errWrapperOrCollection == nil || errWrapperOrCollection.IsEmpty() {
		return it
	}

	it.items["ErrorWrapper-Or-Collection"] = errWrapperOrCollection.
		FullStringWithTraces()

	return it
}

func (it *Collection) FullStringer(title string, fullStringer errcoreinf.FullStringer) loggerinf.MetaAttributesStacker {
	if fullStringer == nil {
		return it
	}

	it.items[title] = fullStringer.FullString()

	return it
}

func (it *Collection) OnlyFullStringer(fullStringer errcoreinf.FullStringer) loggerinf.MetaAttributesStacker {
	if fullStringer == nil {
		return it
	}

	it.items["FullStringer"] = fullStringer.FullString()

	return it
}

func (it *Collection) FullTraceAsAttr(
	title string, attrFullStringWithTraces errcoreinf.FullStringWithTracesGetter,
) loggerinf.MetaAttributesStacker {
	if attrFullStringWithTraces == nil {
		return it
	}

	it.items[title] = attrFullStringWithTraces.FullStringWithTraces()

	return it
}

func (it *Collection) BasicErrWrapper(
	errWrapperOrCollection errcoreinf.BasicErrWrapper,
) loggerinf.MetaAttributesStacker {
	if errWrapperOrCollection == nil || errWrapperOrCollection.IsEmpty() {
		return it
	}

	it.items["BasicError"] = errWrapperOrCollection.FullStringWithTraces()

	return it
}

func (it *Collection) BaseRawErrCollectionDefiner(
	errWrapperOrCollection errcoreinf.BaseRawErrCollectionDefiner,
) loggerinf.MetaAttributesStacker {
	if errWrapperOrCollection == nil || errWrapperOrCollection.IsEmpty() {
		return it
	}

	it.items["ErrorWrapperOrCollection"] = errWrapperOrCollection.FullStringWithTraces()

	return it
}

func (it *Collection) BaseErrorWrapperCollectionDefiner(
	errWrapperCollection errcoreinf.BaseErrorWrapperCollectionDefiner,
) loggerinf.MetaAttributesStacker {
	if errWrapperCollection == nil || errWrapperCollection.IsEmpty() {
		return it
	}

	it.items["ErrorCollection"] = errWrapperCollection.FullStringWithTraces()

	return it
}

func (it *Collection) ErrWrapperOrCollection(
	errWrapperOrCollection errcoreinf.BaseErrorOrCollectionWrapper,
) loggerinf.MetaAttributesStacker {
	if errWrapperOrCollection == nil || errWrapperOrCollection.IsEmpty() {
		return it
	}

	it.items["ErrorWrapperOrCollection"] = errWrapperOrCollection.FullStringWithTraces()

	return it
}

func (it *Collection) CompiledBasicErrWrapper(
	compiler errcoreinf.CompiledBasicErrWrapper,
) loggerinf.MetaAttributesStacker {
	if compiler == nil {
		return it
	}

	basicErrWp := compiler.
		CompiledToGenericBasicErrWrapper()

	if basicErrWp == nil || basicErrWp.IsEmpty() {
		return it
	}

	it.items["ErrorCompiler"] = basicErrWp.FullStringWithTraces()

	return it
}

func (it *Collection) Namer(
	title string, namer enuminf.Namer,
) loggerinf.MetaAttributesStacker {
	if namer == nil {
		return it
	}

	it.items[title] = namer.Name()

	return it
}

func (it *Collection) OnlyNamer(
	namer enuminf.Namer,
) loggerinf.MetaAttributesStacker {
	if namer == nil {
		return it
	}

	it.items[namer.Name()] = namer.Name()

	return it
}

func (it *Collection) EnumTitleEnum(
	title enuminf.SimpleEnumer,
	enum enuminf.BasicEnumer,
) loggerinf.MetaAttributesStacker {
	if enum == nil || title == nil {
		return it
	}

	it.items[title.Name()] = enum.NameValue()

	return it
}

func (it *Collection) SimpleEnumTitleEnum(
	title enuminf.SimpleEnumer,
	enum enuminf.SimpleEnumer,
) loggerinf.MetaAttributesStacker {
	if enum == nil || title == nil {
		return it
	}

	it.items[title.Name()] = enum.String()

	return it
}

func (it *Collection) Enum(
	title string,
	enum enuminf.BasicEnumer,
) loggerinf.MetaAttributesStacker {
	if enum == nil {
		return it
	}

	it.items[title] = enum.NameValue()

	return it
}

func (it *Collection) Enums(
	key string, enums ...enuminf.BasicEnumer,
) loggerinf.MetaAttributesStacker {
	if len(enums) == 0 {
		it.items[key] = ""

		return it
	}

	slice := make([]string, len(enums))
	for i, enum := range enums {
		if enum == nil {
			continue
		}

		slice[i] = enum.NameValue()
	}

	it.items[key] = slice

	return it
}

func (it *Collection) OnlyEnum(
	enum enuminf.BasicEnumer,
) loggerinf.MetaAttributesStacker {
	if enum == nil {
		return it
	}

	it.items["enum"] = enum.NameValue()

	return it
}

func (it *Collection) OnlyEnums(
	enums ...enuminf.BasicEnumer,
) loggerinf.MetaAttributesStacker {
	return it.Enums("enums", enums...)
}

func (it *Collection) OnlyString(
	value string,
) loggerinf.MetaAttributesStacker {
	return it.Title(value)
}

func (it *Collection) OnlyStrings(
	values ...string,
) loggerinf.MetaAttributesStacker {
	return it.Strings("strings", values...)
}

func (it *Collection) OnlyStringer(
	stringer fmt.Stringer,
) loggerinf.MetaAttributesStacker {
	return it.Stringer("stringer", stringer)
}

func (it *Collection) OnlyStringers(
	stringers ...fmt.Stringer,
) loggerinf.MetaAttributesStacker {
	return it.Stringers(
		"stringers", stringers...)
}

func (it *Collection) OnlyIntegers(
	values ...int,
) loggerinf.MetaAttributesStacker {
	return it.Integers(
		"integers", values...)
}

func (it *Collection) OnlyBooleans(
	values ...bool,
) loggerinf.MetaAttributesStacker {
	return it.Booleans(
		"booleans", values...)
}

func (it *Collection) OnlyBytes(
	rawBytes []byte,
) loggerinf.MetaAttributesStacker {
	return it.Bytes(
		"bytes", rawBytes)
}

func (it *Collection) OnlyRawJson(
	rawBytes []byte,
) loggerinf.MetaAttributesStacker {
	return it.Bytes(
		"raw-json", rawBytes)
}

func (it *Collection) OnlyBytesErr(
	rawBytes []byte, err error,
) loggerinf.MetaAttributesStacker {
	jsonResult := corejson.NewResult.Ptr(rawBytes, err, "bytes")

	return it.SimpleBytesResulter(
		"bytes-err",
		jsonResult)
}

func (it *Collection) OnlySimpleBytesResulter(
	result serializerinf.SimpleBytesResulter,
) loggerinf.MetaAttributesStacker {
	return it.SimpleBytesResulter(
		"simple-bytes-result",
		result)
}

func (it *Collection) OnlyBaseJsonResulter(
	result serializerinf.BaseJsonResulter,
) loggerinf.MetaAttributesStacker {
	return it.SimpleBytesResulter(
		"base-json-result",
		result)
}

func (it *Collection) OnlyBasicJsonResulter(
	result serializerinf.BasicJsonResulter,
) loggerinf.MetaAttributesStacker {
	return it.SimpleBytesResulter(
		"basic-json-result",
		result)
}

func (it *Collection) OnlyJsonResulter(
	result serializerinf.JsonResulter,
) loggerinf.MetaAttributesStacker {
	return it.SimpleBytesResulter(
		"json-resulter",
		result)
}

func (it *Collection) OnlyAny(
	anyItem interface{},
) loggerinf.MetaAttributesStacker {
	return it.Any(
		"any",
		anyItem)
}

func (it *Collection) OnlyAnyItems(
	values ...interface{},
) loggerinf.MetaAttributesStacker {
	return it.AnyItems(
		"any-items",
		values...)
}

func (it *Collection) OnlyMetaAttr(
	metaAttr loggerinf.MetaAttributesCompiler,
) loggerinf.MetaAttributesStacker {
	return it.Meta(
		"meta-attr",
		metaAttr)
}

func (it *Collection) OnlyAnyIf(
	isLog bool, anyItem interface{},
) loggerinf.MetaAttributesStacker {
	if !isLog {
		return it
	}

	return it.Any(
		"any",
		anyItem)
}

func (it *Collection) OnlyAnyItemsIf(
	isLog bool, anyItems ...interface{},
) loggerinf.MetaAttributesStacker {
	if !isLog {
		return it
	}

	return it.AnyItems(
		"any-items",
		anyItems...)
}

func (it *Collection) OnlyMapBool(
	mapItems map[string]bool,
) loggerinf.MetaAttributesStacker {
	return it.MapBool(
		"map-string-boolean",
		mapItems)
}

func (it *Collection) OnlyMapInt(
	mapInt map[string]int,
) loggerinf.MetaAttributesStacker {
	return it.MapInt(
		"map-string-int",
		mapInt)
}

func (it *Collection) OnlyMapAny(
	mapAny map[string]interface{},
) loggerinf.MetaAttributesStacker {
	return it.MapAny(
		"map-string-any",
		mapAny)
}

func (it *Collection) OnlyMapIntAny(
	mapAny map[int]interface{},
) loggerinf.MetaAttributesStacker {
	return it.MapIntAny(
		"map-int-any",
		mapAny)
}

func (it *Collection) OnlyMapIntString(
	mapAny map[int]string,
) loggerinf.MetaAttributesStacker {
	return it.MapIntString(
		"map-int-string",
		mapAny)
}

func (it *Collection) OnlyMapJsonResult(
	mapAny map[string]corejson.Result,
) loggerinf.MetaAttributesStacker {
	return it.MapJsonResult(
		"map-json-result",
		mapAny)
}

func (it *Collection) OnlyJson(
	json *corejson.Result,
) loggerinf.MetaAttributesStacker {
	return it.JsonResult(
		"json",
		json)
}

func (it *Collection) OnlyJsons(
	jsons ...*corejson.Result,
) loggerinf.MetaAttributesStacker {
	return it.JsonResultItems(
		"json-results",
		jsons...)
}

func (it *Collection) Booleans(
	title string, isResults ...bool,
) loggerinf.MetaAttributesStacker {
	it.items[title] = isResults

	return it
}

func (it *Collection) Jsoner(
	jsoner corejson.Jsoner,
) loggerinf.MetaAttributesStacker {
	if jsoner == nil {
		return it.JsonResult(
			"jsoner",
			nil)
	}

	return it.JsonResult(
		"jsoner",
		jsoner.JsonPtr())
}

func (it *Collection) Jsoners(
	jsoners ...corejson.Jsoner,
) loggerinf.MetaAttributesStacker {
	if len(jsoners) == 0 {
		return it.JsonResult(
			"jsoners",
			nil)
	}

	jsonResults := corejson.
		NewResultsPtrCollection.
		Jsoners(jsoners...)

	return it.JsonResultItems(
		"jsoners",
		jsonResults.Items...)
}

func (it *Collection) JsonerTitle(
	title string,
	jsoner corejson.Jsoner,
) loggerinf.MetaAttributesStacker {
	if jsoner == nil {
		return it.JsonResult(
			title,
			nil)
	}

	return it.JsonResult(
		title,
		jsoner.JsonPtr())
}

func (it *Collection) JsonerIf(
	isLog bool,
	jsoner corejson.Jsoner,
) loggerinf.MetaAttributesStacker {
	if !isLog {
		return it
	}

	return it.Jsoner(jsoner)
}

func (it *Collection) JsonersIf(
	isLog bool,
	jsoners ...corejson.Jsoner,
) loggerinf.MetaAttributesStacker {
	if !isLog {
		return it
	}

	return it.Jsoners(jsoners...)
}

func (it *Collection) Serializer(
	serializer loggerinf.Serializer,
) loggerinf.MetaAttributesStacker {
	if serializer == nil {
		return it
	}

	allBytes, err := serializer.Serialize()

	return it.BytesWithErr(
		"serializer",
		allBytes,
		err)
}

func (it *Collection) Serializers(
	serializers ...loggerinf.Serializer,
) loggerinf.MetaAttributesStacker {
	if len(serializers) == 0 {
		return it.JsonResulter(
			"serializers",
			nil)
	}

	jsonResults := corejson.
		NewResultsPtrCollection.
		UsingCap(len(serializers))

	for _, serializer := range serializers {
		allBytes, err := serializer.Serialize()

		jsonResult := corejson.NewResult.Ptr(
			allBytes,
			err,
			"serializer")

		jsonResults.Add(jsonResult)
	}

	return it.JsonResultItems(
		"serializers",
		jsonResults.Items...)
}

func (it *Collection) SerializerFunc(
	serializerFunc func() ([]byte, error),
) loggerinf.MetaAttributesStacker {
	if serializerFunc == nil {
		return it.JsonResulter(
			"serializer-func",
			nil)
	}

	allBytes, err := serializerFunc()
	jsonResult := corejson.NewResult.Ptr(
		allBytes,
		err,
		"serializerFunc")

	return it.JsonResult(
		"serializer-func",
		jsonResult)
}

func (it *Collection) SerializerFunctions(
	serializerFunctions ...func() ([]byte, error),
) loggerinf.MetaAttributesStacker {
	if serializerFunctions == nil {
		return it.JsonResulter(
			"serializer-functions",
			nil)
	}

	jsonResults := corejson.
		NewResultsPtrCollection.
		UsingCap(len(serializerFunctions))

	for _, serializer := range serializerFunctions {
		allBytes, err := serializer()

		jsonResult := corejson.NewResult.Ptr(
			allBytes,
			err,
			"serializerFunc")

		jsonResults.Add(jsonResult)
	}

	return it.JsonResultItems(
		"serializers-functions",
		jsonResults.Items...)
}

func (it *Collection) StandardTaskEntityDefiner(
	entity entityinf.StandardTaskEntityDefiner,
) loggerinf.MetaAttributesStacker {
	return it.StandardTaskEntityDefinerTitle(
		"standard-task-entity",
		entity)
}

func (it *Collection) TaskEntityDefiner(
	entity entityinf.TaskEntityDefiner,
) loggerinf.MetaAttributesStacker {
	return it.TaskEntityDefinerTitle(
		"task-entity",
		entity)
}

func (it *Collection) StandardTaskEntityDefinerTitle(
	title string,
	entity entityinf.StandardTaskEntityDefiner,
) loggerinf.MetaAttributesStacker {
	return it.TaskEntityDefinerTitle(
		title,
		entity)
}

func (it *Collection) TaskEntityDefinerTitle(
	title string,
	entity entityinf.TaskEntityDefiner,
) loggerinf.MetaAttributesStacker {
	if entity == nil {
		it.items[title] = constants.NilAngelBracket

		return it
	}

	return it.JsonResult(
		title, entity.JsonPtr())
}

func (it *Collection) LoggerModel(
	loggerModel loggerinf.SingleLogModeler,
) loggerinf.MetaAttributesStacker {
	return it.LoggerModelTitle(
		"logger-model",
		loggerModel)
}

func (it *Collection) LoggerModelTitle(
	title string,
	loggerModel loggerinf.SingleLogModeler,
) loggerinf.MetaAttributesStacker {
	if loggerModel == nil {
		it.items[title] = constants.NilAngelBracket

		return it
	}

	return it.JsonResult(
		title,
		corejson.NewPtr(loggerModel))
}

func (it *Collection) Integers(
	key string,
	integerItems ...int,
) loggerinf.MetaAttributesStacker {
	if integerItems == nil {
		it.items[key] = constants.NilAngelBracket

		return it
	}

	it.items[key] = integerItems

	return it
}

func (it *Collection) Fmt(
	title,
	format string,
	v ...interface{},
) loggerinf.MetaAttributesStacker {
	if format == "" && len(v) == 0 {
		return it
	}

	it.items[title] = fmt.Sprintf(
		format,
		v...)

	return it
}

func (it *Collection) FmtIf(
	isLog bool,
	title,
	format string,
	v ...interface{},
) loggerinf.MetaAttributesStacker {
	if !isLog || format == "" && len(v) == 0 {
		return it
	}

	return it.Fmt(
		title,
		format,
		v...)
}

func (it *Collection) OnlyFmt(
	format string,
	v ...interface{},
) loggerinf.MetaAttributesStacker {
	return it.Fmt(
		"Format",
		format,
		v...)
}

func (it *Collection) OnlyFmtIf(
	isLog bool,
	format string,
	v ...interface{},
) loggerinf.MetaAttributesStacker {
	if !isLog {
		return it
	}

	return it.OnlyFmt(
		format, v...)
}

func (it *Collection) RawPayloadsGetter(
	payloadsGetter loggerinf.RawPayloadsGetter,
) loggerinf.MetaAttributesStacker {
	return it.RawPayloadsGetterTitle(
		"raw-payloads-getter",
		payloadsGetter)
}

func (it *Collection) RawPayloadsGetterTitle(
	title string,
	payloadsGetter loggerinf.RawPayloadsGetter,
) loggerinf.MetaAttributesStacker {
	if payloadsGetter == nil {
		return it.NullKey(title)
	}

	allBytes, err := payloadsGetter.RawPayloads()

	return it.BytesWithErr(
		title,
		allBytes,
		err)
}

func (it *Collection) RawPayloadsGetterIf(
	isLog bool, payloadsGetter loggerinf.RawPayloadsGetter,
) loggerinf.MetaAttributesStacker {
	if !isLog {
		return it
	}

	return it.RawPayloadsGetter(
		payloadsGetter)
}

func (it *Collection) Inject(
	others ...loggerinf.MetaAttributesStacker,
) loggerinf.MetaAttributesStacker {
	if len(others) == 0 {
		return it
	}

	for _, other := range others {
		if other == nil {
			continue
		}

		compiledMap := other.CompileMap()

		if len(compiledMap) == 0 {
			continue
		}

		it.InjectMap(compiledMap)
	}

	return it
}

func (it *Collection) InjectMap(
	compiledMap map[string]interface{},
) loggerinf.MetaAttributesStacker {
	if len(compiledMap) == 0 {
		return it
	}

	for key, val := range compiledMap {
		it.items[key] = val
	}

	return it
}

func (it *Collection) ConcatNew(
	others ...loggerinf.MetaAttributesStacker,
) loggerinf.MetaAttributesStacker {
	if len(others) == 0 {
		return it.Clone()
	}

	// todo improve
	newAttrs := NewCollectionCap(
		it.Length() * len(others))

	newAttrs.InjectMap(it.CompileMap())

	for _, other := range others {
		newAttrs.InjectMap(other.CompileMap())
	}

	return newAttrs
}

func (it *Collection) Dispose() {
	if it == nil {
		return
	}

	it.Clear()
	it.items = nil
}

func (it *Collection) Finalize() string {
	return it.Compile()
}

func (it *Collection) CompileIf(isCompile bool) string {
	if !isCompile {
		return ""
	}

	return it.Compile()
}

func (it *Collection) CompileToJsonResult() *corejson.Result {
	return corejson.NewPtr(it.CompileMap())
}

func (it *Collection) Compile() string {
	result := corejson.New(
		it.items)
	it.Clear()

	return result.JsonString()
}

func (it *Collection) CompileFmt(
	formatter string, v ...interface{},
) string {
	compiledMessage := fmt.Sprintf(
		formatter,
		v...)

	return compiledMessage + " " + it.Compile()
}

func (it *Collection) Commit() {
	it.Finalize()
}

func (it *Collection) CompileAnyTo(toPointer interface{}) error {
	// todo
	panic("implement me, intentional todo")
}

func (it *Collection) CompileAny() interface{} {
	temp := it.items
	it.Clear()

	return temp
}

func (it *Collection) CompileStacks() []string {
	if it.IsEmpty() {
		return []string{}
	}

	slice := it.GetAsStrings()
	it.Clear()

	return slice
}

func (it *Collection) ReflectSetTo(toPointer interface{}) error {
	return coredynamic.ReflectSetFromTo(it, toPointer)
}

func (it *Collection) CompileMap() map[string]interface{} {
	temp := it.items
	it.Clear()

	return temp
}

func (it *Collection) CompileBytes() ([]byte, error) {
	jsonResult := corejson.New(
		it.CompileMap())

	return jsonResult.Raw()
}

func (it *Collection) CompileBytesIf(isCompile bool) ([]byte, error) {
	if !isCompile {
		it.Clear()
	}

	jsonResult := corejson.New(it.CompileMap())
	return jsonResult.Raw()
}

func (it *Collection) CompileBytesMust() []byte {
	jsonResult := corejson.New(it.items)

	return jsonResult.RawMust()
}

func (it *Collection) Id(
	value string,
) loggerinf.MetaAttributesStacker {
	it.items[id] = value

	return it
}

func (it *Collection) IdInt(
	value int,
) loggerinf.MetaAttributesStacker {
	it.items[id] = value

	return it
}

func (it *Collection) IdAny(
	value interface{},
) loggerinf.MetaAttributesStacker {
	it.items[id] = value

	return it
}

func (it *Collection) StrNonEmpty(
	name, value string,
) loggerinf.MetaAttributesStacker {
	if value == "" {
		return it
	}

	it.items[name] = value

	return it
}

func (it *Collection) Stringer(
	name string,
	stringer fmt.Stringer,
) loggerinf.MetaAttributesStacker {
	return it.Str(name, stringer.String())
}

func (it *Collection) Strings(
	name string,
	items ...string,
) loggerinf.MetaAttributesStacker {
	if len(items) == 0 {
		return it
	}

	it.items[name] = items

	return it
}

func (it *Collection) StringArray(
	name string,
	items []string,
) loggerinf.MetaAttributesStacker {
	if len(items) == 0 {
		return it
	}

	it.items[name] = items

	return it
}

func (it *Collection) Int(
	name string, value int,
) loggerinf.MetaAttributesStacker {
	it.items[name] = value

	return it
}

func (it *Collection) IntArray(
	name string, value []int,
) loggerinf.MetaAttributesStacker {
	it.items[name] = value

	return it
}

func (it *Collection) Float32(
	name string, value float32,
) loggerinf.MetaAttributesStacker {
	it.items[name] = value

	return it
}

func (it *Collection) Float32Array(
	name string, value []float32,
) loggerinf.MetaAttributesStacker {
	it.items[name] = value

	return it
}

func (it *Collection) Float64(
	name string, value float64,
) loggerinf.MetaAttributesStacker {
	it.items[name] = value

	return it
}

func (it *Collection) Float64Array(
	name string, value []float64,
) loggerinf.MetaAttributesStacker {
	it.items[name] = value

	return it
}

func (it *Collection) UInt(
	name string, value uint,
) loggerinf.MetaAttributesStacker {
	it.items[name] = value

	return it
}

func (it *Collection) UIntArray(
	name string, value []uint,
) loggerinf.MetaAttributesStacker {
	it.items[name] = value

	return it
}

func (it *Collection) Bool(
	name string, isTrue bool,
) loggerinf.MetaAttributesStacker {
	it.items[name] = isTrue

	return it
}

func (it *Collection) BoolArray(
	name string, rawBooleans []bool,
) loggerinf.MetaAttributesStacker {
	it.items[name] = rawBooleans

	return it
}

func (it *Collection) Byte(
	name string, value byte,
) loggerinf.MetaAttributesStacker {
	it.items[name] = value

	return it
}

func (it *Collection) Bytes(
	name string, rawBytes []byte,
) loggerinf.MetaAttributesStacker {
	if len(rawBytes) == 0 {
		it.items[name] = ""

		return it
	}

	it.items[name] = string(rawBytes)

	return it
}

func (it *Collection) DefaultBytesWithErr(
	rawBytes []byte,
	err error,
) loggerinf.MetaAttributesStacker {
	return it.BytesWithErr(
		"BytesWithErr",
		rawBytes,
		err)
}

func (it *Collection) BytesWithErr(
	name string,
	rawBytes []byte,
	err error,
) loggerinf.MetaAttributesStacker {
	it.Bytes(name, rawBytes)
	it.NameDotErr(name, err)

	return it
}

func (it *Collection) Any(
	name string,
	value interface{},
) loggerinf.MetaAttributesStacker {
	it.items[name] = value

	return it
}

func (it *Collection) AnyItems(
	name string,
	anyItems ...interface{},
) loggerinf.MetaAttributesStacker {
	it.items[name] = anyItems

	return it
}

func (it *Collection) MarshallJson(
	name string,
	anyItem interface{},
) *corejson.Result {
	json := corejson.New(anyItem)
	it.CoreJson(name, &json)

	return &json
}

func (it *Collection) Marshall(
	name string,
	anyItem interface{},
) []byte {
	json := it.MarshallJson(name, anyItem)

	return json.SafeBytes()
}

func (it *Collection) CoreJson(
	name string,
	json *corejson.Result,
) loggerinf.MetaAttributesStacker {
	it.items[name] = json.JsonString()
	it.NameDotErr(name, json.MeaningfulError())

	return it
}

func (it *Collection) NameDotErr(
	name string,
	err error,
) loggerinf.MetaAttributesStacker {
	if err == nil {
		return it
	}

	it.items[name+dotError] = err.Error()

	return it
}

func (it *Collection) AnErr(
	name string,
	err error,
) loggerinf.MetaAttributesStacker {
	if err == nil {
		return it
	}

	it.items[name] = err.Error()

	return it
}

func (it *Collection) Err(
	err error,
) loggerinf.MetaAttributesStacker {
	if err == nil {
		return it
	}

	it.items[errorKey] = err.Error()

	return it
}

func (it *Collection) Errs(
	name string,
	errs ...error,
) loggerinf.MetaAttributesStacker {
	compiledErrString := errcore.MergeErrorsToString(
		constants.DefaultLine,
		errs...)

	if compiledErrString == "" {
		return it
	}

	it.items[name] = compiledErrString

	return it
}

func (it *Collection) RawErrCollection(
	key string,
	rawErrCollection errcoreinf.BaseRawErrCollectionDefiner,
) loggerinf.MetaAttributesStacker {
	if rawErrCollection.IsEmpty() {
		return it
	}

	it.items[key] =
		rawErrCollection.
			String()

	return it
}

func (it *Collection) ErrorWrapper(
	name string,
	errWp *errorwrapper.Wrapper,
) loggerinf.MetaAttributesStacker {
	if errWp.IsEmpty() {
		return it
	}

	it.items[name] = errWp.StackTraceString()

	return it
}

func (it *Collection) ErrorWrapperCollection(
	name string,
	errCollection *errwrappers.Collection,
) loggerinf.MetaAttributesStacker {
	if errCollection.IsEmpty() {
		return it
	}

	it.items[name] = errCollection.
		DisplayStringWithTraces()

	return it
}

func (it *Collection) StackTraces(
	stackSkipIndex int,
	title string,
) loggerinf.MetaAttributesStacker {
	it.items[stackTraceKey] = title +
		constants.DefaultLine +
		codestack.StacksCountString(
			stackSkipIndex+codestack.Skip1,
			codestack.DefaultStackCount*2)

	return it
}

func (it *Collection) StackTracesDefaultCount(
	stackSkipIndex int,
) loggerinf.MetaAttributesStacker {
	it.items[stackTraceKey] = codestack.StacksCountString(
		stackSkipIndex+codestack.Skip1,
		codestack.DefaultStackCount)

	return it
}

func (it *Collection) StackTracesDefault() loggerinf.MetaAttributesStacker {
	it.items[stackTraceKey] = codestack.StacksCountString(
		codestack.Skip1,
		codestack.DefaultStackCount)

	return it
}

func (it *Collection) Length() int {
	if it == nil {
		return 0
	}

	return len(it.items)
}

func (it *Collection) Count() int {
	return it.Length()
}

func (it *Collection) IsEmpty() bool {
	return it.Length() == 0
}

func (it *Collection) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *Collection) LastIndex() int {
	return it.Length() - 1
}

func (it *Collection) HasIndex(index int) bool {
	return index <= it.LastIndex()
}

func (it *Collection) HasKey(key string) bool {
	if it.IsEmpty() {
		return false
	}

	_, has := it.items[key]

	return has
}

func (it *Collection) IsKeyMissing(key string) bool {
	if it.IsEmpty() {
		return true
	}

	_, has := it.items[key]

	return !has
}

func (it *Collection) DeleteKeys(
	keys ...string,
) loggerinf.MetaAttributesStacker {
	if it.IsEmpty() {
		return it
	}

	for _, key := range keys {
		if it.IsKeyMissing(key) {
			continue
		}

		// exist remove
		delete(it.items, key)
	}

	return it
}

func (it *Collection) Clear() {
	it.items = map[string]interface{}{}
}

func (it *Collection) CompileStringNoMessage() string {
	json := it.Compile()

	return json
}

func (it *Collection) String() string {
	return it.CompileStringNoMessage()
}

func (it *Collection) CompileJson(
	message string,
) corejson.Result {
	if message == "" {
		jsonResult := corejson.New(it.items)
		it.Clear()

		return jsonResult
	}

	metaCompile := struct {
		Message  string
		MetaData map[string]interface{}
	}{
		Message:  message,
		MetaData: it.items,
	}

	jsonResult := corejson.New(metaCompile)
	it.Clear()

	return jsonResult
}

func (it *Collection) CompileString(message string) string {
	json := it.CompileJson(message)

	return json.JsonString()
}

func (it *Collection) CompiledAsBasicErr(
	basicErrTyper errcoreinf.BasicErrorTyper,
) errcoreinf.BasicErrWrapper {
	if it.IsEmpty() || it.IsSilent() {
		return nil
	}

	refsCollection := refs.New(it.Length())
	refsCollection.AddMap(it.Items())
	it.Clear()

	return errorwrapper.NewMsgDisplayError(
		codestack.Skip1,
		errtype.NewUsingTyper(basicErrTyper),
		"",
		refsCollection)
}

func (it *Collection) CompiledAsErrorWrapper(
	variant errtype.Variation,
	finalMessage string,
) *errorwrapper.Wrapper {
	if it.IsEmpty() || it.IsSilent() {
		return nil
	}

	refsCollection := refs.New(it.Length())
	refsCollection.AddMap(it.Items())
	it.Clear()

	return errorwrapper.NewMsgDisplayError(
		codestack.Skip1,
		variant,
		finalMessage,
		refsCollection)
}

func (it *Collection) CompiledInjectToErrorCollection(
	errCollection *errwrappers.Collection,
	variant errtype.Variation,
	finalMessage string,
) *errwrappers.Collection {
	errWrap := it.CompiledAsErrorWrapper(
		variant,
		finalMessage)

	return errCollection.AddWrapperPtr(
		errWrap)
}

func (it Collection) AsMetaAttributesStacker() loggerinf.MetaAttributesStacker {
	return &it
}

func (it *Collection) Clone() MetaAttributesCollector {
	if it == nil {
		return nil
	}

	collection := NewCollectionCap(it.Length())

	for key, valueAny := range it.items {
		collection.items[key] = valueAny
	}

	return collection
}

func (it Collection) AsBasicSliceContractsBinder() coreinterface.BasicSlicerContractsBinder {
	return &it
}

func (it Collection) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return &it
}

func (it Collection) AsStandardSlicerContractsBinder() coreinterface.StandardSlicerContractsBinder {
	return &it
}

func (it Collection) AsMetaAttributesCollector() MetaAttributesCollector {
	return &it
}
