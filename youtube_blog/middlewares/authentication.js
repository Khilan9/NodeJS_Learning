const { validateToken } = require("../services/authentication");

function checkForAuthenticationCookie(cookiename) {
    return (req, res, next) => {
        const tokenCookieValue = req.cookies[cookiename];
        if (!tokenCookieValue) {
            console.log("nocookie");
            return next();
        }
        try {
            const userpayload = validateToken(tokenCookieValue);
            console.log("userpayload", userpayload);
            req.user = userpayload;
        }
        catch (error) {}
        return next();
    }
}

module.exports = {
    checkForAuthenticationCookie
}