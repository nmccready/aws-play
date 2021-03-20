const yargs = require('yargs');
const debug = require('../../../debug').spawn('args');

const getArgs = () => {
  const { argv } = yargs
    .option('encoding', {
      alias: 'e',
      type: 'string',
      description: `incoming or outgoing additional encoding / decoding
        IE:
          - encrypt -> will take the decimal output and encode in base64
          - decrypt -> will take base64 input and decode base64 to decimal
`,
      choices: ['base64', 'hex', 'binary'],
    })
    .option('key-id', {
      alias: 'k',
      type: 'string',
      description: 'aws kms id or alias defaults to proccess.env.KMS_ID',
    })
    .option('forceKeyId', {
      alias: 'fk',
      type: 'bool',
      description:
        'for decrypt which defaults to false IE uses first key that works, this is to foce a specific key usage',
    });

  debug(() => argv);
  return argv;
};

module.exports = getArgs;
