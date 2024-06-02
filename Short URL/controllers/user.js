const {v4: uuidv4} = require('uuid')
const User = require("../models/user");
const { setUser } = require('../service/auth');

async function handleUserSignup(req, res){
    const {name, email, password} = req.body;
    await User.create({
        name: name, 
        email: email, 
        password: password});
    return res.render("home");
}

async function handleUserLogin(req, res){
    const {email, password} = req.body;
    const user = await User.findOne({ 
        email: email, 
        password: password});
    if(!user) return res.render("login", { 
        error: "Invalid username or password"
    });
    // const sessionID = uuidv4();
    // setUser(sessionID, user);
    console.log("userfound", user);
    const token = setUser(user)
    console.log("token=",token)
    res.cookie("uid", token);
    return res.redirect("/")
}

module.exports = {
    handleUserSignup,
    handleUserLogin
}
