// pkg/platform/vcs_stars_keyboard.go
// Copyright(c) 2022-2024 vice contributors, licensed under the GNU Public License, Version 3.
// SPDX: GPL-3.0-only

package platform

import (
	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type keyMapping struct {
	key   imgui.Key
	ctrl  bool
	shift bool
}

func (g *glfwPlatform) GetVcsStarsKeyboard() *KeyboardState {
	keyboard := &KeyboardState{
		Pressed:   make(map[imgui.Key]any),
		HeldFKeys: g.heldFKeys,
	}

	buttons := g.vcsStarsKeyboard.GetButtons()
	if buttons == nil {
		return keyboard
	}

	keyMap := map[int]keyMapping{
		// Hold key (CNTR/F1 held)
		// 0: beaconator, handled as hold key
		1: {imgui.KeyF2, true, false}, // CNTR
		2: {imgui.KeyF3, true, false}, // MAPS
		3: {imgui.KeyF4, true, false}, // WX
		// 4: CORE - not mapped yet
		// 5: SIGNON - not mapped yet
		5: {imgui.KeyF5, true, false}, // BRITE
		6: {imgui.KeyF6, true, false}, // LDR
		7: {imgui.KeyF7, true, false}, // CHAR SIZE
		// 9: SHIFT - not mapped yet
		10: {imgui.KeyF8, true, false},  // DCB
		11: {imgui.KeyF10, true, false}, // RNG RING
		12: {imgui.KeyF11, true, false}, // RANGE
		14: {imgui.KeyF12, true, false}, // PREF SET
		18: {imgui.KeyF13, true, false}, // SITE
		// 19: MODE - not mapped yet
		24: {imgui.KeyF1, false, false},          // INIT CNTL
		25: {imgui.KeyF2, false, false},          // TRK RPOS
		26: {imgui.KeyF3, false, false},          // TRK SUSP
		27: {imgui.KeyF4, false, false},          // TERM CNTL
		28: {imgui.KeyF5, false, false},          // HND OFF
		29: {imgui.KeyF6, false, false},          // FLT DATA
		30: {imgui.KeyF7, false, false},          // MULTI FUNC
		31: {imgui.KeyF8, false, false},          // F8
		32: {imgui.KeyGraveAccent, false, false}, // Delta
		34: {imgui.KeyF9, false, false},          // F9
		35: {imgui.KeyF10, false, false},         // F10
		36: {imgui.KeyF11, false, false},         // CA
		37: {imgui.KeyF12, false, false},         // F12
		38: {imgui.KeyF13, false, true},          // F13
		39: {imgui.KeyF14, false, false},         // F14
		40: {imgui.KeyF15, false, false},         // TGT GEN
		41: {imgui.KeyF16, false, false},         // F16
	}

	for idx, action := range buttons {
		if action == glfw.Press {
			if idx == 0 {
				g.heldFKeys[imgui.KeyF1] = nil
				continue
			}
			if mapping, ok := keyMap[idx]; ok {
				keyboard.Pressed[mapping.key] = nil
				if mapping.ctrl {
					keyboard.Pressed[imgui.KeyLeftCtrl] = nil
				}
				if mapping.shift {
					keyboard.Pressed[imgui.KeyLeftShift] = nil
				}
			}
		} else if action == glfw.Release && idx == 0 {
			delete(g.heldFKeys, imgui.KeyF1)
		}
	}

	return keyboard
}
