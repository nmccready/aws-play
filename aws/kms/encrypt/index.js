const through = require('through2');
const { EncryptCommand } = require("@aws-sdk/client-kms");
const kms = require('../..');
const getArgs = require('../args');

const args = getArgs();
const debug = require('../../../debug').spawn('aws:kms:encrypt');

const encoders = {
  default: (data) => data, // pass through
  encode: (encoding) => (data) => Buffer.from(data, 'utf8').toString(encoding),
};

const encryptAsync = async (toEncrypt, keyId = process.env.KMS_ID, cb) => {
  try {
    const client = kms.getClient();
    debug(() => `Encrypting data with key ${keyId}`);
    const data = await client.send(new EncryptCommand({ KeyId: keyId, Plaintext: toEncrypt }));
    debug(() => `Encrypted data with key ${keyId}`);
    const encoder = args.encoding ? encoders.encode(args.encoding) : encoders.default;
    const encrypted = encoder(data.CiphertextBlob);
    cb(null, encrypted);
  } catch (err) {
    debug(() => `Encryption failed: ${err.message}`);
    cb(err);
  }
}

const encrypt = through.obj((toEncrypt, _, cb) => {
  encryptAsync(toEncrypt, args['key-id'], cb)
});


module.exports = {
  encoders,
  encrypt,
  encryptAsync,
};
