package main

import (
	"strings"
	"testing"
)

func TestRenderPages(t *testing.T) {
	for _, p := range pages {
		html := renderPage(p)
		if !strings.Contains(html, "<html") {
			t.Errorf("%s: missing <html>", p.File)
		}
		if !strings.Contains(html, p.Title) {
			t.Errorf("%s: missing title %q", p.File, p.Title)
		}
		if !strings.Contains(html, `class="shell"`) {
			t.Errorf("%s: missing shell layout", p.File)
		}
		if !strings.Contains(html, `class="sidebar"`) {
			t.Errorf("%s: missing sidebar", p.File)
		}
		if !strings.Contains(html, `class="content"`) {
			t.Errorf("%s: missing content", p.File)
		}
	}
}

func TestLandingPage(t *testing.T) {
	html := renderPage(pages[0])
	if !strings.Contains(html, "idonce") {
		t.Error("expected idonce in landing page")
	}
	if !strings.Contains(html, "eIDAS") {
		t.Error("expected eIDAS in landing page")
	}
}

func TestDevelopersPage(t *testing.T) {
	html := renderPage(pages[1])
	if !strings.Contains(html, "Quick Start") {
		t.Error("expected Quick Start section")
	}
	if !strings.Contains(html, "/vp/sessions") {
		t.Error("expected API endpoint reference")
	}
}

func TestPrivacyPage(t *testing.T) {
	html := renderPage(pages[2])
	if !strings.Contains(html, "Privacy") {
		t.Error("expected Privacy content")
	}
}

func TestTermsPage(t *testing.T) {
	html := renderPage(pages[3])
	if !strings.Contains(html, "Terms") {
		t.Error("expected Terms content")
	}
}

func TestLegalNoticePage(t *testing.T) {
	html := renderPage(pages[4])
	if !strings.Contains(html, "Gradient Zero") {
		t.Error("expected Gradient Zero in legal notice")
	}
	if !strings.Contains(html, "Legal") {
		t.Error("expected Legal in page")
	}
}

func TestSidebarActiveState(t *testing.T) {
	// Landing page should have Home active
	html := renderPage(pages[0])
	if !strings.Contains(html, `class="active">Home`) {
		t.Error("expected Home to be active on landing page")
	}

	// Developers page should have Developers active
	html = renderPage(pages[1])
	if !strings.Contains(html, `class="active">Developers`) {
		t.Error("expected Developers to be active on developers page")
	}
}

func TestSidebarSectionLinks(t *testing.T) {
	// Landing should have section links
	html := renderPage(pages[0])
	if !strings.Contains(html, `data-section="how"`) {
		t.Error("expected How it Works section link on landing")
	}

	// Developers should have section links
	html = renderPage(pages[1])
	if !strings.Contains(html, `data-section="quick-start"`) {
		t.Error("expected Quick Start section link on developers")
	}

	// Legal pages should NOT have section links
	html = renderPage(pages[2])
	if strings.Contains(html, "data-section") {
		t.Error("expected no section links on privacy page")
	}
}

func TestAllPagesCount(t *testing.T) {
	if len(pages) != 5 {
		t.Errorf("expected 5 pages, got %d", len(pages))
	}
}
