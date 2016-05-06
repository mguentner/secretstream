var shs = require('secret-handshake')
var fs = require('fs')
var pull = require("pull-stream")
var toPull = require('stream-to-pull-stream')

function readKeyF(fname) {
    var tmpobj = JSON.parse(fs.readFileSync(fname).toString())
    return {
        'publicKey': new Buffer(tmpobj.publicKey, 'base64'),
        'secretKey': new Buffer(tmpobj.secretKey, 'base64'),
    }
}

var appKey = new Buffer('IhrX11txvFiVzm+NurzHLCqUUe3xZXkPfODnp7WlMpk=', 'base64')

var createBob = shs.createServer(readKeyF('key.bob.json'), function(pub, cb) {
    // decide whether to allow access to pub.
    cb(null, true)
}, appKey)

pull(
    toPull.source(process.stdin),
    createBob(function(err, stream) {
        if (err) throw err
        // ...
    }),
    toPull.sink(process.stdout)
)