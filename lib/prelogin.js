var PRELOGIN_RE = /{.*?}/;
var config = require('./config');
var encrypt = require('./encrypt');

module.exports = function(inOptions) {
  return new Promise(function(resolve, reject) {
    var data = {
      entry: 'weibo',
      callback: 'sinaSSOController.preloginCallBack',
      rsakt: 'mod',
      checkpin: 1,
      client: 'ssologin.js(v1.4.15)',
      _: Date.now(),
      su: encrypt.suEncrypt(inOptions.username)
    };

    nx.NodeFetch.get(config.PRELOGIN_URL, data, { responseType: 'text' }).then(
      (res) => {
        var matched = res.match(PRELOGIN_RE);
        resolve(nx.parse(matched[0]));
      }
    );
  });
};
