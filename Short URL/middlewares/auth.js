const { getUser } = require("../service/auth");

async function restrictToLoggedinUserOnly(req, res, next){
    const useruuid = req.cookies?.uid;
    console.log(useruuid)
    if(!useruuid) return res.redirect('login');
    const user = getUser(useruuid);
    if(!useruuid) return res.redirect('login');
    req.user = user;
    next();
}

async function checkAuth(req, res, next){
    const useruuid = req.cookies?.uid;
    console.log("useruuid", useruuid)
    const user = getUser(useruuid);
    req.user = user;
    next();
}

module.exports = {
    restrictToLoggedinUserOnly,
    checkAuth
}