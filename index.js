// next packages:
require('next-js-core2');
require("next-node-base64");
require('next-random');
require('next-param');
require('next-json');


var prelogin = require("./lib/prelogin");
var login = require("./lib/login");
var weibosso = require("./lib/weibo-sso");

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
