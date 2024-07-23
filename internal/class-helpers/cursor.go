package classhelpers

import (
	g "github.com/zaptross/gorgeous"
)

type CursorTypes struct {
	Auto         string
	Default      string
	None         string
	ContextMenu  string
	Help         string
	Pointer      string
	Progress     string
	Wait         string
	Cell         string
	Crosshair    string
	Text         string
	VerticalText string
	Alias        string
	Copy         string
	Move         string
	NoDrop       string
	NotAllowed   string
	Grab         string
	Grabbing     string
	EResize      string
	NResize      string
	NeResize     string
	NwResize     string
	SResize      string
	SeResize     string
	SwResize     string
	WResize      string
	EwResize     string
	NsResize     string
	NeswResize   string
	NwseResize   string
	ColResize    string
	RowResize    string
	AllScroll    string
	ZoomIn       string
	ZoomOut      string
}

var Cursors = CursorTypes{
	Auto:         "auto",
	Default:      "default",
	None:         "none",
	ContextMenu:  "context-menu",
	Help:         "help",
	Pointer:      "pointer",
	Progress:     "progress",
	Wait:         "wait",
	Cell:         "cell",
	Crosshair:    "crosshair",
	Text:         "text",
	VerticalText: "vertical-text",
	Alias:        "alias",
	Copy:         "copy",
	Move:         "move",
	NoDrop:       "no-drop",
	NotAllowed:   "not-allowed",
	Grab:         "grab",
	Grabbing:     "grabbing",
	EResize:      "e-resize",
	NResize:      "n-resize",
	NeResize:     "ne-resize",
	NwResize:     "nw-resize",
	SResize:      "s-resize",
	SeResize:     "se-resize",
	SwResize:     "sw-resize",
	WResize:      "w-resize",
	EwResize:     "ew-resize",
	NsResize:     "ns-resize",
	NeswResize:   "nesw-resize",
	NwseResize:   "nwse-resize",
	ColResize:    "col-resize",
	RowResize:    "row-resize",
	AllScroll:    "all-scroll",
	ZoomIn:       "zoom-in",
	ZoomOut:      "zoom-out",
}

func Cursor(value string) string {
	className := "cursor-" + ClassNameSanitiser(value)

	g.Class(&g.CSSClass{
		Selector: "." + className,
		Props: g.CSSProps{
			"cursor": value,
		},
	})

	return className
}
