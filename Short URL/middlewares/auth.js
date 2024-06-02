const { getUser } = require("../service/auth");

function checkForAuthentication(req, res, next){
    const useruuid = req.cookies?.uid;
    if(!useruuid) return next();
    req.user = null;
    const user = getUser(useruuid);
    req.user = user;
    return next();
}

function restrictTo(roles = []){
    return function(req, res, next){
        if(!req.user){
            return res.redirect("/login");
        }
        console.log("role",req.user.role);
        if(!req.user.role || !roles.includes(req.user.role)){
            return res.end('Unauthorized');
        }
        next();
    }
}

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
    checkAuth,
    checkForAuthentication,
    restrictTo
}