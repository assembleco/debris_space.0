package loader

import (
	"time"

	"github.com/g3n/engine/core"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/loader/obj"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/util/helper"
	"github.com/assembleco/debris_space/engine/app"
)

func init() {
	app.DemoMap["loader.obj"] = &LoaderObj{}
}

type LoaderObj struct {
	prevLoaded core.INode
}

// Start is called once at the start of the demo.
func (t *LoaderObj) Start(a *app.App) {

	address := a.DirData()+"/binary/ring.obj"
	t.load(a, address)

	// Adds white directional front light
	l1 := light.NewDirectional(&math32.Color{1, 1, 1}, 1.0)
	l1.SetPosition(0, 0, 10)
	a.Scene().Add(l1)

	// Adds white directional top light
	l2 := light.NewDirectional(&math32.Color{1, 1, 1}, 1.0)
	l2.SetPosition(0, 10, 0)
	a.Scene().Add(l2)

	// Adds white directional right light
	l3 := light.NewDirectional(&math32.Color{1, 1, 1}, 1.0)
	l3.SetPosition(10, 0, 0)
	a.Scene().Add(l3)

	// Create axes helper
	axes := helper.NewAxes(2)
	a.Scene().Add(axes)
}

func (t *LoaderObj) load(a *app.App, path string) error {

	// Remove previous model from the scene
	if t.prevLoaded != nil {
		a.Scene().Remove(t.prevLoaded)
		t.prevLoaded.Dispose()
		t.prevLoaded = nil
	}

	// Decodes obj file and associated mtl file
	dec, err := obj.Decode(path, "")
	if err != nil {
		return err
	}

	// Creates a new node with all the objects in the decoded file and adds it to the scene
	group, err := dec.NewGroup()
	if err != nil {
		return err
	}
	a.Scene().Add(group)
	t.prevLoaded = group
	return nil
}

// Update is called every frame.
func (t *LoaderObj) Update(a *app.App, deltaTime time.Duration) {}

// Cleanup is called once at the end of the demo.
func (t *LoaderObj) Cleanup(a *app.App) {}
