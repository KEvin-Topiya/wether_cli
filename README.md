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

## 🚀 Installation

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

When prompted, enter the **latitude** and **longitude** of the location you'd like to check the weather for.

Example:

```
Enter latitude: 28.61
Enter longitude: 77.20

🌡️  Temperature: 32.3°C
💨 Wind Speed: 14.2 km/h
🌈 Conditions: Clear sky
⏰ Time: 2025-04-30T12:00
```

---

## 🌐 Supported Locations

Any global location using geographic coordinates (latitude & longitude).

You can use [latlong.net](https://www.latlong.net/) or Google Maps to find coordinates.

---

## 📦 Dependencies

- Standard Go libraries (`net/http`, `encoding/json`, `fmt`, etc.)

---

## 📄 License

MIT License

---

## 🙋‍♂️ Author

**Your Name**  
[GitHub Profile](https://github.com/yourusername)