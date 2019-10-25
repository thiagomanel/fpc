global.Promise = require('bluebird');
import SharedBuffer from './shared-buffer';
 
const buffer = new SharedBuffer();
let count = 0;

function producer() {
    console.log('producer');
    buffer.put(++count)
        .then(_ => {
            const timeout = Math.random() * 1000;
            setTimeout(producer, timeout);
        });
}
 
function consumer() {
    console.log('consumer');
    buffer.get()
        .then(value => {
            console.log('read', value);
 
            const timeout = Math.random() * 1000;
            setTimeout(consumer, timeout);
        });
}
 
producer();
consumer();