var NxHttpRequest = require('next-http-request');
var PRELOGIN_RE = /\{.*?\}/;
var config = require('./config');
var encrypt = require('./encrypt');

//load next package:
require('next-json');
require('next-param');


module.exports = function () {
    return new Promise(function (resolve, reject) {
        var data = {
            entry: 'weibo',
            callback: 'sinaSSOController.preloginCallBack',
            rsakt: 'mod',
            checkpin: 1,
            client: 'ssologin.js(v1.4.15)',
            _: Date.now(),
            su: encrypt.suEncrypt(),
        };
        var url = config.PRELOGIN_URL + '?' + nx.param(data);

        NxHttpRequest.get(url).then(function (response) {
            var matched = response.match(PRELOGIN_RE);
            resolve(nx.parse(matched[0]));
        });
    })
};