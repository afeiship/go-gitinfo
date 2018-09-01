var prelogin = require('./lib/prelogin');
var login = require('./lib/login');
var weibosso = require('./lib/weibosso');

// 1. prelogin:
prelogin().then(function (response){
    login(response).then(function (data){
        console.log(data.crossDomainUrlList);
        weibosso(data.crossDomainUrlList[0]).then(function (response){
            console.log(response);
        });
    });
});


