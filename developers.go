package main

const developersCSS = `<style>
/* Developer page */
.dev-page{max-width:680px;padding:48px 28px 80px}
.dev-page a{color:#7c3aed}
.dev-page h1{font-family:'Space Grotesk',system-ui,sans-serif;font-size:clamp(32px,5vw,48px);font-weight:700;letter-spacing:-1.5px;line-height:1;text-transform:uppercase;margin-bottom:8px}
.dev-page h1 span{color:#7c3aed}
.dev-page .sub{color:#a8a29e;font-size:15px;margin-bottom:48px}
.dev-page h2{font-family:'Space Grotesk',system-ui,sans-serif;font-size:20px;font-weight:700;letter-spacing:-0.5px;text-transform:uppercase;margin:40px 0 12px;padding-top:24px;border-top:1px solid #e7e5e4}
.dev-page h3{font-family:'Space Grotesk',system-ui,sans-serif;font-size:18px;font-weight:700;letter-spacing:-0.3px;text-transform:uppercase;margin:32px 0 12px}
.dev-page p,.dev-page li{font-size:15px;color:#57534e;margin-bottom:12px}
.dev-page strong{color:#0a0a0a}

/* Code */
pre{background:#1e1e1e;color:#d4d4d4;border-radius:12px;padding:24px;overflow-x:auto;font-size:13px;line-height:1.6;margin:16px 0 24px;font-family:'SF Mono',Consolas,monospace}
code{font-family:'SF Mono',Consolas,monospace;font-size:13px}
p code,li code,td code{background:#ede9fe;color:#7c3aed;padding:2px 8px;border-radius:6px;font-size:13px}

/* Steps */
.steps{display:flex;gap:0;margin:32px 0 40px;counter-reset:step}
.step{flex:1;position:relative;padding:24px 20px;border:1.5px solid #e7e5e4;background:#fff}
.step:first-child{border-radius:12px 0 0 12px}
.step:last-child{border-radius:0 12px 12px 0}
.step:not(:first-child){border-left:none}
.step::before{counter-increment:step;content:counter(step);font-family:'Space Grotesk',system-ui,sans-serif;font-size:36px;font-weight:700;color:#e7e5e4;line-height:1;display:block;margin-bottom:8px}
.step h4{font-family:'Space Grotesk',system-ui,sans-serif;font-size:14px;font-weight:700;text-transform:uppercase;letter-spacing:0.5px;margin-bottom:4px}
.step p{font-size:13px;color:#a8a29e;margin:0;line-height:1.5}
@media(max-width:860px){.steps{flex-direction:column}.step{border-radius:0!important}.step:not(:first-child){border-left:1.5px solid #e7e5e4;border-top:none}}

/* Use case cards */
.uc-grid{display:grid;grid-template-columns:repeat(2,1fr);gap:20px;margin:24px 0 40px}
.uc-card{border:1.5px solid #e7e5e4;border-radius:12px;padding:24px;background:#fff}
.uc-card .tag{display:inline-block;font-size:11px;font-weight:600;text-transform:uppercase;letter-spacing:0.5px;padding:4px 10px;border-radius:50px;margin-bottom:12px}
.uc-card h4{font-family:'Space Grotesk',system-ui,sans-serif;font-size:16px;font-weight:700;text-transform:uppercase;letter-spacing:-0.3px;margin-bottom:8px}
.uc-card p{font-size:14px;color:#a8a29e;margin:0;line-height:1.6}
.uc-card .detail{margin-top:12px;font-size:13px;color:#57534e;line-height:1.6}
@media(max-width:860px){.uc-grid{grid-template-columns:1fr}}

/* Pattern cards */
.pat-grid{display:grid;grid-template-columns:repeat(2,1fr);gap:20px;margin:24px 0}
.pat{border:1.5px solid #e7e5e4;border-radius:12px;padding:24px;background:#fff}
.pat h4{font-family:'Space Grotesk',system-ui,sans-serif;font-size:15px;font-weight:700;text-transform:uppercase;letter-spacing:-0.2px;margin-bottom:4px}
.pat .for{font-size:12px;color:#7c3aed;font-weight:500;margin-bottom:8px}
.pat p{font-size:13px;color:#a8a29e;margin:0;line-height:1.5}
@media(max-width:860px){.pat-grid{grid-template-columns:1fr}}

/* Response table */
.resp-table{width:100%;border-collapse:collapse;margin:16px 0 24px;font-size:14px}
.resp-table th{text-align:left;font-weight:600;padding:10px 12px;border-bottom:2px solid #e7e5e4;font-size:13px;text-transform:uppercase;letter-spacing:0.5px;color:#a8a29e}
.resp-table td{padding:10px 12px;border-bottom:1px solid #e7e5e4;color:#57534e}

/* Trust box */
.trust-box{background:#fff;border:1.5px solid #e7e5e4;border-radius:12px;padding:24px;margin:16px 0 24px}
.trust-box p{margin-bottom:8px}
.trust-box p:last-child{margin-bottom:0}

/* CTA */
.dev-cta{text-align:center;padding:60px 0 20px;border-top:1px solid #e7e5e4;margin-top:60px}
.dev-cta h2{border:none;padding:0;margin:0 0 8px}
.dev-cta p{color:#a8a29e;max-width:480px;margin:0 auto 24px}
.dev-cta .btns{display:flex;gap:12px;justify-content:center;flex-wrap:wrap}
.btn{display:inline-flex;align-items:center;gap:8px;font-family:'Inter',system-ui,sans-serif;font-weight:500;font-size:15px;border:none;cursor:pointer;transition:all 0.2s;text-decoration:none;border-radius:50px}
.btn-p{background:#0a0a0a;color:#fafaf9;padding:14px 28px}.btn-p:hover{background:#7c3aed}
.btn-s{background:none;color:#0a0a0a;padding:14px 28px;border:1.5px solid #e7e5e4}.btn-s:hover{border-color:#0a0a0a}
</style>`

const developersBody = `<div class="dev-page">
<h1>Developer<br><span>Guide</span></h1>
<p class="sub">Add human verification to your platform in 5 minutes. No API keys, no SDKs, just HTTP.</p>

<!-- Quick Start -->
<h2 id="quick-start">Quick Start</h2>

<div class="steps">
  <div class="step">
    <h4>Create Session</h4>
    <p>POST to create a verification session. Get a QR code back.</p>
  </div>
  <div class="step">
    <h4>Show QR</h4>
    <p>Display the QR code. User scans with any eIDAS wallet.</p>
  </div>
  <div class="step">
    <h4>Get Result</h4>
    <p>Poll for the result. Get verified claims.</p>
  </div>
</div>

<h3>1. Create a verification session</h3>
<pre>curl -X POST https://www.idonce.com/vp/sessions \
  -H 'Content-Type: application/json' \
  -d '{"client_id": "your-platform"}'</pre>

<p>Response:</p>
<pre>{
  "session_id": "vp_a1b2c3d4...",
  "qr_data": "openid4vp://authorize?request_uri=...",
  "deeplink": "idonce://vp?request_uri=...",
  "poll_url": "https://www.idonce.com/vp/sessions/vp_a1b2c3d4...",
  "expires_at": "2026-04-02T15:05:00Z"
}</pre>

<h3>2. Show the QR code to your user</h3>
<p>Render <code>qr_data</code> as a QR code on your page. On mobile, open <code>deeplink</code> directly. The user scans with the idonce app (or any eIDAS-compatible wallet), confirms with biometrics, and the credential is presented automatically.</p>

<h3>3. Poll for the result</h3>
<pre>curl https://www.idonce.com/vp/sessions/vp_a1b2c3d4...</pre>

<p>When the user completes verification:</p>
<pre>{
  "session_id": "vp_a1b2c3d4...",
  "status": "presented",
  "verified_at": "2026-04-02T15:01:23Z",
  "subject": "a9f3...e7b1",
  "vct": "https://idonce.com/credentials/HumanVerification/v1",
  "disclosed_claims": {
    "biometricConfirmed": true,
    "deviceBound": true,
    "attestationPlatform": "ios"
  }
}</pre>

<p>That's it. <code>biometricConfirmed: true</code> means a real human on a real device. <code>subject</code> is a stable, pseudonymous device hash you can use to detect duplicates.</p>

<!-- Integration Patterns -->
<h2 id="patterns">Integration Patterns</h2>

<div class="pat-grid">
  <div class="pat">
    <h4>QR Code Flow</h4>
    <div class="for">Desktop Web</div>
    <p>Create session, render QR, poll for result. Best for desktop websites where users verify with their phone.</p>
  </div>
  <div class="pat">
    <h4>Deep Link Flow</h4>
    <div class="for">Mobile Web &amp; Apps</div>
    <p>Open the <code>deeplink</code> URL directly. The idonce app handles verification and returns control to your app.</p>
  </div>
  <div class="pat">
    <h4>Pre-Registration</h4>
    <div class="for">Signup Flows</div>
    <p>Require verification before account creation. Only real humans can register. Eliminates bot signups entirely.</p>
  </div>
  <div class="pat">
    <h4>On-Demand</h4>
    <div class="for">Sensitive Actions</div>
    <p>Trigger verification at checkout, before voting, or when posting. Step-up authentication without passwords.</p>
  </div>
</div>

<!-- Use Cases -->
<h2 id="use-cases">Use Cases</h2>

<div class="uc-grid">
  <div class="uc-card">
    <div class="tag" style="background:#ede9fe;color:#7c3aed">Community</div>
    <h4>Forums &amp; Comments</h4>
    <p>Verify before first post. No CAPTCHAs, no email confirmation.</p>
    <div class="detail">Check <code>biometricConfirmed: true</code> and allow posting. Use <code>subject</code> to detect ban evasion across accounts.</div>
  </div>
  <div class="uc-card">
    <div class="tag" style="background:#dcfce7;color:#22c55e">Commerce</div>
    <h4>Fraud Prevention</h4>
    <p>Require verification at checkout for high-value orders.</p>
    <div class="detail">Combine <code>HumanVerification</code> + <code>EmailVerification</code> for verified email from a verified human.</div>
  </div>
  <div class="uc-card">
    <div class="tag" style="background:#dbeafe;color:#2563eb">Democracy</div>
    <h4>Voting &amp; Polls</h4>
    <p>One person, one vote. Without collecting names or IDs.</p>
    <div class="detail">The <code>subject</code> hash is stable per device. Same human = same subject = duplicate vote detected.</div>
  </div>
  <div class="uc-card">
    <div class="tag" style="background:#fef3c7;color:#d97706">API</div>
    <h4>Bot Protection</h4>
    <p>Gate API access behind human verification. Stronger than rate limiting.</p>
    <div class="detail">Verify once per session, then issue your own session token. Or verify per sensitive operation.</div>
  </div>
  <div class="uc-card">
    <div class="tag" style="background:#fce7f3;color:#db2777">Identity</div>
    <h4>Email + Human</h4>
    <p>Know it's a real email from a real person. Without knowing who.</p>
    <div class="detail">Request both: <code>HumanVerification</code> + <code>EmailVerification</code>. Verified email bound to a verified human.</div>
  </div>
  <div class="uc-card">
    <div class="tag" style="background:#f3e8ff;color:#7c3aed">Platform</div>
    <h4>Verified Badges</h4>
    <p>Give users a "Verified Human" badge backed by real verification.</p>
    <div class="detail">Store the <code>subject</code> hash with the user profile. Re-verify periodically or on suspicious activity.</div>
  </div>
</div>

<!-- Requesting Specific Credentials -->
<h2 id="credentials">Request Specific Credentials</h2>

<p>By default, a VP session requests <code>HumanVerification</code>. You can request any credential type:</p>

<pre>curl -X POST https://www.idonce.com/vp/sessions \
  -H 'Content-Type: application/json' \
  -d '{
    "client_id": "your-platform",
    "requested_credentials": [
      {
        "id": "human",
        "vct": "https://idonce.com/credentials/HumanVerification/v1",
        "claims": ["biometricConfirmed", "attestationPlatform"]
      },
      {
        "id": "email",
        "vct": "https://idonce.com/credentials/EmailVerification/v1",
        "claims": ["emailVerified", "email"]
      }
    ]
  }'</pre>

<p>The user chooses which claims to disclose. You only get what they consent to share.</p>

<!-- API Reference -->
<h2 id="api">API Reference</h2>

<h3>POST /vp/sessions</h3>
<p>Create a new verification session.</p>
<table class="resp-table">
  <tr><th>Field</th><th>Type</th><th>Description</th></tr>
  <tr><td><code>client_id</code></td><td>string</td><td>Your platform identifier (optional, defaults to "anonymous")</td></tr>
  <tr><td><code>requested_credentials</code></td><td>array</td><td>Credentials to request (optional, defaults to HumanVerification)</td></tr>
</table>

<h3>GET /vp/sessions/{id}</h3>
<p>Poll for verification result.</p>
<table class="resp-table">
  <tr><th>Field</th><th>Type</th><th>Description</th></tr>
  <tr><td><code>status</code></td><td>string</td><td><code>pending</code>, <code>presented</code>, or <code>expired</code></td></tr>
  <tr><td><code>subject</code></td><td>string</td><td>Stable device hash (only when presented)</td></tr>
  <tr><td><code>disclosed_claims</code></td><td>object</td><td>Verified claims the user consented to share</td></tr>
  <tr><td><code>vct</code></td><td>string</td><td>Credential type URI</td></tr>
  <tr><td><code>verified_at</code></td><td>string</td><td>ISO 8601 timestamp of verification</td></tr>
</table>

<!-- Trust Model -->
<h2 id="trust">Trust Model</h2>

<div class="trust-box">
  <p><strong>How do you know a credential is real?</strong></p>
  <p>Every credential contains an <code>iss</code> (issuer) field signed into the JWT. The issuer is identified by a <code>did:web</code> identifier tied to their domain. When a credential is presented, two things happen automatically:</p>
  <p><strong>1. Cryptographic verification</strong> — The verifier resolves the issuer's public key by fetching <code>https://www.idonce.com/.well-known/jwks.json</code> (derived from the <code>did:web</code> identifier) and verifies the JWT signature. This happens on every presentation, not just once. Nobody can forge credentials for a domain they don't control — same trust model as HTTPS.</p>
  <p><strong>2. Issuer trust</strong> — The verifier checks whether the issuer is trusted. In eIDAS 2.0, this will be handled by EU Trusted Lists — official registries of recognized issuers maintained by member states. Until then, you maintain your own allowlist.</p>
</div>

<p>If you use idonce as your verifier (the <code>/vp/sessions</code> API), step 1 is done for you automatically. You only need to decide which issuers to trust. If you build your own verifier, both steps are your responsibility:</p>

<pre>// Your server-side verification (step 2)
trustedIssuers := []string{"did:web:www.idonce.com"}

if !contains(trustedIssuers, credential.Issuer) {
    return errors.New("untrusted issuer")
}
// Step 1 (signature check via JWKS) must also be implemented
// — or use idonce's /vp/sessions API which handles both</pre>

<!-- Self-Hosting -->
<h2 id="self-hosting">Self-Hosting</h2>

<p>idonce is fully open source. You can run the entire stack yourself — issuer, verifier, or both.</p>

<h3>1. Run the Issuer</h3>

<pre># Docker
docker run -p 8080:8080 -v issuer_data:/app/data gradient0/idonce-issuer

# Or build from source
git clone https://github.com/idonce/issuer
cd issuer && go run .</pre>

<p>On first start, the issuer generates a unique ES256 signing key and persists it to <code>data/server_key.pem</code>. This key is your issuer's identity — back it up. If you lose it, all previously issued credentials become unverifiable.</p>

<h3>2. Configure your domain</h3>

<p>Set <code>BASE_URL</code> to your public domain. This determines your issuer's DID:</p>

<pre># Your issuer becomes did:web:verify.example.com
BASE_URL=https://verify.example.com go run .</pre>

<p>The issuer automatically serves the correct <code>did:web</code> resolution endpoints:</p>
<ul style="font-size:14px;color:#57534e;margin-bottom:16px">
  <li><code>/.well-known/did.json</code> — DID Document with your public key</li>
  <li><code>/.well-known/jwks.json</code> — JWKS for signature verification</li>
  <li><code>/.well-known/openid-credential-issuer</code> — OpenID4VCI metadata</li>
</ul>

<p>Your domain must be reachable via HTTPS for verifiers to resolve your DID. Use a reverse proxy like Caddy or nginx for TLS.</p>

<h3>3. What changes</h3>

<p>Credentials issued by your instance will have:</p>
<ul style="font-size:14px;color:#57534e;margin-bottom:16px">
  <li><code>iss: "did:web:verify.example.com"</code> — your DID, not idonce.com</li>
  <li>Signed with <strong>your</strong> key — only your JWKS can verify them</li>
  <li>Same credential format (SD-JWT-VC), same claims, same eIDAS 2.0 compatibility</li>
</ul>

<h3>4. What stays the same</h3>

<ul style="font-size:14px;color:#57534e;margin-bottom:16px">
  <li>The idonce mobile app accepts credentials from <strong>any</strong> issuer — users just scan your QR code</li>
  <li>The OpenID4VCI flow is standard — any eIDAS wallet works, not just the idonce app</li>
  <li>Credential format, claims, and verification logic are identical</li>
</ul>

<h3>5. Optional: Run the verifier too</h3>

<pre>git clone https://github.com/idonce/web
cd web
BASE_URL=https://verify.example.com go run .</pre>

<p>This gives you the full <code>/vp/sessions</code> API for verifying credentials on your own infrastructure.</p>

<h3>6. Tell verifiers to trust you</h3>

<p>Verifiers who want to accept your credentials need to add your DID to their trust list:</p>

<pre>trustedIssuers := []string{
    "did:web:www.idonce.com",          // official idonce
    "did:web:verify.example.com",      // your self-hosted instance
}</pre>

<p>When eIDAS 2.0 Trusted Lists go live, self-hosted issuers can register there for automatic trust across the EU ecosystem.</p>

<h3>Environment Variables</h3>

<table class="resp-table">
  <tr><th>Variable</th><th>Default</th><th>Description</th></tr>
  <tr><td><code>BASE_URL</code></td><td>http://localhost:8080</td><td>Public URL (determines DID, JWKS, all credential links)</td></tr>
  <tr><td><code>PORT</code></td><td>8080</td><td>Server port</td></tr>
  <tr><td><code>ATTESTATION_STRICT</code></td><td>false</td><td>Enforce Apple/Google hardware attestation</td></tr>
  <tr><td><code>ADMIN_API_KEY</code></td><td>-</td><td>Protect admin endpoints (stats, revocations)</td></tr>
  <tr><td><code>SENDGRID_API_KEY</code></td><td>-</td><td>Email verification OTP (logs code if unset)</td></tr>
  <tr><td><code>ALLOWED_ORIGIN</code></td><td>*</td><td>CORS origin</td></tr>
</table>

<!-- CTA -->
<div class="dev-cta">
  <h2>Ready to integrate?</h2>
  <p>No signup needed. Just start making API calls.</p>
  <div class="btns">
    <a href="https://github.com/idonce" class="btn btn-p" target="_blank" rel="noopener">View on GitHub</a>
    <a href="/#demo" class="btn btn-s">Try the Demo</a>
  </div>
</div>

</div>`

const developersJS = ``
