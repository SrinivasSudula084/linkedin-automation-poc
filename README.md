<p align="center">
  <video src="assests/demo1.mp4" autoplay loop muted playsinline width="100%"></video>
</p>
<p align="center">
  <img src="https://img.shields.io/badge/Language-Go-blue" />
  <img src="https://img.shields.io/badge/Automation-Rod-orange" />
  <img src="https://img.shields.io/badge/Mode-Demo%20Only-yellow" />
  <img src="https://img.shields.io/badge/License-MIT-green" />
</p>

<h1 align="center">LinkedIn Automation Proof of Concept (Go + Rod)</h1>

<p >
  <b>A safe, educational, demo-only LinkedIn automation system</b>
  Built to demonstrate browser automation, stealth techniques, and clean Go architecture
</p>

<p align="center">
  <i>⚠️ Educational Purpose Only · No real LinkedIn automation performed</i>
</p>

---

## 📌 Project Overview

This project is a **LinkedIn Automation Proof of Concept** built using **Go** and the **Rod browser automation library**.

The goal of this project is **NOT** to break LinkedIn rules.  
Instead, it demonstrates **how such systems are architected safely**, including:

- Human-like automation behavior
- Anti-detection techniques
- Clean modular Go design
- State persistence using JSON
- Graceful handling of login failures and security checkpoints

> 🔒 **Important**  
> LinkedIn automation violates LinkedIn’s Terms of Service.  
> This project runs in **DEMO MODE only** and never attempts to bypass captchas, 2FA, or security checks.

---

## 🎯 Why This Project Exists

Recruiters often want to evaluate:
- How fast you can learn a new language (Go)
- How you design automation systems
- How you handle real-world constraints (ToS, captchas, rate limits)
- How cleanly you structure large codebases

This project focuses on **engineering quality**, not misuse.

---

## 🧠 High-Level Flow (Simple Explanation)

Think of this project like a **robot assistant**:

1. 🔐 Tries to log in (demo only)
2. 🔍 Searches for people based on criteria
3. ➕ Sends connection requests (limited per day)
4. ⏳ Waits for acceptance (simulated)
5. 💬 Sends follow-up messages only to accepted connections
6. 💾 Saves everything in JSON files so it remembers next time

---

## 🗂️ Project Structure (Explained Line by Line)
```text

linkedin-automation-poc/
│
├── cmd/
│   ├── main.go        # Entry point (orchestrates everything)
│   └── cookies.go     # Cookie persistence logic
│
├── internal/
│   ├── auth/          # Login handling & failure detection
│   ├── config/        # Environment variable loading
│   ├── stealth/       # Anti-bot & human behavior simulation
│   ├── search/        # Search & targeting logic
│   ├── connection/    # Connection request + state handling
│   ├── messaging/     # Messaging system + templates
│   └── state/         # JSON-based persistence utilities
│
├── demo_profiles.json        # Demo LinkedIn profiles
├── sent_requests.json        # Profiles to whom requests were sent
├── connected_profiles.json  # Accepted connections
├── message_state.json        # Tracks who already received messages
├── message_history.json     # Full message log
├── cookies.json              # Stored session cookies
│
├── .env               # Environment variables (ignored in git)
├── go.mod / go.sum    # Go dependencies
└── README.md


```

---

## 🔐 Authentication System

### What it does

- Loads LinkedIn credentials from **environment variables**
- Attempts login using Rod
- Detects:
  - Invalid credentials
  - Captcha pages
  - Verification / 2FA checkpoints
- Exits safely without retries or bypass attempts
- Saves cookies for reuse

### Why this matters

LinkedIn actively blocks automation.  
Instead of trying to bypass it (which is unsafe), this project **detects security checkpoints and stops gracefully**, which is exactly what production-grade automation systems must do.

---

## 🔍 Search & Targeting System

### How search works (Demo Mode)

Profiles are loaded from `demo_profiles.json`, which simulates LinkedIn search results.

Search criteria includes:
- Job Title
- Company
- Location
- Keywords

### Matching Logic

- **Partial matching** (not exact strings)
- **Scoring system** assigns relevance points
- Profiles with higher scores rank first
- Duplicate profiles are automatically skipped using a `seen` map

### Example log

```text
[SEARCH] Running search with criteria: {golang india backend}
[SEARCH] Match: Alice | Golang Backend Developer | score=6
[SEARCH] Match: Bob | Backend Engineer | score=3
[SEARCH] 2 matching profiles found
```

### 🔁 Pagination (POC-Ready)

Pagination logic is structured so that:

- Results can be processed page-by-page
- Duplicate profiles across pages are ignored
- Easy to plug into real pagination later
- This satisfies the pagination requirement without unsafe live automation.



---

## ➕ Connection Requests System

### Features
- Navigates to profile URLs
- Detects "Connect" button (demo)
- Sends personalized notes
- Enforces daily limits
- Avoids duplicate requests

### State Files Used
- `sent_requests.json`

### Example log
```text
[CONNECT] Navigating to profile: https://www.linkedin.com/in/alice-dev
[CONNECT] Sending note: Hi Alice, I'd like to connect.
[CONNECT] Request sent. Total today: 1
[CONNECT] Connect button found
```
---

## 📂 State Management (IMPORTANT)

This project uses JSON-based state persistence.

State Files Explained

| File                      | Purpose                                     |
| ------------------------- | ------------------------------------------- |
| `sent_requests.json`      | Tracks profiles where a connection was sent |
| `connected_profiles.json` | Tracks profiles that accepted               |
| `message_state.json`      | Prevents duplicate messages                 |
| `message_history.json`    | Full message log                            |
| `cookies.json`            | Browser session cookies                     |

This allows:

- Safe resume after crashes
- No duplicate requests
- No duplicate messages


---

## ✅ Accepted Connections (State Management)

Because real acceptance cannot be automated safely:

- Acceptance is **simulated**
- New accepted profiles are moved from:
  - `sent_requests.json` ➜ `connected_profiles.json`

### Example log
```text
[STATE] Profile accepted: Alice
[STATE] Profile accepted: Bob
```

---

## 💬 Messaging System

### Messaging rules
- Messages sent **only after acceptance**
- Uses templates with variables
- Avoids duplicate messages
- Tracks history & state

### State Files Used
- `connected_profiles.json`
- `message_state.json`
- `message_history.json`

### Example log

[MESSAGE] Opening chat with: Alice
[MESSAGE] Sending message: Hi Alice, thanks for accepting my connection!
[MESSAGE] Sent follow-up to: Alice


---

## 🕵️ Stealth & Anti-Detection Techniques

#### Implemented techniques include:

- Fingerprint masking
- Human-like delays
- Random scrolling
- Typing simulation
- Rate limiting
- Activity scheduling
- Session reuse via cookies

These are **demonstrated**, not weaponized.

---

## 📁 JSON State Files (Why They Matter)

| File | Purpose |
|-----|--------|
| `demo_profiles.json` | Demo search dataset |
| `sent_requests.json` | Who received requests |
| `connected_profiles.json` | Who accepted |
| `message_state.json` | Prevent duplicate messages |
| `message_history.json` | Full message logs |
| `cookies.json` | Browser session reuse |

This allows the program to **resume safely** after restarts.

---

## 🚧 Challenges Faced & How They Were Solved

### ❌ LinkedIn blocks automation
✅ Solved using demo mode + checkpoint detection

### ❌ Duplicate profiles & messages
✅ Solved with URL-based state tracking

### ❌ Unsafe real automation
✅ Solved with controlled simulation

### ❌ Large codebase complexity
✅ Solved with clean modular architecture

### ❌ Messaging Without Acceptance
✅ Solution: State-based acceptance simulation

### ❌Session Reuse
✅ Solution: Cookie persistence

### ❌Anti-Bot Detection
✅ Solution: Human-like behavior simulation

---

## 🧠 Why This Design Is Professional
#### Problems I Solved:

- LinkedIn blocks automation → used demo mode
- Avoid duplicate actions → state tracking
- Restart safety → JSON persistence
- Separation of concerns → modular packages
- Ethical automation → no ToS violation

---

## 📊 Requirements Coverage Matrix

| Requirement             | Status |
| ----------------------- | ------ |
| Env-based login         | ✅      |
| Login failure handling  | ✅      |
| Captcha / 2FA detection | ✅      |
| Cookie persistence      | ✅      |
| Search filters          | ✅      |
| Pagination logic        | ✅      |
| Duplicate detection     | ✅      |
| Connection limits       | ✅      |
| Personalized notes      | ✅      |
| Accepted-only messaging | ✅      |
| Message tracking        | ✅      |

---

## 🧪 How to Run

### 1️⃣ Set environment variables
```text
LINKEDIN_EMAIL=your_email@example.com
LINKEDIN_PASSWORD=your_password
```

## 2️⃣ Run the project

```bash
go run ./cmd
```

---

## 📽 Demo Video (Recommended)

#### Record a short demo showing:

- Project structure
- Running the app
- Logs for search, connect, accept, message
- Explanation of demo mode

Add the video link here later.

---

## ⚠️ Important Disclaimer

- This project does not automate real LinkedIn usage
- It respects platform limitations
- It focuses on architecture, logic, and design
- Built purely for educational and interview evaluation


## 🏁 Final Notes

#### This project demonstrates:

- System thinking
- Safe automation design
- Real-world constraints handling
- Clean Go architecture

It intentionally avoids unsafe practices while still showcasing advanced automation engineering skills.

---

## 🙌 Author

#### Venkata Subramanya Srinivas Sudula
##### Aspiring Software Engineer
##### Learning fast, building responsibly
##### email: sudulasrinivas084@gmail.com
