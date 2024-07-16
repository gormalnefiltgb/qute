   package main

   import (
       "example.com/hajimehoshi/ebiten/v2"
       "example.com/hajimehoshi/ebiten/v2/inpututil"
       "log"
   )

   type Game struct{}

   func (g *Game) Update() error {
       // Check if the 'Space' key was just pressed
       if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
           log.Println("Space key pressed")
       }

       // Check if the 'Escape' key was just pressed
       if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
           log.Println("Escape key pressed. Exiting...")
           return ebiten.Termination
       }

       return nil
   }

   func (g *Game) Draw(screen *ebiten.Image) {
       // Drawing code goes here
   }

   func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
       return outsideWidth, outsideHeight
   }

   func main() {
       game := &Game{}
       ebiten.SetWindowSize(640, 480)
       ebiten.SetWindowTitle("Input Example")
       if err := ebiten.RunGame(game); err != nil {
           log.Fatal(err)
       }
   }
   