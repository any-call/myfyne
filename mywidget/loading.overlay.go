package mywidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/image/colornames"
	"image/color"
)

// LoadingOverlay is a custom Fyne component
type LoadingOverlay struct {
	widget.BaseWidget
	child          fyne.CanvasObject
	loading        *LoadingDots
	loadingVisible bool
	container      *fyne.Container
}

// NewLoadingOverlay creates a new LoadingOverlay with the given child
func NewLoadingOverlay(child fyne.CanvasObject, loadColor color.Color, radius float64) *LoadingOverlay {
	if loadColor == nil {
		loadColor = colornames.Blue
	}

	if radius <= 0 {
		radius = 30
	}

	lo := &LoadingOverlay{
		child:   child,
		loading: NewLoadingDots(loadColor, radius),
	}

	lo.ExtendBaseWidget(lo)
	lo.container = container.New(layout.NewStackLayout(), lo.child, lo.loading)
	return lo
}

func (lo *LoadingOverlay) Start() {
	lo.loadingVisible = true
	lo.loading.Start()
	lo.Refresh()
}

func (lo *LoadingOverlay) Stop() {
	lo.loadingVisible = false
	lo.loading.Stop()
	lo.Refresh()
}

func (lo *LoadingOverlay) GetChild() fyne.CanvasObject {
	return lo.child
}

// CreateRenderer creates and returns a new renderer for the LoadingOverlay
func (lo *LoadingOverlay) CreateRenderer() fyne.WidgetRenderer {
	return &loadingOverlayRenderer{
		overlay: lo,
	}
}

type loadingOverlayRenderer struct {
	overlay *LoadingOverlay
}

func (r *loadingOverlayRenderer) Layout(size fyne.Size) {
	r.overlay.container.Resize(size)
	if r.overlay.loadingVisible {
		r.overlay.loading.Show()
	} else {
		r.overlay.loading.Hide()
	}
}

func (r *loadingOverlayRenderer) MinSize() fyne.Size {
	return r.overlay.container.MinSize()
}

func (r *loadingOverlayRenderer) Refresh() {
	r.Layout(r.overlay.container.Size())
	fyne.Do(func() {
		r.overlay.container.Refresh()
	})
}

func (r *loadingOverlayRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.overlay.container}
}

func (r *loadingOverlayRenderer) Destroy() {
	// No resources to clean up in this example
	r.overlay.loading.Stop()
}
