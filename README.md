# 🌤️ Go Weather CLI Tool

A simple CLI tool written in Go that fetches and displays the **current weather** using the free [Open-Meteo API](https://open-meteo.com/).

---

## ✨ Features

- Get current temperature, wind speed, time, and weather conditions
- Uses free and no-authentication-required Open-Meteo API
- Lightweight and fast terminal application
- Clean output with icons and formatting

---

## 🧰 Requirements

- Go 1.16 or higher
- Internet connection

---

## 🚀 Installation in linux

1. **Clone the repo:**

```bash
git clone https://github.com/yourusername/weather-cli.git
cd weather-cli
```

2. **Build the app:**

```bash
go build -o weather
```

3. **Run it:**

```bash
./weather
```

---

## 🛠️ Usage

Enter city name

Example:

```
./wether mumbai

🌡️  Temperature: 28.3°C
💨 Wind Speed: 0.5 km/h
🌈 Conditions: Partly cloudy
⏰ Time: 2025-04-30T19:15
```

---

## 📦 Dependencies

- Standard Go libraries (`net/http`, `encoding/json`, `fmt`, etc.)

---

## 📄 License

MIT License

---

## 🙋‍♂️ Author

**Kevin**  
[GitHub Profile](https://github.com/KEvin-topiya)