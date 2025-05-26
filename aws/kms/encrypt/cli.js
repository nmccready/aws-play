#!/usr/bin/env node
const { encrypt, encryptAsync } = require('./index');
const getArgs = require('../args');

const args = getArgs();

if (args.data) {
  encryptAsync(args.data, args.keyId, (err, result) => {
    if (err) {
      console.error('Encryption error:', err);
      process.exit(1);
    }
    console.log(result);
    process.exit(0);
  })
}

process.stdin.pipe(encrypt).pipe(process.stdout);
