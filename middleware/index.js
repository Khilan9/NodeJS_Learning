const express = require('express')

function useURLEncoder() {
    console.log("In url encode");
    return express.urlencoded({ extended: false });
}

function Middleware1(req, res, next){
    console.log("Middleware 1 logged");
    next()
}

function Middleware2(req, res, next){
    console.log("Middleware 2 logged");
    next()
}

module.exports = {
    useURLEncoder,
    Middleware1,
    Middleware2
}