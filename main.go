package main

import (
	"fmt"
	"strings"
)

// Item is a component of a window composite tree.
type Item interface {
	// Price returns the total price of the item and all its components.
	Price() float64
	// Display returns an indented tree representation of the item.
	Display() string
}

const (
	framePrice = 50.00
	glassPrice = 30.00
)

// Frame is a leaf item that cannot contain other items.
type Frame struct{}

func NewFrame() *Frame {
	return &Frame{}
}

func (Frame) Price() float64 {
	return framePrice
}

func (f Frame) Display() string {
	return fmt.Sprintf("Frame: %.2f", f.Price())
}

// Glass is a leaf item that cannot contain other items.
type Glass struct{}

func NewGlass() *Glass {
	return &Glass{}
}

func (Glass) Price() float64 {
	return glassPrice
}

func (g Glass) Display() string {
	return fmt.Sprintf("Glass: %.2f", g.Price())
}

// Window is a composite containing either one Glass or four Windows and one Frame.
type Window struct {
	glass    *Glass
	children [4]*Window
	frame    *Frame
}

// NewGlassWindow creates the simple form of a Window.
func NewGlassWindow(glass *Glass) (*Window, error) {
	if glass == nil {
		return nil, fmt.Errorf("a glass window requires a Glass")
	}

	return &Window{glass: glass}, nil
}

// NewCompositeWindow creates the framed form of a Window with exactly four children.
func NewCompositeWindow(children [4]*Window, frame *Frame) (*Window, error) {
	if frame == nil {
		return nil, fmt.Errorf("a composite window requires a Frame")
	}
	for i, child := range children {
		if child == nil {
			return nil, fmt.Errorf("composite window child %d must not be nil", i+1)
		}
	}

	return &Window{children: children, frame: frame}, nil
}

func (w Window) Price() float64 {
	if w.glass != nil {
		return w.glass.Price()
	}

	total := w.frame.Price()
	for _, child := range w.children {
		total += child.Price()
	}
	return total
}

func (w Window) Display() string {
	var builder strings.Builder
	w.writeDisplay(&builder, 0)
	return builder.String()
}

func (w Window) writeDisplay(builder *strings.Builder, level int) {
	indent := strings.Repeat("  ", level)
	builder.WriteString(indent)
	builder.WriteString("Window\n")

	if w.glass != nil {
		builder.WriteString(strings.Repeat("  ", level+1))
		builder.WriteString(w.glass.Display())
		builder.WriteByte('\n')
		return
	}

	for _, child := range w.children {
		child.writeDisplay(builder, level+1)
	}
	builder.WriteString(strings.Repeat("  ", level+1))
	builder.WriteString(w.frame.Display())
	builder.WriteByte('\n')
}

func main() {
	simpleWindow, err := NewGlassWindow(NewGlass())
	if err != nil {
		panic(err)
	}

	firstLevel := [4]*Window{}
	for i := 0; i < 3; i++ {
		firstLevel[i], err = NewGlassWindow(NewGlass())
		if err != nil {
			panic(err)
		}
	}

	secondLevel := [4]*Window{}
	for i := range secondLevel {
		secondLevel[i], err = NewGlassWindow(NewGlass())
		if err != nil {
			panic(err)
		}
	}
	firstLevel[3], err = NewCompositeWindow(secondLevel, NewFrame())
	if err != nil {
		panic(err)
	}

	compositeWindow, err := NewCompositeWindow(firstLevel, NewFrame())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Simple window (price: %.2f):\n%s\n", simpleWindow.Price(), simpleWindow.Display())
	fmt.Printf("Composite window (price: %.2f):\n%s", compositeWindow.Price(), compositeWindow.Display())
}
