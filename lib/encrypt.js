var base64 = require('node-base64');
var sinaSSOEncoder = require('./sinaSSOEncoder.js');


function suEncrypt(inUsername) {
  return base64.encode(encodeURIComponent(inUsername));
}

function spEncrypt(inPassword,inPreloginData) {
  var RSAKey = new sinaSSOEncoder.RSAKey();
  var nonce = inPreloginData.nonce;
  var pubkey = inPreloginData.pubkey;
  var servertime = inPreloginData.servertime;
  RSAKey.setPublic(pubkey, "10001");
  return RSAKey.encrypt([servertime, nonce].join("\t") + "\n" + inPassword);
}

module.exports = {
  suEncrypt: suEncrypt,
  spEncrypt: spEncrypt
};
