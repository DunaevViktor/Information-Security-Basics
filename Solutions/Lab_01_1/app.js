const fs = require('fs');
const argv = require('minimist')(process.argv.slice(2));
const otosc = require('./simple');
//const alphabit1 = "abcdefghijklmnopqrstuvwxyz";

const {file, mode, alg, key} = argv;
if (mode === 'decode' && !key) {
    throw new Error('Provide key for decoding');

}

const fileContent = fs.readFileSync(file, 'utf-8');
let keyText;
if (key) {
    keyText = fs.readFileSync(key, 'utf-8');
}

function run(alg, mode, text, key) {
    if (alg === 'otosc') {
        if (mode === 'encode') {
            const {result, key} = otosc.encode(text);
            writeFile('output.txt', result);
            writeFile('key.txt', key);
        }
        if (mode === 'decode') {
            const result = otosc.decode(text, key);
            writeFile('output.txt', result);
            writeFile('key.txt', key);
        }
    }
}

run(alg, mode, fileContent, keyText);

function writeFile(fileName, content) {
    fs.writeFileSync(fileName, content);
}