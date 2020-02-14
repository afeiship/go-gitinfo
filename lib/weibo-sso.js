module.exports = function(inOptions, inUrl) {
  return new Promise(function(resolve, reject) {
    nx.NodeFetch.get(inUrl, null, { responseType: null }).then(function(response) {
      resolve(response.headers.raw()['set-cookie'][0]);
    });
  });
};
