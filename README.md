# Go Port Scanner Tool

####  Homework #1 Video Link: https://www.youtube.com/watch?v=8k7HNCdTwDo

## 📌 Description

This tool is a concurrent TCP port scanner written in Go. It allows scanning one or more targets using either a custom range of ports or specific ports, and provides advanced features such as:

- Banner grabbing
- Adjustable timeout and worker count
- Progress display
- JSON output
- Multiple targets in one run

---

## ⚙️ How to Build and Run

### ✅ 1. Build

Make sure you have Go installed, then build the tool:

```bash
go build -o portscanner main.go
```

### ✅ 2. Run

You can run the tool using:

```bash
./portscanner [flags]
```

### 🔧 Flags

| Flag | Description | Default |
|------|-------------|---------|
| `-targets` | Comma-separated list of target hostnames or IPs | `scanme.nmap.org` |
| `-start-port` | Starting port for range scan | `1` |
| `-end-port` | Ending port for range scan | `1024` |
| `-ports` | Specific ports to scan (comma-separated) | _none_ |
| `-workers` | Number of concurrent workers | `100` |
| `-timeout` | Connection timeout in seconds | `5` |
| `-json` | Output scan results as JSON | `false` |

---

### 📥 Examples

#### Scan a range of ports:
```bash
./portscanner -targets=scanme.nmap.org -start-port=20 -end-port=25
```

#### Scan specific ports:
```bash
./portscanner -targets=scanme.nmap.org -ports=22,80,443
```

#### Scan multiple targets:
```bash
./portscanner -targets=scanme.nmap.org,google.com -ports=22,80,443
```

#### Output as JSON:
```bash
./portscanner -targets=scanme.nmap.org -ports=80 -json
```

---

## 📤 Sample Output

### Scanning Specific Ports:
```
Scanning port 80 from scanme.nmap.org
Scanning port 80 from google.com
Scanning port 22 from google.com
Scanning port 22 from scanme.nmap.org
----------------------------------------------------------
Connection to scanme.nmap.org:22 was successful
Banner: SSH-2.0-OpenSSH_6.6.1p1 Ubuntu-2ubuntu2.13
----------------------------------------------------------
Connection to google.com:80 was successful (no banner)
Connection to scanme.nmap.org:80 was successful (no banner)
Attempt 1 to google.com:22 failed. Waiting 1s...
Scanning port 22 from google.com
Attempt 2 to google.com:22 failed. Waiting 2s...
Scanning port 22 from google.com
Attempt 3 to google.com:22 failed. Waiting 4s...
Failed to connect to google.com:22 after 3 attempts

---------------------------------
Scan Summary:

Open ports: 3
Total ports scanned: 4
Time taken: 22.044211918s
---------------------------------
```

---

### Scanning A Range Of Ports:
```
Scanning port 25/25 from scanme.nmap.org
Scanning port 23/25 from scanme.nmap.org
Scanning port 20/25 from scanme.nmap.org
Scanning port 20/25 from google.com
Scanning port 21/25 from google.com
Scanning port 23/25 from google.com
Scanning port 22/25 from scanme.nmap.org
Scanning port 24/25 from scanme.nmap.org
Scanning port 21/25 from scanme.nmap.org
Scanning port 22/25 from google.com
Scanning port 24/25 from google.com
Scanning port 25/25 from google.com
----------------------------------------------------------
Connection to scanme.nmap.org:22 was successful
Banner: SSH-2.0-OpenSSH_6.6.1p1 Ubuntu-2ubuntu2.13
----------------------------------------------------------
Attempt 1 to scanme.nmap.org:23 failed. Waiting 1s...
Attempt 1 to scanme.nmap.org:25 failed. Waiting 1s...
Attempt 1 to scanme.nmap.org:20 failed. Waiting 1s...
Attempt 1 to scanme.nmap.org:21 failed. Waiting 1s...
Attempt 1 to scanme.nmap.org:24 failed. Waiting 1s...
Scanning port 20/25 from scanme.nmap.org
Scanning port 23/25 from scanme.nmap.org
Scanning port 21/25 from scanme.nmap.org
Scanning port 25/25 from scanme.nmap.org
Scanning port 24/25 from scanme.nmap.org
Attempt 2 to scanme.nmap.org:21 failed. Waiting 2s...
Attempt 2 to scanme.nmap.org:20 failed. Waiting 2s...
Attempt 2 to scanme.nmap.org:25 failed. Waiting 2s...
Attempt 2 to scanme.nmap.org:23 failed. Waiting 2s...
Attempt 2 to scanme.nmap.org:24 failed. Waiting 2s...
Scanning port 20/25 from scanme.nmap.org
Scanning port 21/25 from scanme.nmap.org
Scanning port 23/25 from scanme.nmap.org
Scanning port 25/25 from scanme.nmap.org
Scanning port 24/25 from scanme.nmap.org
Attempt 3 to scanme.nmap.org:25 failed. Waiting 4s...
Attempt 3 to scanme.nmap.org:20 failed. Waiting 4s...
Attempt 3 to scanme.nmap.org:21 failed. Waiting 4s...
Attempt 3 to scanme.nmap.org:23 failed. Waiting 4s...
Attempt 3 to scanme.nmap.org:24 failed. Waiting 4s...
Attempt 1 to google.com:20 failed. Waiting 1s...
Attempt 1 to google.com:21 failed. Waiting 1s...
Attempt 1 to google.com:22 failed. Waiting 1s...
Attempt 1 to google.com:23 failed. Waiting 1s...
Attempt 1 to google.com:24 failed. Waiting 1s...
Attempt 1 to google.com:25 failed. Waiting 1s...
Scanning port 21/25 from google.com
Scanning port 20/25 from google.com
Scanning port 25/25 from google.com
Scanning port 22/25 from google.com
Scanning port 24/25 from google.com
Scanning port 23/25 from google.com
Failed to connect to scanme.nmap.org:25 after 3 attempts
Failed to connect to scanme.nmap.org:23 after 3 attempts
Failed to connect to scanme.nmap.org:20 after 3 attempts
Failed to connect to scanme.nmap.org:21 after 3 attempts
Failed to connect to scanme.nmap.org:24 after 3 attempts
Attempt 2 to google.com:20 failed. Waiting 2s...
Attempt 2 to google.com:21 failed. Waiting 2s...
Attempt 2 to google.com:23 failed. Waiting 2s...
Attempt 2 to google.com:24 failed. Waiting 2s...
Attempt 2 to google.com:22 failed. Waiting 2s...
Attempt 2 to google.com:25 failed. Waiting 2s...
Scanning port 20/25 from google.com
Scanning port 21/25 from google.com
Scanning port 24/25 from google.com
Scanning port 25/25 from google.com
Scanning port 22/25 from google.com
Scanning port 23/25 from google.com
Attempt 3 to google.com:20 failed. Waiting 4s...
Attempt 3 to google.com:21 failed. Waiting 4s...
Attempt 3 to google.com:24 failed. Waiting 4s...
Attempt 3 to google.com:22 failed. Waiting 4s...
Attempt 3 to google.com:25 failed. Waiting 4s...
Attempt 3 to google.com:23 failed. Waiting 4s...
Failed to connect to google.com:20 after 3 attempts
Failed to connect to google.com:21 after 3 attempts
Failed to connect to google.com:22 after 3 attempts
Failed to connect to google.com:24 after 3 attempts
Failed to connect to google.com:25 after 3 attempts
Failed to connect to google.com:23 after 3 attempts

---------------------------------
Scan Summary:

Open ports: 1
Total ports scanned: 12
Time taken: 22.067061974s
---------------------------------
```

---

### JSON Output If Enabled:
```json
[
  {
    "target": "scanme.nmap.org",
    "port": "22",
    "banner": "SSH-2.0-OpenSSH_6.6.1p1 Ubuntu-2ubuntu2.13\r\n"
  }
]
```
