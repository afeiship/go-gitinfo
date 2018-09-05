var axios = require('axios');

module.exports = function (inOptions, inUrl) {
  return new Promise(function (resolve, reject) {
    axios.get(inUrl).then(function (response) {
      resolve(response.headers['set-cookie'][0]);
    });
  })
};
