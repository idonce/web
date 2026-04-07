package main

const layoutCSS = `<style>
@font-face{font-family:'Inter';font-style:normal;font-weight:400;font-display:swap;src:url(/fonts/inter-400.woff2) format('truetype')}
@font-face{font-family:'Inter';font-style:normal;font-weight:500;font-display:swap;src:url(/fonts/inter-500.woff2) format('truetype')}
@font-face{font-family:'Space Grotesk';font-style:normal;font-weight:700;font-display:swap;src:url(/fonts/space-grotesk-700.woff2) format('truetype')}
*{margin:0;padding:0;box-sizing:border-box}
body{font-family:'Inter',system-ui,sans-serif;color:#0a0a0a;background:#fafaf9;-webkit-font-smoothing:antialiased;line-height:1.7}
a{color:inherit;text-decoration:none}

/* Shell: centered container with sidebar + content */
.shell{display:flex;max-width:1100px;margin:0 auto;min-height:100vh}
.sidebar{position:sticky;top:0;width:200px;flex-shrink:0;height:100vh;padding:40px 24px 40px 0;overflow-y:auto;border-right:1px solid #e7e5e4;display:flex;flex-direction:column}
.sidebar .logo{font-family:'Space Grotesk',system-ui,sans-serif;font-size:24px;font-weight:700;letter-spacing:-1px;margin-bottom:32px;color:#0a0a0a;text-align:center;display:block;padding:16px 12px;border-radius:12px;transition:background 0.15s}
.sidebar .logo:hover{background:#ede9fe}
.sidebar .logo span{color:#7c3aed}
.sidebar .logo small{display:block;font-family:'Inter',system-ui,sans-serif;font-size:10px;font-weight:500;text-transform:uppercase;letter-spacing:1.5px;color:#a8a29e;margin-top:4px}
.sidebar .nav-label{font-size:10px;font-weight:600;text-transform:uppercase;letter-spacing:1px;color:#d6d3d1;padding:12px 12px 4px;margin-top:4px}
.sidebar nav{display:flex;flex-direction:column;gap:1px}
.sidebar nav a{font-size:13px;font-weight:500;color:#a8a29e;padding:6px 12px;border-radius:8px;transition:all 0.15s}
.sidebar nav a:hover{color:#0a0a0a;background:#f5f5f4}
.sidebar nav a.active{color:#7c3aed;background:#ede9fe}
.sidebar nav a.section{font-size:12px;color:#c7c3be;padding:4px 12px 4px 20px}
.sidebar nav a.section:hover{color:#0a0a0a}
.sidebar nav a.section.active{color:#7c3aed;background:#ede9fe}
.sidebar .sep{height:1px;background:#e7e5e4;margin:12px 0}
.sidebar .bottom{margin-top:auto;display:flex;flex-direction:column;gap:1px}
.sidebar .bottom a{font-size:12px;color:#a8a29e;padding:5px 12px;border-radius:8px;transition:all 0.15s}
.sidebar .bottom a:hover{color:#0a0a0a;background:#f5f5f4}
.sidebar .bottom a.active{color:#7c3aed;background:#ede9fe}
.sidebar .copy{font-size:11px;color:#d6d3d1;padding:12px 12px 0;line-height:1.4}
.content{flex:1;min-width:0}

@media(max-width:860px){
  .shell{display:block}
  .sidebar{position:static;width:auto;height:auto;border-right:none;border-bottom:1px solid #e7e5e4;padding:20px 20px 16px;flex-direction:row;flex-wrap:wrap;align-items:center;gap:8px}
  .sidebar .logo{margin-bottom:0;margin-right:auto}
  .sidebar .nav-label{display:none}
  .sidebar nav{flex-direction:row;flex-wrap:wrap;gap:4px;width:100%;margin-top:12px}
  .sidebar nav a{font-size:12px;padding:5px 10px}
  .sidebar nav a.section{padding:5px 10px;font-size:12px}
  .sidebar .sep{display:none}
  .sidebar .bottom{display:none}
  .sidebar .copy{display:none}
}
</style>`

type sidebarLink struct {
	Href    string
	Label   string
	ID      string
	Section bool // indented sub-link for page sections
}

// Page-specific section anchors
var landingSections = []sidebarLink{
	{"#how", "How it Works", "how", true},
	{"#demo", "Try it Now", "demo", true},
	{"#usecases", "Who it's For", "usecases", true},
	{"#security", "Security", "security", true},
	{"#eidas", "eIDAS 2.0", "eidas", true},
	{"#dev", "For Developers", "dev", true},
	{"#compare", "Comparison", "compare", true},
	{"#faq", "FAQ", "faq", true},
}

var developerSections = []sidebarLink{
	{"#quick-start", "Quick Start", "quick-start", true},
	{"#patterns", "Patterns", "patterns", true},
	{"#use-cases", "Use Cases", "use-cases", true},
	{"#credentials", "Credentials", "credentials", true},
	{"#api", "API Reference", "api", true},
	{"#trust", "Trust Model", "trust", true},
	{"#self-hosting", "Self-Hosting", "self-hosting", true},
}

func sidebarHTML(activePage string, sections []sidebarLink) string {
	type link struct {
		href, label, id string
	}
	pages := []link{
		{"/", "Home", "home"},
		{"/developers", "Developers", "developers"},
		{"https://github.com/idonce", "GitHub", "github"},
	}
	bottom := []link{
		{"/privacy", "Privacy", "privacy"},
		{"/terms", "Terms", "terms"},
		{"/legal-notice", "Legal Notice", "legal-notice"},
	}

	s := `<aside class="sidebar">
  <a href="/" class="logo">id<span>once</span><small>Human Verification</small></a>
  <nav>`
	for _, l := range pages {
		cls := ""
		if l.id == activePage {
			cls = ` class="active"`
		}
		s += "\n    <a href=\"" + l.href + "\"" + cls + ">" + l.label + "</a>"
		// Section links appear right after the active page link
		if l.id == activePage && len(sections) > 0 {
			for _, sec := range sections {
				s += "\n    <a href=\"" + sec.Href + "\" class=\"section\" data-section=\"" + sec.ID + "\">" + sec.Label + "</a>"
			}
		}
	}
	s += `
  </nav>
  <div class="sep"></div>
  <div class="bottom">`
	for _, l := range bottom {
		cls := ""
		if l.id == activePage {
			cls = ` class="active"`
		}
		s += "\n    <a href=\"" + l.href + "\"" + cls + ">" + l.label + "</a>"
	}
	s += `
  </div>
  <div class="copy">&copy; 2026 idonce<br>Human verification infrastructure. No passwords. No CAPTCHAs. No personal data.</div>
</aside>`
	return s
}

// JS for scroll-spy on section links
const scrollSpyJS = `<script>
(function(){
  const links = document.querySelectorAll('.sidebar nav a.section');
  if (!links.length) return;
  const ids = Array.from(links).map(a => a.getAttribute('href').replace('#',''));
  const targets = ids.map(id => document.getElementById(id)).filter(Boolean);
  if (!targets.length) return;
  const observer = new IntersectionObserver(entries => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        links.forEach(a => a.classList.remove('active'));
        const link = document.querySelector('.sidebar nav a[href="#' + entry.target.id + '"]');
        if (link) link.classList.add('active');
      }
    });
  }, {rootMargin: '-10% 0px -80% 0px'});
  targets.forEach(t => observer.observe(t));
  // Smooth scroll
  links.forEach(a => a.addEventListener('click', e => {
    e.preventDefault();
    const t = document.querySelector(a.getAttribute('href'));
    if (t) t.scrollIntoView({behavior:'smooth'});
  }));
})();
</script>`
