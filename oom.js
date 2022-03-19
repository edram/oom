var arr = [];

const sleep = (ms) => {
  return new Promise((resolve) => {
    setTimeout(resolve, ms);
  });
};

const run = () => {
  return sleep(0).then(() => {
    arr.push(10);
    return run();
  });
};

console.log(process.pid);

for (let index = 0; index < 10000; index++) {
  run();
}
