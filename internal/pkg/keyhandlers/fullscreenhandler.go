package keyhandlers

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

const FullscreenId = "Fullscreen"

type FullscreenHandler struct {
	List         *views.ListWidget
	IsFullscreen *bool
	Content      *views.ItemWidget
}

func (h FullscreenHandler) Id() string {
	return listBackId
}

func toggle(b bool) bool {
	return !b
}

func (h FullscreenHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		tmp := toggle(*h.IsFullscreen)
		h.IsFullscreen = &tmp // memory leak?
		if *h.IsFullscreen {
			g.Cursor = true
			maxX, maxY := g.Size()
			v, _ := g.SetView("fullscreenContent", 0, 0, maxX, maxY)
			v.Editable = true
			v.Frame = false
			v.Wrap = true
			v.Title = "JSON Response - Fullscreen (CTRL+F to exit)"
			fmt.Fprintf(v, h.Content.GetContent())
			g.SetCurrentView("fullscreenContent")
		} else {
			g.Cursor = false
			g.DeleteView("fullscreenContent")
			g.SetCurrentView("listWidget")
		}
		return nil
	}
}

func (h FullscreenHandler) Widget() string {
	return ""
}

func (h FullscreenHandler) DefaultKey() gocui.Key {
	return gocui.KeyCtrlF
}
