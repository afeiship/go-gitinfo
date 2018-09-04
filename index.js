var prelogin = require('./lib/prelogin');
var login = require('./lib/login');
var weibosso = require('./lib/weibo-sso');

module.exports = function(){
  return new Promise(function(resolve){
    prelogin().then(function (res1) {
      login(res1).then(function (res2) {
        weibosso(res2.crossDomainUrlList[0]).then(function (res3) {
          resolve(res3);
        });
      });
    });
  });
};
