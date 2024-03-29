package keyboard

import (
	"log"

	"github.com/o5h/winapi"
	"github.com/o5h/winapi/user32"
)

func WindowsVKToCode(vk winapi.WPARAM) Code {
	switch vk {

	// VK_LBUTTON             = 0x01
	// VK_RBUTTON             = 0x02
	// VK_CANCEL              = 0x03
	// VK_MBUTTON             = 0x04
	// VK_XBUTTON1            = 0x05
	// VK_XBUTTON2            = 0x06
	// VK_BACK                = 0x08
	// VK_TAB                 = 0x09
	// VK_CLEAR               = 0x0C
	// VK_RETURN              = 0x0D
	// VK_SHIFT               = 0x10
	// VK_CONTROL             = 0x11
	// VK_MENU                = 0x12
	// VK_PAUSE               = 0x13
	// VK_CAPITAL             = 0x14
	// VK_KANA                = 0x15
	// VK_HANGEUL             = 0x15
	// VK_HANGUL              = 0x15
	// VK_JUNJA               = 0x17
	// VK_FINAL               = 0x18
	// VK_HANJA               = 0x19
	// VK_KANJI               = 0x19

	case user32.VK_ESCAPE:
		return Escape // VK_ESCAPE              = 0x1B

		// VK_CONVERT             = 0x1C
		// VK_NONCONVERT          = 0x1D
		// VK_ACCEPT              = 0x1E
		// VK_MODECHANGE          = 0x1F
		// VK_SPACE               = 0x20
		// VK_PRIOR               = 0x21
		// VK_NEXT                = 0x22
		// VK_END                 = 0x23
		// VK_HOME                = 0x24
		// VK_LEFT                = 0x25
		// VK_UP                  = 0x26
		// VK_RIGHT               = 0x27
		// VK_DOWN                = 0x28
		// VK_SELECT              = 0x29
		// VK_PRINT               = 0x2A
		// VK_EXECUTE             = 0x2B
	case user32.VK_SNAPSHOT:
		return PrintScreen //VK_SNAPSHOT = 0x2C
	// VK_INSERT              = 0x2D
	// VK_DELETE              = 0x2E
	// VK_HELP                = 0x2F

	case user32.VK_KEY_0:
		return Code0
	case user32.VK_KEY_1:
		return Code1
	case user32.VK_KEY_2:
		return Code2
	case user32.VK_KEY_3:
		return Code3
	case user32.VK_KEY_4:
		return Code4
	case user32.VK_KEY_5:
		return Code5
	case user32.VK_KEY_6:
		return Code6
	case user32.VK_KEY_7:
		return Code7
	case user32.VK_KEY_8:
		return Code8
	case user32.VK_KEY_9:
		return Code9
	case user32.VK_KEY_A:
		return CodeA
	case user32.VK_KEY_B:
		return CodeB
	case user32.VK_KEY_C:
		return CodeC
	case user32.VK_KEY_D:
		return CodeD
	case user32.VK_KEY_E:
		return CodeE
	case user32.VK_KEY_F:
		return CodeF
	case user32.VK_KEY_G:
		return CodeG
	case user32.VK_KEY_H:
		return CodeH
	case user32.VK_KEY_I:
		return CodeI
	case user32.VK_KEY_J:
		return CodeJ
	case user32.VK_KEY_K:
		return CodeK
	case user32.VK_KEY_L:
		return CodeL
	case user32.VK_KEY_M:
		return CodeM
	case user32.VK_KEY_N:
		return CodeN
	case user32.VK_KEY_O:
		return CodeO
	case user32.VK_KEY_P:
		return CodeP
	case user32.VK_KEY_Q:
		return CodeQ
	case user32.VK_KEY_R:
		return CodeR
	case user32.VK_KEY_S:
		return CodeS
	case user32.VK_KEY_T:
		return CodeT
	case user32.VK_KEY_U:
		return CodeU
	case user32.VK_KEY_V:
		return CodeV
	case user32.VK_KEY_W:
		return CodeW
	case user32.VK_KEY_X:
		return CodeX
	case user32.VK_KEY_Y:
		return CodeY
	case user32.VK_KEY_Z:
		return CodeZ
	// VK_LWIN                = 0x5B
	// VK_RWIN                = 0x5C
	// VK_APPS                = 0x5D
	// VK_SLEEP               = 0x5F

	case user32.VK_NUMPAD0:
		return Keypad0
	case user32.VK_NUMPAD1:
		return Keypad1
	case user32.VK_NUMPAD2:
		return Keypad2
	case user32.VK_NUMPAD3:
		return Keypad3
	case user32.VK_NUMPAD4:
		return Keypad4
	case user32.VK_NUMPAD5:
		return Keypad5
	case user32.VK_NUMPAD6:
		return Keypad6
	case user32.VK_NUMPAD7:
		return Keypad7
	case user32.VK_NUMPAD8:
		return Keypad8
	case user32.VK_NUMPAD9:
		return Keypad9

	case user32.VK_MULTIPLY:
		return KeypadAsterisk // VK_MULTIPLY            = 0x6A
	case user32.VK_ADD:
		return KeypadPlusSign // VK_ADD                 = 0x6B
	// case user32.VK_SEPARATOR:
	// 	return KeypadHyphenMinus //VK_SEPARATOR= 0x6C

	case user32.VK_SUBTRACT:
		return KeypadHyphenMinus // VK_SUBTRACT            = 0x6D

	// VK_DECIMAL             = 0x6E
	// VK_DIVIDE              = 0x6F

	case user32.VK_F1:
		return F1 // VK_F1                  = 0x70
	case user32.VK_F2:
		return F2 // VK_F2                  = 0x71
	case user32.VK_F3:
		return F3 // VK_F3                  = 0x72
	case user32.VK_F4:
		return F4 // VK_F4                  = 0x73
	case user32.VK_F5:
		return F5 // VK_F5                  = 0x74
	case user32.VK_F6:
		return F6 // VK_F6                  = 0x75
	case user32.VK_F7:
		return F7 // VK_F7                  = 0x76
	case user32.VK_F8:
		return F8 // VK_F8                  = 0x77
	case user32.VK_F9:
		return F9 // VK_F9                  = 0x78
	case user32.VK_F10:
		return F10 // VK_F10                 = 0x79
	case user32.VK_F11:
		return F11 // VK_F11                 = 0x7A
	case user32.VK_F12:
		return F12 // VK_F12                 = 0x7B

		// VK_F13                 = 0x7C
		// VK_F14                 = 0x7D
		// VK_F15                 = 0x7E
		// VK_F16                 = 0x7F
		// VK_F17                 = 0x80
		// VK_F18                 = 0x81
		// VK_F19                 = 0x82
		// VK_F20                 = 0x83
		// VK_F21                 = 0x84
		// VK_F22                 = 0x85
		// VK_F23                 = 0x86
		// VK_F24                 = 0x87

		// VK_NUMLOCK             = 0x90
		// VK_SCROLL              = 0x91
		// VK_OEM_NEC_EQUAL       = 0x92
		// VK_OEM_FJ_JISHO        = 0x92
		// VK_OEM_FJ_MASSHOU      = 0x93
		// VK_OEM_FJ_TOUROKU      = 0x94
		// VK_OEM_FJ_LOYA         = 0x95
		// VK_OEM_FJ_ROYA         = 0x96
		// VK_LSHIFT              = 0xA0
		// VK_RSHIFT              = 0xA1
		// VK_LCONTROL            = 0xA2
		// VK_RCONTROL            = 0xA3
		// VK_LMENU               = 0xA4
		// VK_RMENU               = 0xA5
		// VK_BROWSER_BACK        = 0xA6
		// VK_BROWSER_FORWARD     = 0xA7
		// VK_BROWSER_REFRESH     = 0xA8
		// VK_BROWSER_STOP        = 0xA9
		// VK_BROWSER_SEARCH      = 0xAA
		// VK_BROWSER_FAVORITES   = 0xAB
		// VK_BROWSER_HOME        = 0xAC
		// VK_VOLUME_MUTE         = 0xAD
		// VK_VOLUME_DOWN         = 0xAE
		// VK_VOLUME_UP           = 0xAF
		// VK_MEDIA_NEXT_TRACK    = 0xB0
		// VK_MEDIA_PREV_TRACK    = 0xB1
		// VK_MEDIA_STOP          = 0xB2
		// VK_MEDIA_PLAY_PAUSE    = 0xB3
		// VK_LAUNCH_MAIL         = 0xB4
		// VK_LAUNCH_MEDIA_SELECT = 0xB5
		// VK_LAUNCH_APP1         = 0xB6
		// VK_LAUNCH_APP2         = 0xB7
		// VK_OEM_1               = 0xBA
		// VK_OEM_PLUS            = 0xBB
		// VK_OEM_COMMA           = 0xBC
		// VK_OEM_MINUS           = 0xBD
		// VK_OEM_PERIOD          = 0xBE
		// VK_OEM_2               = 0xBF
		// VK_OEM_3               = 0xC0
		// VK_OEM_4               = 0xDB
		// VK_OEM_5               = 0xDC
		// VK_OEM_6               = 0xDD
		// VK_OEM_7               = 0xDE
		// VK_OEM_8               = 0xDF
		// VK_OEM_AX              = 0xE1
		// VK_OEM_102             = 0xE2
		// VK_ICO_HELP            = 0xE3
		// VK_ICO_00              = 0xE4
		// VK_PROCESSKEY          = 0xE5
		// VK_ICO_CLEAR           = 0xE6
		// VK_OEM_RESET           = 0xE9
		// VK_OEM_JUMP            = 0xEA
		// VK_OEM_PA1             = 0xEB
		// VK_OEM_PA2             = 0xEC
		// VK_OEM_PA3             = 0xED
		// VK_OEM_WSCTRL          = 0xEE
		// VK_OEM_CUSEL           = 0xEF
		// VK_OEM_ATTN            = 0xF0
		// VK_OEM_FINISH          = 0xF1
		// VK_OEM_COPY            = 0xF2
		// VK_OEM_AUTO            = 0xF3
		// VK_OEM_ENLW            = 0xF4
		// VK_OEM_BACKTAB         = 0xF5
		// VK_ATTN                = 0xF6
		// VK_CRSEL               = 0xF7
		// VK_EXSEL               = 0xF8
		// VK_EREOF               = 0xF9
		// VK_PLAY                = 0xFA
		// VK_ZOOM                = 0xFB
		// VK_NONAME              = 0xFC
		// VK_PA1                 = 0xFD
		// VK_OEM_CLEAR           = 0xFE
	default:
		log.Println(" Not mapped yet ", vk)
	}
	return Unknown
}
