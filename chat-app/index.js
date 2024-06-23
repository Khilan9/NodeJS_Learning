const express = require('express');
const { createServer } = require('node:http');
const path = require('node:path');
const { Server } = require('socket.io');

const app = express();
const server = createServer(app);

// connectionStateRecovery is feature used to recover when it is disconnected temporary(temporary store the data)
const io = new Server(server, {
    connectionStateRecovery: {}
});

app.use(express.static(path.resolve("./public")));

app.get('/', (req, res) => {
    res.sendFile('./public/index.html');
});

io.on('connection', (socket) => {
    console.log('a user connected', socket.id);
    socket.on('chat message', (msg, callback) => {
        console.log('message: ' + msg);
        io.emit('chat message', msg);
        callback({
            status: 'ok'
        });
    });
    socket.on('disconnect', () => {
        console.log('user disconnected');
    });
});

server.listen(8000, () => {
    console.log('server running at http://localhost:8000');
});