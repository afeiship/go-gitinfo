var weiboSSO = require('../index');


weiboSSO({
  username: 'YOUR_USERNAME',
  password: 'YOUR_PASSWORD'
}).then(
  (resp) => {
    console.log(resp);
  },
  (err) => {
    console.log(err);
  }
);
