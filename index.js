// next packages:
require('@feizheng/next-js-core2');
require('@feizheng/next-node-base64');
require('@feizheng/next-random');
require('@feizheng/next-param');
require('@feizheng/next-json');
require('@feizheng/next-node-fetch');

var prelogin = require('./lib/prelogin');
var login = require('./lib/login');
var weibosso = require('./lib/weibo-sso');

module.exports = function(inOptions) {
  return new Promise(function(resolve, reject) {
    prelogin(inOptions).then(function(res1) {
      login(inOptions, res1).then(
        function(res2) {
          weibosso(inOptions, res2).then(function(res3) {
            resolve(res3);
          });
        },
        function(error) {
          reject(error);
        }
      );
    });
  });
};
