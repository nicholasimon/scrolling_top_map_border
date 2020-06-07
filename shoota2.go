package main

import (
	"math/rand"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	backimg rl.Texture2D
	// level
	levelmap16              = make([]string, levela16)
	levelmap16two           = make([]string, levela16)
	levela16, levellength16 int
	// draw map
	drawblock, drawblocknext, screenw16, screenh16, screena16, halfh16, screenw4, screenh4, screena4, halfh4, widthcount16 int
	mapswitch                                                                                                              bool
	// core
	monh32, monw32                       int32
	monitorh, monitorw, monitornum       int
	grid16on, grid4on, debugon, lrg, sml bool
	framecount                           int
	mousepos                             rl.Vector2
	camera                               rl.Camera2D
)

func timers() { // MARK: timers
	if framecount%2 == 0 {
		drawblocknext++
		widthcount16--
	}
	if widthcount16 == 0 {
		drawblocknext = 0
		widthcount16 = screenw16 - 1
		if mapswitch {
			mapswitch = false
			createlevel()
		} else {
			mapswitch = true
			createleveltwo()
		}

	}
}
func main() { // MARK: main()
	rand.Seed(time.Now().UnixNano()) // random numbers
	rl.SetTraceLog(rl.LogError)      // hides INFO window
	startsettings()
	raylib()
}
func setscreen() { // MARK: setscreen()
	monitornum = rl.GetMonitorCount()
	monitorh = rl.GetScreenHeight()
	monitorw = rl.GetScreenWidth()
	monh32 = int32(monitorh)
	monw32 = int32(monitorw)
	rl.SetWindowSize(monitorw, monitorh)
	setsizes()
	startlevel()
}
func setsizes() { // MARK: setsizes()
	if monitorw >= 1600 {
		lrg = true
		sml = false
	} else if monitorw < 1600 && monitorw >= 1280 {
		lrg = false
		sml = true
	}
	screenh16 = monitorh / 16
	screenw16 = monitorw / 16
	screena16 = screenh16 * screenw16
	halfh16 = screenh16 / 2
	widthcount16 = screenw16 - 1

	screenh4 = monitorh / 4
	screenw4 = monitorw / 4
	screena4 = screenh4 * screenw4
	halfh4 = screenh4 / 2
}
func startsettings() { // MARK: start
	camera.Zoom = 1.0
	camera.Target.X = 0.0
	camera.Target.Y = 0.0
	//debugon = true
	//grid16on = true
	//selectedmenuon = true
}
func startlevel() { // MARK: startlevel()
	levellength16 = screenw16 * 2
	levela16 = levellength16 * screenh16
	levelmap16 = make([]string, levela16)
	levelmap16two = make([]string, levela16)
	createlevel()
	createleveltwo()
}
func createleveltwo() { // MARK: createlevel()
	for a := 0; a < levela16; a++ {
		levelmap16two[a] = "."
	}
	for a := 0; a < levellength16; a++ {
		levelmap16two[a] = "%"
	}

	length := rInt(16, halfh16-10)
	for a := 0; a < levellength16; a++ {
		block := a
		for b := 0; b < length; b++ {
			levelmap16two[block] = "%"
			block += levellength16
		}
		if length < 16 {
			length += rInt(4, 9)
		} else if length < halfh16-10 {
			length += rInt(-8, 9)
		} else if length >= halfh16-10 {
			length -= rInt(4, 9)
		}
	}
}
func createlevel() { // MARK: createlevel()
	for a := 0; a < levela16; a++ {
		levelmap16[a] = "."
	}
	for a := 0; a < levellength16; a++ {
		levelmap16[a] = "#"
	}
	length := rInt(16, halfh16-10)
	for a := 0; a < levellength16; a++ {
		block := a
		for b := 0; b < length; b++ {
			levelmap16[block] = "#"
			block += levellength16
		}
		if length < 16 {
			length += rInt(4, 9)
		} else if length < halfh16-10 {
			length += rInt(-8, 9)
		} else if length >= halfh16-10 {
			length -= rInt(4, 9)
		}
	}
}
func updateall() { // MARK: updateall()
	if grid16on {
		grid16()
	}
	if grid4on {
		grid4()
	}
	if debugon {
		debug()
	}
	timers()
	updatemap()
}
func updatemap() { // MARK: updatemap()

}
func backgroundimg() { // MARK: backimg()
	genimg := rl.GenImageCellular(monitorw, monitorh, 32)
	backimg = rl.LoadTextureFromImage(genimg)
	rl.UnloadImage(genimg)
}
func raylib() { // MARK: raylib()
	rl.InitWindow(monw32, monh32, "spaceshoota turbo VII DX")

	setscreen()
	rl.CloseWindow()
	rl.InitWindow(monw32, monh32, "spaceshoota turbo VII DX")
	backgroundimg()
	setscreen()

	rl.SetExitKey(rl.KeyEnd) // key to end the game and close window
	//	imgs = rl.LoadTexture("imgs.png") // load images
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() { // MARK: WindowShouldClose

		mousepos = rl.GetMousePosition()
		framecount++
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		// rl.DrawTexture(backimg, 0, 0, rl.Red) // MARK: draw backimg
		rl.BeginMode2D(camera)

		drawblock = drawblocknext
		drawx, drawy := int32(0), int32(0)
		linecount := 0
		for a := 0; a < screena16; a++ {

			checklevel16 := levelmap16[drawblock]
			if mapswitch {
				checklevel16 = levelmap16[drawblock]
			} else {
				checklevel16 = levelmap16two[drawblock]
			}

			switch checklevel16 {
			case "#":
				rl.DrawRectangle(drawx, drawy, 15, 15, rl.Green)
			case "%":
				rl.DrawRectangle(drawx, drawy, 15, 15, rl.Red)
			}

			drawblock++
			drawx += 16
			linecount++
			if linecount == screenw16 {
				drawx = 0
				drawy += 16
				drawblock += levellength16 - screenw16
				linecount = 0
			}

		}

		// MARK: draw map layer 1

		// MARK: draw map layer 2
		rl.EndMode2D() // MARK: draw no camera

		rl.EndDrawing()
		input()
		updateall()
	}
	rl.CloseWindow()
}
func debug() { // MARK: debug
	rl.DrawRectangle(monw32-300, 0, 500, monw32, rl.Fade(rl.Black, 0.7))
	rl.DrawFPS(monw32-290, monh32-100)

	levellength16TEXT := strconv.Itoa(levellength16)
	levela16TEXT := strconv.Itoa(levela16)
	screenh16TEXT := strconv.Itoa(screenh16)
	screenw16TEXT := strconv.Itoa(screenw16)
	screena16TEXT := strconv.Itoa(screena16)

	rl.DrawText(screenw16TEXT, monw32-290, 20, 10, rl.White)
	rl.DrawText("screenw16", monw32-200, 20, 10, rl.White)
	rl.DrawText(levellength16TEXT, monw32-290, 30, 10, rl.White)
	rl.DrawText("levellength4", monw32-200, 30, 10, rl.White)
	rl.DrawText(levela16TEXT, monw32-290, 40, 10, rl.White)
	rl.DrawText("levela4", monw32-200, 40, 10, rl.White)
	rl.DrawText(screenh16TEXT, monw32-290, 50, 10, rl.White)
	rl.DrawText("screenh16", monw32-200, 50, 10, rl.White)
	rl.DrawText(screenw16TEXT, monw32-290, 60, 10, rl.White)
	rl.DrawText("screenw16", monw32-200, 60, 10, rl.White)
	rl.DrawText(screena16TEXT, monw32-290, 70, 10, rl.White)
	rl.DrawText("screena16", monw32-200, 70, 10, rl.White)

}
func input() { // MARK: keys input
	if rl.IsKeyPressed(rl.KeyF1) {
		if grid16on {
			grid16on = false
		} else {
			grid16on = true
		}
	}
	if rl.IsKeyPressed(rl.KeyF2) {
		if grid4on {
			grid4on = false
		} else {
			grid4on = true
		}
	}
	if rl.IsKeyPressed(rl.KeyKpDecimal) {
		if debugon {
			debugon = false
		} else {
			debugon = true
		}
	}

}
func grid16() { // MARK: grid16()
	for a := 0; a < monitorw; a += 16 {
		a32 := int32(a)
		rl.DrawLine(a32, 0, a32, monh32, rl.Fade(rl.Green, 0.1))
	}
	for a := 0; a < monitorh; a += 16 {
		a32 := int32(a)
		rl.DrawLine(0, a32, monw32, a32, rl.Fade(rl.Green, 0.1))
	}
}
func grid4() { // MARK: grid4()
	for a := 0; a < monitorw; a += 4 {
		a32 := int32(a)
		rl.DrawLine(a32, 0, a32, monh32, rl.Fade(rl.DarkGreen, 0.1))
	}
	for a := 0; a < monitorh; a += 4 {
		a32 := int32(a)
		rl.DrawLine(0, a32, monw32, a32, rl.Fade(rl.DarkGreen, 0.1))
	}
}

// random numbers
func rInt(min, max int) int {
	return rand.Intn(max-min) + min
}
func rInt32(min, max int) int32 {
	a := int32(rand.Intn(max-min) + min)
	return a
}
func rFloat32(min, max int) float32 {
	a := float32(rand.Intn(max-min) + min)
	return a
}
func flipcoin() bool {
	var b bool
	a := rInt(0, 10001)
	if a < 5000 {
		b = true
	}
	return b
}
func rolldice() int {
	a := rInt(1, 7)
	return a
}
