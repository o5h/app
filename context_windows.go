package app

import (
	"log"
	"time"
	"unsafe"

	"github.com/o5h/app/input/keyboard"
	"github.com/o5h/app/input/mouse"
	"github.com/o5h/egl"
	"github.com/o5h/winapi"
	"github.com/o5h/winapi/kernel32"
	"github.com/o5h/winapi/user32"
	"golang.org/x/sys/windows"
)

func create(title string, app App) *context {
	ctx := &context{App: app}
	wndproc := winapi.WNDPROC(windows.NewCallback(wndProc))
	mh, _ := kernel32.GetModuleHandle(nil)
	myicon, _ := user32.LoadIconW(0, user32.IDI_APPLICATION)
	mycursor, _ := user32.LoadCursorW(0, user32.IDC_ARROW)

	var wc user32.WNDCLASSEX
	wc.Size = uint32(unsafe.Sizeof(wc))
	wc.WndProc = wndproc
	wc.Instance = winapi.HINSTANCE(mh)
	wc.Icon = myicon
	wc.Cursor = mycursor
	wc.Background = user32.COLOR_BTNFACE + 1
	wc.MenuName = nil
	wcname, _ := windows.UTF16PtrFromString("OPENGLES_WindowClass")
	wc.ClassName = wcname
	wc.IconSm = myicon
	user32.RegisterClassExW(&wc)

	windowTitle, _ := windows.UTF16PtrFromString(title)
	user32.CreateWindowExW(
		0,
		wcname,
		windowTitle,
		// No border, no title
		user32.WS_POPUP|user32.WS_CLIPSIBLINGS|user32.WS_CLIPCHILDREN|
			user32.WS_OVERLAPPEDWINDOW,
		user32.CW_USEDEFAULT,
		user32.CW_USEDEFAULT,
		user32.CW_USEDEFAULT,
		user32.CW_USEDEFAULT,
		winapi.HWND(0),
		winapi.HMENU(0),
		winapi.HINSTANCE(mh),
		winapi.LPVOID(ctx))
	return ctx
}

func (ctx *context) mainLoop() {
	hWnd := winapi.HWND(ctx.NativeWindow)
	var message user32.Msg

	for !ctx.Done {
		gotMsg, _ := user32.PeekMessageW(&message, 0, 0, 0, user32.PM_REMOVE)
		if gotMsg == winapi.FALSE {
			user32.SendMessageW(hWnd, user32.WM_PAINT, 0, 0)
		} else {
			user32.TranslateMessage(&message)
			user32.DispatchMessageW(&message)
		}
	}
}

func (ctx *context) initialize(hWnd winapi.HWND) error {

	ctx.NativeWindow = egl.NativeWindow(hWnd)
	dc, _ := user32.GetDC(hWnd)
	ctx.NativeDisplay = egl.NativeDisplay(dc)

	user32.SetWindowLongPtrW(hWnd, user32.GWLP_USERDATA, winapi.LONG_PTR(unsafe.Pointer(ctx)))

	//init opengl
	var err error

	ctx.Context, ctx.Display, ctx.Surface, err = egl.CreateEGLSurface(ctx.NativeDisplay, ctx.NativeWindow)
	if err != nil {
		return err
	}
	err = egl.MakeCurrent(ctx.Display, ctx.Surface, ctx.Surface, ctx.Context)
	if err != nil {
		return err
	}

	err = egl.SwapInterval(ctx.Display, 1)
	if err != nil {
		return err
	}

	ctx.onCreate()

	user32.SetWindowPos(hWnd, user32.HWND_TOP, 0, 0, 1024, 768, user32.SWP_SHOWWINDOW)
	user32.ShowWindow(hWnd, user32.SW_SHOW)
	user32.UpdateWindow(hWnd)
	user32.SetFocus(hWnd)
	//user32.ShowCursor(winapi.FALSE)
	return err
}

func (ctx *context) onKey(msg winapi.UINT, wParam winapi.WPARAM, lParam winapi.LPARAM) {
	code := keyboard.WindowsVKToCode(wParam)
	char := user32.MapVirtualKeyW(winapi.UINT(wParam), user32.MAPVK_VK_TO_CHAR)
	switch msg {
	case user32.WM_KEYDOWN:
		keyboard.Keyboard[code] = keyboard.Press
		event := &keyboard.Event{
			Direction: keyboard.Press,
			Code:      code,
			Rune:      rune(char)}
		ctx.onInput(event)
	case user32.WM_KEYUP:
		keyboard.Keyboard[code] = keyboard.Release
		event := &keyboard.Event{
			Direction: keyboard.Release,
			Code:      code,
			Rune:      rune(char)}
		ctx.onInput(event)
	}
}

func (ctx *context) onMouse(hWnd winapi.HWND, msg winapi.UINT, lParam winapi.LPARAM) {
	var x, y int
	var action mouse.Action
	var btn mouse.Button
	x = int(winapi.GET_X_LPARAM(lParam))
	y = int(winapi.GET_Y_LPARAM(lParam))
	switch msg {
	case user32.WM_LBUTTONDOWN:
		action = mouse.ActionPress
		btn = mouse.ButtonLeft
	}
	ctx.onInput(&mouse.Event{
		Action: action,
		X:      x,
		Y:      y,
		Button: btn})
}

func wndProc(hWnd winapi.HWND, msg winapi.UINT, wParam winapi.WPARAM, lParam winapi.LPARAM) (rc winapi.LRESULT) {
	var ctx *context
	ptr, _ := user32.GetWindowLongPtrW(hWnd, user32.GWLP_USERDATA)
	if ptr != 0 {
		ctx = (*context)(unsafe.Pointer(ptr))
	}
	switch msg {
	case user32.WM_CREATE:
		create := (*user32.CREATESTRUCTW)(unsafe.Pointer(lParam))
		ctx = (*context)(unsafe.Pointer(create.CreateParams))
		err := ctx.initialize(hWnd)
		if err != nil {
			log.Fatal(err)
		}
	case user32.WM_PAINT:
		now := time.Now()
		delta := now.Sub(ctx.LastUpdateTime).Seconds()
		ctx.LastUpdateTime = now
		ctx.onUpdate(delta)
	case user32.WM_SIZE:
		w := int(winapi.LOWORD(winapi.DWORD(lParam)))
		h := int(winapi.HIWORD(winapi.DWORD(lParam)))
		ctx.onResize(w, h)
	case user32.WM_CLOSE:
		ctx.onClose()
	case user32.WM_DESTROY:
		user32.PostQuitMessage(0)
	case user32.WM_KEYDOWN, user32.WM_KEYUP:
		ctx.onKey(msg, wParam, lParam)
	case user32.WM_LBUTTONDOWN,
		user32.WM_LBUTTONUP,
		user32.WM_MOUSEMOVE:
		ctx.onMouse(hWnd, msg, lParam)
	default:
		rc = user32.DefWindowProcW(hWnd, msg, wParam, lParam)
	}
	return
}
