package service

import (
	"fmt"
	"github.com/TheTitanrain/w32"
	"github.com/hnakamur/w32syscall"
	"image"
	"math"
	"os"
	"reflect"
	"regexp"
	"syscall"
	"unsafe"
)

type Position struct {
	X int
	Y int
}
type HandleInfo struct {
	Handle       w32.HWND
	ParentHandle w32.HWND
	SideHandle   w32.HWND
	Title        string
	Rect         *w32.RECT
	Type         string
	Back         Position
	Home         Position
	Recent       Position
}

type PeekabooWin struct {
	Wildcard  string
	HandleMap map[w32.HWND]*HandleInfo
	//HWNDList      []w32.HWND
	//SideHWNDList  []w32.HWND
	//ParentHWNDMap map[w32.HWND]w32.HWND
	//SideHWNDMap   map[w32.HWND]w32.HWND
	//TitleMap      map[w32.HWND]string
	//RectMap       map[w32.HWND]*w32.RECT
}

var peekabooWindowInfo PeekabooWin
var noxSideWindows []HandleInfo

func init() {
	peekabooWindowInfo.HandleMap = make(map[w32.HWND]*HandleInfo)
}

func (p *PeekabooWin) NextHandle(hWnd w32.HWND) (handle w32.HWND) {
	i := 0
	var handleList []w32.HWND

	for k := range p.HandleMap {
		handleList = append(handleList, k)
	}
	for _, e := range handleList {
		i += 1
		if e == hWnd {
			break
		}
	}
	if i == len(handleList) {
		i = 0
	}
	handle = handleList[i]

	return
}

//func (p *PeekabooWin) CallbackMomoChildProcess(hWnd w32.HWND, lParam w32.LPARAM) w32.LRESULT {
//	s := w32.GetWindowText(hWnd)
//	//fmt.Println("MOMO child:", s, "# hWnd:", hWnd, "# lParam:", lParam)
//	if len(s) > 3 {
//		if _, ok := p.HandleMap[hWnd]; ok {
//
//			p.HWNDList = append(p.HWNDList, hWnd)
//			if p.ParentHWNDMap == nil {
//				p.ParentHWNDMap = make(map[w32.HWND]w32.HWND)
//			}
//			p.ParentHWNDMap[hWnd] = w32.HWND(lParam)
//
//			if p.TitleMap == nil {
//				p.TitleMap = make(map[w32.HWND]string)
//			}
//			p.TitleMap[hWnd] = w32.GetWindowText(w32.HWND(lParam))
//
//			rect := w32.GetWindowRect(hWnd)
//			if p.RectMap == nil {
//				p.RectMap = make(map[w32.HWND]*w32.RECT)
//			}
//			p.RectMap[hWnd] = rect
//		}
//	}
//	return 1
//}
//func (p *PeekabooWin) CallbackMemuChildProcess(hWnd w32.HWND, lParam w32.LPARAM) w32.LRESULT {
//	s := w32.GetWindowText(hWnd)
//	//fmt.Println("Memu child:", s, "# hWnd:", hWnd, "# lParam:", lParam)
//	if s == "RenderWindowWindow" {
//		isFound := false
//		for _, v := range p.HWNDList {
//			if v == hWnd {
//				isFound = true
//			}
//		}
//		if !isFound {
//			p.HWNDList = append(p.HWNDList, hWnd)
//			if p.ParentHWNDMap == nil {
//				p.ParentHWNDMap = make(map[w32.HWND]w32.HWND)
//			}
//			p.ParentHWNDMap[hWnd] = w32.HWND(lParam)
//
//			if p.TitleMap == nil {
//				p.TitleMap = make(map[w32.HWND]string)
//			}
//			p.TitleMap[hWnd] = w32.GetWindowText(w32.HWND(lParam))
//
//			rect := w32.GetWindowRect(hWnd)
//			if p.RectMap == nil {
//				p.RectMap = make(map[w32.HWND]*w32.RECT)
//			}
//			p.RectMap[hWnd] = rect
//		}
//	}
//	return 1
//}
//func (p *PeekabooWin) FindWindowWildcard() {
//	//defer timeTrack(time.Now(), GetFunctionName())
//	err := w32syscall.EnumWindows(func(hWnd syscall.Handle, lParam uintptr) bool {
//		var h w32.HWND
//		h = w32.HWND(hWnd)
//		if w32.IsWindowVisible(h) == false {
//			return true
//		}
//
//		rect := w32.GetWindowRect(h)
//		title := w32.GetWindowText(h)
//
//		match, _ := regexp.MatchString(p.Wildcard, title)
//		if match {
//			//fmt.Fprintf(os.Stderr, "DEBUG: %v, %v, %v\n", title, match, *rect)
//			//fmt.Fprintf(os.Stderr, "DEBUG: %v, %v\n", GetConfigSizeWidth(), GetConfigSizeHeight())
//			//fmt.Fprintf(os.Stderr, "DEBUG: (%v, %v)\n", math.Abs(float64(rect.Right-rect.Left)), float64(GetConfigSizeWidth()+38))
//
//			if math.Abs(float64(rect.Right-rect.Left)) == float64(GetConfigSizeWidth()+38) {
//				// MOMO
//				w32.EnumChildWindows(h, p.CallbackMomoChildProcess, w32.LPARAM(h))
//			} else if math.Abs(float64(rect.Right-rect.Left)) == float64(GetConfigSizeWidth()+4) {
//				//fmt.Println("Nox child:", title, "# hWnd:", h)
//
//				if p.TitleMap == nil {
//					p.TitleMap = make(map[w32.HWND]string)
//				}
//				p.TitleMap[h] = title
//
//				if p.RectMap == nil {
//					p.RectMap = make(map[w32.HWND]*w32.RECT)
//				}
//				p.RectMap[h] = rect
//
//				p.HWNDList = append(p.HWNDList, h)
//			} else if math.Abs(float64(rect.Right-rect.Left)) == float64(GetConfigSizeWidth()+40) {
//				w32.EnumChildWindows(h, p.CallbackMemuChildProcess, w32.LPARAM(h))
//			}
//		}
//
//		side1, _ := regexp.MatchString("nox", title)
//		side2, _ := regexp.MatchString("Nox", title)
//		side3, _ := regexp.MatchString("Form", title)
//		if side1 || side2 || side3 {
//			fmt.Println(">>", math.Abs(float64(rect.Right-rect.Left))-float64(36), math.Abs(float64(rect.Bottom-rect.Top-GetConfigSizeHeight())))
//			if math.Abs(float64(rect.Right-rect.Left))-float64(36) < float64(40) &&
//				math.Abs(float64(rect.Bottom-rect.Top-GetConfigSizeHeight())) < float64(40) {
//				isFound := false
//				for _, v := range p.SideHWNDList {
//					if v == h {
//						isFound = true
//					}
//				}
//				if !isFound {
//					p.SideHWNDList = append(p.SideHWNDList, h)
//				}
//			}
//		}
//		return true
//	}, 0)
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "FindWindowWildcard: %v", err)
//	}
//
//	for _, e := range p.SideHWNDList {
//		s := w32.GetWindowRect(e)
//		for _, ee := range p.HWNDList {
//			w := w32.GetWindowRect(ee)
//			if math.Abs(math.Abs(float64(s.Left-w.Left))-float64(GetConfigSizeWidth()+4)) < 40 &&
//				math.Abs(math.Abs(float64(s.Top-w.Top))-float64(30)) < 40 &&
//				math.Abs(math.Abs(float64(s.Right-w.Right))-float64(36)) < 40 &&
//				math.Abs(float64(s.Bottom-w.Bottom)) < 40 {
//				if p.SideHWNDMap == nil {
//					p.SideHWNDMap = make(map[w32.HWND]w32.HWND)
//				}
//				p.SideHWNDMap[ee] = e
//			}
//		}
//	}
//}

func (p *PeekabooWin) GetCommandPosition(handle w32.HWND, sideHandle w32.HWND, emulatorType string, command string) (position Position) {

	switch emulatorType {
	case "LDPlayer":
		rect := w32.GetWindowRect(handle)
		sideRect := w32.GetWindowRect(sideHandle)
		titleHeight := int32(math.Abs(float64(sideRect.Top - rect.Top)))
		adjPosition := GetConfigButtonPosition(emulatorType, titleHeight, command)
		position.X = int(math.Abs(float64(sideRect.Right-sideRect.Left))) + int(adjPosition.X)
		position.Y = int(math.Abs(float64(sideRect.Bottom-sideRect.Top))) + int(adjPosition.Y)

		fmt.Fprintf(os.Stderr, "%v %v Button: %v, %v\n", emulatorType, titleHeight, command, position)
	case "Nox":
		rect := w32.GetWindowRect(handle)
		sideRect := w32.GetWindowRect(sideHandle)
		titleHeight := int32(math.Abs(float64(sideRect.Top - rect.Top)))
		adjPosition := GetConfigButtonPosition(emulatorType, titleHeight, command)
		position.X = int(math.Abs(float64(sideRect.Right-sideRect.Left))) + int(adjPosition.X)
		position.Y = int(math.Abs(float64(sideRect.Bottom-sideRect.Top))) + int(adjPosition.Y)

		fmt.Fprintf(os.Stderr, "%v %v Button: %v, %v\n", emulatorType, titleHeight, command, position)
	}

	return
}

func (p *PeekabooWin) CallbackChildProcess(hWnd w32.HWND, lParam w32.LPARAM) w32.LRESULT {
	if p.HandleMap == nil {
		p.HandleMap = make(map[w32.HWND]*HandleInfo)
	}
	s := w32.GetWindowText(hWnd)
	rect := w32.GetWindowRect(hWnd)
	fmt.Println("+- child:", s, "# hWnd:", hWnd, "# lParam:", lParam, "# title:", s, "# rect:", rect)
	var sideHandle w32.HWND
	var emulatorType string
	var back Position
	var home Position
	var recent Position
	switch s {
	case "TheRender":
		sideHandle = w32.HWND(lParam)
		emulatorType = "LDPlayer"
		back = p.GetCommandPosition(hWnd, w32.HWND(lParam), emulatorType, "back")
		home = p.GetCommandPosition(hWnd, w32.HWND(lParam), emulatorType, "home")
		recent = p.GetCommandPosition(hWnd, w32.HWND(lParam), emulatorType, "recent")
	case "ScreenBoardClassWindow":
		emulatorType = "Nox"
	}

	if len(emulatorType) > 0 {
		if _, ok := p.HandleMap[hWnd]; !ok {
			h := HandleInfo{
				Handle:       hWnd,
				ParentHandle: w32.HWND(lParam),
				Title:        w32.GetWindowText(w32.HWND(lParam)),
				Rect:         rect,
				SideHandle:   sideHandle,
				Type:         emulatorType,
				Back:         back,
				Home:         home,
				Recent:       recent,
			}
			p.HandleMap[hWnd] = &h
		}
	}
	return 1
}
func (p *PeekabooWin) FindWindowWildcard2() {
	//defer timeTrack(time.Now(), GetFunctionName())
	noxSideWindows = []HandleInfo{}
	err := w32syscall.EnumWindows(func(hWnd syscall.Handle, lParam uintptr) bool {
		var h w32.HWND
		h = w32.HWND(hWnd)
		if w32.IsWindowVisible(h) == false {
			return true
		}

		rect := w32.GetWindowRect(h)
		title := w32.GetWindowText(h)

		// Nox 사이드바
		if title == "Form" || title == "Nox" || title == "nox" {
			ok := w32.EnumChildWindows(h, p.CallbackChildProcess, w32.LPARAM(h))
			if !ok {
				fmt.Fprintf(os.Stderr, "녹스 사이드바: %v, %v\n", title, *rect)
				noxSideWindows = append(noxSideWindows,
					HandleInfo{
						Handle: h,
						Rect:   rect,
						Title:  title,
					})
			}
		}

		match, _ := regexp.MatchString(p.Wildcard, title)
		if len(title) > 0 && match {
			if w32.IsWindowVisible(h) {
				fmt.Fprintf(os.Stderr, "+ %v, %v, %v\n", title, match, *rect)
				//fmt.Fprintf(os.Stderr, "DEBUG: %v, %v\n", GetConfigSizeWidth(), GetConfigSizeHeight())
				//fmt.Fprintf(os.Stderr, "DEBUG: (%v, %v)\n", math.Abs(float64(rect.Right-rect.Left)), float64(GetConfigSizeWidth()+38))
				w32.EnumChildWindows(h, p.CallbackChildProcess, w32.LPARAM(h))
			}
		}

		return true
	}, 0)

	for _, v := range p.HandleMap {
		if v.Type == "Nox" {
			for _, e := range noxSideWindows {
				if  math.Abs(float64(v.Rect.Right - e.Rect.Left)) == float64(2) &&
					math.Abs(float64(v.Rect.Bottom - e.Rect.Bottom)) == float64(2) {
					v.SideHandle = e.Handle
					v.Back = p.GetCommandPosition(v.Handle, v.SideHandle, v.Type, "back")
					v.Home = p.GetCommandPosition(v.Handle, v.SideHandle, v.Type, "home")
					v.Recent = p.GetCommandPosition(v.Handle, v.SideHandle, v.Type, "recent")
				}
			}
		}
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "FindWindowWildcard: %v", err)
	}
}
func (p *PeekabooWin) GetWindowScreenShot(hWnd w32.HWND) (img *image.RGBA) {

	//defer timeTrack(time.Now(), GetFunctionName())
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

	defer func() {
		if v := recover(); v != nil {
			img = nil
		}
	}()

	rect := w32.GetClientRect(hWnd)
	w := uint(rect.Right - rect.Left)
	h := uint(rect.Bottom - rect.Top)
	//fmt.Println("#", w, "#", h)

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

	//defer timeTrack(time.Now(), GetFunctionName()+"-1")

	for i := 0; i < len(imageBytes); i += 4 {
		imageBytes[i], imageBytes[i+2], imageBytes[i+1], imageBytes[i+3] = slice[i+2], slice[i], slice[i+1], slice[i+3]
	}

	img = &image.RGBA{imageBytes, int(4 * w), image.Rect(0, 0, int(w), int(h))}

	//f, err := os.Create("img.jpg")
	//if err != nil {
	//	panic(err)
	//}
	//defer f.Close()
	//
	//_ = jpeg.Encode(f, img, nil)

	return

	//hBitmap := w32.CreateCompatibleBitmap(hdcSrc, w, h)
	//hOld := w32.SelectObject(hdcDest, w32.HGDIOBJ(hBitmap))
	//w32.BitBlt(hdcDest, 0, 0, int(w), int(h), hdcSrc, 0, 0, w32.SRCCOPY)
	//w32.SelectObject(hdcDest, hOld)
	//w32.ReleaseDC(hWnd, hdcSrc)

	//w32.DeleteObject(w32.HGDIOBJ(hBitmap))
}
func (p *PeekabooWin) MouseDown(hWnd w32.HWND, x, y float32) {
	if _, ok := p.HandleMap[hWnd]; ok {
		rect := p.HandleMap[hWnd].Rect
		w := math.Abs(float64(rect.Right - rect.Left))
		h := math.Abs(float64(rect.Bottom - rect.Top))

		posX := int(w * float64(x))
		posY := int(h * float64(y))

		p.MouseDown2(hWnd, posX, posY)
	}
}
func (p *PeekabooWin) MouseDown2(hWnd w32.HWND, x, y int) {
	fmt.Fprintf(os.Stderr, "MouseDown2 #%v - X: %v, Y: %v\n", w32.GetWindowText(hWnd), x, y)
	w32.PostMessage(hWnd, w32.WM_LBUTTONDOWN, w32.VK_LBUTTON, MakeLong(x, y))
}
func (p *PeekabooWin) MouseUp(hWnd w32.HWND, x, y float32) {
	if _, ok := p.HandleMap[hWnd]; ok {
		rect := p.HandleMap[hWnd].Rect
		w := math.Abs(float64(rect.Right - rect.Left))
		h := math.Abs(float64(rect.Bottom - rect.Top))

		posX := int(w * float64(x))
		posY := int(h * float64(y))

		p.MouseUp2(hWnd, posX, posY)
	}
}
func (p *PeekabooWin) MouseUp2(hWnd w32.HWND, x, y int) {
	fmt.Fprintf(os.Stderr, "MouseUp2 #%v - X: %v, Y: %v\n", w32.GetWindowText(hWnd), x, y)
	w32.PostMessage(hWnd, w32.WM_LBUTTONUP, 0, MakeLong(x, y))
}
func (p *PeekabooWin) MouseMove(hWnd w32.HWND, x, y float32) {
	if _, ok := p.HandleMap[hWnd]; ok {
		rect := p.HandleMap[hWnd].Rect
		w := math.Abs(float64(rect.Right - rect.Left))
		h := math.Abs(float64(rect.Bottom - rect.Top))

		posX := int(w * float64(x))
		posY := int(h * float64(y))

		p.MouseMove2(hWnd, posX, posY)
	}
}
func (p *PeekabooWin) MouseMove2(hWnd w32.HWND, x, y int) {
	fmt.Fprintf(os.Stderr, "MouseMove2 #%v - X: %v, Y: %v\n", w32.GetWindowText(hWnd), x, y)
	w32.PostMessage(hWnd, w32.WM_MOUSEMOVE, w32.VK_LBUTTON, MakeLong(x, y))
}
func (p *PeekabooWin) Back(hWnd w32.HWND) {
	if _, ok := p.HandleMap[hWnd]; ok {
		sideHandle := p.HandleMap[hWnd].SideHandle
		buttonPosX := p.HandleMap[hWnd].Back.X
		buttonPosY := p.HandleMap[hWnd].Back.Y
		p.MouseClick(sideHandle, buttonPosX, buttonPosY)
	}
}
func (p *PeekabooWin) Home(hWnd w32.HWND) {
	if _, ok := p.HandleMap[hWnd]; ok {
		sideHandle := p.HandleMap[hWnd].SideHandle
		buttonPosX := p.HandleMap[hWnd].Home.X
		buttonPosY := p.HandleMap[hWnd].Home.Y
		p.MouseClick(sideHandle, buttonPosX, buttonPosY)
	}
}
func (p *PeekabooWin) Recent(hWnd w32.HWND) {
	if _, ok := p.HandleMap[hWnd]; ok {
		sideHandle := p.HandleMap[hWnd].SideHandle
		buttonPosX := p.HandleMap[hWnd].Recent.X
		buttonPosY := p.HandleMap[hWnd].Recent.Y
		p.MouseClick(sideHandle, buttonPosX, buttonPosY)
	}
}
func (p *PeekabooWin) MouseClick(hWnd w32.HWND, x, y int) {
	p.MouseDown2(hWnd, x, y)
	p.MouseUp2(hWnd, x, y)
}
func MakeLong(x, y int) uintptr {
	return uintptr(y<<16 | x&0xFFFF)
}
