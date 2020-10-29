const yargs = require('yargs');

const getArgs = () => {
  const { argv } = yargs.option('encoding', {
    alias: 'e',
    type: 'string',
    description: `incoming or outgoing additional encoding / decoding
        IE:
          - encrypt -> will take the decimal outpu and encode in base64
          - decrypt -> will take base64 input and decode base64 to decimal
`,
    choices: ['base64', 'hex', 'binary'],
  });

  return argv;
};

module.exports = getArgs;
