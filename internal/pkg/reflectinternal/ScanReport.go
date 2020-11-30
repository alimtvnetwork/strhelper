package reflectinternal

import (
	"fmt"
	"reflect"
	"unsafe"
)

type ScanReport struct {
	ReflectionValue         *reflect.Value
	IndirectReflectionValue *reflect.Value
	ReflectionType          reflect.Type
	IndirectReflectionType  *reflect.Type
	unsafePointer           *unsafe.ArbitraryType
	VisitedTypes            *[]reflect.Type
	FinalDeductedType       reflect.Type
	IsPointer               bool
	IsInterface             bool
	IsInt                   bool
	IsBool                  bool
	IsByte                  bool
	IsString                bool
	IsSlice                 bool
	IsMap                   bool
	IsFunc                  bool
	IsArray                 bool
	TypeName                string
}

func NewScanReport(any interface{}, maxTries int) ScanReport {
	reflectValueOfAny := reflect.ValueOf(any)
	kind := reflectValueOfAny.Kind()
	isPtr := kind == reflect.Ptr
	reflectValue := &reflectValueOfAny
	reflectionType := reflect.TypeOf(any)
	finalReflectType, visitedTypes := GetElementTypesMaxTry(any, maxTries)
	typeName := fmt.Sprintf("%T", any)

	if isPtr {
		iReflectVal := reflect.Indirect(reflect.ValueOf(any))
		indirectReflectValue := &iReflectVal
		reflectType := iReflectVal.Type()
		iReflectType := &reflectType
		currentKind := reflectType.Kind()

		return ScanReport{
			ReflectionValue:         reflectValue,
			IndirectReflectionValue: indirectReflectValue,
			ReflectionType:          reflectionType,
			IndirectReflectionType:  iReflectType,
			unsafePointer:           unsafe.Pointer(reflectValue.Pointer()),
			VisitedTypes:            &visitedTypes,
			FinalDeductedType:       finalReflectType,
			IsPointer:               isPtr,
			IsInterface:             finalReflectType == AnyType,
			IsInt:                   finalReflectType == IntegerType,
			IsBool:                  finalReflectType == BooleanType,
			IsByte:                  finalReflectType == BytesType,
			IsString:                finalReflectType == StringType,
			IsSlice:                 currentKind == reflect.Slice,
			IsMap:                   currentKind == reflect.Map,
			IsFunc:                  currentKind == reflect.Func,
			IsArray:                 currentKind == reflect.Array,
			TypeName:                typeName,
		}
	}

	currentKind := reflectionType.Kind()

	return ScanReport{
		ReflectionValue:         reflectValue,
		IndirectReflectionValue: nil,
		ReflectionType:          reflectionType,
		IndirectReflectionType:  nil,
		unsafePointer:           nil,
		VisitedTypes:            &visitedTypes,
		FinalDeductedType:       finalReflectType,
		IsPointer:               isPtr,
		IsInterface:             finalReflectType == AnyType,
		IsInt:                   finalReflectType == IntegerType,
		IsBool:                  finalReflectType == BooleanType,
		IsByte:                  finalReflectType == BytesType,
		IsString:                finalReflectType == StringType,
		IsSlice:                 currentKind == reflect.Slice,
		IsMap:                   currentKind == reflect.Map,
		IsFunc:                  currentKind == reflect.Func,
		IsArray:                 currentKind == reflect.Array,
		TypeName:                typeName,
	}
}
