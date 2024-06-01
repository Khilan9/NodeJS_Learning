const jwt = require('jsonwebtoken')
const secret = 'test'

function setUser(user){
    return jwt.sign(user, secret);
}

function getUser(token){
    // return sessionIDtousermap.get(id);
    return jwt.verify(token, secret)
}

module.exports = {
    setUser,
    getUser
}