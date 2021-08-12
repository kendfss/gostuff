package main
import (
    "github.com/charmbracelet/harmonica"
    "time"
)


// A thing we want to animate.
// sprite := struct{
type Sprite struct{
    x, xVelocity float64
    y, yVelocity float64
}

// Where we want to animate it.
const targetX = 50.0
const targetY = 100.0

// Initialize a spring with framerate, angular frequency, and damping values.

func main() {
    sprite := Sprite{}
    spring := harmonica.NewSpring(harmonica.FPS(60), 6.0, 0.5)
    // Animate!
    for {
        sprite.x, sprite.xVelocity = spring.Update(sprite.x, sprite.xVelocity, targetX)
        sprite.y, sprite.yVelocity = spring.Update(sprite.y, sprite.yVelocity, targetY)
        time.Sleep(time.Second/60)
    }

}
