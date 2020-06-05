const Gpio = require('onoff').Gpio;
const argv = require('yargs').argv;

const pins = {
	ch1: 26,
	ch2: 20,
	ch3: 21
};

console.log(argv);

const test = new Gpio(pins[argv.pin], 'out');

test.writeSync(argv.state);

process.on('SIGNINT', _ => {
	test.undexport();
});
