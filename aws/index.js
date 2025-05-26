const { fromSSO } = require('@aws-sdk/credential-providers');
const { KMSClient } = require("@aws-sdk/client-kms");
const { addProxyToClient } = require("aws-sdk-v3-proxy");

const { HTTP_PROXY, HTTPS_PROXY } = process.env;

const debug = require('../debug').spawn('aws:kms:factory');

const initEnv = () => {
  const proxy = HTTPS_PROXY || HTTP_PROXY;
  if (proxy) {
    debug(() => `Using proxy: ${proxy}`);    
    return addProxyToClient(new KMSClient());
  }
  if (process.env.AWS_SSO_SESSION || process.env.AWS_PROFILE) {
    debug(() => 'Using SSO credentials');
    return new KMSClient({
      credentials: fromSSO(),
    });
  }
  debug(() => 'No proxy configured, using default KMS client');
  return new KMSClient();
};

const getClient = initEnv;

module.exports = {
  getClient,
  initEnv,
};
