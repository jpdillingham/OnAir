const Gpio = require('onoff').Gpio;
const express = require('express');

// this is the pinout for the 3-channel relay HAT sold on Amazon by Electronics-Salon
// your board may differ.
const pins = {
	channel1: 26,
	channel2: 20,
	channel3: 21
};

const lightChannel = new Gpio(pins.channel1, 'out');

const app = express();
app.use(express.json());

app.get('/status', (req, res) => {
	res.status(200)
	res.send(JSON.stringify(lightChannel.readSync()));
});

app.post('/on', (req, res) => {
	lightChannel.writeSync(1);
	res.status(200).send();
});

app.post('/off', (req, res) => {
	lightChannel.writeSync(0);
	res.status(200).send();
});

process.on('SIGNINT', _ => {
	channel1.unexport();
});

app.listen(3000, () => console.log('Listening at http://localhost:3000'));