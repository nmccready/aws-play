#!/usr/bin/env node
const { decrypt, decryptAsync } = require('./index');
const getArgs = require('../args');
const args = getArgs();

if (args.data) {
  decryptAsync(args.data, args.keyId, (err, result) => {
    if (err) {
      console.error('Decryption error:', err);
      process.exit(1);
    }
    console.log(result);
    process.exit(0);
  })
}

process.stdin.pipe(decrypt).pipe(process.stdout).once('error', console.error);
