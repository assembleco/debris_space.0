package main

import (
	_ "github.com/assembleco/space_debris/engine/demos/animation"
	_ "github.com/assembleco/space_debris/engine/demos/audio"
	_ "github.com/assembleco/space_debris/engine/demos/experimental/physics"
	_ "github.com/assembleco/space_debris/engine/demos/geometry"
	_ "github.com/assembleco/space_debris/engine/demos/gui"
	_ "github.com/assembleco/space_debris/engine/demos/helper"
	_ "github.com/assembleco/space_debris/engine/demos/light"
	_ "github.com/assembleco/space_debris/engine/demos/loader"
	_ "github.com/assembleco/space_debris/engine/demos/material"
	_ "github.com/assembleco/space_debris/engine/demos/other"
	_ "github.com/assembleco/space_debris/engine/demos/shader"
	_ "github.com/assembleco/space_debris/engine/demos/tests"
	_ "github.com/assembleco/space_debris/engine/demos/texture"

	"github.com/assembleco/space_debris/engine/app"
)

func main() {

	// Create and run application
	app.Create().Run()
}
