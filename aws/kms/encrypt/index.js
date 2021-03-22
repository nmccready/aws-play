const through = require('through2');

const { initEnv } = require('../..');
const getArgs = require('../args');
const { KMS } = initEnv();

const args = getArgs();
const kms = new KMS();

const encoders = {
  default: (data) => data, // pass through
  encode: (encoding) => (data) => Buffer.from(data, 'utf8').toString(encoding),
};

const encrypt = through.obj((text, _, cb) => {
  kms.encrypt(
    {
      KeyId: args['key-id'] || process.env.KMS_ID,
      Plaintext: text,
    },
    (err, data) => {
      if (err) {
        cb(err);
      }

      encoder = args.encoding ? encoders.encode(args.encoding) : encoders.default;
      encrypted = encoder(data.CiphertextBlob);
      cb(null, encrypted);
    }
  );
});

module.exports = {
  encoders,
  encrypt,
};
