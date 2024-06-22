const express = require('express');
const { createServer } = require('node:http');
const path = require('node:path');
const { Server } = require('socket.io');

const app = express();
const server = createServer(app);
const io = new Server(server);

app.use(express.static(path.resolve("./public")));

app.get('/', (req, res) => {
    res.sendFile('./public/index.html');
});

io.on('connection', (socket) => {
    console.log('a user connected', socket.id);
    socket.on('chat message', (msg) => {
        console.log('message: ' + msg);
        io.emit('chat message', msg);
      });
    socket.on('disconnect', () => {
        console.log('user disconnected');
    });
});

server.listen(8000, () => {
    console.log('server running at http://localhost:8000');
});