package main

import (
	"image/color"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func setCamTarget(r, A, B float64, cam *rl.Camera3D) {
	x := float32(math.Sin(A) * math.Cos(B))
	z := -float32(math.Sin(A) * math.Sin(B))
	y := float32(math.Cos(A))
	cam.Target = rl.Vector3Add(rl.Vector3Scale(rl.NewVector3(x, y, z), float32(r)/rl.GetFrameTime()), cam.Position)
	// cam.Target = rl.Vector3Scale(rl.NewVector3(x, y, z), float32(r))
}

func moveF(B float64, F, G float64, cam *rl.Camera3D) {
	f := float32(F * math.Cos(B*rl.Deg2rad))
	l := float32(F * math.Sin(B*rl.Deg2rad))
	g := float32(G * math.Cos((B+90)*rl.Deg2rad))
	k := float32(G * math.Sin((B-90)*rl.Deg2rad))

	cam.Position.X += f - g
	cam.Position.Z -= l + k

	// fmt.Print(cam.Target)
	// fmt.Print(" ")
	// fmPrit(math.Sin((B - 90) * rl.Deg2rad))
	// fmt.Print(" ")
	// fmt.Print(math.Cos((B + 90) * rl.Deg2rad))
	// fmt.Println()
}

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(700, 700, "3d cam")
	defer rl.CloseWindow()
	rl.SetTargetFPS(144)

	targetP := rl.NewVector3(0, 0, 0)

	angleA, angleB := 135.0, 0.0
	// vertical, horizontal
	// b 0 = point to +x (+x +y)right to left 0(+x -z -x +z +x)360
	// a 0 = point to +y (+y +x)top to down   0{+y +x -y -x +y}360

	cam := rl.NewCamera3D(
		rl.NewVector3(-5, 5, 0),
		targetP,
		rl.NewVector3(0, 1, 0),
		50,
		rl.CameraPerspective,
	)
	// cam2 := cam

	// cam.Position = rl.NewVector3(0, 0, 0)
	isCurDis := false
	isSetCurDis := true
	rl.DisableCursor()

	obj1p := rl.NewVector3(2, 0, 0)

	for !rl.WindowShouldClose() {
		// fmt.Println(rl.IsCursorOnScreen())
		if rl.IsKeyPressed(rl.KeyL) {
			isCurDis = !isCurDis
			isSetCurDis = false
		}
		if isCurDis && !isSetCurDis {
			rl.DisableCursor()
			isSetCurDis = true
		} else if !isCurDis && !isSetCurDis {
			rl.EnableCursor()
			isSetCurDis = true
		}

		di := 0.0
		dj := 0.0
		speed := float32(5)

		if rl.IsKeyDown(rl.KeyQ) {
			cam.Position.Y += (speed * rl.GetFrameTime())
		}

		if rl.IsKeyDown(rl.KeyE) {
			cam.Position.Y -= (speed * rl.GetFrameTime())
		}
		if rl.IsKeyDown(rl.KeyW) {
			di = float64(speed * rl.GetFrameTime())
		}

		if rl.IsKeyDown(rl.KeyA) {
			dj = -float64(speed * rl.GetFrameTime())
		}
		if rl.IsKeyDown(rl.KeyS) {
			di = -float64(speed * rl.GetFrameTime())
		}

		if rl.IsKeyDown(rl.KeyD) {
			dj = float64(speed * rl.GetFrameTime())
		}

		md := rl.GetMouseDelta()
		angleA += float64(md.Y) * float64(rl.GetFrameTime()) * 10
		angleB += -float64(md.X) * float64(rl.GetFrameTime()) * 10
		// angleA += 2 * float64(rl.GetFrameTime()) * 10
		// angleB = 45

		if angleA > 360 {
			angleA = 0
		}
		if angleA < 0 {
			angleA = 360
		}
		if angleB > 360 {
			angleB = 0
		}
		if angleB < 0 {
			angleB = 360
		}

		setCamTarget(1, angleA*rl.Deg2rad, angleB*rl.Deg2rad, &cam)
		// fmt.Print(angleA)
		// fmt.Print(" ")
		// fmt.Print(angleB)
		// fmt.Print(" ")
		// fmt.Println(cam.Target)

		moveF(angleB, float64(di), float64(dj), &cam)

		rl.BeginDrawing()
		rl.ClearBackground(rl.SkyBlue)
		rl.BeginMode3D(cam)
		rl.DrawGrid(100, 1)
		rl.DrawSphere(rl.Vector3Scale(rl.Vector3Add(cam.Target, rl.Vector3Scale(cam.Position, -1)), rl.GetFrameTime()), .1, rl.Pink)
		rl.DrawSphere(rl.NewVector3(1, 0, 0), .05, rl.Red)
		rl.DrawSphere(rl.NewVector3(0, 1, 0), .05, rl.Green)
		rl.DrawSphere(rl.NewVector3(0, 0, 1), .05, rl.Blue)
		rl.DrawSphere(rl.NewVector3(0, 0, 0), 1, color.RGBA{90, 90, 90, 50})
		rl.DrawCube(obj1p, 1, 1, 1, rl.Blue)
		rl.EndMode3D()
		rl.EndDrawing()
	}
}
