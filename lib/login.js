var NxHttpRequest = require('next-http-request');
var config = require('./config');
var encrypt = require('./encrypt');
var nx = require('next-js-core2');
var axios = require('axios');
var querystring = require('querystring')
var data = {
    entry: 'weibo',
    gateway: 1,
    from: '',
    savestate: 30,
    useticket: 0,
    pagerefer: '',
    vsnf: 1,
    su: '[:encrypt_username]',
    service: 'sso',
    servertime: '[:servertime]',
    nonce: '[:nonce]',
    pwencode: 'rsa2',
    rsakv: '[:rsakv]',
    sp: '[:encrypt_password]',
    sr: 1920 * 1200,
    encoding: 'UTF-8',
    cdult: 3,
    domain: 'sina.com.cn',
    prelt: '[:prelt]',
    returntype: 'TEXT'
};

//load next package:
require('next-random');
require('next-param');

module.exports = function (inPreloginData) {
    var su = encrypt.suEncrypt();
    var sp = encrypt.spEncrypt(inPreloginData);
    var servertime = inPreloginData.servertime;
    var nonce = inPreloginData.nonce;
    var rsakv = inPreloginData.rsakv;

    var _data = nx.mix(data, {
        su: su,
        sp: sp,
        servertime: servertime,
        nonce: nonce,
        prelt: nx.random(100,1000),
        rsakv: rsakv
    });

    return new Promise(function (resolve, reject) {
        axios.post(config.LOGIN_URL + '?' + nx.param(_data)).then(function (response) {
            resolve(response.data)
        });
    })
};