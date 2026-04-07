package main

const landingCSS = `<style>
:root{--black:#0a0a0a;--white:#fafaf9;--purple:#7c3aed;--purple-light:#ede9fe;--purple-muted:#c4b5fd;--grey:#a8a29e;--grey-light:#e7e5e4;--green:#22c55e;--green-light:#dcfce7;--font-display:'Space Grotesk',system-ui,sans-serif;--font-body:'Inter',system-ui,sans-serif}
.wrap{max-width:1080px;margin:0 auto;padding:0 28px}

/* Display type */
.d-xxl{font-family:var(--font-display);font-size:clamp(48px,8vw,96px);font-weight:700;letter-spacing:-3px;line-height:0.95;text-transform:uppercase}
.d-xl{font-family:var(--font-display);font-size:clamp(36px,6vw,64px);font-weight:700;letter-spacing:-2px;line-height:0.95;text-transform:uppercase}
.d-lg{font-family:var(--font-display);font-size:clamp(24px,4vw,40px);font-weight:700;letter-spacing:-1px;line-height:1;text-transform:uppercase}
.accent{color:var(--purple)}
.txt{color:var(--grey);font-size:17px;line-height:1.6;max-width:480px}
.txt-sm{color:var(--grey);font-size:14px;line-height:1.6}

/* Buttons */
.btn{display:inline-flex;align-items:center;gap:8px;font-family:var(--font-body);font-weight:500;font-size:15px;border:none;cursor:pointer;transition:all 0.2s;text-decoration:none;border-radius:50px}
.btn-p{background:var(--black);color:var(--white);padding:14px 28px}.btn-p:hover{background:var(--purple)}
.btn-s{background:none;color:var(--black);padding:14px 28px;border:1.5px solid var(--grey-light)}.btn-s:hover{border-color:var(--black)}
.btn-t{background:none;color:var(--black);padding:0;border-bottom:1.5px solid var(--black);border-radius:0;padding-bottom:2px;font-size:14px}.btn-t:hover{color:var(--purple);border-color:var(--purple)}
.badge{display:inline-flex;align-items:center;gap:6px;padding:6px 14px;border-radius:50px;font-size:12px;font-weight:500;background:var(--purple-light);color:var(--purple);border:1.5px solid var(--purple-muted)}

/* Sections */
.sec{padding:100px 0}
.sec+.sec{border-top:1px solid var(--grey-light)}
.sec-first{padding-top:48px}

/* Hero */
.hero{text-align:center}
.hero .d-xxl{margin-bottom:24px}
.hero .txt{margin:0 auto 40px;text-align:center}
.hero .btns{display:flex;gap:12px;justify-content:center;flex-wrap:wrap}

/* Steps */
.steps{display:grid;grid-template-columns:repeat(3,1fr);gap:48px;margin-top:64px;text-align:center}
.steps .num{font-family:var(--font-display);font-size:56px;font-weight:700;color:var(--grey-light);line-height:1;margin-bottom:8px}
.steps h3{font-family:var(--font-display);font-size:22px;font-weight:700;text-transform:uppercase;letter-spacing:-0.5px;margin-bottom:8px}
.steps p{color:var(--grey);font-size:14px;line-height:1.6}
@media(max-width:768px){.steps{grid-template-columns:1fr;gap:40px}}

/* Demo */
.demo-box{max-width:420px;margin:48px auto 0;padding:40px;border:1.5px solid var(--grey-light);border-radius:20px;text-align:center}
.demo-box h3{font-family:var(--font-display);font-size:20px;text-transform:uppercase;margin-bottom:4px}
.qr-wrap{background:#fff;border-radius:12px;padding:16px;display:inline-block;margin:20px 0 12px;border:1px solid var(--grey-light)}
.qr-wrap canvas{display:block}
.pulse{display:inline-block;width:8px;height:8px;background:var(--purple);border-radius:50%;margin-right:8px;animation:pulse 1.5s ease infinite}
@keyframes pulse{0%,100%{opacity:1}50%{opacity:.3}}
.v-icon{width:64px;height:64px;background:var(--green-light);border-radius:50%;display:inline-flex;align-items:center;justify-content:center;margin-bottom:16px}
.v-icon svg{width:32px;height:32px;color:var(--green)}
.claim{display:inline-block;background:var(--green-light);color:var(--green);font-size:12px;font-weight:600;padding:4px 12px;border-radius:50px;margin:3px}

/* Use cases */
.uc-grid{display:grid;grid-template-columns:repeat(2,1fr);gap:64px 80px;margin-top:64px}
.uc h3{font-family:var(--font-display);font-size:24px;font-weight:700;text-transform:uppercase;letter-spacing:-0.5px;margin-bottom:8px}
.uc p{color:var(--grey);font-size:15px;line-height:1.6}
@media(max-width:768px){.uc-grid{grid-template-columns:1fr;gap:40px}}

/* Stats */
.stats{display:grid;grid-template-columns:repeat(4,1fr);gap:32px;margin-top:64px;text-align:center}
.stat-n{font-family:var(--font-display);font-size:64px;font-weight:700;letter-spacing:-2px;line-height:1}
.stat-l{color:var(--grey);font-size:14px;margin-top:4px}
@media(max-width:768px){.stats{grid-template-columns:1fr}}

/* eIDAS */
.eidas-grid{display:grid;grid-template-columns:repeat(2,1fr);gap:48px;margin-top:64px}
.eidas-card h3{font-family:var(--font-display);font-size:20px;font-weight:700;text-transform:uppercase;letter-spacing:-0.5px;margin-bottom:4px;margin-top:8px}
.eidas-card p{color:var(--grey);font-size:14px;line-height:1.6}
@media(max-width:768px){.eidas-grid{grid-template-columns:1fr}}

/* Code */
.code{max-width:600px;margin:48px auto 0;background:var(--black);color:#d6d3d1;border-radius:12px;padding:24px;font-family:'SF Mono','Fira Code',monospace;font-size:13px;line-height:1.7;overflow-x:auto;text-align:left}
.code .c{color:#78716c}
.code .s{color:var(--purple-muted)}
.code .k{color:#86efac}

/* Security */
.sec-grid{display:grid;grid-template-columns:repeat(3,1fr);gap:48px;margin-top:64px}
.sec-item h4{font-family:var(--font-display);font-size:16px;font-weight:700;text-transform:uppercase;letter-spacing:-0.3px;margin-bottom:4px;margin-top:12px}
.sec-item p{color:var(--grey);font-size:13px;line-height:1.5}
.sec-item .ico{width:44px;height:44px;border-radius:12px;display:flex;align-items:center;justify-content:center;font-size:20px}
@media(max-width:768px){.sec-grid{grid-template-columns:1fr}}

/* Comparison */
.cmp{overflow-x:auto;margin-top:48px}
.cmp table{width:100%;border-collapse:separate;border-spacing:0;font-size:14px;min-width:640px}
.cmp th{text-align:left;padding:12px 16px;border-bottom:2px solid var(--grey-light);font-weight:600;font-size:13px;text-transform:uppercase;letter-spacing:0.5px;color:var(--grey)}
.cmp th:first-child{color:var(--black)}
.cmp th.hl{background:var(--purple-light);color:var(--purple);border-radius:12px 12px 0 0}
.cmp td{padding:12px 16px;border-bottom:1px solid var(--grey-light);font-size:14px}
.cmp td.hl{background:var(--purple-light)}
.cmp tr:last-child td{border-bottom:none}
.cmp tr:last-child td.hl{border-radius:0 0 12px 12px}
.ok{color:var(--green);font-weight:500}
.no{color:#ef4444;font-weight:500}
.meh{color:#f59e0b;font-weight:500}

/* FAQ */
.faq{max-width:640px;margin:48px auto 0}
.faq-item{display:flex;justify-content:space-between;align-items:center;padding:20px 0;border-bottom:1px solid var(--grey-light);cursor:pointer}
.faq-item:first-child{border-top:1px solid var(--grey-light)}
.faq-q{font-size:16px;font-weight:500;flex:1}
.faq-a{display:none;padding:0 0 20px;color:var(--grey);font-size:15px;line-height:1.6;border-bottom:1px solid var(--grey-light)}
.faq-arrow{font-size:18px;color:var(--grey);transition:transform 0.2s;margin-left:16px}
.faq-item.open .faq-arrow{transform:rotate(180deg)}
.faq-item.open+.faq-a{display:block}

/* CTA */
.cta-sec{text-align:center;padding:160px 0}

</style>`

const landingBody = `<!-- Hero -->
<section class="sec sec-first hero">
<div class="wrap">
  <h1 class="d-xxl">Prove you're<br><span class="accent">human.</span><br>Nothing else.</h1>
  <p class="txt">Free, open-source human verification. No passwords, no CAPTCHAs, no personal data. For everyone.</p>
  <div class="btns">
    <a href="#demo" class="btn btn-p">See it in Action</a>
    <a href="https://github.com/idonce" class="btn btn-s" target="_blank" rel="noopener">&#9733; View on GitHub</a>
  </div>
</div>
</section>

<!-- Stats -->
<section class="sec">
<div class="wrap">
  <div class="stats">
    <div><div class="stat-n">3<span style="font-size:36px">s</span></div><div class="stat-l">Verification time</div></div>
    <div><div class="stat-n accent">0</div><div class="stat-l">Personal data shared</div></div>
    <div><div class="stat-n">100<span style="font-size:36px">%</span></div><div class="stat-l">eIDAS 2.0 compatible</div></div>
    <div><div class="stat-n">100<span style="font-size:36px">%</span></div><div class="stat-l">Free & Open source</div></div>
  </div>
</div>
</section>

<!-- How it works -->
<section class="sec" id="how">
<div class="wrap">
  <div class="d-xl" style="text-align:center">How it<br>works</div>
  <div class="steps">
    <div>
      <div class="num">01</div>
      <h3>Scan</h3>
      <p>A platform shows a QR code. You scan it with the idonce app on your phone.</p>
    </div>
    <div>
      <div class="num">02</div>
      <h3>Confirm</h3>
      <p>Authenticate with FaceID or your fingerprint. Your device creates a cryptographic proof.</p>
    </div>
    <div>
      <div class="num">03</div>
      <h3>Verified</h3>
      <p>The platform knows a real person confirmed. No name, no email, no tracking. Done.</p>
    </div>
  </div>
</div>
</section>

<!-- Demo -->
<section class="sec" id="demo">
<div class="wrap" style="text-align:center">
  <div class="d-xl">Try it<br><span class="accent">now</span></div>
  <p class="txt" style="margin:16px auto 0">See human verification in action. Get the app, then try the live demo.</p>

  <!-- App Store badges -->
  <div style="display:flex;gap:12px;justify-content:center;margin:28px 0 36px;flex-wrap:wrap">
    <a href="#" style="display:inline-block;background:var(--black);color:var(--white);padding:10px 20px;border-radius:10px;text-decoration:none;display:flex;align-items:center;gap:10px">
      <svg width="22" height="22" viewBox="0 0 24 24" fill="white"><path d="M18.71 19.5c-.83 1.24-1.71 2.45-3.05 2.47-1.34.03-1.77-.79-3.29-.79-1.53 0-2 .77-3.27.82-1.31.05-2.3-1.32-3.14-2.53C4.25 17 2.94 12.45 4.7 9.39c.87-1.52 2.43-2.48 4.12-2.51 1.28-.02 2.5.87 3.29.87.78 0 2.26-1.07 3.8-.91.65.03 2.47.26 3.64 1.98-.09.06-2.17 1.28-2.15 3.81.03 3.02 2.65 4.03 2.68 4.04-.03.07-.42 1.44-1.38 2.83M13 3.5c.73-.83 1.94-1.46 2.94-1.5.13 1.17-.34 2.35-1.04 3.19-.69.85-1.83 1.51-2.95 1.42-.15-1.15.41-2.35 1.05-3.11z"/></svg>
      <div style="text-align:left;line-height:1.2"><div style="font-size:10px;opacity:0.7">Download on the</div><div style="font-size:15px;font-weight:600">App Store</div></div>
    </a>
    <a href="#" style="display:inline-block;background:var(--black);color:var(--white);padding:10px 20px;border-radius:10px;text-decoration:none;display:flex;align-items:center;gap:10px">
      <svg width="20" height="22" viewBox="0 0 20 22" fill="none"><path d="M1 1.56v18.88c0 .56.3.73.67.38L11 11.5 1.67 1.18C1.3.83 1 1 1 1.56z" fill="#4285F4"/><path d="M14.5 8L11 11.5l3.5 3.5 4-2.3c.67-.38.67-1.02 0-1.4L14.5 8z" fill="#FBBC04"/><path d="M1.67 1.18L11 11.5l3.5-3.5L5.17.27C4.5-.12 3.83-.05 3.17.35L1.67 1.18z" fill="#34A853"/><path d="M1.67 21.82L11 11.5 14.5 15l-9.33 5.73c-.67.4-1.33.47-2 .1L1.67 21.82z" fill="#EA4335"/></svg>
      <div style="text-align:left;line-height:1.2"><div style="font-size:10px;opacity:0.7">Get it on</div><div style="font-size:15px;font-weight:600">Google Play</div></div>
    </a>
  </div>

  <div class="demo-box">
    <h3>Verify you're human</h3>
    <p class="txt-sm" style="margin-bottom:20px">Scan a QR code, confirm with biometrics, done. No passwords, no personal data.</p>
    <a href="/demo" class="btn btn-p" style="width:100%;justify-content:center" target="_blank" rel="noopener">Try the Live Demo</a>
  </div>
</div>
</section>

<!-- Use Cases -->
<section class="sec" id="usecases">
<div class="wrap">
  <div class="d-xl" style="text-align:center">Who it's<br><span class="accent">for</span></div>
  <div class="uc-grid">
    <div class="uc">
      <h3>Forums &<br>communities</h3>
      <p>Stop bot registrations. Every account is backed by a real device with biometric confirmation.</p>
    </div>
    <div class="uc">
      <h3>E-commerce</h3>
      <p>Prevent fake reviews. Require human verification before a review is published.</p>
    </div>
    <div class="uc">
      <h3>Voting &<br>governance</h3>
      <p>One device, one vote. Sybil-resistant polling without collecting personal data.</p>
    </div>
    <div class="uc">
      <h3>Fintech</h3>
      <p>Real-time transaction confirmation. Proof that a human approved this action, right now.</p>
    </div>
  </div>
</div>
</section>

<!-- Security -->
<section class="sec" id="security">
<div class="wrap">
  <div class="d-xl" style="text-align:center">Security by<br><span class="accent">design</span></div>
  <div class="sec-grid">
    <div class="sec-item">
      <div class="ico" style="background:var(--purple-light)">&#128274;</div>
      <h4>Biometric<br>confirmation</h4>
      <p>FaceID or Fingerprint for every verification. The device owner actively confirms.</p>
    </div>
    <div class="sec-item">
      <div class="ico" style="background:var(--green-light)">&#128241;</div>
      <h4>Device-bound<br>keys</h4>
      <p>Private key in Secure Enclave. Cannot be exported or copied to another device.</p>
    </div>
    <div class="sec-item">
      <div class="ico" style="background:#fef3c7">&#9989;</div>
      <h4>Hardware<br>attestation</h4>
      <p>Apple App Attest and Google Play Integrity verify genuine hardware.</p>
    </div>
    <div class="sec-item">
      <div class="ico" style="background:#ede9fe">&#128065;</div>
      <h4>Selective<br>disclosure</h4>
      <p>You choose claim-by-claim what to share. Nothing more.</p>
    </div>
    <div class="sec-item">
      <div class="ico" style="background:#fce7f3">&#128260;</div>
      <h4>Replay<br>protection</h4>
      <p>Every proof is single-use. Signature tracking prevents reuse.</p>
    </div>
    <div class="sec-item">
      <div class="ico" style="background:#f5f5f4">&#128683;</div>
      <h4>Zero personal<br>data</h4>
      <p>No name. No email. No tracking. Only a pseudonymous device ID.</p>
    </div>
  </div>
</div>
</section>

<!-- eIDAS -->
<section class="sec" id="eidas">
<div class="wrap">
  <div class="d-xl" style="text-align:center">Built on<br><span class="accent">EU standards</span></div>
  <p class="txt" style="text-align:center;margin:16px auto 0">Same technology as the upcoming EU Digital Identity Wallet.</p>
  <div class="eidas-grid">
    <div class="eidas-card">
      <div class="badge">SD-JWT-VC</div>
      <h3>Verifiable Credentials</h3>
      <p>The mandatory credential format for eIDAS 2.0. Selective disclosure built in.</p>
    </div>
    <div class="eidas-card">
      <div class="badge">OpenID4VP</div>
      <h3>Credential Presentation</h3>
      <p>Standard protocol for sharing credentials with any eIDAS-compatible verifier.</p>
    </div>
    <div class="eidas-card">
      <div class="badge">OpenID4VCI</div>
      <h3>Credential Issuance</h3>
      <p>Standard protocol for receiving credentials. Any wallet can use the idonce issuer.</p>
    </div>
    <div class="eidas-card">
      <div class="badge">did:web</div>
      <h3>Decentralized Identity</h3>
      <p>No vendor lock-in. Resolvable via standard DID Document.</p>
    </div>
  </div>
</div>
</section>

<!-- Developers -->
<section class="sec" id="dev">
<div class="wrap" style="text-align:center">
  <div class="d-xl">For<br><span class="accent">developers</span></div>
  <p class="txt" style="margin:16px auto 0">One API call. That's the integration.</p>
  <div class="code">
<span class="c"># Use www.idonce.com or your own hosted issuer</span>
<br>
<br>
<span class="c"># 1. Create verification session</span>
<br>
curl -X POST https://www.idonce.com/vp/sessions \
<br>
  -H <span class="s">'Content-Type: application/json'</span> \
<br>
  -d <span class="s">'{"client_id":"your-platform"}'</span>
<br>
<br>
<span class="c"># 2. Show QR code from response.qr_data to user</span>
<br>
<br>
<span class="c"># 3. Poll for result</span>
<br>
curl https://www.idonce.com/vp/sessions/<span class="s">vp_a1b2c3...</span>
<br>
<br>
<span class="c"># Response when verified:</span>
<br>
{
  <span class="k">"status"</span>: <span class="s">"presented"</span>,
  <span class="k">"disclosed_claims"</span>: {
    <span class="k">"biometricConfirmed"</span>: <span class="s">true</span>,
    <span class="k">"deviceBound"</span>: <span class="s">true</span>
  }
}
  </div>
  <div style="display:flex;gap:12px;justify-content:center;margin-top:24px;flex-wrap:wrap">
    <a href="/developers" class="btn btn-p">Full Documentation</a>
    <a href="https://github.com/idonce" class="btn btn-s" target="_blank" rel="noopener">GitHub</a>
  </div>
</div>
</section>

<!-- Comparison -->
<section class="sec" id="compare">
<div class="wrap">
  <div class="d-xl" style="text-align:center">How id<span class="accent">once</span><br>compares</div>
  <div class="cmp">
  <table>
    <tr><th></th><th class="hl">idonce</th><th>Turnstile</th><th>reCAPTCHA</th><th>World ID</th></tr>
    <tr><td>Mechanism</td><td class="hl">Biometrics + HW</td><td>Fingerprint</td><td>Behavioral ML</td><td>Iris scan</td></tr>
    <tr><td>Guarantee</td><td class="hl ok">Cryptographic</td><td class="meh">Probabilistic</td><td class="meh">Probabilistic</td><td class="ok">Cryptographic</td></tr>
    <tr><td>Privacy</td><td class="hl ok">No tracking</td><td class="no">Cloudflare</td><td class="no">Google tracks</td><td class="meh">Iris database</td></tr>
    <tr><td>App needed</td><td class="hl">Yes</td><td class="ok">No</td><td class="ok">No</td><td>Yes</td></tr>
    <tr><td>eIDAS 2.0</td><td class="hl ok">Yes</td><td class="no">No</td><td class="no">No</td><td class="no">No</td></tr>
    <tr><td>Sybil resistance</td><td class="hl ok">Device-bound</td><td class="no">Weak</td><td class="no">Weak</td><td class="ok">Iris-bound</td></tr>
  </table>
  </div>
</div>
</section>

<!-- FAQ -->
<section class="sec" id="faq">
<div class="wrap">
  <div class="d-xl" style="text-align:center">Frequently<br>asked <span class="accent">questions</span></div>
  <div class="faq">
    <div class="faq-item" onclick="toggleFaq(this)"><span class="faq-q">Do I need to install an app?</span><span class="faq-arrow">&darr;</span></div>
    <div class="faq-a">Yes. idonce uses your device's Secure Enclave and biometrics to create cryptographic proofs. This cannot be done in a browser alone.</div>

    <div class="faq-item" onclick="toggleFaq(this)"><span class="faq-q">What data is shared with platforms?</span><span class="faq-arrow">&darr;</span></div>
    <div class="faq-a">Only pseudonymous claims: "biometric confirmed" and "device bound". No name, email, location, or any personal data.</div>

    <div class="faq-item" onclick="toggleFaq(this)"><span class="faq-q">Is idonce really free?</span><span class="faq-arrow">&darr;</span></div>
    <div class="faq-a">Yes, completely. idonce is open source under a permissive license. The app, the issuer server, and the verifier are all free to use, modify, and self-host. No premium tier, no hidden costs, no vendor lock-in.</div>

    <div class="faq-item" onclick="toggleFaq(this)"><span class="faq-q">How does this relate to eIDAS 2.0?</span><span class="faq-arrow">&darr;</span></div>
    <div class="faq-a">idonce uses the same credential format (SD-JWT-VC) and protocols (OpenID4VP, OpenID4VCI) as the EU Digital Identity Wallet. Any eIDAS system can verify idonce credentials.</div>

    <div class="faq-item" onclick="toggleFaq(this)"><span class="faq-q">Can one person have multiple identities?</span><span class="faq-arrow">&darr;</span></div>
    <div class="faq-a">Yes — one device equals one identity. For most use cases like bot prevention and review verification, device-level identity is sufficient.</div>

    <div class="faq-item" onclick="toggleFaq(this)"><span class="faq-q">What happens if I lose my phone?</span><span class="faq-arrow">&darr;</span></div>
    <div class="faq-a">Your identity is bound to the device and cannot be transferred. On a new phone, you create a new identity. The old one can be revoked.</div>
  </div>
</div>
</section>

<!-- CTA -->
<section class="cta-sec">
<div class="wrap">
  <div class="d-xxl"><span class="accent">Verify human.</span><br>Not identity.</div>
  <p class="txt" style="margin:20px auto 0;text-align:center">Open source. Free forever. No vendor lock-in.</p>
  <div style="margin-top:24px;display:flex;gap:12px;justify-content:center;flex-wrap:wrap">
    <a href="#demo" class="btn btn-p" style="background:var(--purple)">See it in Action</a>
    <a href="https://github.com/idonce" class="btn btn-s" target="_blank" rel="noopener">View on GitHub</a>
  </div>
</div>
</section>`

const landingJS = `<script>
function toggleFaq(el){el.classList.toggle('open')}
</script>`
