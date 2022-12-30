package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

const (
	PADDLE_HEIGHT     = 64
	PADDLE_WIDTH      = 10
	PADDLE_MOVE_SPEED = 100
)

type Paddle struct {
	isMoving bool
	yPos float32
	velocity float32

	color rl.Color
	mesh rl.Rectangle
}

func NewPaddle(x float32, color rl.Color) Paddle {
	return Paddle{
		yPos: 0.0,
		color: color,
		velocity: 0.0,
		mesh: rl.NewRectangle(x, 0.0, PADDLE_WIDTH, PADDLE_HEIGHT),
	}
}

func (p *Paddle) Render(deltaTime float32) {
	p.yPos += p.velocity * deltaTime
	p.yPos = rl.Clamp(p.yPos, 0, 536)
	p.mesh.Y = p.yPos

	rl.DrawRectangleRec(p.mesh, p.color)
}

type application struct {
	camera rl.Camera2D

	leftPaddle Paddle
	rightPaddle Paddle
}

func (app *application) setup() {
	app.camera = rl.NewCamera2D(rl.Vector2Zero(), rl.Vector2Zero(), 0.0, 1.0)

	app.leftPaddle = NewPaddle(0.0, rl.Blue)
	app.rightPaddle = NewPaddle(790.0, rl.Red)
}

func (app *application) loop() {
	if rl.IsKeyDown(rl.KeyW) {
		app.leftPaddle.velocity = PADDLE_MOVE_SPEED
	} else if rl.IsKeyDown(rl.KeyS) {
		app.leftPaddle.velocity = -PADDLE_MOVE_SPEED
	} else {
		app.leftPaddle.velocity = 0
	}

	if rl.IsKeyDown(rl.KeyUp) {
		app.rightPaddle.velocity = PADDLE_MOVE_SPEED
	} else if rl.IsKeyDown(rl.KeyDown) {
		app.rightPaddle.velocity = -PADDLE_MOVE_SPEED
	} else {
		app.rightPaddle.velocity = 0
	}

	deltaTime := rl.GetFrameTime()
	app.leftPaddle.Render(deltaTime)
	app.rightPaddle.Render(deltaTime)
}

func main() {
	app := application{}
	app.setup()

	rl.InitWindow(800, 600, "RayPonGO")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginMode2D(app.camera)
		rl.ClearBackground(rl.RayWhite)

		app.loop()

		rl.GetFrameTime()

		rl.EndMode2D()
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
