package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/dweymouth/supersonic/backend/mediaprovider"
)

type LyricsViewer struct {
	widget.BaseWidget

	noLyricsLabel  widget.Label
	unsyncedViewer *widget.RichText

	container   *container.Scroll
	currentView lyricView
}

type lyricView int

const (
	lyricViewEmpty lyricView = iota
	lyricViewUnsynced
	lyricViewSynced
)

func NewLyricsViewer() *LyricsViewer {
	l := &LyricsViewer{noLyricsLabel: widget.Label{
		Text: "Lyrics not available",
	}}
	l.ExtendBaseWidget(l)
	l.container = container.NewVScroll(&l.noLyricsLabel)
	return l
}

func (l *LyricsViewer) SetLyrics(lyrics *mediaprovider.Lyrics) {
	if lyrics == nil || len(lyrics.Lines) == 0 {
		if l.currentView != lyricViewEmpty {
			l.container.Content = &l.noLyricsLabel
			l.currentView = lyricViewEmpty
			l.Refresh()
		}
		return
	}

	if l.unsyncedViewer == nil {
		l.unsyncedViewer = widget.NewRichText()
		l.unsyncedViewer.Wrapping = fyne.TextWrapWord
	}
	l.unsyncedViewer.Segments = nil
	for _, line := range lyrics.Lines {
		ts := &widget.TextSegment{Text: line.Text}
		ts.Style.Alignment = fyne.TextAlignCenter
		ts.Style.SizeName = widget.RichTextStyleSubHeading.SizeName
		ts.Style.Inline = false
		l.unsyncedViewer.Segments = append(l.unsyncedViewer.Segments, ts)
	}
	l.unsyncedViewer.Refresh()
	if l.currentView != lyricViewUnsynced {
		l.container.Content = l.unsyncedViewer
		l.currentView = lyricViewUnsynced
		l.Refresh()
	}
}

func (l *LyricsViewer) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(l.container)
}
