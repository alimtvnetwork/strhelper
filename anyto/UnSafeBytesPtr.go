package anyto

import (
	"unsafe"

	"gitlab.com/evatix-go/strhelper/internal/reflectinternal"
)

func UnSafeBytesPtr(any interface{}) *[]byte {
	if any == nil {
		return nil
	}

	isBytes, allBytes := reflectinternal.IsBytesOrBytesPointer(any)

	if isBytes {
		return &allBytes
	}

	pointerInfo := reflectinternal.GetPointerInfo(&any)

	if pointerInfo.IsPointer {
		//goland:noinspection GoVetUnsafePointer
		return (*[]byte)(unsafe.Pointer(pointerInfo.Pointer))
	}

	pointerInfo = reflectinternal.GetPointerInfo(&any)

	if pointerInfo.IsPointer {
		//goland:noinspection ALL
		return (*[]byte)(unsafe.Pointer(pointerInfo.Pointer))
	}

	return nil
}
