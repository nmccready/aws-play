#!/usr/bin/env node

const through = require('through2');
const { DecryptCommand } = require("@aws-sdk/client-kms");
const kms = require('../..');
const args = require('../args')();

const debug = require('../../../debug').spawn('aws:kms:decrypt');

const decoders = {
  default: (data) => data, // pass through
  decode: (encoding) => (data) => Buffer.from(data, encoding),
};

const decryptAsync = async (toDecrypt, keyId = process.env.KMS_ID, cb) => {
  try {
    const client = kms.getClient();
    debug(() => `Decrypting data with key ${keyId}`);

    const decoder = args.encoding ? decoders.decode(args.encoding) : decoders.default;
    toDecrypt = decoder(String(toDecrypt));

    const opts = { CiphertextBlob: toDecrypt };

    if (args.forceKeyId) {
      if (keyId) {
        opts.KeyId = keyId;
      }
    }

    const data = await client.send(new DecryptCommand(opts));
    debug(() => `Decrypted data with key ${keyId}`);
    cb(null, Buffer.from(data.Plaintext).toString('utf8'));
  } catch (err) {
    debug(() => `Decryption failed: ${err.message}`);
    cb(err);
  }
}

const decrypt = through.obj((toDecrypt, _, cb) => {
  decryptAsync(toDecrypt, args['key-id'], cb);
});

module.exports = {
  decoders,
  decrypt,
  decryptAsync,
};
