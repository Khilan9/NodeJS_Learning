const { Router } = require('express');
const User = require('../models/user');

const router = Router();

router.get('/signin', (req, res) => {
    return res.render("signin");
});

router.get('/signup', (req, res) => {
    return res.render("signup");
});

router.get("/logout", (req, res) => {
    res.clearCookie("token").redirect("/");
});

router.post('/signin', async (req, res) => {
    try {
        const { email, password } = req.body;
        console.log(email);
        console.log(password);
        const token = await User.matchPasswordandgenerateToken(email, password);
        return res.cookie('token', token).redirect("/");
    }
    catch (error) {
        return res.render("signin", { error: "Incorrect email or password" })
    }
});

router.post('/signup', async (req, res) => {
    console.log(req.body);
    const { fullName, email, password } = req.body;
    await User.create({
        fullName,
        email,
        password
    })
    return res.redirect("/");
});

module.exports = router;

