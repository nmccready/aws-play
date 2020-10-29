#!/usr/bin/env node

const { KMS } = require("aws-sdk");
const through = require("through2");

const kms = new KMS();

const decrypt = through.obj((text, _, cb) => {
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
