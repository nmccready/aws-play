const AWS = require('aws-sdk');
const proxyFact = require('proxy-agent');

const { HTTP_PROXY, HTTPS_PROXY } = process.env;

const init = ({ proxy, ...rest }) => {
  AWS.config.update({
    ...rest,
    httpOptions: {
      agent: proxyFact(proxy),
    },
  });
  return AWS;
};

const initEnv = () => {
  const proxy = HTTPS_PROXY || HTTP_PROXY;
  if (proxy) {
    return init({ proxy });
  }
  return AWS;
};

module.exports = {
  AWS,
  init,
  initEnv,
};
