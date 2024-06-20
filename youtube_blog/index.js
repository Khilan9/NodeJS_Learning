const path = require('path')
const express = require('express')
const userRoute = require('./routes/user');
const mongoose = require('mongoose');
const cookieParser = require('cookie-parser');
const { checkForAuthenticationCookie } = require('./middlewares/authentication');
const app = express()
PORT = 8000

mongoose.connect('mongodb://localhost:27017/blogify').then((e) => {
    console.log("MongoDB connected");
})

app.set('view engine', 'ejs');
app.set("views", path.resolve('./views'));

app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());
app.use(checkForAuthenticationCookie("token"));

app.get('/', (req, res) => {
    return res.render('home', {
        user: req.user
    });
})

app.use('/user', userRoute);

app.listen(PORT, () => {
    console.log(`Server started on PORT ${PORT}`);
})

