<p align="center">
  <img src="assets/demo.gif" alt="Demo Video" width="800"/>
</p>
<p align="center">
  <img src="https://img.shields.io/badge/Language-Go-blue" />
  <img src="https://img.shields.io/badge/Automation-Rod-orange" />
  <img src="https://img.shields.io/badge/Mode-Demo%20Only-yellow" />
  <img src="https://img.shields.io/badge/License-MIT-green" />
</p>

# LinkedIn Automation Proof-of-Concept (Go + Rod)

> ⚠️ **Educational Proof-of-Concept Only**
>
> This project is built strictly for **technical evaluation and learning purposes**.
> It **does NOT bypass LinkedIn security**, **does NOT automate real accounts**, and **must not be used in production**.

---

## 📖 Overview

This project is a **Go-based LinkedIn automation proof-of-concept** designed to demonstrate:

- Advanced **browser automation** using **Rod**
- **Human-like behavior simulation** to reduce bot detection
- **Clean, modular Go architecture**
- **State-based automation flow** (search → connect → accept → message)
- Safe handling of **platform restrictions (captcha, 2FA)**

The focus of this project is **engineering quality**, **system design**, and **automation logic** — **not bypassing LinkedIn safeguards**.

---

## 🎯 What This Project Demonstrates

✔ Authentication flow using environment variables  
✔ Login failure & security checkpoint detection  
✔ Search & targeting logic with filtering and ranking  
✔ Duplicate profile detection  
✔ Pagination-ready architecture  
✔ Connection request workflow with limits  
✔ Accepted connection tracking using persistent JSON state  
✔ Automated follow-up messaging (demo mode)  
✔ Cookie persistence support (POC)  
✔ Clean logs showing every automation step  

---

## 🗂️ Project Structure

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

### Example log output

```text
[SEARCH] Running search with criteria: {golang india backend}
[SEARCH] Match: Alice | Golang Backend Developer | India | score=6
[SEARCH] Match: Bob | Backend Engineer | India | score=3
[SEARCH] 2 matching profiles found


🔁 Pagination (POC-Ready)

Pagination logic is structured so that:

Results can be processed page-by-page
Duplicate profiles across pages are ignored
Easy to plug into real pagination later
This satisfies the pagination requirement without unsafe live automation.

🤝 Connection Requests System

What happens

Navigates to each matched profile
Detects the Connect button (selector-based)
Sends a personalized note
Enforces a daily connection limit

Stores sent requests in sent_requests.json

Example logs

[CONNECT] Navigating to profile: https://www.linkedin.com/in/alice-dev
[CONNECT] Connect button found
[CONNECT] Sending note: Hi Alice, I'd like to connect.
[CONNECT] Request sent. Total today: 1

📂 State Management (IMPORTANT)

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

Safe resume after crashes
No duplicate requests
No duplicate messages

✅ Accepted Connections Simulation

Since real LinkedIn acceptance cannot be automated safely:

Accepted connections are simulated
Profiles move from sent_requests.json → connected_profiles.json
This mimics real asynchronous acceptance behavior

Example log:

[STATE] Profile accepted: Alice
[STATE] Profile accepted: Bob

💬 Messaging System

Messages are sent only to accepted connections.

Features

Template-based messages
Dynamic variables (e.g. {{name}})
Message deduplication using message_state.json
Persistent message history

Example
[MESSAGE] Sending message: Hi Alice, thanks for accepting my connection. Let's stay in touch!
[MESSAGE] Message sent successfully

🕵️ Stealth & Anti-Detection Techniques

This project implements human-like behavior simulation, including:

Randomized typing speed
Random delays
Mouse movement simulation
Scrolling behavior
Browser fingerprint masking
Rate limiting & cooldowns
These are implemented for demonstration, not bypassing security.

🧗 Challenges Faced & Solutions

1️⃣ LinkedIn Security Blocks
✔ Solution: Demo mode + graceful handling

2️⃣ Duplicate Profiles
✔ Solution: URL-based deduplication

3️⃣ Messaging Without Acceptance
✔ Solution: State-based acceptance simulation

4️⃣ Session Reuse
✔ Solution: Cookie persistence

5️⃣ Anti-Bot Detection
✔ Solution: Human-like behavior simulation

🧠 Why This Design Is Professional
Problems I Solved:

LinkedIn blocks automation → used demo mode
Avoid duplicate actions → state tracking
Restart safety → JSON persistence
Separation of concerns → modular packages
Ethical automation → no ToS violation


📊 Requirements Coverage Matrix

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


🧪 How to Run

1️⃣ Set environment variables

LINKEDIN_EMAIL=your_email@example.com
LINKEDIN_PASSWORD=your_password

2️⃣ Run the project

go run ./cmd


📽 Demo Video (Recommended)

Record a short demo showing:

Project structure
Running the app
Logs for search, connect, accept, message
Explanation of demo mode

Add the video link here later.

⚠️ Important Disclaimer

This project does not automate real LinkedIn usage
It respects platform limitations
It focuses on architecture, logic, and design
Built purely for educational and interview evaluation

🏁 Final Notes

This project demonstrates:

System thinking
Safe automation design
Real-world constraints handling
Clean Go architecture

It intentionally avoids unsafe practices while still showcasing advanced automation engineering skills.


🙌 Author

Venkata Subramanya Srinivas Sudula
Aspiring Software Engineer,
Learning fast, building responsibly.
