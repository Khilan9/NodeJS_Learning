const express = require("express");
const URL = require("../models/urls");
const router = express.Router();

router.get('/', async (req, res) => {
    if(!req.user) {
        console.log("redirect to login page as user not found")
        return res.redirect('/login')
    }
    const allurls = await URL.find({createdBy: req.user._id});
    return res.render("home", {
        urls: allurls
    });
})

router.get('/signup', async (req, res) => {
    return res.render("signup")
})

router.get('/login', async (req, res) => {
    return res.render("login")
})

module.exports = router;