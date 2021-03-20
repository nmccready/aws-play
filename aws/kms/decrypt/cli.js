#!/usr/bin/env node
const { decrypt } = require('./index');

process.stdin.pipe(decrypt).pipe(process.stdout).once('error', console.error);
