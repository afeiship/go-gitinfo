var base64 = require('node-base64');
var sinaSSOEncoder = require('./sinaSSOEncoder.js');
var privateConfig = require('../config.json');


function suEncrypt() {
  return base64.encode(encodeURIComponent(privateConfig.username));
}

function spEncrypt(inPreloginData) {
  var RSAKey = new sinaSSOEncoder.RSAKey();
  var nonce = inPreloginData.nonce;
  var pubkey = inPreloginData.pubkey;
  var servertime = inPreloginData.servertime;
  RSAKey.setPublic(pubkey, "10001");
  return RSAKey.encrypt([servertime, nonce].join("\t") + "\n" + privateConfig.password);
}


module.exports = {
  suEncrypt: suEncrypt,
  spEncrypt: spEncrypt
};