var prelogin = require('./lib/prelogin');
var login = require('./lib/login');
var weibosso = require('./lib/weibo-sso');

module.exports = function(inOptions){
  return new Promise(function(resolve){
    prelogin(inOptions).then(function (res1) {
      login(inOptions, res1).then(function (res2) {
        weibosso(inOptions, res2.crossDomainUrlList[0]).then(function (res3) {
          resolve(res3);
        });
      });
    });
  });
};
