# LinkedIn Messenger Automation Tool 🚀

A Golang-based LinkedIn messaging automation tool built as part of the
**LinkedIn Automation Internship Assignment (October Birth Month)**.

> **Note:** I had no prior experience with Golang before starting this assignment.
> The focus was on quickly understanding the existing codebase and implementing
> missing or incomplete features using documentation and AI-assisted
> development tools.

## 📋 Features Implemented

### Core Features ✅
| Feature | Description | Status |
|---------|-------------|--------|
| **Message Templates** | Personalized messages with `{{firstName}}`, `{{lastName}}`, `{{company}}`, `{{title}}` placeholders | ✅ Done |
| **Message Queue** | FIFO queue with configurable delays between messages | ✅ Done |
| **Rate Limiting** | Configurable max messages per hour/day to avoid LinkedIn detection | ✅ Done |
| **Logging System** | Detailed logs with timestamps, success/failure status, and error reasons | ✅ Done |
| **Safe Error Handling** | Auto-stops after consecutive failures to prevent spam/detection | ✅ Done |

### Bonus Features ✅
| Feature | Description | Status |
|---------|-------------|--------|
| **Scheduled Messages** | Schedule messages to be sent at specific times | ✅ Done |
| **Pause/Resume** | Pause and resume automation without losing progress | ✅ Done |
| **CLI Commands** | `start`, `stop`, `stats`, `demo`, `templates` commands | ✅ Done |
| **Queue Persistence** | Queue state saved to disk, survives restarts | ✅ Done |

## 🏗️ Project Structure

```
linkedin-automation/
├── main.go              # Entry point
├── config.json          # Configuration file
├── go.mod               # Go module definition
├── README.md            # This file
│
├── cmd/                 # CLI commands (Cobra)
│   └── root.go          # All CLI command definitions
│
├── config/              # Configuration handling
│   └── config.go        # Load/save config, defaults
│
├── models/              # Data structures
│   └── models.go        # Contact, Message, Stats structs
│
├── templates/           # Message template processing
│   └── templates.go     # Placeholder replacement logic
│
├── queue/               # Message queue management
│   └── queue.go         # Queue operations, rate tracking
│
├── logger/              # Logging system
│   └── logger.go        # Structured logging to file/console
│
└── automation/          # Core automation engine
    └── automation.go    # Main loop, send logic, error handling
```

## 🚀 How to Run

### Prerequisites
- Go 1.21 or higher installed
- Terminal/Command Prompt

### Quick Start

```bash
# 1. Navigate to project folder
cd linkedin-automation

# 2. Download dependencies
go mod tidy

# 3. Run the demo (recommended first time)
go run main.go demo

# 4. Or start with your own contacts
go run main.go start
```

### Available Commands

```bash
# Run demo with sample contacts
go run main.go demo

# Start automation
go run main.go start

# View statistics
go run main.go stats

# List available message templates
go run main.go templates

# Add a contact (demo contact)
go run main.go add

# Get help
go run main.go --help
```

### Configuration

Edit `config.json` to customize behavior:

```json
{
  "max_messages_per_hour": 20,    // Rate limit per hour
  "max_messages_per_day": 100,   // Rate limit per day
  "min_delay_seconds": 30,       // Min delay between messages
  "max_delay_seconds": 90,       // Max delay (randomized)
  "max_consecutive_errors": 3    // Stop after N errors
}
```

## 💡 Key Design Decisions

### 1. Human-like Behavior
- Random delays between messages (30-90 seconds)
- Rate limiting to avoid LinkedIn's spam detection
- Configurable limits based on LinkedIn's known thresholds

### 2. Safety First
- Auto-stop after consecutive failures
- Graceful shutdown with Ctrl+C
- Queue persistence (no lost messages on restart)

### 3. Simple & Readable Code
- Clear comments explaining logic
- Single responsibility per package
- No over-engineering

### 4. Template System
```
Template: "Hi {{firstName}}, I noticed you work at {{company}}..."
Contact:  {FirstName: "John", Company: "Google"}
Result:   "Hi John, I noticed you work at Google..."
```

## 🛠️ Tools Used

- **Language**: Golang 1.21
- **CLI Framework**: Cobra (industry standard)
- **AI Assistance**: Cursor AI (as permitted in assignment)
- **Version Control**: Git

## 📝 What I Learned

1. **Golang Basics**: structs, interfaces, goroutines, channels
2. **Concurrency**: Thread-safe operations with sync.Mutex
3. **CLI Development**: Using Cobra for professional CLI
4. **Rate Limiting**: Implementing hourly/daily limits
5. **Error Handling**: Safe shutdown patterns

## ⚠️ Important Notes

- This is a **simulation** - actual LinkedIn integration would require:
  - LinkedIn's official API (limited access), or
  - Browser automation (Selenium/Playwright)
- The `sendMessage` function simulates sending with 90% success rate
- In production, never store credentials in config files (use env vars)

## 🎯 Assignment Requirements Met

| Requirement | Implementation |
|-------------|----------------|
| Message automation with templates | ✅ `templates/templates.go` |
| Message queue with delays | ✅ `queue/queue.go` |
| Rate limiting | ✅ `queue/queue.go` + `config/config.go` |
| Logging system | ✅ `logger/logger.go` |
| Safe error handling | ✅ `automation/automation.go` |
| Scheduled messages | ✅ `automation/automation.go` |
| Pause/Resume | ✅ `automation/automation.go` |
| CLI commands | ✅ `cmd/root.go` |

---



Built with 💪 for LinkedIn Automation Internship Assignment  
(October Birth Month Candidate)



