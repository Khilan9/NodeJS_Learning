const path = require('path')
const express = require('express')
const userRoute = require('./routes/user');
const mongoose = require('mongoose');

const app = express()
PORT = 8000

mongoose.connect('mongodb://localhost:27017/blogify').then((e) => {
    console.log("MongoDB connected");
})

app.set('view engine', 'ejs');
app.set("views", path.resolve('./views'));

app.use(express.urlencoded({extended: false}));

app.get('/', (req, res) => {
    return res.render('home');
})

app.use('/user', userRoute);

app.listen(PORT, () => {
    console.log(`Server started on PORT ${PORT}`);
})

