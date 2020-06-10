package main

import "C"

import (
	"fmt"
	"net/url"
	"unsafe"

	"github.com/imroc/req"
)

func unmarshallArguments(argc int, argv **C.char) []string {
	slice := (*[1 << 28]*C.char)(unsafe.Pointer(argv))[:argc:argc]
	arr := make([]string, argc)
	for i := range arr {
		arr[i] = C.GoString(slice[i])
	}

	return arr
}

//export sendGetRequest
func sendGetRequest(argc int, argv **C.char) *C.char {
	args := unmarshallArguments(argc, argv)
	
	if len(args) < 2 {
		return C.CString("status=0&error=Invalid+Arguments")
	}
	
	request := req.New()
	requestURL := fmt.Sprintf("%s%s", args[0], args[1])

	var requestHeaders req.Header = nil
	if len(args) >= 4 {
		requestHeaders = make(map[string]string)
		for i := 2; i < len(args); i += 2 {
			requestHeaders[args[i]] = args[i+1]
		}
	}

	r, err := request.Get(requestURL, requestHeaders)
	if err != nil {
		errString := fmt.Sprintf("status=0&error=%s", url.QueryEscape(err.Error()))
		retString := C.CString(errString)
		//defer C.free(unsafe.Pointer(retString))
		return retString
	}

	response := r.Response()
	respBody, _ := r.ToString()
	respString := fmt.Sprintf("status=%d&content=%s", response.StatusCode, url.QueryEscape(respBody))
	retString := C.CString(respString)
	//Let's hope BYOND will free this
	//defer C.free(unsafe.Pointer(retString))

	return retString
}

//export testFunc
func testFunc(argc int, argv **C.char) *C.char {
	args := unmarshallArguments(argc, argv)
	outstring := fmt.Sprintf("Hello %s! %s | args was %d long. argc was %d", args[0], args[1], len(args), argc)
	retstring := C.CString(outstring)
	//defer C.free(unsafe.Pointer(retstring))
	return retstring
}

func main() {
	//Do nothing
}
