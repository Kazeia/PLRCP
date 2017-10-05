"use strict"
var readline = require('linebyline');
var rl = readline('./Test.txt');
rl.on('line', function (line, lineCount, byteCount) {
    console.log(line)
})
.on('error', function (e) {
// something went wrong 
});