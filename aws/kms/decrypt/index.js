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
  kms.decrypt(
    {
      CiphertextBlob: text,
    },
    (err, data) => {
      if (err) {
        cb(err);
      }
      cb(null, data.Plaintext);
    }
  );
});

process.stdin.pipe(decrypt).pipe(process.stdout);
