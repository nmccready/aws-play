#!/usr/bin/env node
const { encrypt } = require('./index');

process.stdin.pipe(encrypt).pipe(process.stdout);
