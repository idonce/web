package main

const legalCSS = `<style>
.content .legal{max-width:680px;padding:48px 28px 80px}
.content .legal a{color:#7c3aed}
.content .legal h1{font-family:'Space Grotesk',system-ui,sans-serif;font-size:clamp(32px,5vw,48px);font-weight:700;letter-spacing:-1.5px;line-height:1;text-transform:uppercase;margin-bottom:8px}
.content .legal h1 span{color:#7c3aed}
.content .legal .sub{color:#a8a29e;font-size:15px;margin-bottom:48px}
.content .legal h2{font-family:'Space Grotesk',system-ui,sans-serif;font-size:20px;font-weight:700;letter-spacing:-0.5px;text-transform:uppercase;margin:40px 0 12px;padding-top:24px;border-top:1px solid #e7e5e4}
.content .legal h2:first-of-type{border-top:none;padding-top:0}
.content .legal p,.content .legal li{font-size:15px;color:#57534e;margin-bottom:12px}
.content .legal ul{margin-left:20px;margin-bottom:16px}
.content .legal li{margin-bottom:6px}
.content .legal strong{color:#0a0a0a}
.content .legal .info-box{background:#f5f5f4;border-radius:12px;padding:20px;margin:20px 0;font-size:14px;line-height:1.6}
.content .legal .info-box strong{display:block;margin-bottom:4px}
</style>`

const privacyBody = `<div class="legal">
<h1>Privacy<br><span>Policy</span></h1>
<p class="sub">Last updated: April 2, 2026</p>

<h2>Who we are</h2>
<p><strong>Gradient Zero Deutschland GmbH</strong><br>Vogesenring 57, 79219 Staufen, Deutschland<br>HRB 722891, Amtsgericht Freiburg<br>USt-IdNr.: DE335929671</p>
<p>Contact: privacy@idonce.com</p>

<h2>What idonce does</h2>
<p>idonce is a human verification system. It allows users to prove that a real person is actively present on a device, without revealing personal identity. idonce is designed with <strong>privacy as a core principle</strong>, not as an afterthought.</p>

<h2>What data we collect</h2>
<div class="info-box">
<strong>Short answer: almost nothing.</strong>
idonce does not collect, store, or process personal data. No name, no email, no phone number, no location, no browsing history, no device identifiers beyond a pseudonymous cryptographic hash.
</div>

<p>Specifically:</p>
<ul>
<li><strong>Device ID:</strong> A SHA-256 hash of your device's public key. This is pseudonymous — it cannot be traced back to your name, phone number, or any other identifier. It is generated locally on your device.</li>
<li><strong>Credentials:</strong> SD-JWT-VC credentials are stored exclusively on your device in encrypted secure storage (iOS Keychain / Android EncryptedSharedPreferences). They are never uploaded to our servers.</li>
<li><strong>Verification history:</strong> Stored locally on your device only. No server has access to your verification log.</li>
<li><strong>IP addresses:</strong> Our servers process your IP address to serve requests. IP addresses are not stored or logged in association with your Device ID.</li>
</ul>

<h2>What data we do NOT collect</h2>
<ul>
<li>Name, email, phone number, or any personal identifiers</li>
<li>Biometric data (FaceID/Fingerprint data never leaves your device's Secure Enclave)</li>
<li>Location data</li>
<li>Browsing history or app usage patterns</li>
<li>Device identifiers (IDFA, GAID, IMEI)</li>
<li>Contacts, photos, or any other on-device data</li>
</ul>

<h2>Credential issuance</h2>
<p>When you request a credential from the idonce Issuer:</p>
<ul>
<li>Your device sends a <strong>cryptographic proof of key possession</strong> (a signed JWT). This proves your device controls a specific key pair — nothing more.</li>
<li>If device attestation is provided (Apple App Attest / Google Play Integrity), the attestation data is verified and discarded. We do not store attestation objects.</li>
<li>The issued credential (SD-JWT-VC) is sent to your device and not retained on our servers.</li>
</ul>

<h2>Credential presentation</h2>
<p>When you share a credential with a platform (verifier):</p>
<ul>
<li>You choose which claims to disclose via <strong>selective disclosure</strong>. The platform only sees the claims you selected.</li>
<li>The presentation is sent directly from your device to the verifier. idonce does not act as an intermediary and does not see or log the presentation.</li>
</ul>

<h2>Third-party services</h2>
<ul>
<li><strong>Apple App Attest / Google Play Integrity:</strong> Used for optional hardware attestation. Subject to Apple's / Google's respective privacy policies.</li>
</ul>
<p>All fonts are self-hosted. We do not use Google Fonts, analytics, tracking pixels, advertising networks, or any other third-party data collection services.</p>

<h2>Data retention</h2>
<p>We do not retain personal data. Server-side data is limited to:</p>
<ul>
<li><strong>Revocation list:</strong> A list of revoked Device IDs (pseudonymous hashes). Retained permanently for security.</li>
<li><strong>Server signing key:</strong> Used to sign credentials. No personal data.</li>
</ul>

<h2>Your rights (GDPR)</h2>
<p>Since we do not process personal data in the GDPR sense, most data subject rights do not apply. However:</p>
<ul>
<li><strong>Right to deletion:</strong> You can delete all local data by removing the app. You can revoke your device identity via the app settings.</li>
<li><strong>Right to information:</strong> This privacy policy describes all data processing.</li>
<li><strong>Complaint:</strong> You may contact the data protection authority of Baden-Württemberg (LfDI).</li>
</ul>

<h2>Children</h2>
<p>idonce does not knowingly collect data from children under 16. The app requires biometric capability, which is typically controlled by parents on children's devices.</p>

<h2>Changes</h2>
<p>We may update this policy. Changes will be posted on this page with an updated date.</p>

<h2>Contact</h2>
<p>Gradient Zero Deutschland GmbH<br>Vogesenring 57, 79219 Staufen, Deutschland<br>privacy@idonce.com</p>
</div>`

const termsBody = `<div class="legal">
<h1>Terms of<br><span>Service</span></h1>
<p class="sub">Last updated: April 2, 2026</p>

<h2>Provider</h2>
<p><strong>Gradient Zero Deutschland GmbH</strong><br>Vogesenring 57, 79219 Staufen, Deutschland<br>HRB 722891, Amtsgericht Freiburg<br>USt-IdNr.: DE335929671</p>

<h2>Scope</h2>
<p>These terms govern your use of the idonce mobile application ("App"), the idonce credential issuance service ("Issuer"), and the idonce website at idonce.com ("Website"). Together, these are referred to as the "Service".</p>

<h2>Description of Service</h2>
<p>idonce provides a human verification system that allows users to cryptographically prove that a real person is actively present on a biometrically-secured device. The Service issues eIDAS 2.0 compatible verifiable credentials (SD-JWT-VC) and enables their presentation to third-party platforms via OpenID4VP.</p>

<h2>Eligibility</h2>
<p>You must be at least 16 years old to use the Service. You must have a compatible device with biometric capabilities (FaceID, TouchID, or Fingerprint).</p>

<h2>User obligations</h2>
<ul>
<li>You will not use the Service for illegal purposes, fraud, or impersonation.</li>
<li>You will not attempt to reverse-engineer, tamper with, or circumvent the security mechanisms of the Service.</li>
<li>You will not create automated scripts or bots to interact with the Service at scale.</li>
<li>You are responsible for keeping your device secure. Loss of your device means loss of your identity.</li>
</ul>

<h2>Credentials</h2>
<ul>
<li>Credentials are issued based on cryptographic proof of device ownership and biometric authentication. They are not a guarantee of personal identity.</li>
<li>Credentials are bound to your device and cannot be transferred.</li>
<li>Credentials may expire. You can request new credentials from any compatible issuer.</li>
<li>We may revoke credentials if we detect misuse or security violations.</li>
</ul>

<h2>No warranty</h2>
<p>The Service is provided "as is" without warranty of any kind. We do not guarantee that:</p>
<ul>
<li>The Service will be available at all times without interruption.</li>
<li>Credentials will be accepted by all platforms.</li>
<li>The Service prevents all forms of bot activity or fraud.</li>
</ul>

<h2>Limitation of liability</h2>
<p>To the extent permitted by German law, Gradient Zero Deutschland GmbH is not liable for:</p>
<ul>
<li>Loss of credentials due to device loss, theft, or damage.</li>
<li>Actions taken by third-party platforms based on verification results.</li>
<li>Indirect, incidental, or consequential damages arising from use of the Service.</li>
</ul>
<p>This does not affect liability for intent or gross negligence under German law (§§ 276, 278 BGB).</p>

<h2>Intellectual property</h2>
<p>The idonce software is open source. Specific license terms are available in each repository. The idonce name, logo, and brand are trademarks of Gradient Zero Deutschland GmbH.</p>

<h2>Governing law</h2>
<p>These terms are governed by the laws of the Federal Republic of Germany. Place of jurisdiction is Freiburg im Breisgau.</p>

<h2>Changes</h2>
<p>We may update these terms. Continued use of the Service after changes constitutes acceptance.</p>

<h2>Contact</h2>
<p>Gradient Zero Deutschland GmbH<br>Vogesenring 57, 79219 Staufen, Deutschland<br>legal@idonce.com</p>
</div>`

const legalNoticeBody = `<div class="legal">
<h1>Legal<br><span>Notice</span></h1>
<p class="sub">Information pursuant to § 5 TMG (German Telemedia Act)</p>

<h2>Provider</h2>
<div class="info-box">
<strong>Gradient Zero Deutschland GmbH</strong>
Vogesenring 57<br>
79219 Staufen, Germany
</div>

<h2>Commercial Register</h2>
<p>Local Court (Amtsgericht) Freiburg<br>HRB 722891</p>

<h2>VAT ID</h2>
<p>VAT identification number pursuant to § 27 a of the German VAT Act:<br><strong>DE335929671</strong></p>

<h2>Managing Directors</h2>
<p>As registered in the commercial register.</p>

<h2>Contact</h2>
<p>Email: info@idonce.com</p>

<h2>Responsible for Content</h2>
<p>Pursuant to § 55 (2) RStV:<br>Gradient Zero Deutschland GmbH<br>Vogesenring 57<br>79219 Staufen, Germany</p>

<h2>Dispute Resolution</h2>
<p>The European Commission provides a platform for online dispute resolution (ODR): <a href="https://ec.europa.eu/consumers/odr" target="_blank">https://ec.europa.eu/consumers/odr</a></p>
<p>We are neither willing nor obliged to participate in dispute resolution proceedings before a consumer arbitration board.</p>

<h2>Liability for Content</h2>
<p>As a service provider, we are responsible for our own content on these pages under general law pursuant to § 7 (1) TMG. However, pursuant to §§ 8 to 10 TMG, we are not obligated to monitor transmitted or stored third-party information or to investigate circumstances that indicate illegal activity.</p>

<h2>Liability for Links</h2>
<p>Our website contains links to external third-party websites over whose content we have no influence. Therefore, we cannot accept any liability for this third-party content. The respective provider or operator of the linked pages is always responsible for the content of those pages.</p>

<h2>Copyright</h2>
<p>The idonce software is open source. The respective license terms can be found in the source code repositories. The name idonce, logo, and brand are property of Gradient Zero Deutschland GmbH.</p>
</div>`
