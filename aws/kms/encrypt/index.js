#!/usr/bin/env node
const { KMS } = require("aws-sdk");
const through = require("through2");

const kms = new KMS();

const encrypt = through.obj((text, _, cb) => {
  kms.encrypt(
    {
      KeyId: process.env.KMS_ID,
      Plaintext: text,
    },
    (err, data) => {
      if (err) {
        cb(err);
      }
      cb(null, data.CiphertextBlob);
    }
  );
});

process.stdin.pipe(encrypt).pipe(process.stdout);
