var weiboSSO = require('../index');


weiboSSO({
  username: '17607171608',
  password: 'Fei0.123'
}).then(
  (resp) => {
    console.log(resp);
  },
  (err) => {
    console.log(err);
  }
);
