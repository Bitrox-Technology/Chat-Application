# **Real-Time Chat Application**

---

## **Overview**
This is a real-time chat application built with a WebSocket server in Go and a React-based frontend. The application enables users to send and receive messages instantly, with all connected users viewing updates in real-time.

---

## **Features**
- Real-time messaging using WebSockets.
- Broadcasts messages to all connected clients.
- Displays chat history dynamically on the frontend.
- Easy-to-setup backend and frontend.

---

## **Tech Stack**
- **Backend**: Go (Golang) with WebSocket protocol.
- **Frontend**: React.js.
- **WebSocket Library**: Native WebSocket API in JavaScript and custom WebSocket implementation in Go.

---

## **Installation**

### **1. Clone the Repository**
```bash
git clone https://github.com/Bitrox-Technology/Chat-Application.git
cd Chat-Application

-setup backend
cd backend
go mod tidy
go run main.go


-setup client
cd frontend
npm install
npm start



