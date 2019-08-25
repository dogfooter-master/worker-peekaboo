package service

import (
	"fmt"
	"github.com/TheTitanrain/w32"
	"github.com/hnakamur/w32syscall"
	"image"
	"image/png"
	"math"
	"os"
	"reflect"
	"regexp"
	"syscall"
	"unsafe"
)

type PeekabooWin struct {
	Wildcard      string
	HWNDList      []w32.HWND
	SideHWNDList  []w32.HWND
	ParentHWNDMap map[w32.HWND]w32.HWND
	SideHWNDMap   map[w32.HWND]w32.HWND
}

func (p *PeekabooWin) CallbackMomoChildProcess(hWnd w32.HWND, lParam w32.LPARAM) w32.LRESULT {
	s := w32.GetWindowText(hWnd)
	fmt.Println("MOMO child:", s, "# hWnd:", hWnd, "# lParam:", lParam)
	if len(s) > 3 {
		isFound := false
		for _, v := range p.HWNDList {
			if v == hWnd {
				isFound = true
			}
		}
		if !isFound {
			p.HWNDList = append(p.HWNDList, hWnd)
			if p.ParentHWNDMap == nil {
				p.ParentHWNDMap = make(map[w32.HWND]w32.HWND)
			}
			p.ParentHWNDMap[hWnd] = w32.HWND(lParam)
		}
	}
	return 1
}
func (p *PeekabooWin) CallbackMemuChildProcess(hWnd w32.HWND, lParam w32.LPARAM) w32.LRESULT {
	s := w32.GetWindowText(hWnd)
	fmt.Println("Memu child:", s, "# hWnd:", hWnd, "# lParam:", lParam)
	if s == "RenderWindowWindow" {
		isFound := false
		for _, v := range p.HWNDList {
			if v == hWnd {
				isFound = true
			}
		}
		if !isFound {
			p.HWNDList = append(p.HWNDList, hWnd)
			if p.ParentHWNDMap == nil {
				p.ParentHWNDMap = make(map[w32.HWND]w32.HWND)
			}
			p.ParentHWNDMap[hWnd] = w32.HWND(lParam)
		}
	}
	return 1
}
func (p *PeekabooWin) FindWindowWildcard() {
	err := w32syscall.EnumWindows(func(hWnd syscall.Handle, lParam uintptr) bool {
		var h w32.HWND
		h = w32.HWND(hWnd)
		if w32.IsWindowVisible(h) == false {
			return true
		}

		rect := w32.GetWindowRect(h)
		title := w32.GetWindowText(h)

		match, _ := regexp.MatchString(p.Wildcard, title)
		if match {
			//fmt.Fprintf(os.Stderr, "DEBUG: %v, %v, %v\n", title, match, *rect)
			//fmt.Fprintf(os.Stderr, "DEBUG: %v, %v\n", GetConfigSizeWidth(), GetConfigSizeHeight())
			//fmt.Fprintf(os.Stderr, "DEBUG: (%v, %v)\n", math.Abs(float64(rect.Right-rect.Left)), float64(GetConfigSizeWidth()+38))
			if math.Abs(float64(rect.Right-rect.Left)) == float64(GetConfigSizeWidth()+38) {
				// MOMO
				w32.EnumChildWindows(h, p.CallbackMomoChildProcess, w32.LPARAM(h))
			} else if math.Abs(float64(rect.Right-rect.Left)) == float64(GetConfigSizeWidth()+4) {
				fmt.Println("Nox child:", title, "# hWnd:", h)
				p.HWNDList = append(p.HWNDList, h)
			} else if math.Abs(float64(rect.Right-rect.Left)) == float64(GetConfigSizeWidth()+40) {
				w32.EnumChildWindows(h, p.CallbackMemuChildProcess, w32.LPARAM(h))
			}
		}

		side1, _ := regexp.MatchString("nox", title)
		side2, _ := regexp.MatchString("Nox", title)
		side3, _ := regexp.MatchString("Form", title)
		if side1 || side2 || side3 {
			fmt.Println(">>", math.Abs(float64(rect.Right-rect.Left))-float64(36), math.Abs(float64(rect.Bottom-rect.Top-GetConfigSizeHeight())))
			if math.Abs(float64(rect.Right-rect.Left))-float64(36) < float64(40) &&
				math.Abs(float64(rect.Bottom-rect.Top-GetConfigSizeHeight())) < float64(40) {
				isFound := false
				for _, v := range p.SideHWNDList {
					if v == h {
						isFound = true
					}
				}
				if !isFound {
					p.SideHWNDList = append(p.SideHWNDList, h)
				}
			}
		}
		return true
	}, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "FindWindowWildcard: %v", err)
	}

	for _, e := range p.SideHWNDList {
		s := w32.GetWindowRect(e)
		for _, ee := range p.HWNDList {
			w := w32.GetWindowRect(ee)
			if math.Abs(math.Abs(float64(s.Left-w.Left))-float64(GetConfigSizeWidth()+4)) < 40 &&
				math.Abs(math.Abs(float64(s.Top-w.Top))-float64(30)) < 40 &&
				math.Abs(math.Abs(float64(s.Right-w.Right))-float64(36)) < 40 &&
				math.Abs(float64(s.Bottom-w.Bottom)) < 40 {
				if p.SideHWNDMap == nil {
					p.SideHWNDMap = make(map[w32.HWND]w32.HWND)
				}
				p.SideHWNDMap[ee] = e
			}
		}
	}
}
func (p *PeekabooWin) GetWindowScreenShot(hWnd w32.HWND) {

	var mod = syscall.NewLazyDLL("user32.dll")
	//var proc = mod.NewProc("MessageBoxW")
	//var MB_YESNOCANCEL = 0x00000003
	//
	//ret, _, _ := proc.Call(0,
	//	uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("This test is Done."))),
	//	uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Done Title"))),
	//	uintptr(MB_YESNOCANCEL))
	//
	//fmt.Println("Return: ", ret)
	var GetWindowDC = mod.NewProc("GetWindowDC")

	rect := w32.GetClientRect(hWnd)
	w := uint(rect.Right - rect.Left)
	h := uint(rect.Bottom - rect.Top)
	fmt.Println("#", w, "#", h)

	hDc, _, _ := GetWindowDC.Call(uintptr(hWnd))
	if hDc == 0 {
		fmt.Println("error GetWindowDC")
		return
	}
	hdcSrc := w32.HDC(hDc)
	defer w32.ReleaseDC(hWnd, hdcSrc)

	hdcDest := w32.CreateCompatibleDC(hdcSrc)
	if hdcDest == 0 {
		fmt.Println("error CreateCompatibleDC")
		return
	}
	defer w32.DeleteDC(hdcDest)

	bt := w32.BITMAPINFO{}
	bt.BmiHeader.BiSize = uint32(reflect.TypeOf(bt.BmiHeader).Size())
	bt.BmiHeader.BiWidth = int32(w)
	bt.BmiHeader.BiHeight = int32(-h)
	bt.BmiHeader.BiPlanes = 1
	bt.BmiHeader.BiBitCount = 32
	bt.BmiHeader.BiCompression = w32.BI_RGB

	ptr := unsafe.Pointer(uintptr(0))

	hBmp := w32.CreateDIBSection(hdcDest, &bt, w32.DIB_RGB_COLORS, &ptr, 0, 0)
	if hBmp == 0 {
		fmt.Println("Could not Create DIB Section")
		return
	}
	if hBmp == w32.InvalidParameter {
		fmt.Println("Could not Create DIB Section")
		return
	}
	defer w32.DeleteObject(w32.HGDIOBJ(hBmp))

	obj := w32.SelectObject(hdcDest, w32.HGDIOBJ(hBmp))
	if obj == 0 {
		fmt.Println("error occurred and the selected object is not a region")
		return
	}
	if obj == 0xffffffff { //GDI_ERROR
		fmt.Println("GDI_ERROR while calling SelectObject")
		return
	}
	defer w32.DeleteObject(obj)

	//Note:BitBlt contains bad error handling, we will just assume it works and if it doesn't it will panic :x
	w32.BitBlt(hdcDest, 0, 0, int(w), int(h), hdcSrc, 0, 0, w32.SRCCOPY)

	var slice []byte
	hDrp := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	hDrp.Data = uintptr(ptr)
	hDrp.Len = int(w * h * 4)
	hDrp.Cap = int(w * h * 4)

	imageBytes := make([]byte, len(slice))

	for i := 0; i < len(imageBytes); i += 4 {
		imageBytes[i], imageBytes[i+2], imageBytes[i+1], imageBytes[i+3] = slice[i+2], slice[i], slice[i+1], slice[i+3]
	}

	img := &image.RGBA{imageBytes, int(4 * w), image.Rect(0, 0, int(w), int(h))}
	f, err := os.Create("img.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_ = png.Encode(f, img)


	//hBitmap := w32.CreateCompatibleBitmap(hdcSrc, w, h)
	//hOld := w32.SelectObject(hdcDest, w32.HGDIOBJ(hBitmap))
	//w32.BitBlt(hdcDest, 0, 0, int(w), int(h), hdcSrc, 0, 0, w32.SRCCOPY)
	//w32.SelectObject(hdcDest, hOld)
	//w32.ReleaseDC(hWnd, hdcSrc)

	//w32.DeleteObject(w32.HGDIOBJ(hBitmap))
}
