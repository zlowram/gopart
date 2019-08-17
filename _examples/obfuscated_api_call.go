package main

import (
	"crypto/sha256"
	"encoding/hex"
	"unsafe"

	"github.com/zlowram/gopart/obfuscation"
	windows "github.com/zlowram/gowin"
)

func hash(str string) string {
	salt := "th1s_1s_a_s4lt!"
	h := sha256.New()
	h.Write(append([]byte(salt), []byte(str)...))
	return hex.EncodeToString(h.Sum(nil)[:])
}

func encryptDecrypt(input string) string {
	key := "sUp3r_3ncr1pti0n_k3y!"
	output := ""
	for i := 0; i < len(input); i++ {
		output += string(input[i] ^ key[i%len(key)])
	}

	return output
}

func main() {
	messageTitle := windows.NewUnicodeString("GOPART")
	message := windows.NewUnicodeString("This call to MessageBoxW was done by finding the function address via PEB")
	wapi := obfuscation.NewWindowsApi(hash, encryptDecrypt)
	wapi.Call(
		encryptDecrypt("user32.dll"),
		hash("MessageBoxW"),
		uintptr(0),
		uintptr(unsafe.Pointer(message.Buffer)),
		uintptr(unsafe.Pointer(messageTitle.Buffer)),
		uintptr(0),
	)
}
