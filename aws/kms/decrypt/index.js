#!/usr/bin/env node

const { KMS } = require('aws-sdk');
const through = require('through2');
const buffer = require('buffer');
const getArgs = require('../args');

const kms = new KMS();

const args = getArgs();

const decoders = {
  default: (data) => data, // pass through
  decode: (encoding) => (data) => Buffer.from(data, encoding),
};

const decrypt = through.obj((text, _, cb) => {
  decoder = args.encoding ? decoders.decode(args.encoding) : decoders.default;
  text = decoder(String(text));

  const opts = { CiphertextBlob: text };

  if (args.forceKeyId) {
    const kmsId = args['key-id'] || process.env.KMS_ID;
    if (kmsId) {
      opts.KeyId = kmsId;
    }
  }

  kms.decrypt(opts, (err, data) => {
    if (err) {
      cb(err);
    }
    cb(null, data.Plaintext);
  });
});

module.exports = {
  decoders,
  decrypt,
};
